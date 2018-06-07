/*
Copyright (c) 2017 SAP SE or an SAP affiliate company. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package infraclient

import (
	"context"
	"errors"
	"io"
	"time"

	pb "github.com/gardener/machine-controller-manager/pkg/grpc/infrapb"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// ExternalDriver structure mediates the communication with the machine-controller-manager
type ExternalDriver struct {
	serverAddr string
	options    []grpc.DialOption
	provider   ExternalDriverProvider
	connection *grpc.ClientConn
	client     pb.InfragrpcClient
	stream     pb.Infragrpc_RegisterClient
}

// NewExternalDriver creates a new Driver instance.
func NewExternalDriver(serverAddr string, options []grpc.DialOption, provider ExternalDriverProvider) *ExternalDriver {
	return &ExternalDriver{
		serverAddr: serverAddr,
		options:    options,
		provider:   provider,
	}
}

// Start starts the external driver.
func (d *ExternalDriver) Start() error {
	conn, err := grpc.Dial(d.serverAddr, d.options...)
	if err != nil {
		glog.Fatalf("fail to dial: %v", err)
		return err
	}
	d.connection = conn
	client := pb.NewInfragrpcClient(conn)
	d.client = client

	go func() {
		d.serveMCM(client)
	}()

	return nil
}

// Stop stops the external driver.
func (d *ExternalDriver) Stop() error {
	stream := d.stream
	//connection := d.connection

	d.stream = nil
	d.connection = nil

	if stream != nil && stream.Context().Err() == nil {
		stream.Send(&pb.DriverSide{
			OperationType: "unregister",
		})
		stream.CloseSend()
	}
	var err error
	/*
		if connection != nil {
			err = connection.Close()
		}
	*/

	return err
}

func (d *ExternalDriver) serveMCM(client pb.InfragrpcClient) error {
	glog.Infof("Registering with MCM...")
	ctx := context.Background()

	stream, err := client.Register(ctx)
	if err != nil {
		glog.Fatalf("%v.Register(_) = _, %v: ", client, err)
		return err
	}

	d.stream = stream

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			// read done.
			return err
		}
		if err != nil {
			glog.Fatalf("Failed to receive: %v", err)
			return err
		}

		glog.Infof("Operation %s", in.OperationType)
		opParams := in.GetOperationparams()
		glog.Infof("create parameters: %v", opParams)

		resp := pb.DriverSide{}
		resp.OperationID = in.OperationID
		resp.OperationType = in.OperationType

		switch in.OperationType {
		case "register":
			machineClassType := d.provider.Register(d)
			pMachineClassType := &machineClassType
			gvk := pMachineClassType.GroupVersionKind()
			resp.Response = &pb.DriverSide_RegisterResp{
				RegisterResp: &pb.DriverSideRegisterationResp{
					Name:    "externalDriver",
					Kind:    gvk.Kind,
					Group:   gvk.Group,
					Version: gvk.Version,
				},
			}
		case "create":
			var machineClass *MachineClassMeta
			if opParams.MachineClassMetaData != nil {
				machineClass = &MachineClassMeta{
					Name:     opParams.MachineClassMetaData.Name,
					Revision: opParams.MachineClassMetaData.Revision,
				}
			}
			providerID, nodename, err := d.provider.Create(machineClass, opParams.Credentials, opParams.MachineID, opParams.MachineName)

			var sErr string
			if err != nil {
				sErr = err.Error()
			}
			resp.Response = &pb.DriverSide_Createresponse{
				Createresponse: &pb.DriverSideCreateResp{
					ProviderID: providerID,
					Nodename:   nodename,
					Error:      sErr,
				},
			}
		case "delete":
			var machineClass *MachineClassMeta
			if opParams.MachineClassMetaData != nil {
				machineClass = &MachineClassMeta{
					Name:     opParams.MachineClassMetaData.Name,
					Revision: opParams.MachineClassMetaData.Revision,
				}
			}
			err := d.provider.Delete(machineClass, opParams.Credentials, opParams.MachineID)
			var sErr string
			if err != nil {
				sErr = err.Error()
			}
			resp.Response = &pb.DriverSide_Deleteresponse{
				Deleteresponse: &pb.DriverSideDeleteResp{
					Error: sErr,
				},
			}
		case "list":
			var machineClass *MachineClassMeta
			if opParams.MachineClassMetaData != nil {
				machineClass = &MachineClassMeta{
					Name:     opParams.MachineClassMetaData.Name,
					Revision: opParams.MachineClassMetaData.Revision,
				}
			}
			vms, err := d.provider.List(machineClass, opParams.Credentials, opParams.MachineID)
			var list []*pb.DriverSideMachine
			var machine *pb.DriverSideMachine

			var sErr string
			if err == nil {
				size := len(vms)
				list = make([]*pb.DriverSideMachine, size)
				i := 0
				for machineID, machineName := range vms {
					machine = new(pb.DriverSideMachine)
					machine.MachineID = machineID
					machine.MachineName = machineName
					list[i] = machine
					i++
				}
			} else {
				sErr = err.Error()
			}
			resp.Response = &pb.DriverSide_Listresponse{
				Listresponse: &pb.DriverSideListResp{
					List:  list,
					Error: sErr,
				},
			}
		}

		stream.Send(&resp)
	}
}

// GetMachineClass contacts the grpc server to get the machine class.
func (d *ExternalDriver) GetMachineClass(machineClassMeta *MachineClassMeta) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resp, err := d.client.GetMachineClass(ctx, &pb.MachineClassMeta{
		Name:     machineClassMeta.Name,
		Revision: machineClassMeta.Revision,
	})

	if err != nil {
		return nil, err
	}

	sErr := resp.Error
	if sErr != "" {
		return nil, errors.New(sErr)
	}

	return resp.Data, nil
}

// GetCloudConfig contacts the grpc server to get the cloud config.
func (d *ExternalDriver) GetCloudConfig(machineClassMeta *MachineClassMeta) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resp, err := d.client.GetCloudConfig(ctx, &pb.CloudConfigMeta{
		MachineClassMeta: &pb.MachineClassMeta{
			Name:     machineClassMeta.Name,
			Revision: machineClassMeta.Revision,
		},
	})

	if err != nil {
		return "", err
	}

	sErr := resp.Error
	if sErr != "" {
		return "", errors.New(sErr)
	}

	return resp.Data, nil
}

// Code generated by go-swagger; DO NOT EDIT.

package partition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new partition API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for partition API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreatePartition creates a partition if the given ID already exists a conflict is returned
*/
func (a *Client) CreatePartition(params *CreatePartitionParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePartitionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePartitionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPartition",
		Method:             "PUT",
		PathPattern:        "/v1/partition",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePartitionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePartitionCreated), nil

}

/*
DeletePartition deletes a partition and returns the deleted entity
*/
func (a *Client) DeletePartition(params *DeletePartitionParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePartitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePartitionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePartition",
		Method:             "DELETE",
		PathPattern:        "/v1/partition/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePartitionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePartitionOK), nil

}

/*
FindPartition gets partition by id
*/
func (a *Client) FindPartition(params *FindPartitionParams, authInfo runtime.ClientAuthInfoWriter) (*FindPartitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPartitionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findPartition",
		Method:             "GET",
		PathPattern:        "/v1/partition/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindPartitionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindPartitionOK), nil

}

/*
ListPartitions gets all partitions
*/
func (a *Client) ListPartitions(params *ListPartitionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPartitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPartitionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPartitions",
		Method:             "GET",
		PathPattern:        "/v1/partition",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPartitionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPartitionsOK), nil

}

/*
PartitionCapacity gets partition capacity
*/
func (a *Client) PartitionCapacity(params *PartitionCapacityParams, authInfo runtime.ClientAuthInfoWriter) (*PartitionCapacityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartitionCapacityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "partitionCapacity",
		Method:             "GET",
		PathPattern:        "/v1/partition/capacity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PartitionCapacityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PartitionCapacityOK), nil

}

/*
UpdatePartition updates a partition if the partition was changed since this one was read a conflict is returned
*/
func (a *Client) UpdatePartition(params *UpdatePartitionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePartitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePartitionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePartition",
		Method:             "POST",
		PathPattern:        "/v1/partition",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePartitionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdatePartitionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

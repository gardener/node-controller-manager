/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	scheme "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PacketMachineClassesGetter has a method to return a PacketMachineClassInterface.
// A group's client should implement this interface.
type PacketMachineClassesGetter interface {
	PacketMachineClasses(namespace string) PacketMachineClassInterface
}

// PacketMachineClassInterface has methods to work with PacketMachineClass resources.
type PacketMachineClassInterface interface {
	Create(*v1alpha1.PacketMachineClass) (*v1alpha1.PacketMachineClass, error)
	Update(*v1alpha1.PacketMachineClass) (*v1alpha1.PacketMachineClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PacketMachineClass, error)
	List(opts v1.ListOptions) (*v1alpha1.PacketMachineClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PacketMachineClass, err error)
	PacketMachineClassExpansion
}

// packetMachineClasses implements PacketMachineClassInterface
type packetMachineClasses struct {
	client rest.Interface
	ns     string
}

// newPacketMachineClasses returns a PacketMachineClasses
func newPacketMachineClasses(c *MachineV1alpha1Client, namespace string) *packetMachineClasses {
	return &packetMachineClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the packetMachineClass, and returns the corresponding packetMachineClass object, and an error if there is any.
func (c *packetMachineClasses) Get(name string, options v1.GetOptions) (result *v1alpha1.PacketMachineClass, err error) {
	result = &v1alpha1.PacketMachineClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("packetmachineclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PacketMachineClasses that match those selectors.
func (c *packetMachineClasses) List(opts v1.ListOptions) (result *v1alpha1.PacketMachineClassList, err error) {
	result = &v1alpha1.PacketMachineClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("packetmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested packetMachineClasses.
func (c *packetMachineClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("packetmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a packetMachineClass and creates it.  Returns the server's representation of the packetMachineClass, and an error, if there is any.
func (c *packetMachineClasses) Create(packetMachineClass *v1alpha1.PacketMachineClass) (result *v1alpha1.PacketMachineClass, err error) {
	result = &v1alpha1.PacketMachineClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("packetmachineclasses").
		Body(packetMachineClass).
		Do().
		Into(result)
	return
}

// Update takes the representation of a packetMachineClass and updates it. Returns the server's representation of the packetMachineClass, and an error, if there is any.
func (c *packetMachineClasses) Update(packetMachineClass *v1alpha1.PacketMachineClass) (result *v1alpha1.PacketMachineClass, err error) {
	result = &v1alpha1.PacketMachineClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("packetmachineclasses").
		Name(packetMachineClass.Name).
		Body(packetMachineClass).
		Do().
		Into(result)
	return
}

// Delete takes name of the packetMachineClass and deletes it. Returns an error if one occurs.
func (c *packetMachineClasses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("packetmachineclasses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *packetMachineClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("packetmachineclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched packetMachineClass.
func (c *packetMachineClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PacketMachineClass, err error) {
	result = &v1alpha1.PacketMachineClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("packetmachineclasses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

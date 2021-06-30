/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GCPMachineClassLister helps list GCPMachineClasses.
// All objects returned here must be treated as read-only.
type GCPMachineClassLister interface {
	// List lists all GCPMachineClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error)
	// GCPMachineClasses returns an object that can list and get GCPMachineClasses.
	GCPMachineClasses(namespace string) GCPMachineClassNamespaceLister
	GCPMachineClassListerExpansion
}

// gCPMachineClassLister implements the GCPMachineClassLister interface.
type gCPMachineClassLister struct {
	indexer cache.Indexer
}

// NewGCPMachineClassLister returns a new GCPMachineClassLister.
func NewGCPMachineClassLister(indexer cache.Indexer) GCPMachineClassLister {
	return &gCPMachineClassLister{indexer: indexer}
}

// List lists all GCPMachineClasses in the indexer.
func (s *gCPMachineClassLister) List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.GCPMachineClass))
	})
	return ret, err
}

// GCPMachineClasses returns an object that can list and get GCPMachineClasses.
func (s *gCPMachineClassLister) GCPMachineClasses(namespace string) GCPMachineClassNamespaceLister {
	return gCPMachineClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GCPMachineClassNamespaceLister helps list and get GCPMachineClasses.
// All objects returned here must be treated as read-only.
type GCPMachineClassNamespaceLister interface {
	// List lists all GCPMachineClasses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error)
	// Get retrieves the GCPMachineClass from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*machine.GCPMachineClass, error)
	GCPMachineClassNamespaceListerExpansion
}

// gCPMachineClassNamespaceLister implements the GCPMachineClassNamespaceLister
// interface.
type gCPMachineClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GCPMachineClasses in the indexer for a given namespace.
func (s gCPMachineClassNamespaceLister) List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.GCPMachineClass))
	})
	return ret, err
}

// Get retrieves the GCPMachineClass from the indexer for a given namespace and name.
func (s gCPMachineClassNamespaceLister) Get(name string) (*machine.GCPMachineClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("gcpmachineclass"), name)
	}
	return obj.(*machine.GCPMachineClass), nil
}

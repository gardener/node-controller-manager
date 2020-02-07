// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	clientsetinternalversion "github.com/gardener/machine-controller-manager/pkg/client/clientset/internalversion"
	internalinterfaces "github.com/gardener/machine-controller-manager/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/gardener/machine-controller-manager/pkg/client/listers/machine/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MetalMachineClassInformer provides access to a shared informer and lister for
// MetalMachineClasses.
type MetalMachineClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.MetalMachineClassLister
}

type metalMachineClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMetalMachineClassInformer constructs a new informer for MetalMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMetalMachineClassInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMetalMachineClassInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMetalMachineClassInformer constructs a new informer for MetalMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMetalMachineClassInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Machine().MetalMachineClasses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Machine().MetalMachineClasses(namespace).Watch(options)
			},
		},
		&machine.MetalMachineClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *metalMachineClassInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMetalMachineClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *metalMachineClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machine.MetalMachineClass{}, f.defaultInformer)
}

func (f *metalMachineClassInformer) Lister() internalversion.MetalMachineClassLister {
	return internalversion.NewMetalMachineClassLister(f.Informer().GetIndexer())
}

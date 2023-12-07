// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	machineconfigurationv1alpha1 "github.com/openshift/api/machineconfiguration/v1alpha1"
	versioned "github.com/openshift/client-go/machineconfiguration/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/machineconfiguration/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/client-go/machineconfiguration/listers/machineconfiguration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MachineConfigNodeInformer provides access to a shared informer and lister for
// MachineConfigNodes.
type MachineConfigNodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineConfigNodeLister
}

type machineConfigNodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewMachineConfigNodeInformer constructs a new informer for MachineConfigNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineConfigNodeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineConfigNodeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredMachineConfigNodeInformer constructs a new informer for MachineConfigNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineConfigNodeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineconfigurationV1alpha1().MachineConfigNodes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineconfigurationV1alpha1().MachineConfigNodes().Watch(context.TODO(), options)
			},
		},
		&machineconfigurationv1alpha1.MachineConfigNode{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineConfigNodeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineConfigNodeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineConfigNodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machineconfigurationv1alpha1.MachineConfigNode{}, f.defaultInformer)
}

func (f *machineConfigNodeInformer) Lister() v1alpha1.MachineConfigNodeLister {
	return v1alpha1.NewMachineConfigNodeLister(f.Informer().GetIndexer())
}
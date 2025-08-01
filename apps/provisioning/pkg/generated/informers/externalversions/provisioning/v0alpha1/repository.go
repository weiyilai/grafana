// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by informer-gen. DO NOT EDIT.

package v0alpha1

import (
	context "context"
	time "time"

	apisprovisioningv0alpha1 "github.com/grafana/grafana/apps/provisioning/pkg/apis/provisioning/v0alpha1"
	versioned "github.com/grafana/grafana/apps/provisioning/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/grafana/grafana/apps/provisioning/pkg/generated/informers/externalversions/internalinterfaces"
	provisioningv0alpha1 "github.com/grafana/grafana/apps/provisioning/pkg/generated/listers/provisioning/v0alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RepositoryInformer provides access to a shared informer and lister for
// Repositories.
type RepositoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() provisioningv0alpha1.RepositoryLister
}

type repositoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRepositoryInformer constructs a new informer for Repository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRepositoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRepositoryInformer constructs a new informer for Repository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProvisioningV0alpha1().Repositories(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProvisioningV0alpha1().Repositories(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProvisioningV0alpha1().Repositories(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProvisioningV0alpha1().Repositories(namespace).Watch(ctx, options)
			},
		},
		&apisprovisioningv0alpha1.Repository{},
		resyncPeriod,
		indexers,
	)
}

func (f *repositoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRepositoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *repositoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisprovisioningv0alpha1.Repository{}, f.defaultInformer)
}

func (f *repositoryInformer) Lister() provisioningv0alpha1.RepositoryLister {
	return provisioningv0alpha1.NewRepositoryLister(f.Informer().GetIndexer())
}

/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha

import (
	"context"
	customv1alpha "custom-apiserver/pkg/apis/custom/v1alpha"
	versioned "custom-apiserver/pkg/generated/clientset/versioned"
	internalinterfaces "custom-apiserver/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha "custom-apiserver/pkg/generated/listers/custom/v1alpha"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// JenkinsServiceInformer provides access to a shared informer and lister for
// JenkinsServices.
type JenkinsServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha.JenkinsServiceLister
}

type jenkinsServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewJenkinsServiceInformer constructs a new informer for JenkinsService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewJenkinsServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredJenkinsServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredJenkinsServiceInformer constructs a new informer for JenkinsService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredJenkinsServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutobusiV1alpha().JenkinsServices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutobusiV1alpha().JenkinsServices(namespace).Watch(context.TODO(), options)
			},
		},
		&customv1alpha.JenkinsService{},
		resyncPeriod,
		indexers,
	)
}

func (f *jenkinsServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredJenkinsServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *jenkinsServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&customv1alpha.JenkinsService{}, f.defaultInformer)
}

func (f *jenkinsServiceInformer) Lister() v1alpha.JenkinsServiceLister {
	return v1alpha.NewJenkinsServiceLister(f.Informer().GetIndexer())
}

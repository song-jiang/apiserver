// Copyright (c) 2019 Tigera, Inc. All rights reserved.

package calico

import (
	"reflect"

	"golang.org/x/net/context"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/storage"
	etcd "k8s.io/apiserver/pkg/storage/etcd3"
	"k8s.io/apiserver/pkg/storage/storagebackend/factory"

	libcalicoapi "github.com/projectcalico/libcalico-go/lib/apis/v3"
	"github.com/projectcalico/libcalico-go/lib/clientv3"
	"github.com/projectcalico/libcalico-go/lib/options"
	"github.com/projectcalico/libcalico-go/lib/watch"

	aapi "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
)

// NewClusterInformationStorage creates a new libcalico-based storage.Interface implementation for ClusterInformation
func NewClusterInformationStorage(opts Options) (registry.DryRunnableStorage, factory.DestroyFunc) {
	c := CreateClientFromConfig()
	getFn := func(ctx context.Context, c clientv3.Interface, ns string, name string, opts clientOpts) (resourceObject, error) {
		ogo := opts.(options.GetOptions)
		return c.ClusterInformation().Get(ctx, name, ogo)
	}
	listFn := func(ctx context.Context, c clientv3.Interface, opts clientOpts) (resourceListObject, error) {
		olo := opts.(options.ListOptions)
		return c.ClusterInformation().List(ctx, olo)
	}
	watchFn := func(ctx context.Context, c clientv3.Interface, opts clientOpts) (watch.Interface, error) {
		olo := opts.(options.ListOptions)
		return c.ClusterInformation().Watch(ctx, olo)
	}
	dryRunnableStorage := registry.DryRunnableStorage{Storage: &resourceStore{
		client:            c,
		codec:             opts.RESTOptions.StorageConfig.Codec,
		versioner:         etcd.APIObjectVersioner{},
		aapiType:          reflect.TypeOf(aapi.ClusterInformation{}),
		aapiListType:      reflect.TypeOf(aapi.ClusterInformationList{}),
		libCalicoType:     reflect.TypeOf(libcalicoapi.ClusterInformation{}),
		libCalicoListType: reflect.TypeOf(libcalicoapi.ClusterInformationList{}),
		isNamespaced:      false,
		get:               getFn,
		list:              listFn,
		watch:             watchFn,
		resourceName:      "ClusterInformation",
		converter:         ClusterInformationConverter{},
	}, Codec: opts.RESTOptions.StorageConfig.Codec}
	return dryRunnableStorage, func() {}
}

type ClusterInformationConverter struct {
}

func (gc ClusterInformationConverter) convertToLibcalico(aapiObj runtime.Object) resourceObject {
	aapiClusterInformation := aapiObj.(*aapi.ClusterInformation)
	lcgClusterInformation := &libcalicoapi.ClusterInformation{}
	lcgClusterInformation.TypeMeta = aapiClusterInformation.TypeMeta
	lcgClusterInformation.ObjectMeta = aapiClusterInformation.ObjectMeta
	lcgClusterInformation.Kind = libcalicoapi.KindClusterInformation
	lcgClusterInformation.APIVersion = libcalicoapi.GroupVersionCurrent
	lcgClusterInformation.Spec = aapiClusterInformation.Spec
	return lcgClusterInformation
}

func (gc ClusterInformationConverter) convertToAAPI(libcalicoObject resourceObject, aapiObj runtime.Object) {
	lcgClusterInformation := libcalicoObject.(*libcalicoapi.ClusterInformation)
	aapiClusterInformation := aapiObj.(*aapi.ClusterInformation)
	aapiClusterInformation.Spec = lcgClusterInformation.Spec
	aapiClusterInformation.TypeMeta = lcgClusterInformation.TypeMeta
	aapiClusterInformation.ObjectMeta = lcgClusterInformation.ObjectMeta
}

func (gc ClusterInformationConverter) convertToAAPIList(libcalicoListObject resourceListObject, aapiListObj runtime.Object, pred storage.SelectionPredicate) {
	lcgClusterInformationList := libcalicoListObject.(*libcalicoapi.ClusterInformationList)
	aapiClusterInformationList := aapiListObj.(*aapi.ClusterInformationList)
	if libcalicoListObject == nil {
		aapiClusterInformationList.Items = []aapi.ClusterInformation{}
		return
	}
	aapiClusterInformationList.TypeMeta = lcgClusterInformationList.TypeMeta
	aapiClusterInformationList.ListMeta = lcgClusterInformationList.ListMeta
	for _, item := range lcgClusterInformationList.Items {
		aapiClusterInformation := aapi.ClusterInformation{}
		gc.convertToAAPI(&item, &aapiClusterInformation)
		if matched, err := pred.Matches(&aapiClusterInformation); err == nil && matched {
			aapiClusterInformationList.Items = append(aapiClusterInformationList.Items, aapiClusterInformation)
		}
	}
}

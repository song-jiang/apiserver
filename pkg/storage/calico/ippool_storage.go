// Copyright (c) 2017-2019 Tigera, Inc. All rights reserved.

package calico

import (
	"reflect"

	"golang.org/x/net/context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/storage"
	etcd "k8s.io/apiserver/pkg/storage/etcd3"
	"k8s.io/apiserver/pkg/storage/storagebackend/factory"

	aapi "github.com/projectcalico/apiserver/pkg/apis/projectcalico"

	libcalicoapi "github.com/projectcalico/libcalico-go/lib/apis/v3"
	"github.com/projectcalico/libcalico-go/lib/clientv3"
	"github.com/projectcalico/libcalico-go/lib/options"
	"github.com/projectcalico/libcalico-go/lib/watch"
)

// NewIPPoolStorage creates a new libcalico-based storage.Interface implementation for IPPools
func NewIPPoolStorage(opts Options) (registry.DryRunnableStorage, factory.DestroyFunc) {
	c := CreateClientFromConfig()
	createFn := func(ctx context.Context, c clientv3.Interface, obj resourceObject, opts clientOpts) (resourceObject, error) {
		oso := opts.(options.SetOptions)
		res := obj.(*libcalicoapi.IPPool)
		return c.IPPools().Create(ctx, res, oso)
	}
	updateFn := func(ctx context.Context, c clientv3.Interface, obj resourceObject, opts clientOpts) (resourceObject, error) {
		oso := opts.(options.SetOptions)
		res := obj.(*libcalicoapi.IPPool)
		return c.IPPools().Update(ctx, res, oso)
	}
	getFn := func(ctx context.Context, c clientv3.Interface, ns string, name string, opts clientOpts) (resourceObject, error) {
		ogo := opts.(options.GetOptions)
		return c.IPPools().Get(ctx, name, ogo)
	}
	deleteFn := func(ctx context.Context, c clientv3.Interface, ns string, name string, opts clientOpts) (resourceObject, error) {
		odo := opts.(options.DeleteOptions)
		return c.IPPools().Delete(ctx, name, odo)
	}
	listFn := func(ctx context.Context, c clientv3.Interface, opts clientOpts) (resourceListObject, error) {
		olo := opts.(options.ListOptions)
		return c.IPPools().List(ctx, olo)
	}
	watchFn := func(ctx context.Context, c clientv3.Interface, opts clientOpts) (watch.Interface, error) {
		olo := opts.(options.ListOptions)
		return c.IPPools().Watch(ctx, olo)
	}
	dryRunnableStorage := registry.DryRunnableStorage{Storage: &resourceStore{
		client:            c,
		codec:             opts.RESTOptions.StorageConfig.Codec,
		versioner:         etcd.APIObjectVersioner{},
		aapiType:          reflect.TypeOf(aapi.IPPool{}),
		aapiListType:      reflect.TypeOf(aapi.IPPoolList{}),
		libCalicoType:     reflect.TypeOf(libcalicoapi.IPPool{}),
		libCalicoListType: reflect.TypeOf(libcalicoapi.IPPoolList{}),
		isNamespaced:      false,
		create:            createFn,
		update:            updateFn,
		get:               getFn,
		delete:            deleteFn,
		list:              listFn,
		watch:             watchFn,
		resourceName:      "IPPool",
		converter:         IPPoolConverter{},
	}, Codec: opts.RESTOptions.StorageConfig.Codec}
	return dryRunnableStorage, func() {}
}

type IPPoolConverter struct {
}

func (gc IPPoolConverter) convertToLibcalico(aapiObj runtime.Object) resourceObject {
	aapiIPPool := aapiObj.(*aapi.IPPool)
	lcgIPPool := &libcalicoapi.IPPool{}
	lcgIPPool.TypeMeta = aapiIPPool.TypeMeta
	lcgIPPool.ObjectMeta = aapiIPPool.ObjectMeta
	lcgIPPool.Kind = libcalicoapi.KindIPPool
	lcgIPPool.APIVersion = libcalicoapi.GroupVersionCurrent
	lcgIPPool.Spec = aapiIPPool.Spec
	return lcgIPPool
}

func (gc IPPoolConverter) convertToAAPI(libcalicoObject resourceObject, aapiObj runtime.Object) {
	lcgIPPool := libcalicoObject.(*libcalicoapi.IPPool)
	aapiIPPool := aapiObj.(*aapi.IPPool)
	aapiIPPool.Spec = lcgIPPool.Spec
	aapiIPPool.TypeMeta = lcgIPPool.TypeMeta
	aapiIPPool.ObjectMeta = lcgIPPool.ObjectMeta
}

func (gc IPPoolConverter) convertToAAPIList(libcalicoListObject resourceListObject, aapiListObj runtime.Object, pred storage.SelectionPredicate) {
	lcgIPPoolList := libcalicoListObject.(*libcalicoapi.IPPoolList)
	aapiIPPoolList := aapiListObj.(*aapi.IPPoolList)
	if libcalicoListObject == nil {
		aapiIPPoolList.Items = []aapi.IPPool{}
		return
	}
	aapiIPPoolList.TypeMeta = lcgIPPoolList.TypeMeta
	aapiIPPoolList.ListMeta = lcgIPPoolList.ListMeta
	for _, item := range lcgIPPoolList.Items {
		aapiIPPool := aapi.IPPool{}
		gc.convertToAAPI(&item, &aapiIPPool)
		if matched, err := pred.Matches(&aapiIPPool); err == nil && matched {
			aapiIPPoolList.Items = append(aapiIPPoolList.Items, aapiIPPool)
		}
	}
}

// Copyright (c) 2021 Tigera, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	projectcalico "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
)

// FakeIPPools implements IPPoolInterface
type FakeIPPools struct {
	Fake *FakeProjectcalico
}

var ippoolsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "", Resource: "ippools"}

var ippoolsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "", Kind: "IPPool"}

// Get takes name of the iPPool, and returns the corresponding iPPool object, and an error if there is any.
func (c *FakeIPPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.IPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ippoolsResource, name), &projectcalico.IPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.IPPool), err
}

// List takes label and field selectors, and returns the list of IPPools that match those selectors.
func (c *FakeIPPools) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.IPPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ippoolsResource, ippoolsKind, opts), &projectcalico.IPPoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &projectcalico.IPPoolList{ListMeta: obj.(*projectcalico.IPPoolList).ListMeta}
	for _, item := range obj.(*projectcalico.IPPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPPools.
func (c *FakeIPPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ippoolsResource, opts))
}

// Create takes the representation of a iPPool and creates it.  Returns the server's representation of the iPPool, and an error, if there is any.
func (c *FakeIPPools) Create(ctx context.Context, iPPool *projectcalico.IPPool, opts v1.CreateOptions) (result *projectcalico.IPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ippoolsResource, iPPool), &projectcalico.IPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.IPPool), err
}

// Update takes the representation of a iPPool and updates it. Returns the server's representation of the iPPool, and an error, if there is any.
func (c *FakeIPPools) Update(ctx context.Context, iPPool *projectcalico.IPPool, opts v1.UpdateOptions) (result *projectcalico.IPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ippoolsResource, iPPool), &projectcalico.IPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.IPPool), err
}

// Delete takes name of the iPPool and deletes it. Returns an error if one occurs.
func (c *FakeIPPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ippoolsResource, name), &projectcalico.IPPool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ippoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &projectcalico.IPPoolList{})
	return err
}

// Patch applies the patch and returns the patched iPPool.
func (c *FakeIPPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.IPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ippoolsResource, name, pt, data, subresources...), &projectcalico.IPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.IPPool), err
}

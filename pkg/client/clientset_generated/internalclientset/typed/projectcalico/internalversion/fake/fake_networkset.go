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

// FakeNetworkSets implements NetworkSetInterface
type FakeNetworkSets struct {
	Fake *FakeProjectcalico
}

var networksetsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "", Resource: "networksets"}

var networksetsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "", Kind: "NetworkSet"}

// Get takes name of the networkSet, and returns the corresponding networkSet object, and an error if there is any.
func (c *FakeNetworkSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.NetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(networksetsResource, name), &projectcalico.NetworkSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.NetworkSet), err
}

// List takes label and field selectors, and returns the list of NetworkSets that match those selectors.
func (c *FakeNetworkSets) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.NetworkSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(networksetsResource, networksetsKind, opts), &projectcalico.NetworkSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &projectcalico.NetworkSetList{ListMeta: obj.(*projectcalico.NetworkSetList).ListMeta}
	for _, item := range obj.(*projectcalico.NetworkSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkSets.
func (c *FakeNetworkSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(networksetsResource, opts))
}

// Create takes the representation of a networkSet and creates it.  Returns the server's representation of the networkSet, and an error, if there is any.
func (c *FakeNetworkSets) Create(ctx context.Context, networkSet *projectcalico.NetworkSet, opts v1.CreateOptions) (result *projectcalico.NetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(networksetsResource, networkSet), &projectcalico.NetworkSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.NetworkSet), err
}

// Update takes the representation of a networkSet and updates it. Returns the server's representation of the networkSet, and an error, if there is any.
func (c *FakeNetworkSets) Update(ctx context.Context, networkSet *projectcalico.NetworkSet, opts v1.UpdateOptions) (result *projectcalico.NetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(networksetsResource, networkSet), &projectcalico.NetworkSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.NetworkSet), err
}

// Delete takes name of the networkSet and deletes it. Returns an error if one occurs.
func (c *FakeNetworkSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(networksetsResource, name), &projectcalico.NetworkSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(networksetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &projectcalico.NetworkSetList{})
	return err
}

// Patch applies the patch and returns the patched networkSet.
func (c *FakeNetworkSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.NetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(networksetsResource, name, pt, data, subresources...), &projectcalico.NetworkSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.NetworkSet), err
}

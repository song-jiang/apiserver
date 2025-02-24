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

// FakeClusterInformations implements ClusterInformationInterface
type FakeClusterInformations struct {
	Fake *FakeProjectcalico
}

var clusterinformationsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "", Resource: "clusterinformations"}

var clusterinformationsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "", Kind: "ClusterInformation"}

// Get takes name of the clusterInformation, and returns the corresponding clusterInformation object, and an error if there is any.
func (c *FakeClusterInformations) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.ClusterInformation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterinformationsResource, name), &projectcalico.ClusterInformation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.ClusterInformation), err
}

// List takes label and field selectors, and returns the list of ClusterInformations that match those selectors.
func (c *FakeClusterInformations) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.ClusterInformationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterinformationsResource, clusterinformationsKind, opts), &projectcalico.ClusterInformationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &projectcalico.ClusterInformationList{ListMeta: obj.(*projectcalico.ClusterInformationList).ListMeta}
	for _, item := range obj.(*projectcalico.ClusterInformationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterInformations.
func (c *FakeClusterInformations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterinformationsResource, opts))
}

// Create takes the representation of a clusterInformation and creates it.  Returns the server's representation of the clusterInformation, and an error, if there is any.
func (c *FakeClusterInformations) Create(ctx context.Context, clusterInformation *projectcalico.ClusterInformation, opts v1.CreateOptions) (result *projectcalico.ClusterInformation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterinformationsResource, clusterInformation), &projectcalico.ClusterInformation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.ClusterInformation), err
}

// Update takes the representation of a clusterInformation and updates it. Returns the server's representation of the clusterInformation, and an error, if there is any.
func (c *FakeClusterInformations) Update(ctx context.Context, clusterInformation *projectcalico.ClusterInformation, opts v1.UpdateOptions) (result *projectcalico.ClusterInformation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterinformationsResource, clusterInformation), &projectcalico.ClusterInformation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.ClusterInformation), err
}

// Delete takes name of the clusterInformation and deletes it. Returns an error if one occurs.
func (c *FakeClusterInformations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterinformationsResource, name), &projectcalico.ClusterInformation{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterInformations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterinformationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &projectcalico.ClusterInformationList{})
	return err
}

// Patch applies the patch and returns the patched clusterInformation.
func (c *FakeClusterInformations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.ClusterInformation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterinformationsResource, name, pt, data, subresources...), &projectcalico.ClusterInformation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.ClusterInformation), err
}

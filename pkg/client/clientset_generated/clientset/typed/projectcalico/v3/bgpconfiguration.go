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

package v3

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v3 "github.com/projectcalico/apiserver/pkg/apis/projectcalico/v3"
	scheme "github.com/projectcalico/apiserver/pkg/client/clientset_generated/clientset/scheme"
)

// BGPConfigurationsGetter has a method to return a BGPConfigurationInterface.
// A group's client should implement this interface.
type BGPConfigurationsGetter interface {
	BGPConfigurations() BGPConfigurationInterface
}

// BGPConfigurationInterface has methods to work with BGPConfiguration resources.
type BGPConfigurationInterface interface {
	Create(ctx context.Context, bGPConfiguration *v3.BGPConfiguration, opts v1.CreateOptions) (*v3.BGPConfiguration, error)
	Update(ctx context.Context, bGPConfiguration *v3.BGPConfiguration, opts v1.UpdateOptions) (*v3.BGPConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.BGPConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.BGPConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.BGPConfiguration, err error)
	BGPConfigurationExpansion
}

// bGPConfigurations implements BGPConfigurationInterface
type bGPConfigurations struct {
	client rest.Interface
}

// newBGPConfigurations returns a BGPConfigurations
func newBGPConfigurations(c *ProjectcalicoV3Client) *bGPConfigurations {
	return &bGPConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the bGPConfiguration, and returns the corresponding bGPConfiguration object, and an error if there is any.
func (c *bGPConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.BGPConfiguration, err error) {
	result = &v3.BGPConfiguration{}
	err = c.client.Get().
		Resource("bgpconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BGPConfigurations that match those selectors.
func (c *bGPConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v3.BGPConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.BGPConfigurationList{}
	err = c.client.Get().
		Resource("bgpconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bGPConfigurations.
func (c *bGPConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("bgpconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bGPConfiguration and creates it.  Returns the server's representation of the bGPConfiguration, and an error, if there is any.
func (c *bGPConfigurations) Create(ctx context.Context, bGPConfiguration *v3.BGPConfiguration, opts v1.CreateOptions) (result *v3.BGPConfiguration, err error) {
	result = &v3.BGPConfiguration{}
	err = c.client.Post().
		Resource("bgpconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bGPConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bGPConfiguration and updates it. Returns the server's representation of the bGPConfiguration, and an error, if there is any.
func (c *bGPConfigurations) Update(ctx context.Context, bGPConfiguration *v3.BGPConfiguration, opts v1.UpdateOptions) (result *v3.BGPConfiguration, err error) {
	result = &v3.BGPConfiguration{}
	err = c.client.Put().
		Resource("bgpconfigurations").
		Name(bGPConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bGPConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bGPConfiguration and deletes it. Returns an error if one occurs.
func (c *bGPConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("bgpconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bGPConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("bgpconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bGPConfiguration.
func (c *bGPConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.BGPConfiguration, err error) {
	result = &v3.BGPConfiguration{}
	err = c.client.Patch(pt).
		Resource("bgpconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

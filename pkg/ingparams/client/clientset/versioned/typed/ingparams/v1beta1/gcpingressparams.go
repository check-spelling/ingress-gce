/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "k8s.io/ingress-gce/pkg/apis/ingparams/v1beta1"
	scheme "k8s.io/ingress-gce/pkg/ingparams/client/clientset/versioned/scheme"
)

// GCPIngressParamsGetter has a method to return a GCPIngressParamsInterface.
// A group's client should implement this interface.
type GCPIngressParamsGetter interface {
	GCPIngressParams() GCPIngressParamsInterface
}

// GCPIngressParamsInterface has methods to work with GCPIngressParams resources.
type GCPIngressParamsInterface interface {
	Create(ctx context.Context, gCPIngressParams *v1beta1.GCPIngressParams, opts v1.CreateOptions) (*v1beta1.GCPIngressParams, error)
	Update(ctx context.Context, gCPIngressParams *v1beta1.GCPIngressParams, opts v1.UpdateOptions) (*v1beta1.GCPIngressParams, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.GCPIngressParams, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.GCPIngressParamsList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.GCPIngressParams, err error)
	GCPIngressParamsExpansion
}

// gCPIngressParams implements GCPIngressParamsInterface
type gCPIngressParams struct {
	client rest.Interface
}

// newGCPIngressParams returns a GCPIngressParams
func newGCPIngressParams(c *NetworkingV1beta1Client) *gCPIngressParams {
	return &gCPIngressParams{
		client: c.RESTClient(),
	}
}

// Get takes name of the gCPIngressParams, and returns the corresponding gCPIngressParams object, and an error if there is any.
func (c *gCPIngressParams) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.GCPIngressParams, err error) {
	result = &v1beta1.GCPIngressParams{}
	err = c.client.Get().
		Resource("gcpingressparams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GCPIngressParams that match those selectors.
func (c *gCPIngressParams) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.GCPIngressParamsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.GCPIngressParamsList{}
	err = c.client.Get().
		Resource("gcpingressparams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gCPIngressParams.
func (c *gCPIngressParams) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("gcpingressparams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gCPIngressParams and creates it.  Returns the server's representation of the gCPIngressParams, and an error, if there is any.
func (c *gCPIngressParams) Create(ctx context.Context, gCPIngressParams *v1beta1.GCPIngressParams, opts v1.CreateOptions) (result *v1beta1.GCPIngressParams, err error) {
	result = &v1beta1.GCPIngressParams{}
	err = c.client.Post().
		Resource("gcpingressparams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gCPIngressParams).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gCPIngressParams and updates it. Returns the server's representation of the gCPIngressParams, and an error, if there is any.
func (c *gCPIngressParams) Update(ctx context.Context, gCPIngressParams *v1beta1.GCPIngressParams, opts v1.UpdateOptions) (result *v1beta1.GCPIngressParams, err error) {
	result = &v1beta1.GCPIngressParams{}
	err = c.client.Put().
		Resource("gcpingressparams").
		Name(gCPIngressParams.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gCPIngressParams).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gCPIngressParams and deletes it. Returns an error if one occurs.
func (c *gCPIngressParams) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("gcpingressparams").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gCPIngressParams) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("gcpingressparams").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gCPIngressParams.
func (c *gCPIngressParams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.GCPIngressParams, err error) {
	result = &v1beta1.GCPIngressParams{}
	err = c.client.Patch(pt).
		Resource("gcpingressparams").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

/*
Copyright The Karbour Authors.

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

	v1beta1 "github.com/KusionStack/karbour/pkg/kubernetes/apis/search/v1beta1"
	scheme "github.com/KusionStack/karbour/pkg/kubernetes/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TransformRulesGetter has a method to return a TransformRuleInterface.
// A group's client should implement this interface.
type TransformRulesGetter interface {
	TransformRules() TransformRuleInterface
}

// TransformRuleInterface has methods to work with TransformRule resources.
type TransformRuleInterface interface {
	Create(ctx context.Context, transformRule *v1beta1.TransformRule, opts v1.CreateOptions) (*v1beta1.TransformRule, error)
	Update(ctx context.Context, transformRule *v1beta1.TransformRule, opts v1.UpdateOptions) (*v1beta1.TransformRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.TransformRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.TransformRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TransformRule, err error)
	TransformRuleExpansion
}

// transformRules implements TransformRuleInterface
type transformRules struct {
	client rest.Interface
}

// newTransformRules returns a TransformRules
func newTransformRules(c *SearchV1beta1Client) *transformRules {
	return &transformRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the transformRule, and returns the corresponding transformRule object, and an error if there is any.
func (c *transformRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.TransformRule, err error) {
	result = &v1beta1.TransformRule{}
	err = c.client.Get().
		Resource("transformrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TransformRules that match those selectors.
func (c *transformRules) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.TransformRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.TransformRuleList{}
	err = c.client.Get().
		Resource("transformrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested transformRules.
func (c *transformRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("transformrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a transformRule and creates it.  Returns the server's representation of the transformRule, and an error, if there is any.
func (c *transformRules) Create(ctx context.Context, transformRule *v1beta1.TransformRule, opts v1.CreateOptions) (result *v1beta1.TransformRule, err error) {
	result = &v1beta1.TransformRule{}
	err = c.client.Post().
		Resource("transformrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(transformRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a transformRule and updates it. Returns the server's representation of the transformRule, and an error, if there is any.
func (c *transformRules) Update(ctx context.Context, transformRule *v1beta1.TransformRule, opts v1.UpdateOptions) (result *v1beta1.TransformRule, err error) {
	result = &v1beta1.TransformRule{}
	err = c.client.Put().
		Resource("transformrules").
		Name(transformRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(transformRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the transformRule and deletes it. Returns an error if one occurs.
func (c *transformRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("transformrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *transformRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("transformrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched transformRule.
func (c *transformRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TransformRule, err error) {
	result = &v1beta1.TransformRule{}
	err = c.client.Patch(pt).
		Resource("transformrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
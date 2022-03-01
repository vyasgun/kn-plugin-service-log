/*
Copyright 2021 The Knative Authors

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

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "knative.dev/serving/pkg/apis/serving/v1"
	scheme "knative.dev/serving/pkg/client/clientset/versioned/scheme"
)

// RevisionsGetter has a method to return a RevisionInterface.
// A group's client should implement this interface.
type RevisionsGetter interface {
	Revisions(namespace string) RevisionInterface
}

// RevisionInterface has methods to work with Revision resources.
type RevisionInterface interface {
	Create(ctx context.Context, revision *v1.Revision, opts metav1.CreateOptions) (*v1.Revision, error)
	Update(ctx context.Context, revision *v1.Revision, opts metav1.UpdateOptions) (*v1.Revision, error)
	UpdateStatus(ctx context.Context, revision *v1.Revision, opts metav1.UpdateOptions) (*v1.Revision, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Revision, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RevisionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Revision, err error)
	RevisionExpansion
}

// revisions implements RevisionInterface
type revisions struct {
	client rest.Interface
	ns     string
}

// newRevisions returns a Revisions
func newRevisions(c *ServingV1Client, namespace string) *revisions {
	return &revisions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the revision, and returns the corresponding revision object, and an error if there is any.
func (c *revisions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Revision, err error) {
	result = &v1.Revision{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("revisions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Revisions that match those selectors.
func (c *revisions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RevisionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RevisionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("revisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested revisions.
func (c *revisions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("revisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a revision and creates it.  Returns the server's representation of the revision, and an error, if there is any.
func (c *revisions) Create(ctx context.Context, revision *v1.Revision, opts metav1.CreateOptions) (result *v1.Revision, err error) {
	result = &v1.Revision{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("revisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(revision).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a revision and updates it. Returns the server's representation of the revision, and an error, if there is any.
func (c *revisions) Update(ctx context.Context, revision *v1.Revision, opts metav1.UpdateOptions) (result *v1.Revision, err error) {
	result = &v1.Revision{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("revisions").
		Name(revision.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(revision).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *revisions) UpdateStatus(ctx context.Context, revision *v1.Revision, opts metav1.UpdateOptions) (result *v1.Revision, err error) {
	result = &v1.Revision{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("revisions").
		Name(revision.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(revision).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the revision and deletes it. Returns an error if one occurs.
func (c *revisions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("revisions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *revisions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("revisions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched revision.
func (c *revisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Revision, err error) {
	result = &v1.Revision{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("revisions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

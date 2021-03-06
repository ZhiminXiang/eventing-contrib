/*
Copyright 2019 The Knative Authors

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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/serving/pkg/apis/networking/v1alpha1"
	scheme "knative.dev/serving/pkg/client/clientset/versioned/scheme"
)

// ServerlessServicesGetter has a method to return a ServerlessServiceInterface.
// A group's client should implement this interface.
type ServerlessServicesGetter interface {
	ServerlessServices(namespace string) ServerlessServiceInterface
}

// ServerlessServiceInterface has methods to work with ServerlessService resources.
type ServerlessServiceInterface interface {
	Create(*v1alpha1.ServerlessService) (*v1alpha1.ServerlessService, error)
	Update(*v1alpha1.ServerlessService) (*v1alpha1.ServerlessService, error)
	UpdateStatus(*v1alpha1.ServerlessService) (*v1alpha1.ServerlessService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServerlessService, error)
	List(opts v1.ListOptions) (*v1alpha1.ServerlessServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServerlessService, err error)
	ServerlessServiceExpansion
}

// serverlessServices implements ServerlessServiceInterface
type serverlessServices struct {
	client rest.Interface
	ns     string
}

// newServerlessServices returns a ServerlessServices
func newServerlessServices(c *NetworkingV1alpha1Client, namespace string) *serverlessServices {
	return &serverlessServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serverlessService, and returns the corresponding serverlessService object, and an error if there is any.
func (c *serverlessServices) Get(name string, options v1.GetOptions) (result *v1alpha1.ServerlessService, err error) {
	result = &v1alpha1.ServerlessService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serverlessservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServerlessServices that match those selectors.
func (c *serverlessServices) List(opts v1.ListOptions) (result *v1alpha1.ServerlessServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServerlessServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serverlessservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serverlessServices.
func (c *serverlessServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serverlessservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a serverlessService and creates it.  Returns the server's representation of the serverlessService, and an error, if there is any.
func (c *serverlessServices) Create(serverlessService *v1alpha1.ServerlessService) (result *v1alpha1.ServerlessService, err error) {
	result = &v1alpha1.ServerlessService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serverlessservices").
		Body(serverlessService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serverlessService and updates it. Returns the server's representation of the serverlessService, and an error, if there is any.
func (c *serverlessServices) Update(serverlessService *v1alpha1.ServerlessService) (result *v1alpha1.ServerlessService, err error) {
	result = &v1alpha1.ServerlessService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serverlessservices").
		Name(serverlessService.Name).
		Body(serverlessService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *serverlessServices) UpdateStatus(serverlessService *v1alpha1.ServerlessService) (result *v1alpha1.ServerlessService, err error) {
	result = &v1alpha1.ServerlessService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serverlessservices").
		Name(serverlessService.Name).
		SubResource("status").
		Body(serverlessService).
		Do().
		Into(result)
	return
}

// Delete takes name of the serverlessService and deletes it. Returns an error if one occurs.
func (c *serverlessServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serverlessservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serverlessServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serverlessservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serverlessService.
func (c *serverlessServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServerlessService, err error) {
	result = &v1alpha1.ServerlessService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serverlessservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

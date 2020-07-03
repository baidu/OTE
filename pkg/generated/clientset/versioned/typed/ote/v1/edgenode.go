/*
Copyright 2019 Baidu, Inc.

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
	"time"

	v1 "github.com/baidu/ote-stack/pkg/apis/ote/v1"
	scheme "github.com/baidu/ote-stack/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EdgeNodesGetter has a method to return a EdgeNodeInterface.
// A group's client should implement this interface.
type EdgeNodesGetter interface {
	EdgeNodes(namespace string) EdgeNodeInterface
}

// EdgeNodeInterface has methods to work with EdgeNode resources.
type EdgeNodeInterface interface {
	Create(*v1.EdgeNode) (*v1.EdgeNode, error)
	Update(*v1.EdgeNode) (*v1.EdgeNode, error)
	UpdateStatus(*v1.EdgeNode) (*v1.EdgeNode, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.EdgeNode, error)
	List(opts metav1.ListOptions) (*v1.EdgeNodeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EdgeNode, err error)
	EdgeNodeExpansion
}

// edgeNodes implements EdgeNodeInterface
type edgeNodes struct {
	client rest.Interface
	ns     string
}

// newEdgeNodes returns a EdgeNodes
func newEdgeNodes(c *OteV1Client, namespace string) *edgeNodes {
	return &edgeNodes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the edgeNode, and returns the corresponding edgeNode object, and an error if there is any.
func (c *edgeNodes) Get(name string, options metav1.GetOptions) (result *v1.EdgeNode, err error) {
	result = &v1.EdgeNode{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("edgenodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EdgeNodes that match those selectors.
func (c *edgeNodes) List(opts metav1.ListOptions) (result *v1.EdgeNodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EdgeNodeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("edgenodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested edgeNodes.
func (c *edgeNodes) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("edgenodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a edgeNode and creates it.  Returns the server's representation of the edgeNode, and an error, if there is any.
func (c *edgeNodes) Create(edgeNode *v1.EdgeNode) (result *v1.EdgeNode, err error) {
	result = &v1.EdgeNode{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("edgenodes").
		Body(edgeNode).
		Do().
		Into(result)
	return
}

// Update takes the representation of a edgeNode and updates it. Returns the server's representation of the edgeNode, and an error, if there is any.
func (c *edgeNodes) Update(edgeNode *v1.EdgeNode) (result *v1.EdgeNode, err error) {
	result = &v1.EdgeNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("edgenodes").
		Name(edgeNode.Name).
		Body(edgeNode).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *edgeNodes) UpdateStatus(edgeNode *v1.EdgeNode) (result *v1.EdgeNode, err error) {
	result = &v1.EdgeNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("edgenodes").
		Name(edgeNode.Name).
		SubResource("status").
		Body(edgeNode).
		Do().
		Into(result)
	return
}

// Delete takes name of the edgeNode and deletes it. Returns an error if one occurs.
func (c *edgeNodes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("edgenodes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *edgeNodes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("edgenodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched edgeNode.
func (c *edgeNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EdgeNode, err error) {
	result = &v1.EdgeNode{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("edgenodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

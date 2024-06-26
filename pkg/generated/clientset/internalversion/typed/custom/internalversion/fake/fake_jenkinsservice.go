/*
Copyright The Kubernetes Authors.

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

package fake

import (
	"context"
	custom "custom-apiserver/pkg/apis/custom"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJenkinsServices implements JenkinsServiceInterface
type FakeJenkinsServices struct {
	Fake *FakeAutobusi
	ns   string
}

var jenkinsservicesResource = schema.GroupVersionResource{Group: "autobusi.group.custom", Version: "", Resource: "jenkinsservices"}

var jenkinsservicesKind = schema.GroupVersionKind{Group: "autobusi.group.custom", Version: "", Kind: "JenkinsService"}

// Get takes name of the jenkinsService, and returns the corresponding jenkinsService object, and an error if there is any.
func (c *FakeJenkinsServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *custom.JenkinsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jenkinsservicesResource, c.ns, name), &custom.JenkinsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*custom.JenkinsService), err
}

// List takes label and field selectors, and returns the list of JenkinsServices that match those selectors.
func (c *FakeJenkinsServices) List(ctx context.Context, opts v1.ListOptions) (result *custom.JenkinsServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jenkinsservicesResource, jenkinsservicesKind, c.ns, opts), &custom.JenkinsServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &custom.JenkinsServiceList{ListMeta: obj.(*custom.JenkinsServiceList).ListMeta}
	for _, item := range obj.(*custom.JenkinsServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jenkinsServices.
func (c *FakeJenkinsServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jenkinsservicesResource, c.ns, opts))

}

// Create takes the representation of a jenkinsService and creates it.  Returns the server's representation of the jenkinsService, and an error, if there is any.
func (c *FakeJenkinsServices) Create(ctx context.Context, jenkinsService *custom.JenkinsService, opts v1.CreateOptions) (result *custom.JenkinsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jenkinsservicesResource, c.ns, jenkinsService), &custom.JenkinsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*custom.JenkinsService), err
}

// Update takes the representation of a jenkinsService and updates it. Returns the server's representation of the jenkinsService, and an error, if there is any.
func (c *FakeJenkinsServices) Update(ctx context.Context, jenkinsService *custom.JenkinsService, opts v1.UpdateOptions) (result *custom.JenkinsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jenkinsservicesResource, c.ns, jenkinsService), &custom.JenkinsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*custom.JenkinsService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJenkinsServices) UpdateStatus(ctx context.Context, jenkinsService *custom.JenkinsService, opts v1.UpdateOptions) (*custom.JenkinsService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(jenkinsservicesResource, "status", c.ns, jenkinsService), &custom.JenkinsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*custom.JenkinsService), err
}

// Delete takes name of the jenkinsService and deletes it. Returns an error if one occurs.
func (c *FakeJenkinsServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(jenkinsservicesResource, c.ns, name, opts), &custom.JenkinsService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJenkinsServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jenkinsservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &custom.JenkinsServiceList{})
	return err
}

// Patch applies the patch and returns the patched jenkinsService.
func (c *FakeJenkinsServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *custom.JenkinsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jenkinsservicesResource, c.ns, name, pt, data, subresources...), &custom.JenkinsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*custom.JenkinsService), err
}

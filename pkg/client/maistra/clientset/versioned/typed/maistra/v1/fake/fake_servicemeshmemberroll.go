// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	maistrav1 "github.com/maistra/istio-operator/pkg/apis/maistra/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceMeshMemberRolls implements ServiceMeshMemberRollInterface
type FakeServiceMeshMemberRolls struct {
	Fake *FakeMaistraV1
	ns   string
}

var servicemeshmemberrollsResource = schema.GroupVersionResource{Group: "maistra.io", Version: "v1", Resource: "servicemeshmemberrolls"}

var servicemeshmemberrollsKind = schema.GroupVersionKind{Group: "maistra.io", Version: "v1", Kind: "ServiceMeshMemberRoll"}

// Get takes name of the serviceMeshMemberRoll, and returns the corresponding serviceMeshMemberRoll object, and an error if there is any.
func (c *FakeServiceMeshMemberRolls) Get(name string, options v1.GetOptions) (result *maistrav1.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicemeshmemberrollsResource, c.ns, name), &maistrav1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshMemberRoll), err
}

// List takes label and field selectors, and returns the list of ServiceMeshMemberRolls that match those selectors.
func (c *FakeServiceMeshMemberRolls) List(opts v1.ListOptions) (result *maistrav1.ServiceMeshMemberRollList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicemeshmemberrollsResource, servicemeshmemberrollsKind, c.ns, opts), &maistrav1.ServiceMeshMemberRollList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &maistrav1.ServiceMeshMemberRollList{ListMeta: obj.(*maistrav1.ServiceMeshMemberRollList).ListMeta}
	for _, item := range obj.(*maistrav1.ServiceMeshMemberRollList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceMeshMemberRolls.
func (c *FakeServiceMeshMemberRolls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicemeshmemberrollsResource, c.ns, opts))

}

// Create takes the representation of a serviceMeshMemberRoll and creates it.  Returns the server's representation of the serviceMeshMemberRoll, and an error, if there is any.
func (c *FakeServiceMeshMemberRolls) Create(serviceMeshMemberRoll *maistrav1.ServiceMeshMemberRoll) (result *maistrav1.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicemeshmemberrollsResource, c.ns, serviceMeshMemberRoll), &maistrav1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshMemberRoll), err
}

// Update takes the representation of a serviceMeshMemberRoll and updates it. Returns the server's representation of the serviceMeshMemberRoll, and an error, if there is any.
func (c *FakeServiceMeshMemberRolls) Update(serviceMeshMemberRoll *maistrav1.ServiceMeshMemberRoll) (result *maistrav1.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicemeshmemberrollsResource, c.ns, serviceMeshMemberRoll), &maistrav1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshMemberRoll), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceMeshMemberRolls) UpdateStatus(serviceMeshMemberRoll *maistrav1.ServiceMeshMemberRoll) (*maistrav1.ServiceMeshMemberRoll, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicemeshmemberrollsResource, "status", c.ns, serviceMeshMemberRoll), &maistrav1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshMemberRoll), err
}

// Delete takes name of the serviceMeshMemberRoll and deletes it. Returns an error if one occurs.
func (c *FakeServiceMeshMemberRolls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicemeshmemberrollsResource, c.ns, name), &maistrav1.ServiceMeshMemberRoll{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceMeshMemberRolls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicemeshmemberrollsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &maistrav1.ServiceMeshMemberRollList{})
	return err
}

// Patch applies the patch and returns the patched serviceMeshMemberRoll.
func (c *FakeServiceMeshMemberRolls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *maistrav1.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicemeshmemberrollsResource, c.ns, name, pt, data, subresources...), &maistrav1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshMemberRoll), err
}
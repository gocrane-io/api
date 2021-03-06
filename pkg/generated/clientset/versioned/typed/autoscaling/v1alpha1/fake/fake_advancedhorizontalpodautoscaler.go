// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gocrane-io/api/autoscaling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAdvancedHorizontalPodAutoscalers implements AdvancedHorizontalPodAutoscalerInterface
type FakeAdvancedHorizontalPodAutoscalers struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var advancedhorizontalpodautoscalersResource = schema.GroupVersionResource{Group: "autoscaling.crane.io", Version: "v1alpha1", Resource: "advancedhorizontalpodautoscalers"}

var advancedhorizontalpodautoscalersKind = schema.GroupVersionKind{Group: "autoscaling.crane.io", Version: "v1alpha1", Kind: "AdvancedHorizontalPodAutoscaler"}

// Get takes name of the advancedHorizontalPodAutoscaler, and returns the corresponding advancedHorizontalPodAutoscaler object, and an error if there is any.
func (c *FakeAdvancedHorizontalPodAutoscalers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AdvancedHorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(advancedhorizontalpodautoscalersResource, c.ns, name), &v1alpha1.AdvancedHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedHorizontalPodAutoscaler), err
}

// List takes label and field selectors, and returns the list of AdvancedHorizontalPodAutoscalers that match those selectors.
func (c *FakeAdvancedHorizontalPodAutoscalers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AdvancedHorizontalPodAutoscalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(advancedhorizontalpodautoscalersResource, advancedhorizontalpodautoscalersKind, c.ns, opts), &v1alpha1.AdvancedHorizontalPodAutoscalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AdvancedHorizontalPodAutoscalerList{ListMeta: obj.(*v1alpha1.AdvancedHorizontalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AdvancedHorizontalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested advancedHorizontalPodAutoscalers.
func (c *FakeAdvancedHorizontalPodAutoscalers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(advancedhorizontalpodautoscalersResource, c.ns, opts))

}

// Create takes the representation of a advancedHorizontalPodAutoscaler and creates it.  Returns the server's representation of the advancedHorizontalPodAutoscaler, and an error, if there is any.
func (c *FakeAdvancedHorizontalPodAutoscalers) Create(ctx context.Context, advancedHorizontalPodAutoscaler *v1alpha1.AdvancedHorizontalPodAutoscaler, opts v1.CreateOptions) (result *v1alpha1.AdvancedHorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(advancedhorizontalpodautoscalersResource, c.ns, advancedHorizontalPodAutoscaler), &v1alpha1.AdvancedHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedHorizontalPodAutoscaler), err
}

// Update takes the representation of a advancedHorizontalPodAutoscaler and updates it. Returns the server's representation of the advancedHorizontalPodAutoscaler, and an error, if there is any.
func (c *FakeAdvancedHorizontalPodAutoscalers) Update(ctx context.Context, advancedHorizontalPodAutoscaler *v1alpha1.AdvancedHorizontalPodAutoscaler, opts v1.UpdateOptions) (result *v1alpha1.AdvancedHorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(advancedhorizontalpodautoscalersResource, c.ns, advancedHorizontalPodAutoscaler), &v1alpha1.AdvancedHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedHorizontalPodAutoscaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAdvancedHorizontalPodAutoscalers) UpdateStatus(ctx context.Context, advancedHorizontalPodAutoscaler *v1alpha1.AdvancedHorizontalPodAutoscaler, opts v1.UpdateOptions) (*v1alpha1.AdvancedHorizontalPodAutoscaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(advancedhorizontalpodautoscalersResource, "status", c.ns, advancedHorizontalPodAutoscaler), &v1alpha1.AdvancedHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedHorizontalPodAutoscaler), err
}

// Delete takes name of the advancedHorizontalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *FakeAdvancedHorizontalPodAutoscalers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(advancedhorizontalpodautoscalersResource, c.ns, name), &v1alpha1.AdvancedHorizontalPodAutoscaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAdvancedHorizontalPodAutoscalers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(advancedhorizontalpodautoscalersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AdvancedHorizontalPodAutoscalerList{})
	return err
}

// Patch applies the patch and returns the patched advancedHorizontalPodAutoscaler.
func (c *FakeAdvancedHorizontalPodAutoscalers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AdvancedHorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(advancedhorizontalpodautoscalersResource, c.ns, name, pt, data, subresources...), &v1alpha1.AdvancedHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedHorizontalPodAutoscaler), err
}

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/gocrane-io/api/autoscaling/v1alpha1"
	ensurancev1alpha1 "github.com/gocrane-io/api/ensurance/v1alpha1"
	predictionv1alpha1 "github.com/gocrane-io/api/prediction/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=autoscaling.crane.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("advancedhorizontalpodautoscalers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Autoscaling().V1alpha1().AdvancedHorizontalPodAutoscalers().Informer()}, nil

		// Group=ensurance.crane.io, Version=v1alpha1
	case ensurancev1alpha1.SchemeGroupVersion.WithResource("nodeqosensurancepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ensurance().V1alpha1().NodeQOSEnsurancePolicies().Informer()}, nil
	case ensurancev1alpha1.SchemeGroupVersion.WithResource("podqosensurancepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ensurance().V1alpha1().PodQOSEnsurancePolicies().Informer()}, nil

		// Group=prediction.crane.io, Version=v1alpha1
	case predictionv1alpha1.SchemeGroupVersion.WithResource("nodepredictions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Prediction().V1alpha1().NodePredictions().Informer()}, nil
	case predictionv1alpha1.SchemeGroupVersion.WithResource("podgrouppredictions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Prediction().V1alpha1().PodGroupPredictions().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

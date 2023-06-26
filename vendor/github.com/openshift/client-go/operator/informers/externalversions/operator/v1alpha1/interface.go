// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ImageContentSourcePolicies returns a ImageContentSourcePolicyInformer.
	ImageContentSourcePolicies() ImageContentSourcePolicyInformer
	// OLMs returns a OLMInformer.
	OLMs() OLMInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ImageContentSourcePolicies returns a ImageContentSourcePolicyInformer.
func (v *version) ImageContentSourcePolicies() ImageContentSourcePolicyInformer {
	return &imageContentSourcePolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OLMs returns a OLMInformer.
func (v *version) OLMs() OLMInformer {
	return &oLMInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
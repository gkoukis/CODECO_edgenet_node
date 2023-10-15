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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/EdgeNet-project/edgenet/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ManagerCaches returns a ManagerCacheInformer.
	ManagerCaches() ManagerCacheInformer
	// SelectiveDeploymentAnchors returns a SelectiveDeploymentAnchorInformer.
	SelectiveDeploymentAnchors() SelectiveDeploymentAnchorInformer
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

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerCaches returns a ManagerCacheInformer.
func (v *version) ManagerCaches() ManagerCacheInformer {
	return &managerCacheInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// SelectiveDeploymentAnchors returns a SelectiveDeploymentAnchorInformer.
func (v *version) SelectiveDeploymentAnchors() SelectiveDeploymentAnchorInformer {
	return &selectiveDeploymentAnchorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

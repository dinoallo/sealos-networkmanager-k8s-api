/*
Copyright 2024.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Backend struct {
	Name    string                    `json:"name,omitempty"`
	PodRefs []*corev1.ObjectReference `json:"podRefs,omitempty"`
	Port    int32                     `json:"port,omitempty"`
}

// IngressPodMappingSpec defines the desired state of IngressPodMapping
type IngressPodMappingSpec struct {
	Ingress  string    `json:"ingress,omitempty"`
	Backends []Backend `json:"backends,omitempty"`
}

// IngressPodMappingStatus defines the observed state of IngressPodMapping
type IngressPodMappingStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName="ipm"
//+kubebuilder:printcolumn:name="ingress",type=string,JSONPath=`.spec.ingress`

// IngressPodMapping is the Schema for the ingresspodmappings API
type IngressPodMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IngressPodMappingSpec   `json:"spec,omitempty"`
	Status IngressPodMappingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IngressPodMappingList contains a list of IngressPodMapping
type IngressPodMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IngressPodMapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IngressPodMapping{}, &IngressPodMappingList{})
}

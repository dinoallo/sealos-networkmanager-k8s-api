/*
Copyright 2023.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type POD_TYPE int64

const (
	POD_TYPE_UNKNOWN POD_TYPE = iota
	POD_TYPE_DB
	POD_TYPE_APP
	POD_TYPE_TERMINAL
	POD_TYPE_JOB
	POD_TYPE_OTHER
	POD_TYPE_OBJECTSTORAGE

	CHECK_DB_LABEL_KEY       = "apps.kubeblocks.io/component-name"
	CHECK_TERMINAL_LABEL_KEY = "TerminalID"
	CHECK_APP_LABEL_KEY      = "app"
	CHECK_JOB_LABEL_KEY      = "job-name"
	DB_TYPE_LABEL_KEY        = "apps.kubernetes.io/instance"
	APP_TYPE_LABEL_KEY       = "app"
	JOB_TYPE_LABEL_KEY       = "job-name"
)

// TrafficSyncRequestSpec defines the desired state of TrafficSyncRequest
type TrafficSyncRequestSpec struct {
	AssociatedNamespace string          `json:"associatedNamespace,omitempty"`
	AssociatedPod       string          `json:"associatedPod,omitempty"`
	PodType             POD_TYPE        `json:"podType,omitempty"`
	PodTypeName         string          `json:"podTypeName,omitempty"`
	NodeIP              string          `json:"nodeIP,omitempty"`
	Address             string          `json:"address,omitempty"`
	Tag                 string          `json:"tag,omitempty"`
	SyncPeriod          metav1.Duration `json:"syncPeriod,omitempty"`
}

// TrafficSyncRequestStatus defines the observed state of TrafficSyncRequest
type TrafficSyncRequestStatus struct {
	LastSyncTime metav1.Time `json:"lastSyncTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Pod",type=string,JSONPath=`.spec.associatedPod`
//+kubebuilder:printcolumn:name="AddressToSync",type=string,JSONPath=`.spec.address`
//+kubebuilder:printcolumn:name="TagToSync",type=string,JSONPath=`.spec.tag`
//+kubebuilder:resource:shortName="tsr"

// TrafficSyncRequest is the Schema for the trafficsyncrequests API
type TrafficSyncRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficSyncRequestSpec   `json:"spec,omitempty"`
	Status TrafficSyncRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TrafficSyncRequestList contains a list of TrafficSyncRequest
type TrafficSyncRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficSyncRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficSyncRequest{}, &TrafficSyncRequestList{})
}

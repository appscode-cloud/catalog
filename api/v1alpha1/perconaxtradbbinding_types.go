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
	kmapi "kmodules.xyz/client-go/api/v1"
)

// PerconaXtraDBBindingSpec defines the desired state of PerconaXtraDBBinding
type PerconaXtraDBBindingSpec struct {
	// SourceRef refers to the source app instance.
	SourceRef kmapi.ObjectReference `json:"sourceRef"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PerconaXtraDBBinding is the Schema for the perconaxtradbbindings API
type PerconaXtraDBBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerconaXtraDBBindingSpec `json:"spec,omitempty"`
	Status AppStatus                `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PerconaXtraDBBindingList contains a list of PerconaXtraDBBinding
type PerconaXtraDBBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerconaXtraDBBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerconaXtraDBBinding{}, &PerconaXtraDBBindingList{})
}

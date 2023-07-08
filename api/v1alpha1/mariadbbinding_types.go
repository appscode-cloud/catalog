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

// MariaDBBindingSpec defines the desired state of MariaDBBinding
type MariaDBBindingSpec struct {
	// SourceRef refers to the source app instance.
	SourceRef kmapi.ObjectReference `json:"sourceRef"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MariaDBBinding is the Schema for the mariadbbindings API
type MariaDBBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MariaDBBindingSpec `json:"spec,omitempty"`
	Status AppStatus          `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MariaDBBindingList contains a list of MariaDBBinding
type MariaDBBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MariaDBBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MariaDBBinding{}, &MariaDBBindingList{})
}

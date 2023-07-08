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
	"k8s.io/apimachinery/pkg/runtime"
	kmapi "kmodules.xyz/client-go/api/v1"
)

// AppConditionType is a type of condition for a route.
type AppConditionType string

// AppConditionReason is a reason for a route condition.
type AppConditionReason string

const (
	// This condition indicates whether the route has been accepted or rejected
	// by a Gateway, and why.
	//
	// Possible reasons for this condition to be true are:
	//
	// * "Accepted"
	//
	// Possible reasons for this condition to be False are:
	//
	// * "NoMatchingSource"
	// * "UnsupportedValue"
	//
	// Possible reasons for this condition to be Unknown are:
	//
	// * "Pending"
	//
	// Controllers may raise this condition with other reasons,
	// but should prefer to use the reasons listed above to improve
	// interoperability.
	AppConditionTypeSourceAccepted AppConditionType = "Accepted"

	// This reason is used with the "Accepted" condition when the App has been
	// accepted by the Gateway.
	AppConditionReasonSourceAccepted AppConditionReason = "Accepted"

	// This reason is used with the "Accepted" condition when there are
	// no matching Sources. In the case of Gateways, this can occur when
	// an App SourceRef specifies a Port and/or SectionName that does not
	// match any Listeners in the Gateway.

	// This reason is used with the "Accepted" condition when there are
	// no matching Vault.

	AppConditionTypeFinalizerAdded          AppConditionType   = "FinalizerAdded"
	AppConditionReasonPatchFinalizerFailed  AppConditionReason = "PatchFinalizerFailed"
	AppConditionReasonPatchFinalizerSucceed AppConditionReason = "PatchFinalizerSucceed"

	AppConditionTypeVaultReady          AppConditionType   = "VaultReady"
	AppConditionReasonVaultReady        AppConditionReason = "VaultReady"
	AppConditionReasonVaultNotCreated   AppConditionReason = "VaultNotCreated"
	AppConditionReasonVaultProvisioning AppConditionReason = "VaultProvisioning"

	AppConditionTypeServiceAccountReady        AppConditionType   = "ServiceAccountReady"
	AppConditionReasonServiceAccountNotCreated AppConditionReason = "ServiceAccountNotCreated"
	AppConditionReasonServiceAccountReady      AppConditionReason = "ServiceAccountReady"

	AppConditionTypeSecretEngineReady        AppConditionType   = "SecretEngineReady"
	AppConditionReasonSecretEngineReady      AppConditionReason = "SecretEngineReady"
	AppConditionReasonSecretEngineNotReady   AppConditionReason = "SecretEngineNotReady"
	AppConditionReasonSecretEngineNotCreated AppConditionReason = "SecretEngineNotCreated"

	AppConditionTypeRoleReady      AppConditionReason = "RoleReady"
	AppConditionReasonRoleNotReady AppConditionReason = "RoleNotReady"
	AppConditionReasonRoleReady    AppConditionReason = "RoleReady"

	AppConditionTypeSecretAccessRequestReady      AppConditionType   = "SecretAccessRequestReady"
	AppConditionReasonSecretAccessRequestNotReady AppConditionReason = "SecretAccessRequestNotReady"
	AppConditionReasonSecretAccessRequestExpired  AppConditionReason = "SecretRequestExpired"
	AppConditionReasonSecretAccessRequestApproved AppConditionReason = "SecretRequestApproved"

	// This reason is used with the "Accepted" condition when a value for an Enum
	// is not recognized.
	AppReasonUnsupportedValue AppConditionReason = "UnsupportedValue"

	// This reason is used with the "Accepted" when a controller has not yet
	// reconciled the route.
	AppReasonPending AppConditionReason = "Pending"

	// This condition indicates whether the controller was able to resolve all
	// the object references for the App.
	//
	// Possible reasons for this condition to be true are:
	//
	// * "ResolvedRefs"
	//
	// Possible reasons for this condition to be false are:
	//
	// * "RefNotPermitted"
	// * "InvalidKind"
	//
	// Controllers may raise this condition with other reasons,
	// but should prefer to use the reasons listed above to improve
	// interoperability.
	AppConditionResolvedRefs AppConditionType = "ResolvedRefs"

	// This reason is used with the "ResolvedRefs" condition when the condition
	// is true.
	AppReasonResolvedRefs AppConditionReason = "ResolvedRefs"

	// This reason is used with the "ResolvedRefs" condition when
	// one of the Listener's Apps has a BackendRef to an object in
	// another namespace, where the object in the other namespace does
	// not have a ReferenceGrant explicitly allowing the reference.
	AppReasonRefNotPermitted AppConditionReason = "RefNotPermitted"

	// This reason is used with the "ResolvedRefs" condition when
	// one of the App's rules has a reference to an unknown or unsupported
	// Group and/or Kind.
	AppReasonInvalidKind AppConditionReason = "InvalidKind"
)

// +kubebuilder:validation:Enum=Pending;InProgress;Terminating;Current;Failed;Expired
type AppPhase string

const (
	AppPhasePending     AppPhase = "Pending"
	AppPhaseInProgress  AppPhase = "InProgress"
	AppPhaseTerminating AppPhase = "Terminating"
	AppPhaseCurrent     AppPhase = "Current"
	AppPhaseFailed      AppPhase = "Failed"
	AppPhaseExpired     AppPhase = "Expired"
)

/*Pending : If MongoDB / VaultServer not found
InProgress: If role or accessReq is not ensured yet. Or their phase is not determined yet
Current: all ok, secret is valid
Expired: all ok, secret is expired
Failed: role or accessReq failed for some reason*/

// AppStatus defines the observed state of App
type AppStatus struct {
	// Conditions describe the current conditions of the Gateway.
	//
	// Implementations should prefer to express Gateway conditions
	// using the `GatewayConditionType` and `GatewayConditionReason`
	// constants so that operators and tools can converge on a common
	// vocabulary to describe Gateway state.
	//
	// Known condition types are:
	//
	// * "Accepted"
	// * "Ready"
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Accepted", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"},{type: "Programmed", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []kmapi.Condition `json:"conditions,omitempty"`

	// Specifies the current phase of the App
	// +optional
	Phase AppPhase `json:"phase,omitempty"`

	// +optional
	Source *runtime.RawExtension `json:"source,omitempty"`
}

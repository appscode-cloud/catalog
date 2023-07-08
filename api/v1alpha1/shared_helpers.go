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
	kmapi "kmodules.xyz/client-go/api/v1"
)

func IsAccessRequestExpired(conditions []kmapi.Condition) bool {
	idx, cond := kmapi.GetCondition(conditions, string(AppConditionTypeSecretAccessRequestReady))
	if idx == -1 {
		return false
	}
	return cond.Reason == string(AppConditionReasonSecretAccessRequestExpired)
}

func IsEngineAPIResourcesReady(conditions []kmapi.Condition) bool {
	se := kmapi.IsConditionTrue(conditions, string(AppConditionTypeSecretEngineReady))
	sar := kmapi.IsConditionTrue(conditions, string(AppConditionTypeSecretAccessRequestReady))
	role := kmapi.IsConditionTrue(conditions, string(AppConditionTypeRoleReady))

	return sar && role && se
}

func IsEngineAPIResourcesConditionDetermined(conditions []kmapi.Condition) bool {
	roleIndex, _ := kmapi.GetCondition(conditions, string(AppConditionTypeRoleReady))
	sarIndex, _ := kmapi.GetCondition(conditions, string(AppConditionTypeSecretAccessRequestReady))

	return (roleIndex != -1) && (sarIndex != -1)
}

func GetFinalizer() string {
	return GroupVersion.Group
}

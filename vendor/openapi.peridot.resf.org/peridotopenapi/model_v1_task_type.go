/*
 * peridot/proto/v1/batch.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peridotopenapi

import (
	"encoding/json"
	"fmt"
)

// V1TaskType the model 'V1TaskType'
type V1TaskType string

// List of v1TaskType
const (
	UNKNOWN V1TaskType = "TASK_TYPE_UNKNOWN"
	IMPORT V1TaskType = "TASK_TYPE_IMPORT"
	IMPORT_SRC_GIT V1TaskType = "TASK_TYPE_IMPORT_SRC_GIT"
	IMPORT_SRC_GIT_TO_DIST_GIT V1TaskType = "TASK_TYPE_IMPORT_SRC_GIT_TO_DIST_GIT"
	IMPORT_DOWNSTREAM V1TaskType = "TASK_TYPE_IMPORT_DOWNSTREAM"
	IMPORT_UPSTREAM V1TaskType = "TASK_TYPE_IMPORT_UPSTREAM"
	BUILD V1TaskType = "TASK_TYPE_BUILD"
	BUILD_SRPM V1TaskType = "TASK_TYPE_BUILD_SRPM"
	BUILD_ARCH V1TaskType = "TASK_TYPE_BUILD_ARCH"
	BUILD_SRPM_UPLOAD V1TaskType = "TASK_TYPE_BUILD_SRPM_UPLOAD"
	BUILD_ARCH_UPLOAD V1TaskType = "TASK_TYPE_BUILD_ARCH_UPLOAD"
	WORKER_PROVISION V1TaskType = "TASK_TYPE_WORKER_PROVISION"
	WORKER_DESTROY V1TaskType = "TASK_TYPE_WORKER_DESTROY"
	YUMREPOFS_UPDATE V1TaskType = "TASK_TYPE_YUMREPOFS_UPDATE"
	KEYKEEPER_SIGN_ARTIFACT V1TaskType = "TASK_TYPE_KEYKEEPER_SIGN_ARTIFACT"
	SYNC_CATALOG V1TaskType = "TASK_TYPE_SYNC_CATALOG"
	RPM_IMPORT V1TaskType = "TASK_TYPE_RPM_IMPORT"
	CREATE_HASHED_REPOSITORIES V1TaskType = "TASK_TYPE_CREATE_HASHED_REPOSITORIES"
	LOOKASIDE_FILE_UPLOAD V1TaskType = "TASK_TYPE_LOOKASIDE_FILE_UPLOAD"
	RPM_LOOKASIDE_BATCH_IMPORT V1TaskType = "TASK_TYPE_RPM_LOOKASIDE_BATCH_IMPORT"
)

func (v *V1TaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1TaskType(value)
	for _, existing := range []V1TaskType{ "TASK_TYPE_UNKNOWN", "TASK_TYPE_IMPORT", "TASK_TYPE_IMPORT_SRC_GIT", "TASK_TYPE_IMPORT_SRC_GIT_TO_DIST_GIT", "TASK_TYPE_IMPORT_DOWNSTREAM", "TASK_TYPE_IMPORT_UPSTREAM", "TASK_TYPE_BUILD", "TASK_TYPE_BUILD_SRPM", "TASK_TYPE_BUILD_ARCH", "TASK_TYPE_BUILD_SRPM_UPLOAD", "TASK_TYPE_BUILD_ARCH_UPLOAD", "TASK_TYPE_WORKER_PROVISION", "TASK_TYPE_WORKER_DESTROY", "TASK_TYPE_YUMREPOFS_UPDATE", "TASK_TYPE_KEYKEEPER_SIGN_ARTIFACT", "TASK_TYPE_SYNC_CATALOG", "TASK_TYPE_RPM_IMPORT", "TASK_TYPE_CREATE_HASHED_REPOSITORIES", "TASK_TYPE_LOOKASIDE_FILE_UPLOAD", "TASK_TYPE_RPM_LOOKASIDE_BATCH_IMPORT",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1TaskType", value)
}

// Ptr returns reference to v1TaskType value
func (v V1TaskType) Ptr() *V1TaskType {
	return &v
}

type NullableV1TaskType struct {
	value *V1TaskType
	isSet bool
}

func (v NullableV1TaskType) Get() *V1TaskType {
	return v.value
}

func (v *NullableV1TaskType) Set(val *V1TaskType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TaskType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TaskType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TaskType(val *V1TaskType) *NullableV1TaskType {
	return &NullableV1TaskType{value: val, isSet: true}
}

func (v NullableV1TaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TaskType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


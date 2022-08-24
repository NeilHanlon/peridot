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
)

// V1GetBuildResponse struct for V1GetBuildResponse
type V1GetBuildResponse struct {
	Build *V1Build `json:"build,omitempty"`
}

// NewV1GetBuildResponse instantiates a new V1GetBuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GetBuildResponse() *V1GetBuildResponse {
	this := V1GetBuildResponse{}
	return &this
}

// NewV1GetBuildResponseWithDefaults instantiates a new V1GetBuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GetBuildResponseWithDefaults() *V1GetBuildResponse {
	this := V1GetBuildResponse{}
	return &this
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *V1GetBuildResponse) GetBuild() V1Build {
	if o == nil || o.Build == nil {
		var ret V1Build
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GetBuildResponse) GetBuildOk() (*V1Build, bool) {
	if o == nil || o.Build == nil {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *V1GetBuildResponse) HasBuild() bool {
	if o != nil && o.Build != nil {
		return true
	}

	return false
}

// SetBuild gets a reference to the given V1Build and assigns it to the Build field.
func (o *V1GetBuildResponse) SetBuild(v V1Build) {
	o.Build = &v
}

func (o V1GetBuildResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Build != nil {
		toSerialize["build"] = o.Build
	}
	return json.Marshal(toSerialize)
}

type NullableV1GetBuildResponse struct {
	value *V1GetBuildResponse
	isSet bool
}

func (v NullableV1GetBuildResponse) Get() *V1GetBuildResponse {
	return v.value
}

func (v *NullableV1GetBuildResponse) Set(val *V1GetBuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GetBuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GetBuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GetBuildResponse(val *V1GetBuildResponse) *NullableV1GetBuildResponse {
	return &NullableV1GetBuildResponse{value: val, isSet: true}
}

func (v NullableV1GetBuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GetBuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



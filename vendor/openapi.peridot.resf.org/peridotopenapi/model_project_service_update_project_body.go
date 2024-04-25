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

// ProjectServiceUpdateProjectBody struct for ProjectServiceUpdateProjectBody
type ProjectServiceUpdateProjectBody struct {
	Project *V1Project `json:"project,omitempty"`
}

// NewProjectServiceUpdateProjectBody instantiates a new ProjectServiceUpdateProjectBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectServiceUpdateProjectBody() *ProjectServiceUpdateProjectBody {
	this := ProjectServiceUpdateProjectBody{}
	return &this
}

// NewProjectServiceUpdateProjectBodyWithDefaults instantiates a new ProjectServiceUpdateProjectBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectServiceUpdateProjectBodyWithDefaults() *ProjectServiceUpdateProjectBody {
	this := ProjectServiceUpdateProjectBody{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectServiceUpdateProjectBody) GetProject() V1Project {
	if o == nil || o.Project == nil {
		var ret V1Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectServiceUpdateProjectBody) GetProjectOk() (*V1Project, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectServiceUpdateProjectBody) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given V1Project and assigns it to the Project field.
func (o *ProjectServiceUpdateProjectBody) SetProject(v V1Project) {
	o.Project = &v
}

func (o ProjectServiceUpdateProjectBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableProjectServiceUpdateProjectBody struct {
	value *ProjectServiceUpdateProjectBody
	isSet bool
}

func (v NullableProjectServiceUpdateProjectBody) Get() *ProjectServiceUpdateProjectBody {
	return v.value
}

func (v *NullableProjectServiceUpdateProjectBody) Set(val *ProjectServiceUpdateProjectBody) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectServiceUpdateProjectBody) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectServiceUpdateProjectBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectServiceUpdateProjectBody(val *ProjectServiceUpdateProjectBody) *NullableProjectServiceUpdateProjectBody {
	return &NullableProjectServiceUpdateProjectBody{value: val, isSet: true}
}

func (v NullableProjectServiceUpdateProjectBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectServiceUpdateProjectBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


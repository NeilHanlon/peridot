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

// ImportServiceImportPackageBatchBody struct for ImportServiceImportPackageBatchBody
type ImportServiceImportPackageBatchBody struct {
	Imports *[]V1ImportPackageRequest `json:"imports,omitempty"`
}

// NewImportServiceImportPackageBatchBody instantiates a new ImportServiceImportPackageBatchBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportServiceImportPackageBatchBody() *ImportServiceImportPackageBatchBody {
	this := ImportServiceImportPackageBatchBody{}
	return &this
}

// NewImportServiceImportPackageBatchBodyWithDefaults instantiates a new ImportServiceImportPackageBatchBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportServiceImportPackageBatchBodyWithDefaults() *ImportServiceImportPackageBatchBody {
	this := ImportServiceImportPackageBatchBody{}
	return &this
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *ImportServiceImportPackageBatchBody) GetImports() []V1ImportPackageRequest {
	if o == nil || o.Imports == nil {
		var ret []V1ImportPackageRequest
		return ret
	}
	return *o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportServiceImportPackageBatchBody) GetImportsOk() (*[]V1ImportPackageRequest, bool) {
	if o == nil || o.Imports == nil {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *ImportServiceImportPackageBatchBody) HasImports() bool {
	if o != nil && o.Imports != nil {
		return true
	}

	return false
}

// SetImports gets a reference to the given []V1ImportPackageRequest and assigns it to the Imports field.
func (o *ImportServiceImportPackageBatchBody) SetImports(v []V1ImportPackageRequest) {
	o.Imports = &v
}

func (o ImportServiceImportPackageBatchBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	return json.Marshal(toSerialize)
}

type NullableImportServiceImportPackageBatchBody struct {
	value *ImportServiceImportPackageBatchBody
	isSet bool
}

func (v NullableImportServiceImportPackageBatchBody) Get() *ImportServiceImportPackageBatchBody {
	return v.value
}

func (v *NullableImportServiceImportPackageBatchBody) Set(val *ImportServiceImportPackageBatchBody) {
	v.value = val
	v.isSet = true
}

func (v NullableImportServiceImportPackageBatchBody) IsSet() bool {
	return v.isSet
}

func (v *NullableImportServiceImportPackageBatchBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportServiceImportPackageBatchBody(val *ImportServiceImportPackageBatchBody) *NullableImportServiceImportPackageBatchBody {
	return &NullableImportServiceImportPackageBatchBody{value: val, isSet: true}
}

func (v NullableImportServiceImportPackageBatchBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportServiceImportPackageBatchBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


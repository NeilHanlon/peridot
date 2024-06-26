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

// V1ImportBatchRetryFailedResponse struct for V1ImportBatchRetryFailedResponse
type V1ImportBatchRetryFailedResponse struct {
	ImportBatchId *string `json:"importBatchId,omitempty"`
}

// NewV1ImportBatchRetryFailedResponse instantiates a new V1ImportBatchRetryFailedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImportBatchRetryFailedResponse() *V1ImportBatchRetryFailedResponse {
	this := V1ImportBatchRetryFailedResponse{}
	return &this
}

// NewV1ImportBatchRetryFailedResponseWithDefaults instantiates a new V1ImportBatchRetryFailedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImportBatchRetryFailedResponseWithDefaults() *V1ImportBatchRetryFailedResponse {
	this := V1ImportBatchRetryFailedResponse{}
	return &this
}

// GetImportBatchId returns the ImportBatchId field value if set, zero value otherwise.
func (o *V1ImportBatchRetryFailedResponse) GetImportBatchId() string {
	if o == nil || o.ImportBatchId == nil {
		var ret string
		return ret
	}
	return *o.ImportBatchId
}

// GetImportBatchIdOk returns a tuple with the ImportBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImportBatchRetryFailedResponse) GetImportBatchIdOk() (*string, bool) {
	if o == nil || o.ImportBatchId == nil {
		return nil, false
	}
	return o.ImportBatchId, true
}

// HasImportBatchId returns a boolean if a field has been set.
func (o *V1ImportBatchRetryFailedResponse) HasImportBatchId() bool {
	if o != nil && o.ImportBatchId != nil {
		return true
	}

	return false
}

// SetImportBatchId gets a reference to the given string and assigns it to the ImportBatchId field.
func (o *V1ImportBatchRetryFailedResponse) SetImportBatchId(v string) {
	o.ImportBatchId = &v
}

func (o V1ImportBatchRetryFailedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportBatchId != nil {
		toSerialize["importBatchId"] = o.ImportBatchId
	}
	return json.Marshal(toSerialize)
}

type NullableV1ImportBatchRetryFailedResponse struct {
	value *V1ImportBatchRetryFailedResponse
	isSet bool
}

func (v NullableV1ImportBatchRetryFailedResponse) Get() *V1ImportBatchRetryFailedResponse {
	return v.value
}

func (v *NullableV1ImportBatchRetryFailedResponse) Set(val *V1ImportBatchRetryFailedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImportBatchRetryFailedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImportBatchRetryFailedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImportBatchRetryFailedResponse(val *V1ImportBatchRetryFailedResponse) *NullableV1ImportBatchRetryFailedResponse {
	return &NullableV1ImportBatchRetryFailedResponse{value: val, isSet: true}
}

func (v NullableV1ImportBatchRetryFailedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImportBatchRetryFailedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



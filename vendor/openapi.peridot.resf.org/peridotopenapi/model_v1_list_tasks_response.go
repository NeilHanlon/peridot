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

// V1ListTasksResponse struct for V1ListTasksResponse
type V1ListTasksResponse struct {
	Tasks *[]V1AsyncTask `json:"tasks,omitempty"`
	Total *string `json:"total,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Page *int32 `json:"page,omitempty"`
}

// NewV1ListTasksResponse instantiates a new V1ListTasksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListTasksResponse() *V1ListTasksResponse {
	this := V1ListTasksResponse{}
	return &this
}

// NewV1ListTasksResponseWithDefaults instantiates a new V1ListTasksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListTasksResponseWithDefaults() *V1ListTasksResponse {
	this := V1ListTasksResponse{}
	return &this
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *V1ListTasksResponse) GetTasks() []V1AsyncTask {
	if o == nil || o.Tasks == nil {
		var ret []V1AsyncTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListTasksResponse) GetTasksOk() (*[]V1AsyncTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *V1ListTasksResponse) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []V1AsyncTask and assigns it to the Tasks field.
func (o *V1ListTasksResponse) SetTasks(v []V1AsyncTask) {
	o.Tasks = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V1ListTasksResponse) GetTotal() string {
	if o == nil || o.Total == nil {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListTasksResponse) GetTotalOk() (*string, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V1ListTasksResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *V1ListTasksResponse) SetTotal(v string) {
	o.Total = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *V1ListTasksResponse) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListTasksResponse) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *V1ListTasksResponse) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *V1ListTasksResponse) SetSize(v int32) {
	o.Size = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *V1ListTasksResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListTasksResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *V1ListTasksResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *V1ListTasksResponse) SetPage(v int32) {
	o.Page = &v
}

func (o V1ListTasksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableV1ListTasksResponse struct {
	value *V1ListTasksResponse
	isSet bool
}

func (v NullableV1ListTasksResponse) Get() *V1ListTasksResponse {
	return v.value
}

func (v *NullableV1ListTasksResponse) Set(val *V1ListTasksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListTasksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListTasksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListTasksResponse(val *V1ListTasksResponse) *NullableV1ListTasksResponse {
	return &NullableV1ListTasksResponse{value: val, isSet: true}
}

func (v NullableV1ListTasksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListTasksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



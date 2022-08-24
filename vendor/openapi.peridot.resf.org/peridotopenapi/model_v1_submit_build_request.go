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

// V1SubmitBuildRequest struct for V1SubmitBuildRequest
type V1SubmitBuildRequest struct {
	// Project ID that we want this build to be assigned to All build requests need a project id, however after the initial import, sharing the VRE in an inter-project way is possible.
	ProjectId *string `json:"projectId,omitempty"`
	PackageName *string `json:"packageName,omitempty"`
	PackageId *string `json:"packageId,omitempty"`
	ScmHash *string `json:"scmHash,omitempty"`
	DisableChecks *bool `json:"disableChecks,omitempty"`
	Branches *[]string `json:"branches,omitempty"`
	ModuleVariant *bool `json:"moduleVariant,omitempty"`
	SideNvrs *[]string `json:"sideNvrs,omitempty"`
	SetInactive *bool `json:"setInactive,omitempty"`
}

// NewV1SubmitBuildRequest instantiates a new V1SubmitBuildRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SubmitBuildRequest() *V1SubmitBuildRequest {
	this := V1SubmitBuildRequest{}
	return &this
}

// NewV1SubmitBuildRequestWithDefaults instantiates a new V1SubmitBuildRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SubmitBuildRequestWithDefaults() *V1SubmitBuildRequest {
	this := V1SubmitBuildRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *V1SubmitBuildRequest) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *V1SubmitBuildRequest) SetPackageName(v string) {
	o.PackageName = &v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetPackageId() string {
	if o == nil || o.PackageId == nil {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetPackageIdOk() (*string, bool) {
	if o == nil || o.PackageId == nil {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasPackageId() bool {
	if o != nil && o.PackageId != nil {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *V1SubmitBuildRequest) SetPackageId(v string) {
	o.PackageId = &v
}

// GetScmHash returns the ScmHash field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetScmHash() string {
	if o == nil || o.ScmHash == nil {
		var ret string
		return ret
	}
	return *o.ScmHash
}

// GetScmHashOk returns a tuple with the ScmHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetScmHashOk() (*string, bool) {
	if o == nil || o.ScmHash == nil {
		return nil, false
	}
	return o.ScmHash, true
}

// HasScmHash returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasScmHash() bool {
	if o != nil && o.ScmHash != nil {
		return true
	}

	return false
}

// SetScmHash gets a reference to the given string and assigns it to the ScmHash field.
func (o *V1SubmitBuildRequest) SetScmHash(v string) {
	o.ScmHash = &v
}

// GetDisableChecks returns the DisableChecks field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetDisableChecks() bool {
	if o == nil || o.DisableChecks == nil {
		var ret bool
		return ret
	}
	return *o.DisableChecks
}

// GetDisableChecksOk returns a tuple with the DisableChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetDisableChecksOk() (*bool, bool) {
	if o == nil || o.DisableChecks == nil {
		return nil, false
	}
	return o.DisableChecks, true
}

// HasDisableChecks returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasDisableChecks() bool {
	if o != nil && o.DisableChecks != nil {
		return true
	}

	return false
}

// SetDisableChecks gets a reference to the given bool and assigns it to the DisableChecks field.
func (o *V1SubmitBuildRequest) SetDisableChecks(v bool) {
	o.DisableChecks = &v
}

// GetBranches returns the Branches field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetBranches() []string {
	if o == nil || o.Branches == nil {
		var ret []string
		return ret
	}
	return *o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetBranchesOk() (*[]string, bool) {
	if o == nil || o.Branches == nil {
		return nil, false
	}
	return o.Branches, true
}

// HasBranches returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasBranches() bool {
	if o != nil && o.Branches != nil {
		return true
	}

	return false
}

// SetBranches gets a reference to the given []string and assigns it to the Branches field.
func (o *V1SubmitBuildRequest) SetBranches(v []string) {
	o.Branches = &v
}

// GetModuleVariant returns the ModuleVariant field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetModuleVariant() bool {
	if o == nil || o.ModuleVariant == nil {
		var ret bool
		return ret
	}
	return *o.ModuleVariant
}

// GetModuleVariantOk returns a tuple with the ModuleVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetModuleVariantOk() (*bool, bool) {
	if o == nil || o.ModuleVariant == nil {
		return nil, false
	}
	return o.ModuleVariant, true
}

// HasModuleVariant returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasModuleVariant() bool {
	if o != nil && o.ModuleVariant != nil {
		return true
	}

	return false
}

// SetModuleVariant gets a reference to the given bool and assigns it to the ModuleVariant field.
func (o *V1SubmitBuildRequest) SetModuleVariant(v bool) {
	o.ModuleVariant = &v
}

// GetSideNvrs returns the SideNvrs field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetSideNvrs() []string {
	if o == nil || o.SideNvrs == nil {
		var ret []string
		return ret
	}
	return *o.SideNvrs
}

// GetSideNvrsOk returns a tuple with the SideNvrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetSideNvrsOk() (*[]string, bool) {
	if o == nil || o.SideNvrs == nil {
		return nil, false
	}
	return o.SideNvrs, true
}

// HasSideNvrs returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasSideNvrs() bool {
	if o != nil && o.SideNvrs != nil {
		return true
	}

	return false
}

// SetSideNvrs gets a reference to the given []string and assigns it to the SideNvrs field.
func (o *V1SubmitBuildRequest) SetSideNvrs(v []string) {
	o.SideNvrs = &v
}

// GetSetInactive returns the SetInactive field value if set, zero value otherwise.
func (o *V1SubmitBuildRequest) GetSetInactive() bool {
	if o == nil || o.SetInactive == nil {
		var ret bool
		return ret
	}
	return *o.SetInactive
}

// GetSetInactiveOk returns a tuple with the SetInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SubmitBuildRequest) GetSetInactiveOk() (*bool, bool) {
	if o == nil || o.SetInactive == nil {
		return nil, false
	}
	return o.SetInactive, true
}

// HasSetInactive returns a boolean if a field has been set.
func (o *V1SubmitBuildRequest) HasSetInactive() bool {
	if o != nil && o.SetInactive != nil {
		return true
	}

	return false
}

// SetSetInactive gets a reference to the given bool and assigns it to the SetInactive field.
func (o *V1SubmitBuildRequest) SetSetInactive(v bool) {
	o.SetInactive = &v
}

func (o V1SubmitBuildRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.PackageName != nil {
		toSerialize["packageName"] = o.PackageName
	}
	if o.PackageId != nil {
		toSerialize["packageId"] = o.PackageId
	}
	if o.ScmHash != nil {
		toSerialize["scmHash"] = o.ScmHash
	}
	if o.DisableChecks != nil {
		toSerialize["disableChecks"] = o.DisableChecks
	}
	if o.Branches != nil {
		toSerialize["branches"] = o.Branches
	}
	if o.ModuleVariant != nil {
		toSerialize["moduleVariant"] = o.ModuleVariant
	}
	if o.SideNvrs != nil {
		toSerialize["sideNvrs"] = o.SideNvrs
	}
	if o.SetInactive != nil {
		toSerialize["setInactive"] = o.SetInactive
	}
	return json.Marshal(toSerialize)
}

type NullableV1SubmitBuildRequest struct {
	value *V1SubmitBuildRequest
	isSet bool
}

func (v NullableV1SubmitBuildRequest) Get() *V1SubmitBuildRequest {
	return v.value
}

func (v *NullableV1SubmitBuildRequest) Set(val *V1SubmitBuildRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SubmitBuildRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SubmitBuildRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SubmitBuildRequest(val *V1SubmitBuildRequest) *NullableV1SubmitBuildRequest {
	return &NullableV1SubmitBuildRequest{value: val, isSet: true}
}

func (v NullableV1SubmitBuildRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SubmitBuildRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



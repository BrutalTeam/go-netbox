/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Circuit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Circuit{}

// Circuit Adds support for custom fields and tags.
type Circuit struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// Unique circuit ID
	Cid             string                       `json:"cid"`
	Provider        BriefProvider                `json:"provider"`
	ProviderAccount NullableBriefProviderAccount `json:"provider_account,omitempty"`
	Type            BriefCircuitType             `json:"type"`
	Status          *CircuitStatus               `json:"status,omitempty"`
	Tenant          NullableBriefTenant          `json:"tenant,omitempty"`
	InstallDate     NullableString               `json:"install_date,omitempty"`
	TerminationDate NullableString               `json:"termination_date,omitempty"`
	// Committed rate
	CommitRate           NullableInt32                     `json:"commit_rate,omitempty"`
	Description          *string                           `json:"description,omitempty"`
	TerminationA         NullableCircuitCircuitTermination `json:"termination_a,omitempty"`
	TerminationZ         NullableCircuitCircuitTermination `json:"termination_z,omitempty"`
	Comments             *string                           `json:"comments,omitempty"`
	Tags                 []NestedTag                       `json:"tags,omitempty"`
	CustomFields         map[string]interface{}            `json:"custom_fields,omitempty"`
	Created              NullableTime                      `json:"created,omitempty"`
	LastUpdated          NullableTime                      `json:"last_updated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Circuit Circuit

// NewCircuit instantiates a new Circuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuit(id int32, url string, display string, cid string, provider BriefProvider, type_ BriefCircuitType) *Circuit {
	this := Circuit{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Cid = cid
	this.Provider = provider
	this.Type = type_
	return &this
}

// NewCircuitWithDefaults instantiates a new Circuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitWithDefaults() *Circuit {
	this := Circuit{}
	return &this
}

// GetId returns the Id field value
func (o *Circuit) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Circuit) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Circuit) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Circuit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Circuit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Circuit) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *Circuit) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Circuit) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Circuit) SetDisplay(v string) {
	o.Display = v
}

// GetCid returns the Cid field value
func (o *Circuit) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Circuit) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Circuit) SetCid(v string) {
	o.Cid = v
}

// GetProvider returns the Provider field value
func (o *Circuit) GetProvider() BriefProvider {
	if o == nil {
		var ret BriefProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Circuit) GetProviderOk() (*BriefProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *Circuit) SetProvider(v BriefProvider) {
	o.Provider = v
}

// GetProviderAccount returns the ProviderAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetProviderAccount() BriefProviderAccount {
	if o == nil || IsNil(o.ProviderAccount.Get()) {
		var ret BriefProviderAccount
		return ret
	}
	return *o.ProviderAccount.Get()
}

// GetProviderAccountOk returns a tuple with the ProviderAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetProviderAccountOk() (*BriefProviderAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderAccount.Get(), o.ProviderAccount.IsSet()
}

// HasProviderAccount returns a boolean if a field has been set.
func (o *Circuit) HasProviderAccount() bool {
	if o != nil && o.ProviderAccount.IsSet() {
		return true
	}

	return false
}

// SetProviderAccount gets a reference to the given NullableBriefProviderAccount and assigns it to the ProviderAccount field.
func (o *Circuit) SetProviderAccount(v BriefProviderAccount) {
	o.ProviderAccount.Set(&v)
}

// SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil
func (o *Circuit) SetProviderAccountNil() {
	o.ProviderAccount.Set(nil)
}

// UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
func (o *Circuit) UnsetProviderAccount() {
	o.ProviderAccount.Unset()
}

// GetType returns the Type field value
func (o *Circuit) GetType() BriefCircuitType {
	if o == nil {
		var ret BriefCircuitType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Circuit) GetTypeOk() (*BriefCircuitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Circuit) SetType(v BriefCircuitType) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Circuit) GetStatus() CircuitStatus {
	if o == nil || IsNil(o.Status) {
		var ret CircuitStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Circuit) GetStatusOk() (*CircuitStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Circuit) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CircuitStatus and assigns it to the Status field.
func (o *Circuit) SetStatus(v CircuitStatus) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Circuit) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *Circuit) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Circuit) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Circuit) UnsetTenant() {
	o.Tenant.Unset()
}

// GetInstallDate returns the InstallDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetInstallDate() string {
	if o == nil || IsNil(o.InstallDate.Get()) {
		var ret string
		return ret
	}
	return *o.InstallDate.Get()
}

// GetInstallDateOk returns a tuple with the InstallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetInstallDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallDate.Get(), o.InstallDate.IsSet()
}

// HasInstallDate returns a boolean if a field has been set.
func (o *Circuit) HasInstallDate() bool {
	if o != nil && o.InstallDate.IsSet() {
		return true
	}

	return false
}

// SetInstallDate gets a reference to the given NullableString and assigns it to the InstallDate field.
func (o *Circuit) SetInstallDate(v string) {
	o.InstallDate.Set(&v)
}

// SetInstallDateNil sets the value for InstallDate to be an explicit nil
func (o *Circuit) SetInstallDateNil() {
	o.InstallDate.Set(nil)
}

// UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
func (o *Circuit) UnsetInstallDate() {
	o.InstallDate.Unset()
}

// GetTerminationDate returns the TerminationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetTerminationDate() string {
	if o == nil || IsNil(o.TerminationDate.Get()) {
		var ret string
		return ret
	}
	return *o.TerminationDate.Get()
}

// GetTerminationDateOk returns a tuple with the TerminationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetTerminationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationDate.Get(), o.TerminationDate.IsSet()
}

// HasTerminationDate returns a boolean if a field has been set.
func (o *Circuit) HasTerminationDate() bool {
	if o != nil && o.TerminationDate.IsSet() {
		return true
	}

	return false
}

// SetTerminationDate gets a reference to the given NullableString and assigns it to the TerminationDate field.
func (o *Circuit) SetTerminationDate(v string) {
	o.TerminationDate.Set(&v)
}

// SetTerminationDateNil sets the value for TerminationDate to be an explicit nil
func (o *Circuit) SetTerminationDateNil() {
	o.TerminationDate.Set(nil)
}

// UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
func (o *Circuit) UnsetTerminationDate() {
	o.TerminationDate.Unset()
}

// GetCommitRate returns the CommitRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetCommitRate() int32 {
	if o == nil || IsNil(o.CommitRate.Get()) {
		var ret int32
		return ret
	}
	return *o.CommitRate.Get()
}

// GetCommitRateOk returns a tuple with the CommitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetCommitRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitRate.Get(), o.CommitRate.IsSet()
}

// HasCommitRate returns a boolean if a field has been set.
func (o *Circuit) HasCommitRate() bool {
	if o != nil && o.CommitRate.IsSet() {
		return true
	}

	return false
}

// SetCommitRate gets a reference to the given NullableInt32 and assigns it to the CommitRate field.
func (o *Circuit) SetCommitRate(v int32) {
	o.CommitRate.Set(&v)
}

// SetCommitRateNil sets the value for CommitRate to be an explicit nil
func (o *Circuit) SetCommitRateNil() {
	o.CommitRate.Set(nil)
}

// UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
func (o *Circuit) UnsetCommitRate() {
	o.CommitRate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Circuit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Circuit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Circuit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Circuit) SetDescription(v string) {
	o.Description = &v
}

// GetTerminationA returns the TerminationA field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetTerminationA() CircuitCircuitTermination {
	if o == nil || IsNil(o.TerminationA.Get()) {
		var ret CircuitCircuitTermination
		return ret
	}
	return *o.TerminationA.Get()
}

// GetTerminationAOk returns a tuple with the TerminationA field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetTerminationAOk() (*CircuitCircuitTermination, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationA.Get(), o.TerminationA.IsSet()
}

// HasTerminationA returns a boolean if a field has been set.
func (o *Circuit) HasTerminationA() bool {
	if o != nil && o.TerminationA.IsSet() {
		return true
	}

	return false
}

// SetTerminationA gets a reference to the given NullableCircuitCircuitTermination and assigns it to the TerminationA field.
func (o *Circuit) SetTerminationA(v CircuitCircuitTermination) {
	o.TerminationA.Set(&v)
}

// SetTerminationANil sets the value for TerminationA to be an explicit nil
func (o *Circuit) SetTerminationANil() {
	o.TerminationA.Set(nil)
}

// UnsetTerminationA ensures that no value is present for TerminationA, not even an explicit nil
func (o *Circuit) UnsetTerminationA() {
	o.TerminationA.Unset()
}

// GetTerminationZ returns the TerminationZ field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetTerminationZ() CircuitCircuitTermination {
	if o == nil || IsNil(o.TerminationZ.Get()) {
		var ret CircuitCircuitTermination
		return ret
	}
	return *o.TerminationZ.Get()
}

// GetTerminationZOk returns a tuple with the TerminationZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetTerminationZOk() (*CircuitCircuitTermination, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationZ.Get(), o.TerminationZ.IsSet()
}

// HasTerminationZ returns a boolean if a field has been set.
func (o *Circuit) HasTerminationZ() bool {
	if o != nil && o.TerminationZ.IsSet() {
		return true
	}

	return false
}

// SetTerminationZ gets a reference to the given NullableCircuitCircuitTermination and assigns it to the TerminationZ field.
func (o *Circuit) SetTerminationZ(v CircuitCircuitTermination) {
	o.TerminationZ.Set(&v)
}

// SetTerminationZNil sets the value for TerminationZ to be an explicit nil
func (o *Circuit) SetTerminationZNil() {
	o.TerminationZ.Set(nil)
}

// UnsetTerminationZ ensures that no value is present for TerminationZ, not even an explicit nil
func (o *Circuit) UnsetTerminationZ() {
	o.TerminationZ.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Circuit) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Circuit) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Circuit) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Circuit) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Circuit) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Circuit) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Circuit) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Circuit) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Circuit) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Circuit) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Circuit) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Circuit) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *Circuit) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *Circuit) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *Circuit) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *Circuit) UnsetCreated() {
	o.Created.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Circuit) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Circuit) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Circuit) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *Circuit) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *Circuit) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *Circuit) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

func (o Circuit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Circuit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["cid"] = o.Cid
	toSerialize["provider"] = o.Provider
	if o.ProviderAccount.IsSet() {
		toSerialize["provider_account"] = o.ProviderAccount.Get()
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.InstallDate.IsSet() {
		toSerialize["install_date"] = o.InstallDate.Get()
	}
	if o.TerminationDate.IsSet() {
		toSerialize["termination_date"] = o.TerminationDate.Get()
	}
	if o.CommitRate.IsSet() {
		toSerialize["commit_rate"] = o.CommitRate.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.TerminationA.IsSet() {
		toSerialize["termination_a"] = o.TerminationA.Get()
	}
	if o.TerminationZ.IsSet() {
		toSerialize["termination_z"] = o.TerminationZ.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.LastUpdated.IsSet() {
		toSerialize["last_updated"] = o.LastUpdated.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Circuit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"cid",
		"provider",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCircuit := _Circuit{}

	err = json.Unmarshal(data, &varCircuit)

	if err != nil {
		return err
	}

	*o = Circuit(varCircuit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "cid")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "provider_account")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "install_date")
		delete(additionalProperties, "termination_date")
		delete(additionalProperties, "commit_rate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "termination_a")
		delete(additionalProperties, "termination_z")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCircuit struct {
	value *Circuit
	isSet bool
}

func (v NullableCircuit) Get() *Circuit {
	return v.value
}

func (v *NullableCircuit) Set(val *Circuit) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuit(val *Circuit) *NullableCircuit {
	return &NullableCircuit{value: val, isSet: true}
}

func (v NullableCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

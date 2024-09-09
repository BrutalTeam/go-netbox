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

// checks if the VirtualDeviceContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualDeviceContext{}

// VirtualDeviceContext Adds support for custom fields and tags.
type VirtualDeviceContext struct {
	Id                   int32                      `json:"id"`
	Url                  string                     `json:"url"`
	Display              string                     `json:"display"`
	Name                 string                     `json:"name"`
	Device               BriefDevice                `json:"device"`
	Identifier           NullableInt32              `json:"identifier,omitempty"`
	Tenant               NullableBriefTenant        `json:"tenant,omitempty"`
	PrimaryIp            NullableBriefIPAddress     `json:"primary_ip,omitempty"`
	PrimaryIp4           NullableBriefIPAddress     `json:"primary_ip4,omitempty"`
	PrimaryIp6           NullableBriefIPAddress     `json:"primary_ip6,omitempty"`
	Status               VirtualDeviceContextStatus `json:"status"`
	Description          *string                    `json:"description,omitempty"`
	Comments             *string                    `json:"comments,omitempty"`
	Tags                 []NestedTag                `json:"tags,omitempty"`
	CustomFields         map[string]interface{}     `json:"custom_fields,omitempty"`
	Created              NullableTime               `json:"created,omitempty"`
	LastUpdated          NullableTime               `json:"last_updated,omitempty"`
	InterfaceCount       int64                      `json:"interface_count"`
	AdditionalProperties map[string]interface{}
}

type _VirtualDeviceContext VirtualDeviceContext

// NewVirtualDeviceContext instantiates a new VirtualDeviceContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualDeviceContext(id int32, url string, display string, name string, device BriefDevice, status VirtualDeviceContextStatus, interfaceCount int64) *VirtualDeviceContext {
	this := VirtualDeviceContext{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Device = device
	this.Status = status
	this.InterfaceCount = interfaceCount
	return &this
}

// NewVirtualDeviceContextWithDefaults instantiates a new VirtualDeviceContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualDeviceContextWithDefaults() *VirtualDeviceContext {
	this := VirtualDeviceContext{}
	return &this
}

// GetId returns the Id field value
func (o *VirtualDeviceContext) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VirtualDeviceContext) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VirtualDeviceContext) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VirtualDeviceContext) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *VirtualDeviceContext) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VirtualDeviceContext) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *VirtualDeviceContext) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualDeviceContext) SetName(v string) {
	o.Name = v
}

// GetDevice returns the Device field value
func (o *VirtualDeviceContext) GetDevice() BriefDevice {
	if o == nil {
		var ret BriefDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *VirtualDeviceContext) SetDevice(v BriefDevice) {
	o.Device = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDeviceContext) GetIdentifier() int32 {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret int32
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDeviceContext) GetIdentifierOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableInt32 and assigns it to the Identifier field.
func (o *VirtualDeviceContext) SetIdentifier(v int32) {
	o.Identifier.Set(&v)
}

// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *VirtualDeviceContext) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *VirtualDeviceContext) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDeviceContext) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDeviceContext) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *VirtualDeviceContext) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *VirtualDeviceContext) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *VirtualDeviceContext) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPrimaryIp returns the PrimaryIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDeviceContext) GetPrimaryIp() BriefIPAddress {
	if o == nil || IsNil(o.PrimaryIp.Get()) {
		var ret BriefIPAddress
		return ret
	}
	return *o.PrimaryIp.Get()
}

// GetPrimaryIpOk returns a tuple with the PrimaryIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDeviceContext) GetPrimaryIpOk() (*BriefIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp.Get(), o.PrimaryIp.IsSet()
}

// HasPrimaryIp returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasPrimaryIp() bool {
	if o != nil && o.PrimaryIp.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp gets a reference to the given NullableBriefIPAddress and assigns it to the PrimaryIp field.
func (o *VirtualDeviceContext) SetPrimaryIp(v BriefIPAddress) {
	o.PrimaryIp.Set(&v)
}

// SetPrimaryIpNil sets the value for PrimaryIp to be an explicit nil
func (o *VirtualDeviceContext) SetPrimaryIpNil() {
	o.PrimaryIp.Set(nil)
}

// UnsetPrimaryIp ensures that no value is present for PrimaryIp, not even an explicit nil
func (o *VirtualDeviceContext) UnsetPrimaryIp() {
	o.PrimaryIp.Unset()
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDeviceContext) GetPrimaryIp4() BriefIPAddress {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret BriefIPAddress
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDeviceContext) GetPrimaryIp4Ok() (*BriefIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableBriefIPAddress and assigns it to the PrimaryIp4 field.
func (o *VirtualDeviceContext) SetPrimaryIp4(v BriefIPAddress) {
	o.PrimaryIp4.Set(&v)
}

// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *VirtualDeviceContext) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *VirtualDeviceContext) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDeviceContext) GetPrimaryIp6() BriefIPAddress {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret BriefIPAddress
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDeviceContext) GetPrimaryIp6Ok() (*BriefIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableBriefIPAddress and assigns it to the PrimaryIp6 field.
func (o *VirtualDeviceContext) SetPrimaryIp6(v BriefIPAddress) {
	o.PrimaryIp6.Set(&v)
}

// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *VirtualDeviceContext) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *VirtualDeviceContext) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetStatus returns the Status field value
func (o *VirtualDeviceContext) GetStatus() VirtualDeviceContextStatus {
	if o == nil {
		var ret VirtualDeviceContextStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetStatusOk() (*VirtualDeviceContextStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VirtualDeviceContext) SetStatus(v VirtualDeviceContextStatus) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualDeviceContext) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualDeviceContext) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VirtualDeviceContext) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VirtualDeviceContext) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualDeviceContext) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *VirtualDeviceContext) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VirtualDeviceContext) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VirtualDeviceContext) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDeviceContext) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDeviceContext) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *VirtualDeviceContext) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *VirtualDeviceContext) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *VirtualDeviceContext) UnsetCreated() {
	o.Created.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDeviceContext) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDeviceContext) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *VirtualDeviceContext) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *VirtualDeviceContext) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *VirtualDeviceContext) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *VirtualDeviceContext) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

// GetInterfaceCount returns the InterfaceCount field value
func (o *VirtualDeviceContext) GetInterfaceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InterfaceCount
}

// GetInterfaceCountOk returns a tuple with the InterfaceCount field value
// and a boolean to check if the value has been set.
func (o *VirtualDeviceContext) GetInterfaceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceCount, true
}

// SetInterfaceCount sets field value
func (o *VirtualDeviceContext) SetInterfaceCount(v int64) {
	o.InterfaceCount = v
}

func (o VirtualDeviceContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualDeviceContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["device"] = o.Device
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.PrimaryIp.IsSet() {
		toSerialize["primary_ip"] = o.PrimaryIp.Get()
	}
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	toSerialize["interface_count"] = o.InterfaceCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualDeviceContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"device",
		"status",
		"interface_count",
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

	varVirtualDeviceContext := _VirtualDeviceContext{}

	err = json.Unmarshal(data, &varVirtualDeviceContext)

	if err != nil {
		return err
	}

	*o = VirtualDeviceContext(varVirtualDeviceContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "device")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "primary_ip")
		delete(additionalProperties, "primary_ip4")
		delete(additionalProperties, "primary_ip6")
		delete(additionalProperties, "status")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "interface_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualDeviceContext struct {
	value *VirtualDeviceContext
	isSet bool
}

func (v NullableVirtualDeviceContext) Get() *VirtualDeviceContext {
	return v.value
}

func (v *NullableVirtualDeviceContext) Set(val *VirtualDeviceContext) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualDeviceContext) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualDeviceContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualDeviceContext(val *VirtualDeviceContext) *NullableVirtualDeviceContext {
	return &NullableVirtualDeviceContext{value: val, isSet: true}
}

func (v NullableVirtualDeviceContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualDeviceContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

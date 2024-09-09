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

// checks if the L2VPN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L2VPN{}

// L2VPN Adds support for custom fields and tags.
type L2VPN struct {
	Id                   int32                  `json:"id"`
	Url                  string                 `json:"url"`
	Display              string                 `json:"display"`
	Identifier           NullableInt64          `json:"identifier,omitempty"`
	Name                 string                 `json:"name"`
	Slug                 string                 `json:"slug"`
	Type                 *BriefL2VPNType        `json:"type,omitempty"`
	ImportTargets        []RouteTarget          `json:"import_targets,omitempty"`
	ExportTargets        []RouteTarget          `json:"export_targets,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tenant               NullableBriefTenant    `json:"tenant,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created,omitempty"`
	LastUpdated          NullableTime           `json:"last_updated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _L2VPN L2VPN

// NewL2VPN instantiates a new L2VPN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL2VPN(id int32, url string, display string, name string, slug string) *L2VPN {
	this := L2VPN{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewL2VPNWithDefaults instantiates a new L2VPN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL2VPNWithDefaults() *L2VPN {
	this := L2VPN{}
	return &this
}

// GetId returns the Id field value
func (o *L2VPN) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *L2VPN) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *L2VPN) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *L2VPN) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *L2VPN) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *L2VPN) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *L2VPN) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *L2VPN) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *L2VPN) SetDisplay(v string) {
	o.Display = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *L2VPN) GetIdentifier() int64 {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret int64
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPN) GetIdentifierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *L2VPN) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableInt64 and assigns it to the Identifier field.
func (o *L2VPN) SetIdentifier(v int64) {
	o.Identifier.Set(&v)
}

// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *L2VPN) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *L2VPN) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetName returns the Name field value
func (o *L2VPN) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *L2VPN) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *L2VPN) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *L2VPN) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *L2VPN) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *L2VPN) SetSlug(v string) {
	o.Slug = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *L2VPN) GetType() BriefL2VPNType {
	if o == nil || IsNil(o.Type) {
		var ret BriefL2VPNType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPN) GetTypeOk() (*BriefL2VPNType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *L2VPN) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given BriefL2VPNType and assigns it to the Type field.
func (o *L2VPN) SetType(v BriefL2VPNType) {
	o.Type = &v
}

// GetImportTargets returns the ImportTargets field value if set, zero value otherwise.
func (o *L2VPN) GetImportTargets() []RouteTarget {
	if o == nil || IsNil(o.ImportTargets) {
		var ret []RouteTarget
		return ret
	}
	return o.ImportTargets
}

// GetImportTargetsOk returns a tuple with the ImportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPN) GetImportTargetsOk() ([]RouteTarget, bool) {
	if o == nil || IsNil(o.ImportTargets) {
		return nil, false
	}
	return o.ImportTargets, true
}

// HasImportTargets returns a boolean if a field has been set.
func (o *L2VPN) HasImportTargets() bool {
	if o != nil && !IsNil(o.ImportTargets) {
		return true
	}

	return false
}

// SetImportTargets gets a reference to the given []RouteTarget and assigns it to the ImportTargets field.
func (o *L2VPN) SetImportTargets(v []RouteTarget) {
	o.ImportTargets = v
}

// GetExportTargets returns the ExportTargets field value if set, zero value otherwise.
func (o *L2VPN) GetExportTargets() []RouteTarget {
	if o == nil || IsNil(o.ExportTargets) {
		var ret []RouteTarget
		return ret
	}
	return o.ExportTargets
}

// GetExportTargetsOk returns a tuple with the ExportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPN) GetExportTargetsOk() ([]RouteTarget, bool) {
	if o == nil || IsNil(o.ExportTargets) {
		return nil, false
	}
	return o.ExportTargets, true
}

// HasExportTargets returns a boolean if a field has been set.
func (o *L2VPN) HasExportTargets() bool {
	if o != nil && !IsNil(o.ExportTargets) {
		return true
	}

	return false
}

// SetExportTargets gets a reference to the given []RouteTarget and assigns it to the ExportTargets field.
func (o *L2VPN) SetExportTargets(v []RouteTarget) {
	o.ExportTargets = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *L2VPN) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPN) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *L2VPN) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *L2VPN) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *L2VPN) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPN) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *L2VPN) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *L2VPN) SetComments(v string) {
	o.Comments = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *L2VPN) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPN) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *L2VPN) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *L2VPN) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *L2VPN) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *L2VPN) UnsetTenant() {
	o.Tenant.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *L2VPN) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPN) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *L2VPN) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *L2VPN) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *L2VPN) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPN) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *L2VPN) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *L2VPN) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *L2VPN) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPN) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *L2VPN) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *L2VPN) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *L2VPN) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *L2VPN) UnsetCreated() {
	o.Created.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *L2VPN) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPN) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *L2VPN) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *L2VPN) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *L2VPN) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *L2VPN) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

func (o L2VPN) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L2VPN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ImportTargets) {
		toSerialize["import_targets"] = o.ImportTargets
	}
	if !IsNil(o.ExportTargets) {
		toSerialize["export_targets"] = o.ExportTargets
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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

func (o *L2VPN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
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

	varL2VPN := _L2VPN{}

	err = json.Unmarshal(data, &varL2VPN)

	if err != nil {
		return err
	}

	*o = L2VPN(varL2VPN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "type")
		delete(additionalProperties, "import_targets")
		delete(additionalProperties, "export_targets")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableL2VPN struct {
	value *L2VPN
	isSet bool
}

func (v NullableL2VPN) Get() *L2VPN {
	return v.value
}

func (v *NullableL2VPN) Set(val *L2VPN) {
	v.value = val
	v.isSet = true
}

func (v NullableL2VPN) IsSet() bool {
	return v.isSet
}

func (v *NullableL2VPN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2VPN(val *L2VPN) *NullableL2VPN {
	return &NullableL2VPN{value: val, isSet: true}
}

func (v NullableL2VPN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2VPN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

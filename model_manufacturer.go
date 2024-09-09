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

// checks if the Manufacturer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Manufacturer{}

// Manufacturer Adds support for custom fields and tags.
type Manufacturer struct {
	Id                   int32                  `json:"id"`
	Url                  string                 `json:"url"`
	Display              string                 `json:"display"`
	Name                 string                 `json:"name"`
	Slug                 string                 `json:"slug"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created,omitempty"`
	LastUpdated          NullableTime           `json:"last_updated,omitempty"`
	DevicetypeCount      *int64                 `json:"devicetype_count,omitempty"`
	InventoryitemCount   int64                  `json:"inventoryitem_count"`
	PlatformCount        int64                  `json:"platform_count"`
	AdditionalProperties map[string]interface{}
}

type _Manufacturer Manufacturer

// NewManufacturer instantiates a new Manufacturer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManufacturer(id int32, url string, display string, name string, slug string, inventoryitemCount int64, platformCount int64) *Manufacturer {
	this := Manufacturer{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.InventoryitemCount = inventoryitemCount
	this.PlatformCount = platformCount
	return &this
}

// NewManufacturerWithDefaults instantiates a new Manufacturer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManufacturerWithDefaults() *Manufacturer {
	this := Manufacturer{}
	return &this
}

// GetId returns the Id field value
func (o *Manufacturer) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Manufacturer) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Manufacturer) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Manufacturer) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *Manufacturer) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Manufacturer) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *Manufacturer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Manufacturer) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Manufacturer) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Manufacturer) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Manufacturer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Manufacturer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Manufacturer) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Manufacturer) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Manufacturer) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Manufacturer) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Manufacturer) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Manufacturer) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Manufacturer) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Manufacturer) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Manufacturer) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *Manufacturer) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *Manufacturer) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *Manufacturer) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *Manufacturer) UnsetCreated() {
	o.Created.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Manufacturer) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Manufacturer) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Manufacturer) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *Manufacturer) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *Manufacturer) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *Manufacturer) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

// GetDevicetypeCount returns the DevicetypeCount field value if set, zero value otherwise.
func (o *Manufacturer) GetDevicetypeCount() int64 {
	if o == nil || IsNil(o.DevicetypeCount) {
		var ret int64
		return ret
	}
	return *o.DevicetypeCount
}

// GetDevicetypeCountOk returns a tuple with the DevicetypeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetDevicetypeCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DevicetypeCount) {
		return nil, false
	}
	return o.DevicetypeCount, true
}

// HasDevicetypeCount returns a boolean if a field has been set.
func (o *Manufacturer) HasDevicetypeCount() bool {
	if o != nil && !IsNil(o.DevicetypeCount) {
		return true
	}

	return false
}

// SetDevicetypeCount gets a reference to the given int64 and assigns it to the DevicetypeCount field.
func (o *Manufacturer) SetDevicetypeCount(v int64) {
	o.DevicetypeCount = &v
}

// GetInventoryitemCount returns the InventoryitemCount field value
func (o *Manufacturer) GetInventoryitemCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InventoryitemCount
}

// GetInventoryitemCountOk returns a tuple with the InventoryitemCount field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetInventoryitemCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryitemCount, true
}

// SetInventoryitemCount sets field value
func (o *Manufacturer) SetInventoryitemCount(v int64) {
	o.InventoryitemCount = v
}

// GetPlatformCount returns the PlatformCount field value
func (o *Manufacturer) GetPlatformCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlatformCount
}

// GetPlatformCountOk returns a tuple with the PlatformCount field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetPlatformCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformCount, true
}

// SetPlatformCount sets field value
func (o *Manufacturer) SetPlatformCount(v int64) {
	o.PlatformCount = v
}

func (o Manufacturer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Manufacturer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	if !IsNil(o.DevicetypeCount) {
		toSerialize["devicetype_count"] = o.DevicetypeCount
	}
	toSerialize["inventoryitem_count"] = o.InventoryitemCount
	toSerialize["platform_count"] = o.PlatformCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Manufacturer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
		"inventoryitem_count",
		"platform_count",
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

	varManufacturer := _Manufacturer{}

	err = json.Unmarshal(data, &varManufacturer)

	if err != nil {
		return err
	}

	*o = Manufacturer(varManufacturer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "devicetype_count")
		delete(additionalProperties, "inventoryitem_count")
		delete(additionalProperties, "platform_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManufacturer struct {
	value *Manufacturer
	isSet bool
}

func (v NullableManufacturer) Get() *Manufacturer {
	return v.value
}

func (v *NullableManufacturer) Set(val *Manufacturer) {
	v.value = val
	v.isSet = true
}

func (v NullableManufacturer) IsSet() bool {
	return v.isSet
}

func (v *NullableManufacturer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManufacturer(val *Manufacturer) *NullableManufacturer {
	return &NullableManufacturer{value: val, isSet: true}
}

func (v NullableManufacturer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManufacturer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

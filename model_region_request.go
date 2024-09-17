/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the RegionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionRequest{}

// RegionRequest Extends PrimaryModelSerializer to include MPTT support.
type RegionRequest struct {
	Name                 *string                     `json:"name,omitempty"`
	Slug                 *string                     `json:"slug,omitempty"`
	Parent               NullableNestedRegionRequest `json:"parent,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	Tags                 []NestedTagRequest          `json:"tags,omitempty"`
	CustomFields         map[string]interface{}      `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegionRequest RegionRequest

// NewRegionRequest instantiates a new RegionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionRequest() *RegionRequest {
	this := RegionRequest{}
	return &this
}

// NewRegionRequestWithDefaults instantiates a new RegionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionRequestWithDefaults() *RegionRequest {
	this := RegionRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegionRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegionRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegionRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *RegionRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *RegionRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *RegionRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegionRequest) GetParent() NestedRegionRequest {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret NestedRegionRequest
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionRequest) GetParentOk() (*NestedRegionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *RegionRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableNestedRegionRequest and assigns it to the Parent field.
func (o *RegionRequest) SetParent(v NestedRegionRequest) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *RegionRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *RegionRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RegionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RegionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RegionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RegionRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RegionRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *RegionRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RegionRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RegionRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *RegionRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o RegionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegionRequest) UnmarshalJSON(data []byte) (err error) {
	varRegionRequest := _RegionRequest{}

	err = json.Unmarshal(data, &varRegionRequest)

	if err != nil {
		return err
	}

	*o = RegionRequest(varRegionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegionRequest struct {
	value *RegionRequest
	isSet bool
}

func (v NullableRegionRequest) Get() *RegionRequest {
	return v.value
}

func (v *NullableRegionRequest) Set(val *RegionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionRequest(val *RegionRequest) *NullableRegionRequest {
	return &NullableRegionRequest{value: val, isSet: true}
}

func (v NullableRegionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

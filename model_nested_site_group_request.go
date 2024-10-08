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

// checks if the NestedSiteGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedSiteGroupRequest{}

// NestedSiteGroupRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedSiteGroupRequest struct {
	Name                 *string `json:"name,omitempty"`
	Slug                 *string `json:"slug,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedSiteGroupRequest NestedSiteGroupRequest

// NewNestedSiteGroupRequest instantiates a new NestedSiteGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedSiteGroupRequest() *NestedSiteGroupRequest {
	this := NestedSiteGroupRequest{}
	return &this
}

// NewNestedSiteGroupRequestWithDefaults instantiates a new NestedSiteGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedSiteGroupRequestWithDefaults() *NestedSiteGroupRequest {
	this := NestedSiteGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NestedSiteGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedSiteGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NestedSiteGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NestedSiteGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *NestedSiteGroupRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedSiteGroupRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *NestedSiteGroupRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *NestedSiteGroupRequest) SetSlug(v string) {
	o.Slug = &v
}

func (o NestedSiteGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedSiteGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedSiteGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varNestedSiteGroupRequest := _NestedSiteGroupRequest{}

	err = json.Unmarshal(data, &varNestedSiteGroupRequest)

	if err != nil {
		return err
	}

	*o = NestedSiteGroupRequest(varNestedSiteGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedSiteGroupRequest struct {
	value *NestedSiteGroupRequest
	isSet bool
}

func (v NullableNestedSiteGroupRequest) Get() *NestedSiteGroupRequest {
	return v.value
}

func (v *NullableNestedSiteGroupRequest) Set(val *NestedSiteGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedSiteGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedSiteGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedSiteGroupRequest(val *NestedSiteGroupRequest) *NullableNestedSiteGroupRequest {
	return &NullableNestedSiteGroupRequest{value: val, isSet: true}
}

func (v NullableNestedSiteGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedSiteGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

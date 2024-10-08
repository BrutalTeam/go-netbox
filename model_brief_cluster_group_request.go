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

// checks if the BriefClusterGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefClusterGroupRequest{}

// BriefClusterGroupRequest Adds support for custom fields and tags.
type BriefClusterGroupRequest struct {
	Name                 *string `json:"name,omitempty"`
	Slug                 *string `json:"slug,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefClusterGroupRequest BriefClusterGroupRequest

// NewBriefClusterGroupRequest instantiates a new BriefClusterGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefClusterGroupRequest() *BriefClusterGroupRequest {
	this := BriefClusterGroupRequest{}
	return &this
}

// NewBriefClusterGroupRequestWithDefaults instantiates a new BriefClusterGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefClusterGroupRequestWithDefaults() *BriefClusterGroupRequest {
	this := BriefClusterGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BriefClusterGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefClusterGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BriefClusterGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BriefClusterGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *BriefClusterGroupRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefClusterGroupRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *BriefClusterGroupRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *BriefClusterGroupRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefClusterGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefClusterGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefClusterGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefClusterGroupRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefClusterGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefClusterGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefClusterGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varBriefClusterGroupRequest := _BriefClusterGroupRequest{}

	err = json.Unmarshal(data, &varBriefClusterGroupRequest)

	if err != nil {
		return err
	}

	*o = BriefClusterGroupRequest(varBriefClusterGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefClusterGroupRequest struct {
	value *BriefClusterGroupRequest
	isSet bool
}

func (v NullableBriefClusterGroupRequest) Get() *BriefClusterGroupRequest {
	return v.value
}

func (v *NullableBriefClusterGroupRequest) Set(val *BriefClusterGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefClusterGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefClusterGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefClusterGroupRequest(val *BriefClusterGroupRequest) *NullableBriefClusterGroupRequest {
	return &NullableBriefClusterGroupRequest{value: val, isSet: true}
}

func (v NullableBriefClusterGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefClusterGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the BriefIPSecProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefIPSecProfileRequest{}

// BriefIPSecProfileRequest Adds support for custom fields and tags.
type BriefIPSecProfileRequest struct {
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefIPSecProfileRequest BriefIPSecProfileRequest

// NewBriefIPSecProfileRequest instantiates a new BriefIPSecProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefIPSecProfileRequest() *BriefIPSecProfileRequest {
	this := BriefIPSecProfileRequest{}
	return &this
}

// NewBriefIPSecProfileRequestWithDefaults instantiates a new BriefIPSecProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefIPSecProfileRequestWithDefaults() *BriefIPSecProfileRequest {
	this := BriefIPSecProfileRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BriefIPSecProfileRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefIPSecProfileRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BriefIPSecProfileRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BriefIPSecProfileRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefIPSecProfileRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefIPSecProfileRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefIPSecProfileRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefIPSecProfileRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefIPSecProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefIPSecProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefIPSecProfileRequest) UnmarshalJSON(data []byte) (err error) {
	varBriefIPSecProfileRequest := _BriefIPSecProfileRequest{}

	err = json.Unmarshal(data, &varBriefIPSecProfileRequest)

	if err != nil {
		return err
	}

	*o = BriefIPSecProfileRequest(varBriefIPSecProfileRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefIPSecProfileRequest struct {
	value *BriefIPSecProfileRequest
	isSet bool
}

func (v NullableBriefIPSecProfileRequest) Get() *BriefIPSecProfileRequest {
	return v.value
}

func (v *NullableBriefIPSecProfileRequest) Set(val *BriefIPSecProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefIPSecProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefIPSecProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefIPSecProfileRequest(val *BriefIPSecProfileRequest) *NullableBriefIPSecProfileRequest {
	return &NullableBriefIPSecProfileRequest{value: val, isSet: true}
}

func (v NullableBriefIPSecProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefIPSecProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

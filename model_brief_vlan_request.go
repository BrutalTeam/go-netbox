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

// checks if the BriefVLANRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefVLANRequest{}

// BriefVLANRequest Adds support for custom fields and tags.
type BriefVLANRequest struct {
	// Numeric VLAN ID (1-4094)
	Vid                  *int32  `json:"vid,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefVLANRequest BriefVLANRequest

// NewBriefVLANRequest instantiates a new BriefVLANRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefVLANRequest() *BriefVLANRequest {
	this := BriefVLANRequest{}
	return &this
}

// NewBriefVLANRequestWithDefaults instantiates a new BriefVLANRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefVLANRequestWithDefaults() *BriefVLANRequest {
	this := BriefVLANRequest{}
	return &this
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *BriefVLANRequest) GetVid() int32 {
	if o == nil || IsNil(o.Vid) {
		var ret int32
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVLANRequest) GetVidOk() (*int32, bool) {
	if o == nil || IsNil(o.Vid) {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *BriefVLANRequest) HasVid() bool {
	if o != nil && !IsNil(o.Vid) {
		return true
	}

	return false
}

// SetVid gets a reference to the given int32 and assigns it to the Vid field.
func (o *BriefVLANRequest) SetVid(v int32) {
	o.Vid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BriefVLANRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVLANRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BriefVLANRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BriefVLANRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefVLANRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVLANRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefVLANRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefVLANRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefVLANRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefVLANRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vid) {
		toSerialize["vid"] = o.Vid
	}
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

func (o *BriefVLANRequest) UnmarshalJSON(data []byte) (err error) {
	varBriefVLANRequest := _BriefVLANRequest{}

	err = json.Unmarshal(data, &varBriefVLANRequest)

	if err != nil {
		return err
	}

	*o = BriefVLANRequest(varBriefVLANRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefVLANRequest struct {
	value *BriefVLANRequest
	isSet bool
}

func (v NullableBriefVLANRequest) Get() *BriefVLANRequest {
	return v.value
}

func (v *NullableBriefVLANRequest) Set(val *BriefVLANRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefVLANRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefVLANRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefVLANRequest(val *BriefVLANRequest) *NullableBriefVLANRequest {
	return &NullableBriefVLANRequest{value: val, isSet: true}
}

func (v NullableBriefVLANRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefVLANRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

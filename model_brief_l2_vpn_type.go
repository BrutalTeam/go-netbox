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

// checks if the BriefL2VPNType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefL2VPNType{}

// BriefL2VPNType struct for BriefL2VPNType
type BriefL2VPNType struct {
	Value                *BriefL2VPNTypeValue `json:"value,omitempty"`
	Label                *BriefL2VPNTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefL2VPNType BriefL2VPNType

// NewBriefL2VPNType instantiates a new BriefL2VPNType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefL2VPNType() *BriefL2VPNType {
	this := BriefL2VPNType{}
	return &this
}

// NewBriefL2VPNTypeWithDefaults instantiates a new BriefL2VPNType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefL2VPNTypeWithDefaults() *BriefL2VPNType {
	this := BriefL2VPNType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BriefL2VPNType) GetValue() BriefL2VPNTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret BriefL2VPNTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefL2VPNType) GetValueOk() (*BriefL2VPNTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BriefL2VPNType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given BriefL2VPNTypeValue and assigns it to the Value field.
func (o *BriefL2VPNType) SetValue(v BriefL2VPNTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BriefL2VPNType) GetLabel() BriefL2VPNTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret BriefL2VPNTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefL2VPNType) GetLabelOk() (*BriefL2VPNTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BriefL2VPNType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given BriefL2VPNTypeLabel and assigns it to the Label field.
func (o *BriefL2VPNType) SetLabel(v BriefL2VPNTypeLabel) {
	o.Label = &v
}

func (o BriefL2VPNType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefL2VPNType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefL2VPNType) UnmarshalJSON(data []byte) (err error) {
	varBriefL2VPNType := _BriefL2VPNType{}

	err = json.Unmarshal(data, &varBriefL2VPNType)

	if err != nil {
		return err
	}

	*o = BriefL2VPNType(varBriefL2VPNType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefL2VPNType struct {
	value *BriefL2VPNType
	isSet bool
}

func (v NullableBriefL2VPNType) Get() *BriefL2VPNType {
	return v.value
}

func (v *NullableBriefL2VPNType) Set(val *BriefL2VPNType) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefL2VPNType) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefL2VPNType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefL2VPNType(val *BriefL2VPNType) *NullableBriefL2VPNType {
	return &NullableBriefL2VPNType{value: val, isSet: true}
}

func (v NullableBriefL2VPNType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefL2VPNType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

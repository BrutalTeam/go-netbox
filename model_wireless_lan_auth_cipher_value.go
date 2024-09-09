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
)

// WirelessLANAuthCipherValue * `auto` - Auto * `tkip` - TKIP * `aes` - AES
type WirelessLANAuthCipherValue string

// List of WirelessLAN_auth_cipher_value
const (
	WIRELESSLANAUTHCIPHERVALUE_AUTO  WirelessLANAuthCipherValue = "auto"
	WIRELESSLANAUTHCIPHERVALUE_TKIP  WirelessLANAuthCipherValue = "tkip"
	WIRELESSLANAUTHCIPHERVALUE_AES   WirelessLANAuthCipherValue = "aes"
	WIRELESSLANAUTHCIPHERVALUE_EMPTY WirelessLANAuthCipherValue = ""
)

// All allowed values of WirelessLANAuthCipherValue enum
var AllowedWirelessLANAuthCipherValueEnumValues = []WirelessLANAuthCipherValue{
	"auto",
	"tkip",
	"aes",
	"",
}

func (v *WirelessLANAuthCipherValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessLANAuthCipherValue(value)
	for _, existing := range AllowedWirelessLANAuthCipherValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessLANAuthCipherValue", value)
}

// NewWirelessLANAuthCipherValueFromValue returns a pointer to a valid WirelessLANAuthCipherValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessLANAuthCipherValueFromValue(v string) (*WirelessLANAuthCipherValue, error) {
	ev := WirelessLANAuthCipherValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessLANAuthCipherValue: valid values are %v", v, AllowedWirelessLANAuthCipherValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessLANAuthCipherValue) IsValid() bool {
	for _, existing := range AllowedWirelessLANAuthCipherValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WirelessLAN_auth_cipher_value value
func (v WirelessLANAuthCipherValue) Ptr() *WirelessLANAuthCipherValue {
	return &v
}

type NullableWirelessLANAuthCipherValue struct {
	value *WirelessLANAuthCipherValue
	isSet bool
}

func (v NullableWirelessLANAuthCipherValue) Get() *WirelessLANAuthCipherValue {
	return v.value
}

func (v *NullableWirelessLANAuthCipherValue) Set(val *WirelessLANAuthCipherValue) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLANAuthCipherValue) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLANAuthCipherValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLANAuthCipherValue(val *WirelessLANAuthCipherValue) *NullableWirelessLANAuthCipherValue {
	return &NullableWirelessLANAuthCipherValue{value: val, isSet: true}
}

func (v NullableWirelessLANAuthCipherValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLANAuthCipherValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

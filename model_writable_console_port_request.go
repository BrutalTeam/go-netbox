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

// checks if the WritableConsolePortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableConsolePortRequest{}

// WritableConsolePortRequest Adds support for custom fields and tags.
type WritableConsolePortRequest struct {
	Device BriefDeviceRequest         `json:"device"`
	Module NullableBriefModuleRequest `json:"module,omitempty"`
	Name   *string                    `json:"name,omitempty"`
	// Physical label
	Label       *string                                        `json:"label,omitempty"`
	Type        *PatchedWritableConsolePortRequestType         `json:"type,omitempty"`
	Speed       NullablePatchedWritableConsolePortRequestSpeed `json:"speed,omitempty"`
	Description *string                                        `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableConsolePortRequest WritableConsolePortRequest

// NewWritableConsolePortRequest instantiates a new WritableConsolePortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableConsolePortRequest(device BriefDeviceRequest) *WritableConsolePortRequest {
	this := WritableConsolePortRequest{}
	this.Device = device
	return &this
}

// NewWritableConsolePortRequestWithDefaults instantiates a new WritableConsolePortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableConsolePortRequestWithDefaults() *WritableConsolePortRequest {
	this := WritableConsolePortRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *WritableConsolePortRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *WritableConsolePortRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *WritableConsolePortRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableConsolePortRequest) GetModule() BriefModuleRequest {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModuleRequest
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConsolePortRequest) GetModuleOk() (*BriefModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModuleRequest and assigns it to the Module field.
func (o *WritableConsolePortRequest) SetModule(v BriefModuleRequest) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *WritableConsolePortRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *WritableConsolePortRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WritableConsolePortRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsolePortRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WritableConsolePortRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableConsolePortRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsolePortRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableConsolePortRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritableConsolePortRequest) GetType() PatchedWritableConsolePortRequestType {
	if o == nil || IsNil(o.Type) {
		var ret PatchedWritableConsolePortRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsolePortRequest) GetTypeOk() (*PatchedWritableConsolePortRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PatchedWritableConsolePortRequestType and assigns it to the Type field.
func (o *WritableConsolePortRequest) SetType(v PatchedWritableConsolePortRequestType) {
	o.Type = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableConsolePortRequest) GetSpeed() PatchedWritableConsolePortRequestSpeed {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret PatchedWritableConsolePortRequestSpeed
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConsolePortRequest) GetSpeedOk() (*PatchedWritableConsolePortRequestSpeed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullablePatchedWritableConsolePortRequestSpeed and assigns it to the Speed field.
func (o *WritableConsolePortRequest) SetSpeed(v PatchedWritableConsolePortRequestSpeed) {
	o.Speed.Set(&v)
}

// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *WritableConsolePortRequest) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *WritableConsolePortRequest) UnsetSpeed() {
	o.Speed.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableConsolePortRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsolePortRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableConsolePortRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *WritableConsolePortRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsolePortRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *WritableConsolePortRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableConsolePortRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsolePortRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableConsolePortRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableConsolePortRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsolePortRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableConsolePortRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableConsolePortRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableConsolePortRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableConsolePortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
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

func (o *WritableConsolePortRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
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

	varWritableConsolePortRequest := _WritableConsolePortRequest{}

	err = json.Unmarshal(data, &varWritableConsolePortRequest)

	if err != nil {
		return err
	}

	*o = WritableConsolePortRequest(varWritableConsolePortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableConsolePortRequest struct {
	value *WritableConsolePortRequest
	isSet bool
}

func (v NullableWritableConsolePortRequest) Get() *WritableConsolePortRequest {
	return v.value
}

func (v *NullableWritableConsolePortRequest) Set(val *WritableConsolePortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableConsolePortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableConsolePortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableConsolePortRequest(val *WritableConsolePortRequest) *NullableWritableConsolePortRequest {
	return &NullableWritableConsolePortRequest{value: val, isSet: true}
}

func (v NullableWritableConsolePortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableConsolePortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

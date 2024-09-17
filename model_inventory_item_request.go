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

// checks if the InventoryItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryItemRequest{}

// InventoryItemRequest Adds support for custom fields and tags.
type InventoryItemRequest struct {
	Device BriefDeviceRequest `json:"device"`
	Parent NullableInt32      `json:"parent,omitempty"`
	Name   *string            `json:"name,omitempty"`
	// Physical label
	Label        *string                               `json:"label,omitempty"`
	Role         NullableBriefInventoryItemRoleRequest `json:"role,omitempty"`
	Manufacturer NullableBriefManufacturerRequest      `json:"manufacturer,omitempty"`
	// Manufacturer-assigned part identifier
	PartId *string `json:"part_id,omitempty"`
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this item
	AssetTag NullableString `json:"asset_tag,omitempty"`
	// This item was automatically discovered
	Discovered           *bool                  `json:"discovered,omitempty"`
	Description          *string                `json:"description,omitempty"`
	ComponentType        NullableString         `json:"component_type,omitempty"`
	ComponentId          NullableInt64          `json:"component_id,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryItemRequest InventoryItemRequest

// NewInventoryItemRequest instantiates a new InventoryItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItemRequest(device BriefDeviceRequest) *InventoryItemRequest {
	this := InventoryItemRequest{}
	this.Device = device
	return &this
}

// NewInventoryItemRequestWithDefaults instantiates a new InventoryItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemRequestWithDefaults() *InventoryItemRequest {
	this := InventoryItemRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *InventoryItemRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *InventoryItemRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemRequest) GetParent() int32 {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret int32
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemRequest) GetParentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableInt32 and assigns it to the Parent field.
func (o *InventoryItemRequest) SetParent(v int32) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *InventoryItemRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *InventoryItemRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InventoryItemRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InventoryItemRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InventoryItemRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InventoryItemRequest) SetLabel(v string) {
	o.Label = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemRequest) GetRole() BriefInventoryItemRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefInventoryItemRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemRequest) GetRoleOk() (*BriefInventoryItemRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefInventoryItemRoleRequest and assigns it to the Role field.
func (o *InventoryItemRequest) SetRole(v BriefInventoryItemRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *InventoryItemRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *InventoryItemRequest) UnsetRole() {
	o.Role.Unset()
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil || IsNil(o.Manufacturer.Get()) {
		var ret BriefManufacturerRequest
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableBriefManufacturerRequest and assigns it to the Manufacturer field.
func (o *InventoryItemRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer.Set(&v)
}

// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *InventoryItemRequest) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *InventoryItemRequest) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *InventoryItemRequest) GetPartId() string {
	if o == nil || IsNil(o.PartId) {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetPartIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartId) {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasPartId() bool {
	if o != nil && !IsNil(o.PartId) {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *InventoryItemRequest) SetPartId(v string) {
	o.PartId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InventoryItemRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InventoryItemRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *InventoryItemRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *InventoryItemRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *InventoryItemRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *InventoryItemRequest) GetDiscovered() bool {
	if o == nil || IsNil(o.Discovered) {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetDiscoveredOk() (*bool, bool) {
	if o == nil || IsNil(o.Discovered) {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasDiscovered() bool {
	if o != nil && !IsNil(o.Discovered) {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *InventoryItemRequest) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InventoryItemRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InventoryItemRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemRequest) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentType.Get()
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemRequest) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentType.Get(), o.ComponentType.IsSet()
}

// HasComponentType returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasComponentType() bool {
	if o != nil && o.ComponentType.IsSet() {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given NullableString and assigns it to the ComponentType field.
func (o *InventoryItemRequest) SetComponentType(v string) {
	o.ComponentType.Set(&v)
}

// SetComponentTypeNil sets the value for ComponentType to be an explicit nil
func (o *InventoryItemRequest) SetComponentTypeNil() {
	o.ComponentType.Set(nil)
}

// UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
func (o *InventoryItemRequest) UnsetComponentType() {
	o.ComponentType.Unset()
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemRequest) GetComponentId() int64 {
	if o == nil || IsNil(o.ComponentId.Get()) {
		var ret int64
		return ret
	}
	return *o.ComponentId.Get()
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemRequest) GetComponentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentId.Get(), o.ComponentId.IsSet()
}

// HasComponentId returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasComponentId() bool {
	if o != nil && o.ComponentId.IsSet() {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given NullableInt64 and assigns it to the ComponentId field.
func (o *InventoryItemRequest) SetComponentId(v int64) {
	o.ComponentId.Set(&v)
}

// SetComponentIdNil sets the value for ComponentId to be an explicit nil
func (o *InventoryItemRequest) SetComponentIdNil() {
	o.ComponentId.Set(nil)
}

// UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
func (o *InventoryItemRequest) UnsetComponentId() {
	o.ComponentId.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InventoryItemRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *InventoryItemRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InventoryItemRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InventoryItemRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *InventoryItemRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o InventoryItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if !IsNil(o.PartId) {
		toSerialize["part_id"] = o.PartId
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if !IsNil(o.Discovered) {
		toSerialize["discovered"] = o.Discovered
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.ComponentType.IsSet() {
		toSerialize["component_type"] = o.ComponentType.Get()
	}
	if o.ComponentId.IsSet() {
		toSerialize["component_id"] = o.ComponentId.Get()
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

func (o *InventoryItemRequest) UnmarshalJSON(data []byte) (err error) {
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

	varInventoryItemRequest := _InventoryItemRequest{}

	err = json.Unmarshal(data, &varInventoryItemRequest)

	if err != nil {
		return err
	}

	*o = InventoryItemRequest(varInventoryItemRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "role")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "part_id")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "discovered")
		delete(additionalProperties, "description")
		delete(additionalProperties, "component_type")
		delete(additionalProperties, "component_id")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryItemRequest struct {
	value *InventoryItemRequest
	isSet bool
}

func (v NullableInventoryItemRequest) Get() *InventoryItemRequest {
	return v.value
}

func (v *NullableInventoryItemRequest) Set(val *InventoryItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItemRequest(val *InventoryItemRequest) *NullableInventoryItemRequest {
	return &NullableInventoryItemRequest{value: val, isSet: true}
}

func (v NullableInventoryItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

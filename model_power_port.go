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

// checks if the PowerPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerPort{}

// PowerPort Adds support for custom fields and tags.
type PowerPort struct {
	Id      int32               `json:"id"`
	Url     string              `json:"url"`
	Display string              `json:"display"`
	Device  BriefDevice         `json:"device"`
	Module  NullableBriefModule `json:"module,omitempty"`
	Name    string              `json:"name"`
	// Physical label
	Label *string               `json:"label,omitempty"`
	Type  NullablePowerPortType `json:"type,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw NullableInt32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw NullableInt32 `json:"allocated_draw,omitempty"`
	Description   *string       `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool              `json:"mark_connected,omitempty"`
	Cable         NullableBriefCable `json:"cable,omitempty"`
	CableEnd      string             `json:"cable_end"`
	LinkPeers     []interface{}      `json:"link_peers"`
	// Return the type of the peer link terminations, or None.
	LinkPeersType               NullableString         `json:"link_peers_type,omitempty"`
	ConnectedEndpoints          []interface{}          `json:"connected_endpoints,omitempty"`
	ConnectedEndpointsType      NullableString         `json:"connected_endpoints_type,omitempty"`
	ConnectedEndpointsReachable bool                   `json:"connected_endpoints_reachable"`
	Tags                        []NestedTag            `json:"tags,omitempty"`
	CustomFields                map[string]interface{} `json:"custom_fields,omitempty"`
	Created                     NullableTime           `json:"created,omitempty"`
	LastUpdated                 NullableTime           `json:"last_updated,omitempty"`
	Occupied                    bool                   `json:"_occupied"`
	AdditionalProperties        map[string]interface{}
}

type _PowerPort PowerPort

// NewPowerPort instantiates a new PowerPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerPort(id int32, url string, display string, device BriefDevice, name string, cableEnd string, linkPeers []interface{}, connectedEndpointsReachable bool, occupied bool) *PowerPort {
	this := PowerPort{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Device = device
	this.Name = name
	this.CableEnd = cableEnd
	this.LinkPeers = linkPeers
	this.ConnectedEndpointsReachable = connectedEndpointsReachable
	this.Occupied = occupied
	return &this
}

// NewPowerPortWithDefaults instantiates a new PowerPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerPortWithDefaults() *PowerPort {
	this := PowerPort{}
	return &this
}

// GetId returns the Id field value
func (o *PowerPort) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PowerPort) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *PowerPort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PowerPort) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *PowerPort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *PowerPort) SetDisplay(v string) {
	o.Display = v
}

// GetDevice returns the Device field value
func (o *PowerPort) GetDevice() BriefDevice {
	if o == nil {
		var ret BriefDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *PowerPort) SetDevice(v BriefDevice) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetModule() BriefModule {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModule
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetModuleOk() (*BriefModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *PowerPort) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModule and assigns it to the Module field.
func (o *PowerPort) SetModule(v BriefModule) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *PowerPort) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *PowerPort) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value
func (o *PowerPort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerPort) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerPort) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPort) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerPort) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PowerPort) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetType() PowerPortType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret PowerPortType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetTypeOk() (*PowerPortType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PowerPort) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullablePowerPortType and assigns it to the Type field.
func (o *PowerPort) SetType(v PowerPortType) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *PowerPort) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PowerPort) UnsetType() {
	o.Type.Unset()
}

// GetMaximumDraw returns the MaximumDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetMaximumDraw() int32 {
	if o == nil || IsNil(o.MaximumDraw.Get()) {
		var ret int32
		return ret
	}
	return *o.MaximumDraw.Get()
}

// GetMaximumDrawOk returns a tuple with the MaximumDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetMaximumDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumDraw.Get(), o.MaximumDraw.IsSet()
}

// HasMaximumDraw returns a boolean if a field has been set.
func (o *PowerPort) HasMaximumDraw() bool {
	if o != nil && o.MaximumDraw.IsSet() {
		return true
	}

	return false
}

// SetMaximumDraw gets a reference to the given NullableInt32 and assigns it to the MaximumDraw field.
func (o *PowerPort) SetMaximumDraw(v int32) {
	o.MaximumDraw.Set(&v)
}

// SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil
func (o *PowerPort) SetMaximumDrawNil() {
	o.MaximumDraw.Set(nil)
}

// UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
func (o *PowerPort) UnsetMaximumDraw() {
	o.MaximumDraw.Unset()
}

// GetAllocatedDraw returns the AllocatedDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetAllocatedDraw() int32 {
	if o == nil || IsNil(o.AllocatedDraw.Get()) {
		var ret int32
		return ret
	}
	return *o.AllocatedDraw.Get()
}

// GetAllocatedDrawOk returns a tuple with the AllocatedDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetAllocatedDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllocatedDraw.Get(), o.AllocatedDraw.IsSet()
}

// HasAllocatedDraw returns a boolean if a field has been set.
func (o *PowerPort) HasAllocatedDraw() bool {
	if o != nil && o.AllocatedDraw.IsSet() {
		return true
	}

	return false
}

// SetAllocatedDraw gets a reference to the given NullableInt32 and assigns it to the AllocatedDraw field.
func (o *PowerPort) SetAllocatedDraw(v int32) {
	o.AllocatedDraw.Set(&v)
}

// SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil
func (o *PowerPort) SetAllocatedDrawNil() {
	o.AllocatedDraw.Set(nil)
}

// UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
func (o *PowerPort) UnsetAllocatedDraw() {
	o.AllocatedDraw.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PowerPort) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPort) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PowerPort) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PowerPort) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PowerPort) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPort) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PowerPort) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PowerPort) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetCable returns the Cable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetCable() BriefCable {
	if o == nil || IsNil(o.Cable.Get()) {
		var ret BriefCable
		return ret
	}
	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetCableOk() (*BriefCable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// HasCable returns a boolean if a field has been set.
func (o *PowerPort) HasCable() bool {
	if o != nil && o.Cable.IsSet() {
		return true
	}

	return false
}

// SetCable gets a reference to the given NullableBriefCable and assigns it to the Cable field.
func (o *PowerPort) SetCable(v BriefCable) {
	o.Cable.Set(&v)
}

// SetCableNil sets the value for Cable to be an explicit nil
func (o *PowerPort) SetCableNil() {
	o.Cable.Set(nil)
}

// UnsetCable ensures that no value is present for Cable, not even an explicit nil
func (o *PowerPort) UnsetCable() {
	o.Cable.Unset()
}

// GetCableEnd returns the CableEnd field value
func (o *PowerPort) GetCableEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CableEnd
}

// GetCableEndOk returns a tuple with the CableEnd field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetCableEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CableEnd, true
}

// SetCableEnd sets field value
func (o *PowerPort) SetCableEnd(v string) {
	o.CableEnd = v
}

// GetLinkPeers returns the LinkPeers field value
func (o *PowerPort) GetLinkPeers() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.LinkPeers
}

// GetLinkPeersOk returns a tuple with the LinkPeers field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetLinkPeersOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkPeers, true
}

// SetLinkPeers sets field value
func (o *PowerPort) SetLinkPeers(v []interface{}) {
	o.LinkPeers = v
}

// GetLinkPeersType returns the LinkPeersType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetLinkPeersType() string {
	if o == nil || IsNil(o.LinkPeersType.Get()) {
		var ret string
		return ret
	}
	return *o.LinkPeersType.Get()
}

// GetLinkPeersTypeOk returns a tuple with the LinkPeersType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetLinkPeersTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkPeersType.Get(), o.LinkPeersType.IsSet()
}

// HasLinkPeersType returns a boolean if a field has been set.
func (o *PowerPort) HasLinkPeersType() bool {
	if o != nil && o.LinkPeersType.IsSet() {
		return true
	}

	return false
}

// SetLinkPeersType gets a reference to the given NullableString and assigns it to the LinkPeersType field.
func (o *PowerPort) SetLinkPeersType(v string) {
	o.LinkPeersType.Set(&v)
}

// SetLinkPeersTypeNil sets the value for LinkPeersType to be an explicit nil
func (o *PowerPort) SetLinkPeersTypeNil() {
	o.LinkPeersType.Set(nil)
}

// UnsetLinkPeersType ensures that no value is present for LinkPeersType, not even an explicit nil
func (o *PowerPort) UnsetLinkPeersType() {
	o.LinkPeersType.Unset()
}

// GetConnectedEndpoints returns the ConnectedEndpoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetConnectedEndpoints() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.ConnectedEndpoints
}

// GetConnectedEndpointsOk returns a tuple with the ConnectedEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetConnectedEndpointsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ConnectedEndpoints) {
		return nil, false
	}
	return o.ConnectedEndpoints, true
}

// HasConnectedEndpoints returns a boolean if a field has been set.
func (o *PowerPort) HasConnectedEndpoints() bool {
	if o != nil && !IsNil(o.ConnectedEndpoints) {
		return true
	}

	return false
}

// SetConnectedEndpoints gets a reference to the given []interface{} and assigns it to the ConnectedEndpoints field.
func (o *PowerPort) SetConnectedEndpoints(v []interface{}) {
	o.ConnectedEndpoints = v
}

// GetConnectedEndpointsType returns the ConnectedEndpointsType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetConnectedEndpointsType() string {
	if o == nil || IsNil(o.ConnectedEndpointsType.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectedEndpointsType.Get()
}

// GetConnectedEndpointsTypeOk returns a tuple with the ConnectedEndpointsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetConnectedEndpointsTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointsType.Get(), o.ConnectedEndpointsType.IsSet()
}

// HasConnectedEndpointsType returns a boolean if a field has been set.
func (o *PowerPort) HasConnectedEndpointsType() bool {
	if o != nil && o.ConnectedEndpointsType.IsSet() {
		return true
	}

	return false
}

// SetConnectedEndpointsType gets a reference to the given NullableString and assigns it to the ConnectedEndpointsType field.
func (o *PowerPort) SetConnectedEndpointsType(v string) {
	o.ConnectedEndpointsType.Set(&v)
}

// SetConnectedEndpointsTypeNil sets the value for ConnectedEndpointsType to be an explicit nil
func (o *PowerPort) SetConnectedEndpointsTypeNil() {
	o.ConnectedEndpointsType.Set(nil)
}

// UnsetConnectedEndpointsType ensures that no value is present for ConnectedEndpointsType, not even an explicit nil
func (o *PowerPort) UnsetConnectedEndpointsType() {
	o.ConnectedEndpointsType.Unset()
}

// GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable field value
func (o *PowerPort) GetConnectedEndpointsReachable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ConnectedEndpointsReachable
}

// GetConnectedEndpointsReachableOk returns a tuple with the ConnectedEndpointsReachable field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetConnectedEndpointsReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectedEndpointsReachable, true
}

// SetConnectedEndpointsReachable sets field value
func (o *PowerPort) SetConnectedEndpointsReachable(v bool) {
	o.ConnectedEndpointsReachable = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PowerPort) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPort) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PowerPort) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *PowerPort) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PowerPort) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PowerPort) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PowerPort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *PowerPort) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *PowerPort) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *PowerPort) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *PowerPort) UnsetCreated() {
	o.Created.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPort) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPort) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PowerPort) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *PowerPort) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *PowerPort) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *PowerPort) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

// GetOccupied returns the Occupied field value
func (o *PowerPort) GetOccupied() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value
// and a boolean to check if the value has been set.
func (o *PowerPort) GetOccupiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occupied, true
}

// SetOccupied sets field value
func (o *PowerPort) SetOccupied(v bool) {
	o.Occupied = v
}

func (o PowerPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.MaximumDraw.IsSet() {
		toSerialize["maximum_draw"] = o.MaximumDraw.Get()
	}
	if o.AllocatedDraw.IsSet() {
		toSerialize["allocated_draw"] = o.AllocatedDraw.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	if o.Cable.IsSet() {
		toSerialize["cable"] = o.Cable.Get()
	}
	toSerialize["cable_end"] = o.CableEnd
	toSerialize["link_peers"] = o.LinkPeers
	if o.LinkPeersType.IsSet() {
		toSerialize["link_peers_type"] = o.LinkPeersType.Get()
	}
	if o.ConnectedEndpoints != nil {
		toSerialize["connected_endpoints"] = o.ConnectedEndpoints
	}
	if o.ConnectedEndpointsType.IsSet() {
		toSerialize["connected_endpoints_type"] = o.ConnectedEndpointsType.Get()
	}
	toSerialize["connected_endpoints_reachable"] = o.ConnectedEndpointsReachable
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
	toSerialize["_occupied"] = o.Occupied

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PowerPort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"device",
		"name",
		"cable_end",
		"link_peers",
		"connected_endpoints_reachable",
		"_occupied",
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

	varPowerPort := _PowerPort{}

	err = json.Unmarshal(data, &varPowerPort)

	if err != nil {
		return err
	}

	*o = PowerPort(varPowerPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "maximum_draw")
		delete(additionalProperties, "allocated_draw")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "cable")
		delete(additionalProperties, "cable_end")
		delete(additionalProperties, "link_peers")
		delete(additionalProperties, "link_peers_type")
		delete(additionalProperties, "connected_endpoints")
		delete(additionalProperties, "connected_endpoints_type")
		delete(additionalProperties, "connected_endpoints_reachable")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "_occupied")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerPort struct {
	value *PowerPort
	isSet bool
}

func (v NullablePowerPort) Get() *PowerPort {
	return v.value
}

func (v *NullablePowerPort) Set(val *PowerPort) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerPort) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerPort(val *PowerPort) *NullablePowerPort {
	return &NullablePowerPort{value: val, isSet: true}
}

func (v NullablePowerPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

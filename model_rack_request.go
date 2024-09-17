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

// checks if the RackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackRequest{}

// RackRequest Adds support for custom fields and tags.
type RackRequest struct {
	Name       *string                           `json:"name,omitempty"`
	FacilityId NullableString                    `json:"facility_id,omitempty"`
	Site       BriefSiteRequest                  `json:"site"`
	Location   NullableBriefLocationRequest      `json:"location,omitempty"`
	Tenant     NullableBriefTenantRequest        `json:"tenant,omitempty"`
	Status     *PatchedWritableRackRequestStatus `json:"status,omitempty"`
	Role       NullableBriefRackRoleRequest      `json:"role,omitempty"`
	Serial     *string                           `json:"serial,omitempty"`
	// A unique tag used to identify this rack
	AssetTag NullableString          `json:"asset_tag,omitempty"`
	Type     NullableRackRequestType `json:"type,omitempty"`
	Width    *RackWidthValue         `json:"width,omitempty"`
	// Height in rack units
	UHeight *int32 `json:"u_height,omitempty"`
	// Starting unit for rack
	StartingUnit *int32          `json:"starting_unit,omitempty"`
	Weight       NullableFloat64 `json:"weight,omitempty"`
	// Maximum load capacity for the rack
	MaxWeight  NullableInt32                       `json:"max_weight,omitempty"`
	WeightUnit NullableDeviceTypeRequestWeightUnit `json:"weight_unit,omitempty"`
	// Units are numbered top-to-bottom
	DescUnits *bool `json:"desc_units,omitempty"`
	// Outer dimension of rack (width)
	OuterWidth NullableInt32 `json:"outer_width,omitempty"`
	// Outer dimension of rack (depth)
	OuterDepth NullableInt32                `json:"outer_depth,omitempty"`
	OuterUnit  NullableRackRequestOuterUnit `json:"outer_unit,omitempty"`
	// Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	MountingDepth        NullableInt32          `json:"mounting_depth,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackRequest RackRequest

// NewRackRequest instantiates a new RackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackRequest(site BriefSiteRequest) *RackRequest {
	this := RackRequest{}
	this.Site = site
	return &this
}

// NewRackRequestWithDefaults instantiates a new RackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackRequestWithDefaults() *RackRequest {
	this := RackRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RackRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RackRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RackRequest) SetName(v string) {
	o.Name = &v
}

// GetFacilityId returns the FacilityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetFacilityId() string {
	if o == nil || IsNil(o.FacilityId.Get()) {
		var ret string
		return ret
	}
	return *o.FacilityId.Get()
}

// GetFacilityIdOk returns a tuple with the FacilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetFacilityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacilityId.Get(), o.FacilityId.IsSet()
}

// HasFacilityId returns a boolean if a field has been set.
func (o *RackRequest) HasFacilityId() bool {
	if o != nil && o.FacilityId.IsSet() {
		return true
	}

	return false
}

// SetFacilityId gets a reference to the given NullableString and assigns it to the FacilityId field.
func (o *RackRequest) SetFacilityId(v string) {
	o.FacilityId.Set(&v)
}

// SetFacilityIdNil sets the value for FacilityId to be an explicit nil
func (o *RackRequest) SetFacilityIdNil() {
	o.FacilityId.Set(nil)
}

// UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
func (o *RackRequest) UnsetFacilityId() {
	o.FacilityId.Unset()
}

// GetSite returns the Site field value
func (o *RackRequest) GetSite() BriefSiteRequest {
	if o == nil {
		var ret BriefSiteRequest
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *RackRequest) GetSiteOk() (*BriefSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *RackRequest) SetSite(v BriefSiteRequest) {
	o.Site = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetLocation() BriefLocationRequest {
	if o == nil || IsNil(o.Location.Get()) {
		var ret BriefLocationRequest
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetLocationOk() (*BriefLocationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *RackRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableBriefLocationRequest and assigns it to the Location field.
func (o *RackRequest) SetLocation(v BriefLocationRequest) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *RackRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *RackRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *RackRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *RackRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *RackRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *RackRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RackRequest) GetStatus() PatchedWritableRackRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableRackRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetStatusOk() (*PatchedWritableRackRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RackRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableRackRequestStatus and assigns it to the Status field.
func (o *RackRequest) SetStatus(v PatchedWritableRackRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetRole() BriefRackRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefRackRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetRoleOk() (*BriefRackRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *RackRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefRackRoleRequest and assigns it to the Role field.
func (o *RackRequest) SetRole(v BriefRackRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *RackRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *RackRequest) UnsetRole() {
	o.Role.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *RackRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *RackRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *RackRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *RackRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *RackRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *RackRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *RackRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetType() RackRequestType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret RackRequestType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetTypeOk() (*RackRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *RackRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableRackRequestType and assigns it to the Type field.
func (o *RackRequest) SetType(v RackRequestType) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *RackRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *RackRequest) UnsetType() {
	o.Type.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *RackRequest) GetWidth() RackWidthValue {
	if o == nil || IsNil(o.Width) {
		var ret RackWidthValue
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetWidthOk() (*RackWidthValue, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *RackRequest) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given RackWidthValue and assigns it to the Width field.
func (o *RackRequest) SetWidth(v RackWidthValue) {
	o.Width = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *RackRequest) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *RackRequest) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *RackRequest) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetStartingUnit returns the StartingUnit field value if set, zero value otherwise.
func (o *RackRequest) GetStartingUnit() int32 {
	if o == nil || IsNil(o.StartingUnit) {
		var ret int32
		return ret
	}
	return *o.StartingUnit
}

// GetStartingUnitOk returns a tuple with the StartingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetStartingUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.StartingUnit) {
		return nil, false
	}
	return o.StartingUnit, true
}

// HasStartingUnit returns a boolean if a field has been set.
func (o *RackRequest) HasStartingUnit() bool {
	if o != nil && !IsNil(o.StartingUnit) {
		return true
	}

	return false
}

// SetStartingUnit gets a reference to the given int32 and assigns it to the StartingUnit field.
func (o *RackRequest) SetStartingUnit(v int32) {
	o.StartingUnit = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *RackRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *RackRequest) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *RackRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *RackRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetMaxWeight() int32 {
	if o == nil || IsNil(o.MaxWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxWeight.Get()
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetMaxWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxWeight.Get(), o.MaxWeight.IsSet()
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *RackRequest) HasMaxWeight() bool {
	if o != nil && o.MaxWeight.IsSet() {
		return true
	}

	return false
}

// SetMaxWeight gets a reference to the given NullableInt32 and assigns it to the MaxWeight field.
func (o *RackRequest) SetMaxWeight(v int32) {
	o.MaxWeight.Set(&v)
}

// SetMaxWeightNil sets the value for MaxWeight to be an explicit nil
func (o *RackRequest) SetMaxWeightNil() {
	o.MaxWeight.Set(nil)
}

// UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
func (o *RackRequest) UnsetMaxWeight() {
	o.MaxWeight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetWeightUnit() DeviceTypeRequestWeightUnit {
	if o == nil || IsNil(o.WeightUnit.Get()) {
		var ret DeviceTypeRequestWeightUnit
		return ret
	}
	return *o.WeightUnit.Get()
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetWeightUnitOk() (*DeviceTypeRequestWeightUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightUnit.Get(), o.WeightUnit.IsSet()
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *RackRequest) HasWeightUnit() bool {
	if o != nil && o.WeightUnit.IsSet() {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given NullableDeviceTypeRequestWeightUnit and assigns it to the WeightUnit field.
func (o *RackRequest) SetWeightUnit(v DeviceTypeRequestWeightUnit) {
	o.WeightUnit.Set(&v)
}

// SetWeightUnitNil sets the value for WeightUnit to be an explicit nil
func (o *RackRequest) SetWeightUnitNil() {
	o.WeightUnit.Set(nil)
}

// UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
func (o *RackRequest) UnsetWeightUnit() {
	o.WeightUnit.Unset()
}

// GetDescUnits returns the DescUnits field value if set, zero value otherwise.
func (o *RackRequest) GetDescUnits() bool {
	if o == nil || IsNil(o.DescUnits) {
		var ret bool
		return ret
	}
	return *o.DescUnits
}

// GetDescUnitsOk returns a tuple with the DescUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetDescUnitsOk() (*bool, bool) {
	if o == nil || IsNil(o.DescUnits) {
		return nil, false
	}
	return o.DescUnits, true
}

// HasDescUnits returns a boolean if a field has been set.
func (o *RackRequest) HasDescUnits() bool {
	if o != nil && !IsNil(o.DescUnits) {
		return true
	}

	return false
}

// SetDescUnits gets a reference to the given bool and assigns it to the DescUnits field.
func (o *RackRequest) SetDescUnits(v bool) {
	o.DescUnits = &v
}

// GetOuterWidth returns the OuterWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetOuterWidth() int32 {
	if o == nil || IsNil(o.OuterWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterWidth.Get()
}

// GetOuterWidthOk returns a tuple with the OuterWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetOuterWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterWidth.Get(), o.OuterWidth.IsSet()
}

// HasOuterWidth returns a boolean if a field has been set.
func (o *RackRequest) HasOuterWidth() bool {
	if o != nil && o.OuterWidth.IsSet() {
		return true
	}

	return false
}

// SetOuterWidth gets a reference to the given NullableInt32 and assigns it to the OuterWidth field.
func (o *RackRequest) SetOuterWidth(v int32) {
	o.OuterWidth.Set(&v)
}

// SetOuterWidthNil sets the value for OuterWidth to be an explicit nil
func (o *RackRequest) SetOuterWidthNil() {
	o.OuterWidth.Set(nil)
}

// UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
func (o *RackRequest) UnsetOuterWidth() {
	o.OuterWidth.Unset()
}

// GetOuterDepth returns the OuterDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetOuterDepth() int32 {
	if o == nil || IsNil(o.OuterDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterDepth.Get()
}

// GetOuterDepthOk returns a tuple with the OuterDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetOuterDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterDepth.Get(), o.OuterDepth.IsSet()
}

// HasOuterDepth returns a boolean if a field has been set.
func (o *RackRequest) HasOuterDepth() bool {
	if o != nil && o.OuterDepth.IsSet() {
		return true
	}

	return false
}

// SetOuterDepth gets a reference to the given NullableInt32 and assigns it to the OuterDepth field.
func (o *RackRequest) SetOuterDepth(v int32) {
	o.OuterDepth.Set(&v)
}

// SetOuterDepthNil sets the value for OuterDepth to be an explicit nil
func (o *RackRequest) SetOuterDepthNil() {
	o.OuterDepth.Set(nil)
}

// UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
func (o *RackRequest) UnsetOuterDepth() {
	o.OuterDepth.Unset()
}

// GetOuterUnit returns the OuterUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetOuterUnit() RackRequestOuterUnit {
	if o == nil || IsNil(o.OuterUnit.Get()) {
		var ret RackRequestOuterUnit
		return ret
	}
	return *o.OuterUnit.Get()
}

// GetOuterUnitOk returns a tuple with the OuterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetOuterUnitOk() (*RackRequestOuterUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterUnit.Get(), o.OuterUnit.IsSet()
}

// HasOuterUnit returns a boolean if a field has been set.
func (o *RackRequest) HasOuterUnit() bool {
	if o != nil && o.OuterUnit.IsSet() {
		return true
	}

	return false
}

// SetOuterUnit gets a reference to the given NullableRackRequestOuterUnit and assigns it to the OuterUnit field.
func (o *RackRequest) SetOuterUnit(v RackRequestOuterUnit) {
	o.OuterUnit.Set(&v)
}

// SetOuterUnitNil sets the value for OuterUnit to be an explicit nil
func (o *RackRequest) SetOuterUnitNil() {
	o.OuterUnit.Set(nil)
}

// UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
func (o *RackRequest) UnsetOuterUnit() {
	o.OuterUnit.Unset()
}

// GetMountingDepth returns the MountingDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackRequest) GetMountingDepth() int32 {
	if o == nil || IsNil(o.MountingDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.MountingDepth.Get()
}

// GetMountingDepthOk returns a tuple with the MountingDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRequest) GetMountingDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountingDepth.Get(), o.MountingDepth.IsSet()
}

// HasMountingDepth returns a boolean if a field has been set.
func (o *RackRequest) HasMountingDepth() bool {
	if o != nil && o.MountingDepth.IsSet() {
		return true
	}

	return false
}

// SetMountingDepth gets a reference to the given NullableInt32 and assigns it to the MountingDepth field.
func (o *RackRequest) SetMountingDepth(v int32) {
	o.MountingDepth.Set(&v)
}

// SetMountingDepthNil sets the value for MountingDepth to be an explicit nil
func (o *RackRequest) SetMountingDepthNil() {
	o.MountingDepth.Set(nil)
}

// UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
func (o *RackRequest) UnsetMountingDepth() {
	o.MountingDepth.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RackRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RackRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RackRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *RackRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *RackRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *RackRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RackRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RackRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *RackRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RackRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RackRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *RackRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o RackRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.FacilityId.IsSet() {
		toSerialize["facility_id"] = o.FacilityId.Get()
	}
	toSerialize["site"] = o.Site
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.StartingUnit) {
		toSerialize["starting_unit"] = o.StartingUnit
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.MaxWeight.IsSet() {
		toSerialize["max_weight"] = o.MaxWeight.Get()
	}
	if o.WeightUnit.IsSet() {
		toSerialize["weight_unit"] = o.WeightUnit.Get()
	}
	if !IsNil(o.DescUnits) {
		toSerialize["desc_units"] = o.DescUnits
	}
	if o.OuterWidth.IsSet() {
		toSerialize["outer_width"] = o.OuterWidth.Get()
	}
	if o.OuterDepth.IsSet() {
		toSerialize["outer_depth"] = o.OuterDepth.Get()
	}
	if o.OuterUnit.IsSet() {
		toSerialize["outer_unit"] = o.OuterUnit.Get()
	}
	if o.MountingDepth.IsSet() {
		toSerialize["mounting_depth"] = o.MountingDepth.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *RackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"site",
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

	varRackRequest := _RackRequest{}

	err = json.Unmarshal(data, &varRackRequest)

	if err != nil {
		return err
	}

	*o = RackRequest(varRackRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "facility_id")
		delete(additionalProperties, "site")
		delete(additionalProperties, "location")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "type")
		delete(additionalProperties, "width")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "starting_unit")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "max_weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "desc_units")
		delete(additionalProperties, "outer_width")
		delete(additionalProperties, "outer_depth")
		delete(additionalProperties, "outer_unit")
		delete(additionalProperties, "mounting_depth")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackRequest struct {
	value *RackRequest
	isSet bool
}

func (v NullableRackRequest) Get() *RackRequest {
	return v.value
}

func (v *NullableRackRequest) Set(val *RackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackRequest(val *RackRequest) *NullableRackRequest {
	return &NullableRackRequest{value: val, isSet: true}
}

func (v NullableRackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

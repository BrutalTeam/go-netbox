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

// checks if the WritableVirtualMachineWithConfigContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableVirtualMachineWithConfigContextRequest{}

// WritableVirtualMachineWithConfigContextRequest Adds support for custom fields and tags.
type WritableVirtualMachineWithConfigContextRequest struct {
	Name           *string                            `json:"name,omitempty"`
	Status         *ModuleStatusValue                 `json:"status,omitempty"`
	Site           NullableBriefSiteRequest           `json:"site,omitempty"`
	Cluster        NullableBriefClusterRequest        `json:"cluster,omitempty"`
	Device         NullableBriefDeviceRequest         `json:"device,omitempty"`
	Role           NullableBriefDeviceRoleRequest     `json:"role,omitempty"`
	Tenant         NullableBriefTenantRequest         `json:"tenant,omitempty"`
	Platform       NullableBriefPlatformRequest       `json:"platform,omitempty"`
	PrimaryIp4     NullableBriefIPAddressRequest      `json:"primary_ip4,omitempty"`
	PrimaryIp6     NullableBriefIPAddressRequest      `json:"primary_ip6,omitempty"`
	Vcpus          NullableFloat64                    `json:"vcpus,omitempty"`
	Memory         NullableInt32                      `json:"memory,omitempty"`
	Disk           NullableInt32                      `json:"disk,omitempty"`
	Description    *string                            `json:"description,omitempty"`
	Comments       *string                            `json:"comments,omitempty"`
	ConfigTemplate NullableBriefConfigTemplateRequest `json:"config_template,omitempty"`
	// Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData     interface{}            `json:"local_context_data,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableVirtualMachineWithConfigContextRequest WritableVirtualMachineWithConfigContextRequest

// NewWritableVirtualMachineWithConfigContextRequest instantiates a new WritableVirtualMachineWithConfigContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVirtualMachineWithConfigContextRequest() *WritableVirtualMachineWithConfigContextRequest {
	this := WritableVirtualMachineWithConfigContextRequest{}
	return &this
}

// NewWritableVirtualMachineWithConfigContextRequestWithDefaults instantiates a new WritableVirtualMachineWithConfigContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVirtualMachineWithConfigContextRequestWithDefaults() *WritableVirtualMachineWithConfigContextRequest {
	this := WritableVirtualMachineWithConfigContextRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContextRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContextRequest) GetStatus() ModuleStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret ModuleStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) GetStatusOk() (*ModuleStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ModuleStatusValue and assigns it to the Status field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetStatus(v ModuleStatusValue) {
	o.Status = &v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetSite() BriefSiteRequest {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSiteRequest
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetSiteOk() (*BriefSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSiteRequest and assigns it to the Site field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetSite(v BriefSiteRequest) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetSite() {
	o.Site.Unset()
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetCluster() BriefClusterRequest {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret BriefClusterRequest
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetClusterOk() (*BriefClusterRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableBriefClusterRequest and assigns it to the Cluster field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetCluster(v BriefClusterRequest) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetCluster() {
	o.Cluster.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetDevice() BriefDeviceRequest {
	if o == nil || IsNil(o.Device.Get()) {
		var ret BriefDeviceRequest
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableBriefDeviceRequest and assigns it to the Device field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetDevice(v BriefDeviceRequest) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetRole() BriefDeviceRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefDeviceRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetRoleOk() (*BriefDeviceRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefDeviceRoleRequest and assigns it to the Role field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetRole(v BriefDeviceRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetRole() {
	o.Role.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetPlatform() BriefPlatformRequest {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret BriefPlatformRequest
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetPlatformOk() (*BriefPlatformRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableBriefPlatformRequest and assigns it to the Platform field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetPlatform(v BriefPlatformRequest) {
	o.Platform.Set(&v)
}

// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetPlatform() {
	o.Platform.Unset()
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetPrimaryIp4() BriefIPAddressRequest {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret BriefIPAddressRequest
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetPrimaryIp4Ok() (*BriefIPAddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableBriefIPAddressRequest and assigns it to the PrimaryIp4 field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetPrimaryIp4(v BriefIPAddressRequest) {
	o.PrimaryIp4.Set(&v)
}

// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetPrimaryIp6() BriefIPAddressRequest {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret BriefIPAddressRequest
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetPrimaryIp6Ok() (*BriefIPAddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableBriefIPAddressRequest and assigns it to the PrimaryIp6 field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetPrimaryIp6(v BriefIPAddressRequest) {
	o.PrimaryIp6.Set(&v)
}

// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetVcpus() float64 {
	if o == nil || IsNil(o.Vcpus.Get()) {
		var ret float64
		return ret
	}
	return *o.Vcpus.Get()
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetVcpusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vcpus.Get(), o.Vcpus.IsSet()
}

// HasVcpus returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasVcpus() bool {
	if o != nil && o.Vcpus.IsSet() {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given NullableFloat64 and assigns it to the Vcpus field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetVcpus(v float64) {
	o.Vcpus.Set(&v)
}

// SetVcpusNil sets the value for Vcpus to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetVcpusNil() {
	o.Vcpus.Set(nil)
}

// UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetVcpus() {
	o.Vcpus.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetMemory() {
	o.Memory.Unset()
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetDisk() int32 {
	if o == nil || IsNil(o.Disk.Get()) {
		var ret int32
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetDiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableInt32 and assigns it to the Disk field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetDisk(v int32) {
	o.Disk.Set(&v)
}

// SetDiskNil sets the value for Disk to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetDisk() {
	o.Disk.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContextRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContextRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetComments(v string) {
	o.Comments = &v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetConfigTemplate() BriefConfigTemplateRequest {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret BriefConfigTemplateRequest
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetConfigTemplateOk() (*BriefConfigTemplateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableBriefConfigTemplateRequest and assigns it to the ConfigTemplate field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetConfigTemplate(v BriefConfigTemplateRequest) {
	o.ConfigTemplate.Set(&v)
}

// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContextRequest) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetLocalContextData returns the LocalContextData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContextRequest) GetLocalContextData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LocalContextData
}

// GetLocalContextDataOk returns a tuple with the LocalContextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContextRequest) GetLocalContextDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LocalContextData) {
		return nil, false
	}
	return &o.LocalContextData, true
}

// HasLocalContextData returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasLocalContextData() bool {
	if o != nil && !IsNil(o.LocalContextData) {
		return true
	}

	return false
}

// SetLocalContextData gets a reference to the given interface{} and assigns it to the LocalContextData field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetLocalContextData(v interface{}) {
	o.LocalContextData = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContextRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContextRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContextRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableVirtualMachineWithConfigContextRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableVirtualMachineWithConfigContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableVirtualMachineWithConfigContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Cluster.IsSet() {
		toSerialize["cluster"] = o.Cluster.Get()
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	if o.Vcpus.IsSet() {
		toSerialize["vcpus"] = o.Vcpus.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Disk.IsSet() {
		toSerialize["disk"] = o.Disk.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.ConfigTemplate.IsSet() {
		toSerialize["config_template"] = o.ConfigTemplate.Get()
	}
	if o.LocalContextData != nil {
		toSerialize["local_context_data"] = o.LocalContextData
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

func (o *WritableVirtualMachineWithConfigContextRequest) UnmarshalJSON(data []byte) (err error) {
	varWritableVirtualMachineWithConfigContextRequest := _WritableVirtualMachineWithConfigContextRequest{}

	err = json.Unmarshal(data, &varWritableVirtualMachineWithConfigContextRequest)

	if err != nil {
		return err
	}

	*o = WritableVirtualMachineWithConfigContextRequest(varWritableVirtualMachineWithConfigContextRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "site")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "device")
		delete(additionalProperties, "role")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "primary_ip4")
		delete(additionalProperties, "primary_ip6")
		delete(additionalProperties, "vcpus")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "disk")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "local_context_data")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableVirtualMachineWithConfigContextRequest struct {
	value *WritableVirtualMachineWithConfigContextRequest
	isSet bool
}

func (v NullableWritableVirtualMachineWithConfigContextRequest) Get() *WritableVirtualMachineWithConfigContextRequest {
	return v.value
}

func (v *NullableWritableVirtualMachineWithConfigContextRequest) Set(val *WritableVirtualMachineWithConfigContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVirtualMachineWithConfigContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVirtualMachineWithConfigContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVirtualMachineWithConfigContextRequest(val *WritableVirtualMachineWithConfigContextRequest) *NullableWritableVirtualMachineWithConfigContextRequest {
	return &NullableWritableVirtualMachineWithConfigContextRequest{value: val, isSet: true}
}

func (v NullableWritableVirtualMachineWithConfigContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVirtualMachineWithConfigContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

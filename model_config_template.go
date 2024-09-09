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

// checks if the ConfigTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigTemplate{}

// ConfigTemplate Introduces support for Tag assignment. Adds `tags` serialization, and handles tag assignment on create() and update().
type ConfigTemplate struct {
	Id          int32   `json:"id"`
	Url         string  `json:"url"`
	Display     string  `json:"display"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Any <a href=\"https://jinja.palletsprojects.com/en/3.1.x/api/#jinja2.Environment\">additional parameters</a> to pass when constructing the Jinja2 environment.
	EnvironmentParams interface{} `json:"environment_params,omitempty"`
	// Jinja2 template code.
	TemplateCode string           `json:"template_code"`
	DataSource   *BriefDataSource `json:"data_source,omitempty"`
	// Path to remote file (relative to data source root)
	DataPath             string         `json:"data_path"`
	DataFile             *BriefDataFile `json:"data_file,omitempty"`
	DataSynced           NullableTime   `json:"data_synced"`
	Tags                 []NestedTag    `json:"tags,omitempty"`
	Created              NullableTime   `json:"created"`
	LastUpdated          NullableTime   `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _ConfigTemplate ConfigTemplate

// NewConfigTemplate instantiates a new ConfigTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigTemplate(id int32, url string, display string, name string, templateCode string, dataPath string, dataSynced NullableTime, created NullableTime, lastUpdated NullableTime) *ConfigTemplate {
	this := ConfigTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.TemplateCode = templateCode
	this.DataPath = dataPath
	this.DataSynced = dataSynced
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewConfigTemplateWithDefaults instantiates a new ConfigTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigTemplateWithDefaults() *ConfigTemplate {
	this := ConfigTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ConfigTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ConfigTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ConfigTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConfigTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *ConfigTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironmentParams returns the EnvironmentParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigTemplate) GetEnvironmentParams() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnvironmentParams
}

// GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigTemplate) GetEnvironmentParamsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnvironmentParams) {
		return nil, false
	}
	return &o.EnvironmentParams, true
}

// HasEnvironmentParams returns a boolean if a field has been set.
func (o *ConfigTemplate) HasEnvironmentParams() bool {
	if o != nil && !IsNil(o.EnvironmentParams) {
		return true
	}

	return false
}

// SetEnvironmentParams gets a reference to the given interface{} and assigns it to the EnvironmentParams field.
func (o *ConfigTemplate) SetEnvironmentParams(v interface{}) {
	o.EnvironmentParams = v
}

// GetTemplateCode returns the TemplateCode field value
func (o *ConfigTemplate) GetTemplateCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetTemplateCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateCode, true
}

// SetTemplateCode sets field value
func (o *ConfigTemplate) SetTemplateCode(v string) {
	o.TemplateCode = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ConfigTemplate) GetDataSource() BriefDataSource {
	if o == nil || IsNil(o.DataSource) {
		var ret BriefDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetDataSourceOk() (*BriefDataSource, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ConfigTemplate) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given BriefDataSource and assigns it to the DataSource field.
func (o *ConfigTemplate) SetDataSource(v BriefDataSource) {
	o.DataSource = &v
}

// GetDataPath returns the DataPath field value
func (o *ConfigTemplate) GetDataPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataPath
}

// GetDataPathOk returns a tuple with the DataPath field value
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetDataPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataPath, true
}

// SetDataPath sets field value
func (o *ConfigTemplate) SetDataPath(v string) {
	o.DataPath = v
}

// GetDataFile returns the DataFile field value if set, zero value otherwise.
func (o *ConfigTemplate) GetDataFile() BriefDataFile {
	if o == nil || IsNil(o.DataFile) {
		var ret BriefDataFile
		return ret
	}
	return *o.DataFile
}

// GetDataFileOk returns a tuple with the DataFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetDataFileOk() (*BriefDataFile, bool) {
	if o == nil || IsNil(o.DataFile) {
		return nil, false
	}
	return o.DataFile, true
}

// HasDataFile returns a boolean if a field has been set.
func (o *ConfigTemplate) HasDataFile() bool {
	if o != nil && !IsNil(o.DataFile) {
		return true
	}

	return false
}

// SetDataFile gets a reference to the given BriefDataFile and assigns it to the DataFile field.
func (o *ConfigTemplate) SetDataFile(v BriefDataFile) {
	o.DataFile = &v
}

// GetDataSynced returns the DataSynced field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConfigTemplate) GetDataSynced() time.Time {
	if o == nil || o.DataSynced.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DataSynced.Get()
}

// GetDataSyncedOk returns a tuple with the DataSynced field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigTemplate) GetDataSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSynced.Get(), o.DataSynced.IsSet()
}

// SetDataSynced sets field value
func (o *ConfigTemplate) SetDataSynced(v time.Time) {
	o.DataSynced.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigTemplate) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigTemplate) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigTemplate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *ConfigTemplate) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConfigTemplate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ConfigTemplate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConfigTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ConfigTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o ConfigTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.EnvironmentParams != nil {
		toSerialize["environment_params"] = o.EnvironmentParams
	}
	toSerialize["template_code"] = o.TemplateCode
	if !IsNil(o.DataSource) {
		toSerialize["data_source"] = o.DataSource
	}
	toSerialize["data_path"] = o.DataPath
	if !IsNil(o.DataFile) {
		toSerialize["data_file"] = o.DataFile
	}
	toSerialize["data_synced"] = o.DataSynced.Get()
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"template_code",
		"data_path",
		"data_synced",
		"created",
		"last_updated",
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

	varConfigTemplate := _ConfigTemplate{}

	err = json.Unmarshal(data, &varConfigTemplate)

	if err != nil {
		return err
	}

	*o = ConfigTemplate(varConfigTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "environment_params")
		delete(additionalProperties, "template_code")
		delete(additionalProperties, "data_source")
		delete(additionalProperties, "data_path")
		delete(additionalProperties, "data_file")
		delete(additionalProperties, "data_synced")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigTemplate struct {
	value *ConfigTemplate
	isSet bool
}

func (v NullableConfigTemplate) Get() *ConfigTemplate {
	return v.value
}

func (v *NullableConfigTemplate) Set(val *ConfigTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigTemplate(val *ConfigTemplate) *NullableConfigTemplate {
	return &NullableConfigTemplate{value: val, isSet: true}
}

func (v NullableConfigTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

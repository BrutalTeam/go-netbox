# BriefVLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vid** | Pointer to **int32** | Numeric VLAN ID (1-4094) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefVLANRequest

`func NewBriefVLANRequest() *BriefVLANRequest`

NewBriefVLANRequest instantiates a new BriefVLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefVLANRequestWithDefaults

`func NewBriefVLANRequestWithDefaults() *BriefVLANRequest`

NewBriefVLANRequestWithDefaults instantiates a new BriefVLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVid

`func (o *BriefVLANRequest) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *BriefVLANRequest) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *BriefVLANRequest) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *BriefVLANRequest) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetName

`func (o *BriefVLANRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefVLANRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefVLANRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BriefVLANRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BriefVLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefVLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefVLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefVLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



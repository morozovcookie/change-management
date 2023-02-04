# Crq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**SelfLink**](SelfLink.md) |  | 
**Id** | **string** |  | 
**Summary** | **string** | The short-form description of change | 
**Description** | **string** | The long-form description of change | 
**Type** | **string** | The type of change | [default to "crq"]
**IsAutoClose** | **bool** | This flag indicates that change request should be closed automatically | [default to false]
**CreatedAt** | **int64** |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewCrq

`func NewCrq(links SelfLink, id string, summary string, description string, type_ string, isAutoClose bool, createdAt int64, ) *Crq`

NewCrq instantiates a new Crq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrqWithDefaults

`func NewCrqWithDefaults() *Crq`

NewCrqWithDefaults instantiates a new Crq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Crq) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Crq) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Crq) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *Crq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Crq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Crq) SetId(v string)`

SetId sets Id field to given value.


### GetSummary

`func (o *Crq) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Crq) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Crq) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *Crq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Crq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Crq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *Crq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Crq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Crq) SetType(v string)`

SetType sets Type field to given value.


### GetIsAutoClose

`func (o *Crq) GetIsAutoClose() bool`

GetIsAutoClose returns the IsAutoClose field if non-nil, zero value otherwise.

### GetIsAutoCloseOk

`func (o *Crq) GetIsAutoCloseOk() (*bool, bool)`

GetIsAutoCloseOk returns a tuple with the IsAutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoClose

`func (o *Crq) SetIsAutoClose(v bool)`

SetIsAutoClose sets IsAutoClose field to given value.


### GetCreatedAt

`func (o *Crq) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Crq) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Crq) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Crq) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Crq) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Crq) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Crq) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



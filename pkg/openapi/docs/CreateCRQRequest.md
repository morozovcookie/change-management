# CreateCRQRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | **string** | The short-form description of change | 
**Description** | **string** | The long-form description of change | 
**Type** | Pointer to **string** | The type of change | [optional] [default to "crq"]
**IsAutoClose** | Pointer to **bool** | This flag indicates that change request should be closed  automatically | [optional] [default to false]

## Methods

### NewCreateCRQRequest

`func NewCreateCRQRequest(summary string, description string, ) *CreateCRQRequest`

NewCreateCRQRequest instantiates a new CreateCRQRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCRQRequestWithDefaults

`func NewCreateCRQRequestWithDefaults() *CreateCRQRequest`

NewCreateCRQRequestWithDefaults instantiates a new CreateCRQRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *CreateCRQRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CreateCRQRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CreateCRQRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *CreateCRQRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCRQRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCRQRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *CreateCRQRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCRQRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCRQRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateCRQRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsAutoClose

`func (o *CreateCRQRequest) GetIsAutoClose() bool`

GetIsAutoClose returns the IsAutoClose field if non-nil, zero value otherwise.

### GetIsAutoCloseOk

`func (o *CreateCRQRequest) GetIsAutoCloseOk() (*bool, bool)`

GetIsAutoCloseOk returns a tuple with the IsAutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoClose

`func (o *CreateCRQRequest) SetIsAutoClose(v bool)`

SetIsAutoClose sets IsAutoClose field to given value.

### HasIsAutoClose

`func (o *CreateCRQRequest) HasIsAutoClose() bool`

HasIsAutoClose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateIncidentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | **string** | The short-form description of change | 
**Description** | **string** | The long-form description of change | 

## Methods

### NewCreateIncidentRequest

`func NewCreateIncidentRequest(summary string, description string, ) *CreateIncidentRequest`

NewCreateIncidentRequest instantiates a new CreateIncidentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIncidentRequestWithDefaults

`func NewCreateIncidentRequestWithDefaults() *CreateIncidentRequest`

NewCreateIncidentRequestWithDefaults instantiates a new CreateIncidentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *CreateIncidentRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CreateIncidentRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CreateIncidentRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *CreateIncidentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIncidentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIncidentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



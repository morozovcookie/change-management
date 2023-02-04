# Incident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**SelfLink**](SelfLink.md) |  | 
**Id** | **string** |  | 
**Summary** | **string** | The short-form description of change | 
**Description** | **string** | The long-form description of change | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewIncident

`func NewIncident(links SelfLink, id string, summary string, description string, createdAt int64, ) *Incident`

NewIncident instantiates a new Incident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentWithDefaults

`func NewIncidentWithDefaults() *Incident`

NewIncidentWithDefaults instantiates a new Incident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Incident) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Incident) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Incident) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *Incident) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Incident) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Incident) SetId(v string)`

SetId sets Id field to given value.


### GetSummary

`func (o *Incident) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Incident) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Incident) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *Incident) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Incident) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Incident) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *Incident) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Incident) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Incident) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Incident) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Incident) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Incident) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Incident) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



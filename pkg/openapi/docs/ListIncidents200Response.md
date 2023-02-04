# ListIncidents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**Links**](Links.md) |  | 
**Limit** | **int32** |  | [default to 20]
**Start** | **int32** |  | [default to 0]
**Total** | **int32** |  | 
**Data** | [**[]Incident**](Incident.md) |  | 

## Methods

### NewListIncidents200Response

`func NewListIncidents200Response(links Links, limit int32, start int32, total int32, data []Incident, ) *ListIncidents200Response`

NewListIncidents200Response instantiates a new ListIncidents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIncidents200ResponseWithDefaults

`func NewListIncidents200ResponseWithDefaults() *ListIncidents200Response`

NewListIncidents200ResponseWithDefaults instantiates a new ListIncidents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListIncidents200Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListIncidents200Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListIncidents200Response) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetLimit

`func (o *ListIncidents200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListIncidents200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListIncidents200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetStart

`func (o *ListIncidents200Response) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListIncidents200Response) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListIncidents200Response) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ListIncidents200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListIncidents200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListIncidents200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetData

`func (o *ListIncidents200Response) GetData() []Incident`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListIncidents200Response) GetDataOk() (*[]Incident, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListIncidents200Response) SetData(v []Incident)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



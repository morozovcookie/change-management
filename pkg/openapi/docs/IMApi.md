# \IMApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncident**](IMApi.md#CreateIncident) | **Post** /api/v1/im | Create a single incident
[**GetIncidentById**](IMApi.md#GetIncidentById) | **Get** /api/v1/im/{incidentId} | Retrieve a single incident by id
[**ListIncidents**](IMApi.md#ListIncidents) | **Get** /api/v1/im | Get list of incidents



## CreateIncident

> CreateIncident(ctx).CreateIncidentRequest(createIncidentRequest).Execute()

Create a single incident



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createIncidentRequest := *openapiclient.NewCreateIncidentRequest("Incident on payroll system", "Since the release of a new Version the Data in the database is all  scrambled up and the System wont start.") // CreateIncidentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IMApi.CreateIncident(context.Background()).CreateIncidentRequest(createIncidentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IMApi.CreateIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIncidentRequest** | [**CreateIncidentRequest**](CreateIncidentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentById

> Incident GetIncidentById(ctx, incidentId).Execute()

Retrieve a single incident by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incidentId := "7bu21gwa31fk8fwtfk36444m4mvukgg0" // string | Incident unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IMApi.GetIncidentById(context.Background(), incidentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IMApi.GetIncidentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentById`: Incident
    fmt.Fprintf(os.Stdout, "Response from `IMApi.GetIncidentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | Incident unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Incident**](Incident.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncidents

> ListIncidents200Response ListIncidents(ctx).Start(start).Limit(limit).Execute()

Get list of incidents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    start := int32(0) // int32 | The count of records that should be skipped (optional) (default to 0)
    limit := int32(50) // int32 | The maximum records that should be returned (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IMApi.ListIncidents(context.Background()).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IMApi.ListIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncidents`: ListIncidents200Response
    fmt.Fprintf(os.Stdout, "Response from `IMApi.ListIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | The count of records that should be skipped | [default to 0]
 **limit** | **int32** | The maximum records that should be returned | [default to 20]

### Return type

[**ListIncidents200Response**](ListIncidents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


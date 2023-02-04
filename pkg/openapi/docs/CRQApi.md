# \CRQApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCRQ**](CRQApi.md#CreateCRQ) | **Post** /api/v1/crq | Create a single change request
[**GetCRQById**](CRQApi.md#GetCRQById) | **Get** /api/v1/crq/{crqId} | Retrieve a single change request by id
[**ListCRQ**](CRQApi.md#ListCRQ) | **Get** /api/v1/crq | Get list of change requests



## CreateCRQ

> CreateCRQ(ctx).CreateCRQRequest(createCRQRequest).Execute()

Create a single change request



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
    createCRQRequest := *openapiclient.NewCreateCRQRequest("Bump service <service-name> version", "Bump <service-name> version

CHANGELOG:
  - minor fixes
") // CreateCRQRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CRQApi.CreateCRQ(context.Background()).CreateCRQRequest(createCRQRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRQApi.CreateCRQ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCRQRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCRQRequest** | [**CreateCRQRequest**](CreateCRQRequest.md) |  | 

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


## GetCRQById

> Crq GetCRQById(ctx, crqId).Execute()

Retrieve a single change request by id



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
    crqId := "8x8cqdf0shscsy4rtrndwliqtaynrvct" // string | Change request unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CRQApi.GetCRQById(context.Background(), crqId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRQApi.GetCRQById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCRQById`: Crq
    fmt.Fprintf(os.Stdout, "Response from `CRQApi.GetCRQById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crqId** | **string** | Change request unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCRQByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Crq**](Crq.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCRQ

> ListCRQ200Response ListCRQ(ctx).Start(start).Limit(limit).Execute()

Get list of change requests



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
    resp, r, err := apiClient.CRQApi.ListCRQ(context.Background()).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRQApi.ListCRQ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCRQ`: ListCRQ200Response
    fmt.Fprintf(os.Stdout, "Response from `CRQApi.ListCRQ`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCRQRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | The count of records that should be skipped | [default to 0]
 **limit** | **int32** | The maximum records that should be returned | [default to 20]

### Return type

[**ListCRQ200Response**](ListCRQ200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


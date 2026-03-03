# \SandboxesAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxesGet**](SandboxesAPI.md#SandboxesGet) | **Get** /sandboxes | List sandboxes
[**SandboxesPost**](SandboxesAPI.md#SandboxesPost) | **Post** /sandboxes | Create a sandbox from a container image
[**SandboxesSandboxIdDelete**](SandboxesAPI.md#SandboxesSandboxIdDelete) | **Delete** /sandboxes/{sandboxId} | Delete a sandbox
[**SandboxesSandboxIdEndpointsPortGet**](SandboxesAPI.md#SandboxesSandboxIdEndpointsPortGet) | **Get** /sandboxes/{sandboxId}/endpoints/{port} | Get sandbox access endpoint
[**SandboxesSandboxIdGet**](SandboxesAPI.md#SandboxesSandboxIdGet) | **Get** /sandboxes/{sandboxId} | Fetch a sandbox by id
[**SandboxesSandboxIdPausePost**](SandboxesAPI.md#SandboxesSandboxIdPausePost) | **Post** /sandboxes/{sandboxId}/pause | Pause execution while retaining state
[**SandboxesSandboxIdRenewExpirationPost**](SandboxesAPI.md#SandboxesSandboxIdRenewExpirationPost) | **Post** /sandboxes/{sandboxId}/renew-expiration | Renew sandbox expiration
[**SandboxesSandboxIdResumePost**](SandboxesAPI.md#SandboxesSandboxIdResumePost) | **Post** /sandboxes/{sandboxId}/resume | Resume a paused sandbox



## SandboxesGet

> ListSandboxesResponse SandboxesGet(ctx).State(state).Metadata(metadata).Page(page).PageSize(pageSize).Execute()

List sandboxes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alibaba/opensandbox-client-go"
)

func main() {
	state := []string{"Inner_example"} // []string | Filter by lifecycle state. Pass multiple times for OR logic. Example: `?state=Running&state=Paused`  (optional)
	metadata := "metadata_example" // string | Arbitrary metadata key-value pairs for filtering，keys and values must be url encoded Example: To filter by `project=Apollo` and `note=Demo Test`: `?metadata=project%3DApollo%26note%3DDemo%252520Test`  (optional)
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.SandboxesGet(context.Background()).State(state).Metadata(metadata).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.SandboxesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxesGet`: ListSandboxesResponse
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.SandboxesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **[]string** | Filter by lifecycle state. Pass multiple times for OR logic. Example: &#x60;?state&#x3D;Running&amp;state&#x3D;Paused&#x60;  | 
 **metadata** | **string** | Arbitrary metadata key-value pairs for filtering，keys and values must be url encoded Example: To filter by &#x60;project&#x3D;Apollo&#x60; and &#x60;note&#x3D;Demo Test&#x60;: &#x60;?metadata&#x3D;project%3DApollo%26note%3DDemo%252520Test&#x60;  | 
 **page** | **int32** | Page number for pagination | [default to 1]
 **pageSize** | **int32** | Number of items per page | [default to 20]

### Return type

[**ListSandboxesResponse**](ListSandboxesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxesPost

> CreateSandboxResponse SandboxesPost(ctx).CreateSandboxRequest(createSandboxRequest).Execute()

Create a sandbox from a container image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alibaba/opensandbox-client-go"
)

func main() {
	createSandboxRequest := *openapiclient.NewCreateSandboxRequest(*openapiclient.NewImageSpec("Uri_example"), int32(123), map[string]string{"key": "Inner_example"}, []string{"Entrypoint_example"}) // CreateSandboxRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.SandboxesPost(context.Background()).CreateSandboxRequest(createSandboxRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.SandboxesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxesPost`: CreateSandboxResponse
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.SandboxesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSandboxRequest** | [**CreateSandboxRequest**](CreateSandboxRequest.md) |  | 

### Return type

[**CreateSandboxResponse**](CreateSandboxResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxesSandboxIdDelete

> SandboxesSandboxIdDelete(ctx, sandboxId).Execute()

Delete a sandbox



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alibaba/opensandbox-client-go"
)

func main() {
	sandboxId := "sandboxId_example" // string | Unique sandbox identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxesAPI.SandboxesSandboxIdDelete(context.Background(), sandboxId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.SandboxesSandboxIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** | Unique sandbox identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxesSandboxIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxesSandboxIdEndpointsPortGet

> Endpoint SandboxesSandboxIdEndpointsPortGet(ctx, sandboxId, port).UseServerProxy(useServerProxy).Execute()

Get sandbox access endpoint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alibaba/opensandbox-client-go"
)

func main() {
	sandboxId := "sandboxId_example" // string | Unique sandbox identifier
	port := int32(56) // int32 | Port number where the service is listening inside the sandbox
	useServerProxy := true // bool | Whether to return a server-proxied URL (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.SandboxesSandboxIdEndpointsPortGet(context.Background(), sandboxId, port).UseServerProxy(useServerProxy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.SandboxesSandboxIdEndpointsPortGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxesSandboxIdEndpointsPortGet`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.SandboxesSandboxIdEndpointsPortGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** | Unique sandbox identifier | 
**port** | **int32** | Port number where the service is listening inside the sandbox | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxesSandboxIdEndpointsPortGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **useServerProxy** | **bool** | Whether to return a server-proxied URL | [default to false]

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxesSandboxIdGet

> Sandbox SandboxesSandboxIdGet(ctx, sandboxId).Execute()

Fetch a sandbox by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alibaba/opensandbox-client-go"
)

func main() {
	sandboxId := "sandboxId_example" // string | Unique sandbox identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.SandboxesSandboxIdGet(context.Background(), sandboxId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.SandboxesSandboxIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxesSandboxIdGet`: Sandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.SandboxesSandboxIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** | Unique sandbox identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxesSandboxIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sandbox**](Sandbox.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxesSandboxIdPausePost

> SandboxesSandboxIdPausePost(ctx, sandboxId).Execute()

Pause execution while retaining state



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alibaba/opensandbox-client-go"
)

func main() {
	sandboxId := "sandboxId_example" // string | Unique sandbox identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxesAPI.SandboxesSandboxIdPausePost(context.Background(), sandboxId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.SandboxesSandboxIdPausePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** | Unique sandbox identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxesSandboxIdPausePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxesSandboxIdRenewExpirationPost

> RenewSandboxExpirationResponse SandboxesSandboxIdRenewExpirationPost(ctx, sandboxId).RenewSandboxExpirationRequest(renewSandboxExpirationRequest).Execute()

Renew sandbox expiration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/alibaba/opensandbox-client-go"
)

func main() {
	sandboxId := "sandboxId_example" // string | Unique sandbox identifier
	renewSandboxExpirationRequest := *openapiclient.NewRenewSandboxExpirationRequest(time.Now()) // RenewSandboxExpirationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.SandboxesSandboxIdRenewExpirationPost(context.Background(), sandboxId).RenewSandboxExpirationRequest(renewSandboxExpirationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.SandboxesSandboxIdRenewExpirationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxesSandboxIdRenewExpirationPost`: RenewSandboxExpirationResponse
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.SandboxesSandboxIdRenewExpirationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** | Unique sandbox identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxesSandboxIdRenewExpirationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renewSandboxExpirationRequest** | [**RenewSandboxExpirationRequest**](RenewSandboxExpirationRequest.md) |  | 

### Return type

[**RenewSandboxExpirationResponse**](RenewSandboxExpirationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxesSandboxIdResumePost

> SandboxesSandboxIdResumePost(ctx, sandboxId).Execute()

Resume a paused sandbox



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alibaba/opensandbox-client-go"
)

func main() {
	sandboxId := "sandboxId_example" // string | Unique sandbox identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxesAPI.SandboxesSandboxIdResumePost(context.Background(), sandboxId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.SandboxesSandboxIdResumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** | Unique sandbox identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxesSandboxIdResumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


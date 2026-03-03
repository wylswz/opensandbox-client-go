# \CommandAPI

All URIs are relative to *http://localhost:44772*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackgroundCommandLogs**](CommandAPI.md#GetBackgroundCommandLogs) | **Get** /command/{id}/logs | Get background command stdout/stderr (non-streamed)
[**GetCommandStatus**](CommandAPI.md#GetCommandStatus) | **Get** /command/status/{id} | Get command running status
[**InterruptCommand**](CommandAPI.md#InterruptCommand) | **Delete** /command | Interrupt command execution
[**RunCommand**](CommandAPI.md#RunCommand) | **Post** /command | Execute shell command



## GetBackgroundCommandLogs

> string GetBackgroundCommandLogs(ctx, id).Cursor(cursor).Execute()

Get background command stdout/stderr (non-streamed)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wylswz/opensandbox-client-go"
)

func main() {
	id := "cmd-abc123" // string | Command ID returned by RunCommand
	cursor := int64(120) // int64 | Optional 0-based line cursor (behaves like a file seek). When provided, only stdout/stderr lines after this line are returned. The response includes the latest line index (`cursor`) so the client can request incremental output on subsequent calls. If omitted, the full log is returned.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommandAPI.GetBackgroundCommandLogs(context.Background(), id).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandAPI.GetBackgroundCommandLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackgroundCommandLogs`: string
	fmt.Fprintf(os.Stdout, "Response from `CommandAPI.GetBackgroundCommandLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Command ID returned by RunCommand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackgroundCommandLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **int64** | Optional 0-based line cursor (behaves like a file seek). When provided, only stdout/stderr lines after this line are returned. The response includes the latest line index (&#x60;cursor&#x60;) so the client can request incremental output on subsequent calls. If omitted, the full log is returned.  | 

### Return type

**string**

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommandStatus

> CommandStatusResponse GetCommandStatus(ctx, id).Execute()

Get command running status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wylswz/opensandbox-client-go"
)

func main() {
	id := "cmd-abc123" // string | Command ID returned by RunCommand

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommandAPI.GetCommandStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandAPI.GetCommandStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommandStatus`: CommandStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `CommandAPI.GetCommandStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Command ID returned by RunCommand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommandStatusResponse**](CommandStatusResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InterruptCommand

> InterruptCommand(ctx).Id(id).Execute()

Interrupt command execution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wylswz/opensandbox-client-go"
)

func main() {
	id := "session-456" // string | Session ID of the execution context to interrupt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommandAPI.InterruptCommand(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandAPI.InterruptCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInterruptCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Session ID of the execution context to interrupt | 

### Return type

 (empty response body)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunCommand

> ServerStreamEvent RunCommand(ctx).RunCommandRequest(runCommandRequest).Execute()

Execute shell command



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wylswz/opensandbox-client-go"
)

func main() {
	runCommandRequest := *openapiclient.NewRunCommandRequest("ls -la /workspace") // RunCommandRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommandAPI.RunCommand(context.Background()).RunCommandRequest(runCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandAPI.RunCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunCommand`: ServerStreamEvent
	fmt.Fprintf(os.Stdout, "Response from `CommandAPI.RunCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCommandRequest** | [**RunCommandRequest**](RunCommandRequest.md) |  | 

### Return type

[**ServerStreamEvent**](ServerStreamEvent.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


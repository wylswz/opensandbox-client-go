# \CodeInterpretingAPI

All URIs are relative to *http://localhost:44772*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCodeContext**](CodeInterpretingAPI.md#CreateCodeContext) | **Post** /code/context | Create code execution context
[**DeleteContext**](CodeInterpretingAPI.md#DeleteContext) | **Delete** /code/contexts/{context_id} | Delete a code execution context by id
[**DeleteContextsByLanguage**](CodeInterpretingAPI.md#DeleteContextsByLanguage) | **Delete** /code/contexts | Delete all contexts under a language
[**GetContext**](CodeInterpretingAPI.md#GetContext) | **Get** /code/contexts/{context_id} | Get a code execution context by id
[**InterruptCode**](CodeInterpretingAPI.md#InterruptCode) | **Delete** /code | Interrupt code execution
[**ListContexts**](CodeInterpretingAPI.md#ListContexts) | **Get** /code/contexts | List active code execution contexts
[**RunCode**](CodeInterpretingAPI.md#RunCode) | **Post** /code | Execute code in context



## CreateCodeContext

> CodeContext CreateCodeContext(ctx).CodeContextRequest(codeContextRequest).Execute()

Create code execution context



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
	codeContextRequest := *openapiclient.NewCodeContextRequest() // CodeContextRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeInterpretingAPI.CreateCodeContext(context.Background()).CodeContextRequest(codeContextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeInterpretingAPI.CreateCodeContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCodeContext`: CodeContext
	fmt.Fprintf(os.Stdout, "Response from `CodeInterpretingAPI.CreateCodeContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCodeContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeContextRequest** | [**CodeContextRequest**](CodeContextRequest.md) |  | 

### Return type

[**CodeContext**](CodeContext.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContext

> DeleteContext(ctx, contextId).Execute()

Delete a code execution context by id



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
	contextId := "session-abc123" // string | Session/context id to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CodeInterpretingAPI.DeleteContext(context.Background(), contextId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeInterpretingAPI.DeleteContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | Session/context id to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteContextsByLanguage

> DeleteContextsByLanguage(ctx).Language(language).Execute()

Delete all contexts under a language



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
	language := "python" // string | Target execution runtime whose contexts should be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CodeInterpretingAPI.DeleteContextsByLanguage(context.Background()).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeInterpretingAPI.DeleteContextsByLanguage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContextsByLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** | Target execution runtime whose contexts should be deleted | 

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


## GetContext

> CodeContext GetContext(ctx, contextId).Execute()

Get a code execution context by id



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
	contextId := "session-abc123" // string | Session/context id to get

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeInterpretingAPI.GetContext(context.Background(), contextId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeInterpretingAPI.GetContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContext`: CodeContext
	fmt.Fprintf(os.Stdout, "Response from `CodeInterpretingAPI.GetContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | Session/context id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CodeContext**](CodeContext.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InterruptCode

> InterruptCode(ctx).Id(id).Execute()

Interrupt code execution



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
	id := "session-123" // string | Session ID of the execution context to interrupt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CodeInterpretingAPI.InterruptCode(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeInterpretingAPI.InterruptCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInterruptCodeRequest struct via the builder pattern


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


## ListContexts

> []CodeContext ListContexts(ctx).Language(language).Execute()

List active code execution contexts



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
	language := "python" // string | Filter contexts by execution runtime (python, bash, java, etc.)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeInterpretingAPI.ListContexts(context.Background()).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeInterpretingAPI.ListContexts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContexts`: []CodeContext
	fmt.Fprintf(os.Stdout, "Response from `CodeInterpretingAPI.ListContexts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** | Filter contexts by execution runtime (python, bash, java, etc.) | 

### Return type

[**[]CodeContext**](CodeContext.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunCode

> ServerStreamEvent RunCode(ctx).RunCodeRequest(runCodeRequest).Execute()

Execute code in context



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
	runCodeRequest := *openapiclient.NewRunCodeRequest("import numpy as np
result = np.array([1, 2, 3])
print(result)
") // RunCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeInterpretingAPI.RunCode(context.Background()).RunCodeRequest(runCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeInterpretingAPI.RunCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunCode`: ServerStreamEvent
	fmt.Fprintf(os.Stdout, "Response from `CodeInterpretingAPI.RunCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCodeRequest** | [**RunCodeRequest**](RunCodeRequest.md) |  | 

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


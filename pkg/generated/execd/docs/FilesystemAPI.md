# \FilesystemAPI

All URIs are relative to *http://localhost:44772*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChmodFiles**](FilesystemAPI.md#ChmodFiles) | **Post** /files/permissions | Change file permissions
[**DownloadFile**](FilesystemAPI.md#DownloadFile) | **Get** /files/download | Download file from sandbox
[**GetFilesInfo**](FilesystemAPI.md#GetFilesInfo) | **Get** /files/info | Get file metadata
[**MakeDirs**](FilesystemAPI.md#MakeDirs) | **Post** /directories | Create directories
[**RemoveDirs**](FilesystemAPI.md#RemoveDirs) | **Delete** /directories | Delete directories
[**RemoveFiles**](FilesystemAPI.md#RemoveFiles) | **Delete** /files | Delete files
[**RenameFiles**](FilesystemAPI.md#RenameFiles) | **Post** /files/mv | Rename or move files
[**ReplaceContent**](FilesystemAPI.md#ReplaceContent) | **Post** /files/replace | Replace file content
[**SearchFiles**](FilesystemAPI.md#SearchFiles) | **Get** /files/search | Search for files
[**UploadFile**](FilesystemAPI.md#UploadFile) | **Post** /files/upload | Upload files to sandbox



## ChmodFiles

> ChmodFiles(ctx).RequestBody(requestBody).Execute()

Change file permissions



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
	requestBody := map[string]Permission{"key": *openapiclient.NewPermission(int32(755))} // map[string]Permission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesystemAPI.ChmodFiles(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.ChmodFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChmodFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string]Permission**](Permission.md) |  | 

### Return type

 (empty response body)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFile

> *os.File DownloadFile(ctx).Path(path).Range_(range_).Execute()

Download file from sandbox



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
	path := "/workspace/data.csv" // string | Absolute or relative path of the file to download
	range_ := "bytes=0-1023" // string | HTTP Range header for partial content requests (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesystemAPI.DownloadFile(context.Background()).Path(path).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.DownloadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FilesystemAPI.DownloadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Absolute or relative path of the file to download | 
 **range_** | **string** | HTTP Range header for partial content requests | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesInfo

> map[string]FileInfo GetFilesInfo(ctx).Path(path).Execute()

Get file metadata



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
	path := []string{"Inner_example"} // []string | File path(s) to get info for (can be specified multiple times)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesystemAPI.GetFilesInfo(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.GetFilesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesInfo`: map[string]FileInfo
	fmt.Fprintf(os.Stdout, "Response from `FilesystemAPI.GetFilesInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **[]string** | File path(s) to get info for (can be specified multiple times) | 

### Return type

[**map[string]FileInfo**](FileInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeDirs

> MakeDirs(ctx).RequestBody(requestBody).Execute()

Create directories



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
	requestBody := map[string]Permission{"key": *openapiclient.NewPermission(int32(755))} // map[string]Permission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesystemAPI.MakeDirs(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.MakeDirs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMakeDirsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string]Permission**](Permission.md) |  | 

### Return type

 (empty response body)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDirs

> RemoveDirs(ctx).Path(path).Execute()

Delete directories



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
	path := []string{"Inner_example"} // []string | Directory path(s) to delete (can be specified multiple times)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesystemAPI.RemoveDirs(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.RemoveDirs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDirsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **[]string** | Directory path(s) to delete (can be specified multiple times) | 

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


## RemoveFiles

> RemoveFiles(ctx).Path(path).Execute()

Delete files



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
	path := []string{"Inner_example"} // []string | File path(s) to delete (can be specified multiple times)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesystemAPI.RemoveFiles(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.RemoveFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **[]string** | File path(s) to delete (can be specified multiple times) | 

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


## RenameFiles

> RenameFiles(ctx).RenameFileItem(renameFileItem).Execute()

Rename or move files



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
	renameFileItem := []openapiclient.RenameFileItem{*openapiclient.NewRenameFileItem("/workspace/old.txt", "/workspace/new.txt")} // []RenameFileItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesystemAPI.RenameFiles(context.Background()).RenameFileItem(renameFileItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.RenameFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameFileItem** | [**[]RenameFileItem**](RenameFileItem.md) |  | 

### Return type

 (empty response body)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceContent

> ReplaceContent(ctx).RequestBody(requestBody).Execute()

Replace file content



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
	requestBody := map[string]ReplaceFileContentItem{"key": *openapiclient.NewReplaceFileContentItem("localhost", "0.0.0.0")} // map[string]ReplaceFileContentItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesystemAPI.ReplaceContent(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.ReplaceContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string]ReplaceFileContentItem**](ReplaceFileContentItem.md) |  | 

### Return type

 (empty response body)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFiles

> []FileInfo SearchFiles(ctx).Path(path).Pattern(pattern).Execute()

Search for files



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
	path := "path_example" // string | Root directory path to search in
	pattern := "pattern_example" // string | Glob pattern to match files (default is **) (optional) (default to "**")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesystemAPI.SearchFiles(context.Background()).Path(path).Pattern(pattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.SearchFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFiles`: []FileInfo
	fmt.Fprintf(os.Stdout, "Response from `FilesystemAPI.SearchFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Root directory path to search in | 
 **pattern** | **string** | Glob pattern to match files (default is **) | [default to &quot;**&quot;]

### Return type

[**[]FileInfo**](FileInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> UploadFile(ctx).Metadata(metadata).File(file).Execute()

Upload files to sandbox



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
	metadata := "metadata_example" // string | JSON-encoded file metadata (FileMetadata object) (optional)
	file := os.NewFile(1234, "some_file") // *os.File | File to upload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesystemAPI.UploadFile(context.Background()).Metadata(metadata).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemAPI.UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | **string** | JSON-encoded file metadata (FileMetadata object) | 
 **file** | ***os.File** | File to upload | 

### Return type

 (empty response body)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


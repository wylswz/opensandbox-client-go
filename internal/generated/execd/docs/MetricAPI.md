# \MetricAPI

All URIs are relative to *http://localhost:44772*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricAPI.md#GetMetrics) | **Get** /metrics | Get system metrics
[**WatchMetrics**](MetricAPI.md#WatchMetrics) | **Get** /metrics/watch | Watch system metrics in real-time



## GetMetrics

> Metrics GetMetrics(ctx).Execute()

Get system metrics



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricAPI.GetMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.GetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `MetricAPI.GetMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


### Return type

[**Metrics**](Metrics.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchMetrics

> Metrics WatchMetrics(ctx).Execute()

Watch system metrics in real-time



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricAPI.WatchMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.WatchMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WatchMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `MetricAPI.WatchMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWatchMetricsRequest struct via the builder pattern


### Return type

[**Metrics**](Metrics.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


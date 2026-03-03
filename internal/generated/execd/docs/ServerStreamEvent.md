# ServerStreamEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Event type for client-side handling | [optional] 
**Text** | Pointer to **string** | Textual data for status, init, and stream events | [optional] 
**ExecutionCount** | Pointer to **int32** | Cell execution number in the session | [optional] 
**ExecutionTime** | Pointer to **int64** | Execution duration in milliseconds | [optional] 
**Timestamp** | Pointer to **int64** | When the event was generated (Unix milliseconds) | [optional] 
**Results** | Pointer to **map[string]interface{}** | Execution output in various MIME types (e.g., \&quot;text/plain\&quot;, \&quot;text/html\&quot;) | [optional] 
**Error** | Pointer to [**ServerStreamEventError**](ServerStreamEventError.md) |  | [optional] 

## Methods

### NewServerStreamEvent

`func NewServerStreamEvent() *ServerStreamEvent`

NewServerStreamEvent instantiates a new ServerStreamEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStreamEventWithDefaults

`func NewServerStreamEventWithDefaults() *ServerStreamEvent`

NewServerStreamEventWithDefaults instantiates a new ServerStreamEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServerStreamEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerStreamEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerStreamEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServerStreamEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *ServerStreamEvent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ServerStreamEvent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ServerStreamEvent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ServerStreamEvent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetExecutionCount

`func (o *ServerStreamEvent) GetExecutionCount() int32`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *ServerStreamEvent) GetExecutionCountOk() (*int32, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *ServerStreamEvent) SetExecutionCount(v int32)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *ServerStreamEvent) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.

### GetExecutionTime

`func (o *ServerStreamEvent) GetExecutionTime() int64`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *ServerStreamEvent) GetExecutionTimeOk() (*int64, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *ServerStreamEvent) SetExecutionTime(v int64)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *ServerStreamEvent) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetTimestamp

`func (o *ServerStreamEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServerStreamEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServerStreamEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServerStreamEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetResults

`func (o *ServerStreamEvent) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ServerStreamEvent) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ServerStreamEvent) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *ServerStreamEvent) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetError

`func (o *ServerStreamEvent) GetError() ServerStreamEventError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ServerStreamEvent) GetErrorOk() (*ServerStreamEventError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ServerStreamEvent) SetError(v ServerStreamEventError)`

SetError sets Error field to given value.

### HasError

`func (o *ServerStreamEvent) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



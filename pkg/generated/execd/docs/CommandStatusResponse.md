# CommandStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Command ID returned by RunCommand | [optional] 
**Content** | Pointer to **string** | Original command content | [optional] 
**Running** | Pointer to **bool** | Whether the command is still running | [optional] 
**ExitCode** | Pointer to **int32** | Exit code if the command has finished | [optional] 
**Error** | Pointer to **string** | Error message if the command failed | [optional] 
**StartedAt** | Pointer to **time.Time** | Start time in RFC3339 format | [optional] 
**FinishedAt** | Pointer to **time.Time** | Finish time in RFC3339 format (null if still running) | [optional] 

## Methods

### NewCommandStatusResponse

`func NewCommandStatusResponse() *CommandStatusResponse`

NewCommandStatusResponse instantiates a new CommandStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandStatusResponseWithDefaults

`func NewCommandStatusResponseWithDefaults() *CommandStatusResponse`

NewCommandStatusResponseWithDefaults instantiates a new CommandStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommandStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommandStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommandStatusResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommandStatusResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *CommandStatusResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CommandStatusResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CommandStatusResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CommandStatusResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRunning

`func (o *CommandStatusResponse) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *CommandStatusResponse) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *CommandStatusResponse) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *CommandStatusResponse) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetExitCode

`func (o *CommandStatusResponse) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *CommandStatusResponse) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *CommandStatusResponse) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *CommandStatusResponse) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetError

`func (o *CommandStatusResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CommandStatusResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CommandStatusResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CommandStatusResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartedAt

`func (o *CommandStatusResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CommandStatusResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CommandStatusResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CommandStatusResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *CommandStatusResponse) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *CommandStatusResponse) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *CommandStatusResponse) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *CommandStatusResponse) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



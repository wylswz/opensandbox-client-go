# RunCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | Shell command to execute | 
**Cwd** | Pointer to **string** | Working directory for command execution | [optional] 
**Background** | Pointer to **bool** | Whether to run command in detached mode | [optional] [default to false]
**Timeout** | Pointer to **int64** | Maximum allowed execution time in milliseconds before the command is forcefully terminated by the server. If omitted, the server will not enforce any timeout. | [optional] 

## Methods

### NewRunCommandRequest

`func NewRunCommandRequest(command string, ) *RunCommandRequest`

NewRunCommandRequest instantiates a new RunCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCommandRequestWithDefaults

`func NewRunCommandRequestWithDefaults() *RunCommandRequest`

NewRunCommandRequestWithDefaults instantiates a new RunCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *RunCommandRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *RunCommandRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *RunCommandRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetCwd

`func (o *RunCommandRequest) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *RunCommandRequest) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *RunCommandRequest) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *RunCommandRequest) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetBackground

`func (o *RunCommandRequest) GetBackground() bool`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *RunCommandRequest) GetBackgroundOk() (*bool, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *RunCommandRequest) SetBackground(v bool)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *RunCommandRequest) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetTimeout

`func (o *RunCommandRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RunCommandRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RunCommandRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *RunCommandRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



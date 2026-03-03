# SandboxStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | High-level lifecycle state of the sandbox.  Common state values: - Pending: Sandbox is being provisioned - Running: Sandbox is running and ready to accept requests - Pausing: Sandbox is in the process of pausing - Paused: Sandbox has been paused while retaining its state - Stopping: Sandbox is being terminated - Terminated: Sandbox has been successfully terminated - Failed: Sandbox encountered a critical error  State transitions: - Pending → Running (after creation completes) - Running → Pausing (when pause is requested) - Pausing → Paused (pause operation completes) - Paused → Running (when resume is requested) - Running/Paused → Stopping (when kill is requested or TTL expires) - Stopping → Terminated (kill/timeout operation completes) - Pending/Running/Paused → Failed (on error)  Note: New state values may be added in future versions. Clients should handle unknown state values gracefully.  | 
**Reason** | Pointer to **string** | Short machine-readable reason code for the current state. Examples: \&quot;user_delete\&quot;, \&quot;ttl_expiry\&quot;, \&quot;provision_timeout\&quot;, \&quot;runtime_error\&quot;  | [optional] 
**Message** | Pointer to **string** | Human-readable message describing the current state or reason for state transition | [optional] 
**LastTransitionAt** | Pointer to **time.Time** | Timestamp of the last state transition | [optional] 

## Methods

### NewSandboxStatus

`func NewSandboxStatus(state string, ) *SandboxStatus`

NewSandboxStatus instantiates a new SandboxStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxStatusWithDefaults

`func NewSandboxStatusWithDefaults() *SandboxStatus`

NewSandboxStatusWithDefaults instantiates a new SandboxStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SandboxStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SandboxStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SandboxStatus) SetState(v string)`

SetState sets State field to given value.


### GetReason

`func (o *SandboxStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SandboxStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SandboxStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SandboxStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *SandboxStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SandboxStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SandboxStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SandboxStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLastTransitionAt

`func (o *SandboxStatus) GetLastTransitionAt() time.Time`

GetLastTransitionAt returns the LastTransitionAt field if non-nil, zero value otherwise.

### GetLastTransitionAtOk

`func (o *SandboxStatus) GetLastTransitionAtOk() (*time.Time, bool)`

GetLastTransitionAtOk returns a tuple with the LastTransitionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionAt

`func (o *SandboxStatus) SetLastTransitionAt(v time.Time)`

SetLastTransitionAt sets LastTransitionAt field to given value.

### HasLastTransitionAt

`func (o *SandboxStatus) HasLastTransitionAt() bool`

HasLastTransitionAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



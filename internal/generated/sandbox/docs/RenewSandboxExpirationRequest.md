# RenewSandboxExpirationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | **time.Time** | New absolute expiration time in UTC (RFC 3339 format). Must be in the future and after the current expiresAt time.  Example: \&quot;2025-11-16T14:30:45Z\&quot;  | 

## Methods

### NewRenewSandboxExpirationRequest

`func NewRenewSandboxExpirationRequest(expiresAt time.Time, ) *RenewSandboxExpirationRequest`

NewRenewSandboxExpirationRequest instantiates a new RenewSandboxExpirationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewSandboxExpirationRequestWithDefaults

`func NewRenewSandboxExpirationRequestWithDefaults() *RenewSandboxExpirationRequest`

NewRenewSandboxExpirationRequestWithDefaults instantiates a new RenewSandboxExpirationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *RenewSandboxExpirationRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RenewSandboxExpirationRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RenewSandboxExpirationRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



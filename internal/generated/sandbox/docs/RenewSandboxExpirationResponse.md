# RenewSandboxExpirationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | **time.Time** | The new absolute expiration time in UTC (RFC 3339 format).  Example: \&quot;2025-11-16T14:30:45Z\&quot;  | 

## Methods

### NewRenewSandboxExpirationResponse

`func NewRenewSandboxExpirationResponse(expiresAt time.Time, ) *RenewSandboxExpirationResponse`

NewRenewSandboxExpirationResponse instantiates a new RenewSandboxExpirationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewSandboxExpirationResponseWithDefaults

`func NewRenewSandboxExpirationResponseWithDefaults() *RenewSandboxExpirationResponse`

NewRenewSandboxExpirationResponseWithDefaults instantiates a new RenewSandboxExpirationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *RenewSandboxExpirationResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RenewSandboxExpirationResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RenewSandboxExpirationResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



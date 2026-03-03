# CreateSandboxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique sandbox identifier | 
**Status** | [**SandboxStatus**](SandboxStatus.md) | Current lifecycle status and detailed state information | 
**Metadata** | Pointer to **map[string]string** | Custom metadata from creation request | [optional] 
**ExpiresAt** | **time.Time** | Timestamp when sandbox will auto-terminate | 
**CreatedAt** | **time.Time** | Sandbox creation timestamp | 
**Entrypoint** | **[]string** | Entry process specification from creation request | 

## Methods

### NewCreateSandboxResponse

`func NewCreateSandboxResponse(id string, status SandboxStatus, expiresAt time.Time, createdAt time.Time, entrypoint []string, ) *CreateSandboxResponse`

NewCreateSandboxResponse instantiates a new CreateSandboxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSandboxResponseWithDefaults

`func NewCreateSandboxResponseWithDefaults() *CreateSandboxResponse`

NewCreateSandboxResponseWithDefaults instantiates a new CreateSandboxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateSandboxResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateSandboxResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateSandboxResponse) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *CreateSandboxResponse) GetStatus() SandboxStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSandboxResponse) GetStatusOk() (*SandboxStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSandboxResponse) SetStatus(v SandboxStatus)`

SetStatus sets Status field to given value.


### GetMetadata

`func (o *CreateSandboxResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSandboxResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSandboxResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSandboxResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateSandboxResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateSandboxResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateSandboxResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCreatedAt

`func (o *CreateSandboxResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateSandboxResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateSandboxResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEntrypoint

`func (o *CreateSandboxResponse) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *CreateSandboxResponse) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *CreateSandboxResponse) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



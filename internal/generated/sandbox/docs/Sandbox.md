# Sandbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique sandbox identifier | 
**Image** | [**ImageSpec**](ImageSpec.md) | Container image specification used to provision this sandbox. Only present in responses for GET/LIST operations. Not returned in createSandbox response.  | 
**Status** | [**SandboxStatus**](SandboxStatus.md) | Current lifecycle status and detailed state information | 
**Metadata** | Pointer to **map[string]string** | Custom metadata from creation request | [optional] 
**Entrypoint** | **[]string** | The command to execute as the sandbox&#39;s entry process. Always present in responses since entrypoint is required in creation requests.  | 
**ExpiresAt** | **time.Time** | Timestamp when sandbox will auto-terminate | 
**CreatedAt** | **time.Time** | Sandbox creation timestamp | 

## Methods

### NewSandbox

`func NewSandbox(id string, image ImageSpec, status SandboxStatus, entrypoint []string, expiresAt time.Time, createdAt time.Time, ) *Sandbox`

NewSandbox instantiates a new Sandbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxWithDefaults

`func NewSandboxWithDefaults() *Sandbox`

NewSandboxWithDefaults instantiates a new Sandbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Sandbox) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sandbox) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sandbox) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *Sandbox) GetImage() ImageSpec`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Sandbox) GetImageOk() (*ImageSpec, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Sandbox) SetImage(v ImageSpec)`

SetImage sets Image field to given value.


### GetStatus

`func (o *Sandbox) GetStatus() SandboxStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Sandbox) GetStatusOk() (*SandboxStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Sandbox) SetStatus(v SandboxStatus)`

SetStatus sets Status field to given value.


### GetMetadata

`func (o *Sandbox) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Sandbox) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Sandbox) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Sandbox) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntrypoint

`func (o *Sandbox) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *Sandbox) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *Sandbox) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.


### GetExpiresAt

`func (o *Sandbox) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Sandbox) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Sandbox) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCreatedAt

`func (o *Sandbox) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Sandbox) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Sandbox) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique identifier for the volume within the sandbox. Must be a valid DNS label (lowercase alphanumeric, hyphens allowed, max 63 chars).  | 
**Host** | Pointer to [**Host**](Host.md) |  | [optional] 
**Pvc** | Pointer to [**PVC**](PVC.md) |  | [optional] 
**MountPath** | **string** | Absolute path inside the container where the volume is mounted. Must start with &#39;/&#39;.  | 
**ReadOnly** | Pointer to **bool** | If true, the volume is mounted as read-only. Defaults to false (read-write).  | [optional] [default to false]
**SubPath** | Pointer to **string** | Optional subdirectory under the backend path to mount. Must be a relative path without &#39;..&#39; components.  | [optional] 

## Methods

### NewVolume

`func NewVolume(name string, mountPath string, ) *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *Volume) GetHost() Host`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Volume) GetHostOk() (*Host, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Volume) SetHost(v Host)`

SetHost sets Host field to given value.

### HasHost

`func (o *Volume) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPvc

`func (o *Volume) GetPvc() PVC`

GetPvc returns the Pvc field if non-nil, zero value otherwise.

### GetPvcOk

`func (o *Volume) GetPvcOk() (*PVC, bool)`

GetPvcOk returns a tuple with the Pvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvc

`func (o *Volume) SetPvc(v PVC)`

SetPvc sets Pvc field to given value.

### HasPvc

`func (o *Volume) HasPvc() bool`

HasPvc returns a boolean if a field has been set.

### GetMountPath

`func (o *Volume) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *Volume) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *Volume) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetReadOnly

`func (o *Volume) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *Volume) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *Volume) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *Volume) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSubPath

`func (o *Volume) GetSubPath() string`

GetSubPath returns the SubPath field if non-nil, zero value otherwise.

### GetSubPathOk

`func (o *Volume) GetSubPathOk() (*string, bool)`

GetSubPathOk returns a tuple with the SubPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPath

`func (o *Volume) SetSubPath(v string)`

SetSubPath sets SubPath field to given value.

### HasSubPath

`func (o *Volume) HasSubPath() bool`

HasSubPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



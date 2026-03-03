# FileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Absolute file path | 
**Size** | **int64** | File size in bytes | 
**ModifiedAt** | **time.Time** | Last modification time | 
**CreatedAt** | **time.Time** | File creation time | 
**Owner** | **string** | File owner username | 
**Group** | **string** | File group name | 
**Mode** | **int32** | File permissions in octal format | 

## Methods

### NewFileInfo

`func NewFileInfo(path string, size int64, modifiedAt time.Time, createdAt time.Time, owner string, group string, mode int32, ) *FileInfo`

NewFileInfo instantiates a new FileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoWithDefaults

`func NewFileInfoWithDefaults() *FileInfo`

NewFileInfoWithDefaults instantiates a new FileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FileInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileInfo) SetPath(v string)`

SetPath sets Path field to given value.


### GetSize

`func (o *FileInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileInfo) SetSize(v int64)`

SetSize sets Size field to given value.


### GetModifiedAt

`func (o *FileInfo) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *FileInfo) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *FileInfo) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCreatedAt

`func (o *FileInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetOwner

`func (o *FileInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FileInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FileInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetGroup

`func (o *FileInfo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FileInfo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FileInfo) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetMode

`func (o *FileInfo) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FileInfo) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FileInfo) SetMode(v int32)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



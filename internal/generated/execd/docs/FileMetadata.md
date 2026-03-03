# FileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Target file path | [optional] 
**Owner** | Pointer to **string** | File owner | [optional] 
**Group** | Pointer to **string** | File group | [optional] 
**Mode** | Pointer to **int32** | File permissions in octal | [optional] 

## Methods

### NewFileMetadata

`func NewFileMetadata() *FileMetadata`

NewFileMetadata instantiates a new FileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetadataWithDefaults

`func NewFileMetadataWithDefaults() *FileMetadata`

NewFileMetadataWithDefaults instantiates a new FileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FileMetadata) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileMetadata) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileMetadata) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileMetadata) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetOwner

`func (o *FileMetadata) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FileMetadata) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FileMetadata) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FileMetadata) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroup

`func (o *FileMetadata) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FileMetadata) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FileMetadata) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *FileMetadata) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMode

`func (o *FileMetadata) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FileMetadata) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FileMetadata) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FileMetadata) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



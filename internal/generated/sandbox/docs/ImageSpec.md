# ImageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Container image URI in standard format.  Examples:   - \&quot;python:3.11\&quot; (Docker Hub)   - \&quot;ubuntu:22.04\&quot;   - \&quot;gcr.io/my-project/model-server:v1.0\&quot;   - \&quot;private-registry.company.com:5000/app:latest\&quot;  | 
**Auth** | Pointer to [**ImageSpecAuth**](ImageSpecAuth.md) |  | [optional] 

## Methods

### NewImageSpec

`func NewImageSpec(uri string, ) *ImageSpec`

NewImageSpec instantiates a new ImageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageSpecWithDefaults

`func NewImageSpecWithDefaults() *ImageSpec`

NewImageSpecWithDefaults instantiates a new ImageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *ImageSpec) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ImageSpec) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ImageSpec) SetUri(v string)`

SetUri sets Uri field to given value.


### GetAuth

`func (o *ImageSpec) GetAuth() ImageSpecAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ImageSpec) GetAuthOk() (*ImageSpecAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ImageSpec) SetAuth(v ImageSpecAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ImageSpec) HasAuth() bool`

HasAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



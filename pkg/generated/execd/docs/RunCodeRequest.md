# RunCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**CodeContext**](CodeContext.md) |  | [optional] 
**Code** | **string** | Source code to execute | 

## Methods

### NewRunCodeRequest

`func NewRunCodeRequest(code string, ) *RunCodeRequest`

NewRunCodeRequest instantiates a new RunCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCodeRequestWithDefaults

`func NewRunCodeRequestWithDefaults() *RunCodeRequest`

NewRunCodeRequestWithDefaults instantiates a new RunCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *RunCodeRequest) GetContext() CodeContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RunCodeRequest) GetContextOk() (*CodeContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RunCodeRequest) SetContext(v CodeContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *RunCodeRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCode

`func (o *RunCodeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RunCodeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RunCodeRequest) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



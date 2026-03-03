# CodeContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique session identifier returned by CreateContext | [optional] 
**Language** | **string** | Execution runtime | 

## Methods

### NewCodeContext

`func NewCodeContext(language string, ) *CodeContext`

NewCodeContext instantiates a new CodeContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeContextWithDefaults

`func NewCodeContextWithDefaults() *CodeContext`

NewCodeContextWithDefaults instantiates a new CodeContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CodeContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CodeContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *CodeContext) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CodeContext) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CodeContext) SetLanguage(v string)`

SetLanguage sets Language field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



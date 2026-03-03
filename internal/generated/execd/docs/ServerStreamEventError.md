# ServerStreamEventError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ename** | Pointer to **string** | Error name/type | [optional] 
**Evalue** | Pointer to **string** | Error value/message | [optional] 
**Traceback** | Pointer to **[]string** | Stack trace lines | [optional] 

## Methods

### NewServerStreamEventError

`func NewServerStreamEventError() *ServerStreamEventError`

NewServerStreamEventError instantiates a new ServerStreamEventError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStreamEventErrorWithDefaults

`func NewServerStreamEventErrorWithDefaults() *ServerStreamEventError`

NewServerStreamEventErrorWithDefaults instantiates a new ServerStreamEventError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEname

`func (o *ServerStreamEventError) GetEname() string`

GetEname returns the Ename field if non-nil, zero value otherwise.

### GetEnameOk

`func (o *ServerStreamEventError) GetEnameOk() (*string, bool)`

GetEnameOk returns a tuple with the Ename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEname

`func (o *ServerStreamEventError) SetEname(v string)`

SetEname sets Ename field to given value.

### HasEname

`func (o *ServerStreamEventError) HasEname() bool`

HasEname returns a boolean if a field has been set.

### GetEvalue

`func (o *ServerStreamEventError) GetEvalue() string`

GetEvalue returns the Evalue field if non-nil, zero value otherwise.

### GetEvalueOk

`func (o *ServerStreamEventError) GetEvalueOk() (*string, bool)`

GetEvalueOk returns a tuple with the Evalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalue

`func (o *ServerStreamEventError) SetEvalue(v string)`

SetEvalue sets Evalue field to given value.

### HasEvalue

`func (o *ServerStreamEventError) HasEvalue() bool`

HasEvalue returns a boolean if a field has been set.

### GetTraceback

`func (o *ServerStreamEventError) GetTraceback() []string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *ServerStreamEventError) GetTracebackOk() (*[]string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *ServerStreamEventError) SetTraceback(v []string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *ServerStreamEventError) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



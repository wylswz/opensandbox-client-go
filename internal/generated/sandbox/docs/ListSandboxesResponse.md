# ListSandboxesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Sandbox**](Sandbox.md) |  | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewListSandboxesResponse

`func NewListSandboxesResponse(items []Sandbox, pagination PaginationInfo, ) *ListSandboxesResponse`

NewListSandboxesResponse instantiates a new ListSandboxesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSandboxesResponseWithDefaults

`func NewListSandboxesResponseWithDefaults() *ListSandboxesResponse`

NewListSandboxesResponseWithDefaults instantiates a new ListSandboxesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListSandboxesResponse) GetItems() []Sandbox`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSandboxesResponse) GetItemsOk() (*[]Sandbox, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSandboxesResponse) SetItems(v []Sandbox)`

SetItems sets Items field to given value.


### GetPagination

`func (o *ListSandboxesResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSandboxesResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSandboxesResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



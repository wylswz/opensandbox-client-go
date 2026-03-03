# PaginationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** | Current page number | 
**PageSize** | **int32** | Number of items per page | 
**TotalItems** | **int32** | Total number of items matching the filter | 
**TotalPages** | **int32** | Total number of pages | 
**HasNextPage** | **bool** | Whether there are more pages after the current one | 

## Methods

### NewPaginationInfo

`func NewPaginationInfo(page int32, pageSize int32, totalItems int32, totalPages int32, hasNextPage bool, ) *PaginationInfo`

NewPaginationInfo instantiates a new PaginationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationInfoWithDefaults

`func NewPaginationInfoWithDefaults() *PaginationInfo`

NewPaginationInfoWithDefaults instantiates a new PaginationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PaginationInfo) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginationInfo) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginationInfo) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *PaginationInfo) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginationInfo) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginationInfo) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalItems

`func (o *PaginationInfo) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PaginationInfo) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PaginationInfo) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.


### GetTotalPages

`func (o *PaginationInfo) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginationInfo) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginationInfo) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.


### GetHasNextPage

`func (o *PaginationInfo) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *PaginationInfo) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *PaginationInfo) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



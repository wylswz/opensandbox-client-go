# PVC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimName** | **string** | Name of the volume on the target platform. In Kubernetes this is the PVC name; in Docker this is the named volume name. Must be a valid DNS label.  | 

## Methods

### NewPVC

`func NewPVC(claimName string, ) *PVC`

NewPVC instantiates a new PVC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPVCWithDefaults

`func NewPVCWithDefaults() *PVC`

NewPVCWithDefaults instantiates a new PVC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimName

`func (o *PVC) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *PVC) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *PVC) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



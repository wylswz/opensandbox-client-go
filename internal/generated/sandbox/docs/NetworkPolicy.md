# NetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultAction** | Pointer to **string** | Default action when no egress rule matches. Defaults to \&quot;deny\&quot;. | [optional] 
**Egress** | Pointer to [**[]NetworkRule**](NetworkRule.md) | List of egress rules evaluated in order. | [optional] 

## Methods

### NewNetworkPolicy

`func NewNetworkPolicy() *NetworkPolicy`

NewNetworkPolicy instantiates a new NetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyWithDefaults

`func NewNetworkPolicyWithDefaults() *NetworkPolicy`

NewNetworkPolicyWithDefaults instantiates a new NetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultAction

`func (o *NetworkPolicy) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *NetworkPolicy) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *NetworkPolicy) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *NetworkPolicy) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetEgress

`func (o *NetworkPolicy) GetEgress() []NetworkRule`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *NetworkPolicy) GetEgressOk() (*[]NetworkRule, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *NetworkPolicy) SetEgress(v []NetworkRule)`

SetEgress sets Egress field to given value.

### HasEgress

`func (o *NetworkPolicy) HasEgress() bool`

HasEgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



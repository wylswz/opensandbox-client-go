# NetworkRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Whether to allow or deny matching targets. | 
**Target** | **string** | FQDN or wildcard domain (e.g., \&quot;example.com\&quot;, \&quot;*.example.com\&quot;). IP/CIDR not yet supported in the egress MVP.  | 

## Methods

### NewNetworkRule

`func NewNetworkRule(action string, target string, ) *NetworkRule`

NewNetworkRule instantiates a new NetworkRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRuleWithDefaults

`func NewNetworkRuleWithDefaults() *NetworkRule`

NewNetworkRuleWithDefaults instantiates a new NetworkRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *NetworkRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NetworkRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NetworkRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetTarget

`func (o *NetworkRule) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *NetworkRule) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *NetworkRule) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



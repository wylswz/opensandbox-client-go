# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | **float32** | Number of CPU cores | 
**CpuUsedPct** | **float32** | CPU usage percentage | 
**MemTotalMib** | **float32** | Total memory in MiB | 
**MemUsedMib** | **float32** | Used memory in MiB | 
**Timestamp** | **int64** | Timestamp when metrics were collected (Unix milliseconds) | 

## Methods

### NewMetrics

`func NewMetrics(cpuCount float32, cpuUsedPct float32, memTotalMib float32, memUsedMib float32, timestamp int64, ) *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *Metrics) GetCpuCount() float32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *Metrics) GetCpuCountOk() (*float32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *Metrics) SetCpuCount(v float32)`

SetCpuCount sets CpuCount field to given value.


### GetCpuUsedPct

`func (o *Metrics) GetCpuUsedPct() float32`

GetCpuUsedPct returns the CpuUsedPct field if non-nil, zero value otherwise.

### GetCpuUsedPctOk

`func (o *Metrics) GetCpuUsedPctOk() (*float32, bool)`

GetCpuUsedPctOk returns a tuple with the CpuUsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsedPct

`func (o *Metrics) SetCpuUsedPct(v float32)`

SetCpuUsedPct sets CpuUsedPct field to given value.


### GetMemTotalMib

`func (o *Metrics) GetMemTotalMib() float32`

GetMemTotalMib returns the MemTotalMib field if non-nil, zero value otherwise.

### GetMemTotalMibOk

`func (o *Metrics) GetMemTotalMibOk() (*float32, bool)`

GetMemTotalMibOk returns a tuple with the MemTotalMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemTotalMib

`func (o *Metrics) SetMemTotalMib(v float32)`

SetMemTotalMib sets MemTotalMib field to given value.


### GetMemUsedMib

`func (o *Metrics) GetMemUsedMib() float32`

GetMemUsedMib returns the MemUsedMib field if non-nil, zero value otherwise.

### GetMemUsedMibOk

`func (o *Metrics) GetMemUsedMibOk() (*float32, bool)`

GetMemUsedMibOk returns a tuple with the MemUsedMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsedMib

`func (o *Metrics) SetMemUsedMib(v float32)`

SetMemUsedMib sets MemUsedMib field to given value.


### GetTimestamp

`func (o *Metrics) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Metrics) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Metrics) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



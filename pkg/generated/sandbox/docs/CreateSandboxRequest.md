# CreateSandboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | [**ImageSpec**](ImageSpec.md) | Container image specification for the sandbox | 
**Timeout** | **int32** | Sandbox timeout in seconds. The sandbox will automatically terminate after this duration. SDK clients should provide a default value (e.g., 3600 seconds / 1 hour).  | 
**ResourceLimits** | **map[string]string** | Runtime resource constraints as key-value pairs. Similar to Kubernetes resource specifications, allows flexible definition of resource limits. Common resource types include: - &#x60;cpu&#x60;: CPU allocation in millicores (e.g., \&quot;250m\&quot; for 0.25 CPU cores) - &#x60;memory&#x60;: Memory allocation in bytes or human-readable format (e.g., \&quot;512Mi\&quot;, \&quot;1Gi\&quot;) - &#x60;gpu&#x60;: Number of GPU devices (e.g., \&quot;1\&quot;)  New resource types can be added without API changes.  | 
**Env** | Pointer to **map[string]string** | Environment variables to inject into the sandbox runtime. | [optional] 
**Metadata** | Pointer to **map[string]string** | Custom key-value metadata for management, filtering, and tagging. Use \&quot;name\&quot; key for a human-readable identifier.  | [optional] 
**Entrypoint** | **[]string** | The command to execute as the sandbox&#39;s entry process (required).  Explicitly specifies the user&#39;s expected main process, allowing the sandbox management service to reliably inject control processes before executing this command.  Format: [executable, arg1, arg2, ...]  Examples: - [\&quot;python\&quot;, \&quot;/app/main.py\&quot;] - [\&quot;/bin/bash\&quot;] - [\&quot;java\&quot;, \&quot;-jar\&quot;, \&quot;/app/app.jar\&quot;] - [\&quot;node\&quot;, \&quot;server.js\&quot;]  | 
**NetworkPolicy** | Pointer to [**NetworkPolicy**](NetworkPolicy.md) | Optional outbound network policy for the sandbox. Shape matches the sidecar &#x60;/policy&#x60; endpoint. If omitted or empty, the sidecar starts in allow-all mode until updated.  | [optional] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) | Storage mounts for the sandbox. Each volume entry specifies a named backend-specific storage source and common mount settings. Exactly one backend type must be specified per volume entry.  | [optional] 
**Extensions** | Pointer to **map[string]string** | Opaque container for provider-specific or transient parameters not supported by the core API.  **Note**: This field is reserved for internal features, experimental flags, or temporary behaviors. Standard parameters should be proposed as core API fields.  **Best Practices**: - **Namespacing**: Use prefixed keys (e.g., &#x60;storage.id&#x60;) to prevent collisions. - **Pass-through**: SDKs and middleware must treat this object as opaque and pass it through transparently.  | [optional] 

## Methods

### NewCreateSandboxRequest

`func NewCreateSandboxRequest(image ImageSpec, timeout int32, resourceLimits map[string]string, entrypoint []string, ) *CreateSandboxRequest`

NewCreateSandboxRequest instantiates a new CreateSandboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSandboxRequestWithDefaults

`func NewCreateSandboxRequestWithDefaults() *CreateSandboxRequest`

NewCreateSandboxRequestWithDefaults instantiates a new CreateSandboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *CreateSandboxRequest) GetImage() ImageSpec`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateSandboxRequest) GetImageOk() (*ImageSpec, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateSandboxRequest) SetImage(v ImageSpec)`

SetImage sets Image field to given value.


### GetTimeout

`func (o *CreateSandboxRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreateSandboxRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreateSandboxRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetResourceLimits

`func (o *CreateSandboxRequest) GetResourceLimits() map[string]string`

GetResourceLimits returns the ResourceLimits field if non-nil, zero value otherwise.

### GetResourceLimitsOk

`func (o *CreateSandboxRequest) GetResourceLimitsOk() (*map[string]string, bool)`

GetResourceLimitsOk returns a tuple with the ResourceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLimits

`func (o *CreateSandboxRequest) SetResourceLimits(v map[string]string)`

SetResourceLimits sets ResourceLimits field to given value.


### GetEnv

`func (o *CreateSandboxRequest) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CreateSandboxRequest) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CreateSandboxRequest) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CreateSandboxRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSandboxRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSandboxRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSandboxRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSandboxRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntrypoint

`func (o *CreateSandboxRequest) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *CreateSandboxRequest) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *CreateSandboxRequest) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.


### GetNetworkPolicy

`func (o *CreateSandboxRequest) GetNetworkPolicy() NetworkPolicy`

GetNetworkPolicy returns the NetworkPolicy field if non-nil, zero value otherwise.

### GetNetworkPolicyOk

`func (o *CreateSandboxRequest) GetNetworkPolicyOk() (*NetworkPolicy, bool)`

GetNetworkPolicyOk returns a tuple with the NetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicy

`func (o *CreateSandboxRequest) SetNetworkPolicy(v NetworkPolicy)`

SetNetworkPolicy sets NetworkPolicy field to given value.

### HasNetworkPolicy

`func (o *CreateSandboxRequest) HasNetworkPolicy() bool`

HasNetworkPolicy returns a boolean if a field has been set.

### GetVolumes

`func (o *CreateSandboxRequest) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CreateSandboxRequest) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CreateSandboxRequest) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CreateSandboxRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetExtensions

`func (o *CreateSandboxRequest) GetExtensions() map[string]string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CreateSandboxRequest) GetExtensionsOk() (*map[string]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CreateSandboxRequest) SetExtensions(v map[string]string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CreateSandboxRequest) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



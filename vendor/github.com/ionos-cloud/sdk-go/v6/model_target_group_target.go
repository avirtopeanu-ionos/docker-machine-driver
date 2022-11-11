/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// TargetGroupTarget struct for TargetGroupTarget
type TargetGroupTarget struct {
	// The IP of the balanced target VM.
	Ip *string `json:"ip"`
	// The port of the balanced target service; valid range is 1 to 65535.
	Port *int32 `json:"port"`
	// Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments.
	Weight *int32 `json:"weight"`
	// When the health check is enabled, the target is available only when it accepts regular TCP or HTTP connection attempts for state checking. The state check consists of one connection attempt with the target's address and port. The default value is 'TRUE'.
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty"`
	// When the maintenance mode is enabled, the target is prevented from receiving traffic; the default value is 'FALSE'.
	MaintenanceEnabled *bool `json:"maintenanceEnabled,omitempty"`
}

// NewTargetGroupTarget instantiates a new TargetGroupTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetGroupTarget(ip string, port int32, weight int32) *TargetGroupTarget {
	this := TargetGroupTarget{}

	this.Ip = &ip
	this.Port = &port
	this.Weight = &weight

	return &this
}

// NewTargetGroupTargetWithDefaults instantiates a new TargetGroupTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetGroupTargetWithDefaults() *TargetGroupTarget {
	this := TargetGroupTarget{}
	return &this
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TargetGroupTarget) GetIp() *string {
	if o == nil {
		return nil
	}

	return o.Ip

}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupTarget) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Ip, true
}

// SetIp sets field value
func (o *TargetGroupTarget) SetIp(v string) {

	o.Ip = &v

}

// HasIp returns a boolean if a field has been set.
func (o *TargetGroupTarget) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// GetPort returns the Port field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TargetGroupTarget) GetPort() *int32 {
	if o == nil {
		return nil
	}

	return o.Port

}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupTarget) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Port, true
}

// SetPort sets field value
func (o *TargetGroupTarget) SetPort(v int32) {

	o.Port = &v

}

// HasPort returns a boolean if a field has been set.
func (o *TargetGroupTarget) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// GetWeight returns the Weight field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TargetGroupTarget) GetWeight() *int32 {
	if o == nil {
		return nil
	}

	return o.Weight

}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupTarget) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Weight, true
}

// SetWeight sets field value
func (o *TargetGroupTarget) SetWeight(v int32) {

	o.Weight = &v

}

// HasWeight returns a boolean if a field has been set.
func (o *TargetGroupTarget) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// GetHealthCheckEnabled returns the HealthCheckEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *TargetGroupTarget) GetHealthCheckEnabled() *bool {
	if o == nil {
		return nil
	}

	return o.HealthCheckEnabled

}

// GetHealthCheckEnabledOk returns a tuple with the HealthCheckEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupTarget) GetHealthCheckEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.HealthCheckEnabled, true
}

// SetHealthCheckEnabled sets field value
func (o *TargetGroupTarget) SetHealthCheckEnabled(v bool) {

	o.HealthCheckEnabled = &v

}

// HasHealthCheckEnabled returns a boolean if a field has been set.
func (o *TargetGroupTarget) HasHealthCheckEnabled() bool {
	if o != nil && o.HealthCheckEnabled != nil {
		return true
	}

	return false
}

// GetMaintenanceEnabled returns the MaintenanceEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *TargetGroupTarget) GetMaintenanceEnabled() *bool {
	if o == nil {
		return nil
	}

	return o.MaintenanceEnabled

}

// GetMaintenanceEnabledOk returns a tuple with the MaintenanceEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupTarget) GetMaintenanceEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.MaintenanceEnabled, true
}

// SetMaintenanceEnabled sets field value
func (o *TargetGroupTarget) SetMaintenanceEnabled(v bool) {

	o.MaintenanceEnabled = &v

}

// HasMaintenanceEnabled returns a boolean if a field has been set.
func (o *TargetGroupTarget) HasMaintenanceEnabled() bool {
	if o != nil && o.MaintenanceEnabled != nil {
		return true
	}

	return false
}

func (o TargetGroupTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.HealthCheckEnabled != nil {
		toSerialize["healthCheckEnabled"] = o.HealthCheckEnabled
	}
	if o.MaintenanceEnabled != nil {
		toSerialize["maintenanceEnabled"] = o.MaintenanceEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableTargetGroupTarget struct {
	value *TargetGroupTarget
	isSet bool
}

func (v NullableTargetGroupTarget) Get() *TargetGroupTarget {
	return v.value
}

func (v *NullableTargetGroupTarget) Set(val *TargetGroupTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetGroupTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetGroupTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetGroupTarget(val *TargetGroupTarget) *NullableTargetGroupTarget {
	return &NullableTargetGroupTarget{value: val, isSet: true}
}

func (v NullableTargetGroupTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetGroupTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

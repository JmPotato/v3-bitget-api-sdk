/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MarginIsolatedFinFlowResult struct for MarginIsolatedFinFlowResult
type MarginIsolatedFinFlowResult struct {
	MaxId                *string                     `json:"maxId,omitempty"`
	MinId                *string                     `json:"minId,omitempty"`
	ResultList           []MarginIsolatedFinFlowInfo `json:"resultList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginIsolatedFinFlowResult MarginIsolatedFinFlowResult

// NewMarginIsolatedFinFlowResult instantiates a new MarginIsolatedFinFlowResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginIsolatedFinFlowResult() *MarginIsolatedFinFlowResult {
	this := MarginIsolatedFinFlowResult{}
	return &this
}

// NewMarginIsolatedFinFlowResultWithDefaults instantiates a new MarginIsolatedFinFlowResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginIsolatedFinFlowResultWithDefaults() *MarginIsolatedFinFlowResult {
	this := MarginIsolatedFinFlowResult{}
	return &this
}

// GetMaxId returns the MaxId field value if set, zero value otherwise.
func (o *MarginIsolatedFinFlowResult) GetMaxId() string {
	if o == nil || isNil(o.MaxId) {
		var ret string
		return ret
	}
	return *o.MaxId
}

// GetMaxIdOk returns a tuple with the MaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedFinFlowResult) GetMaxIdOk() (*string, bool) {
	if o == nil || isNil(o.MaxId) {
		return nil, false
	}
	return o.MaxId, true
}

// HasMaxId returns a boolean if a field has been set.
func (o *MarginIsolatedFinFlowResult) HasMaxId() bool {
	if o != nil && !isNil(o.MaxId) {
		return true
	}

	return false
}

// SetMaxId gets a reference to the given string and assigns it to the MaxId field.
func (o *MarginIsolatedFinFlowResult) SetMaxId(v string) {
	o.MaxId = &v
}

// GetMinId returns the MinId field value if set, zero value otherwise.
func (o *MarginIsolatedFinFlowResult) GetMinId() string {
	if o == nil || isNil(o.MinId) {
		var ret string
		return ret
	}
	return *o.MinId
}

// GetMinIdOk returns a tuple with the MinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedFinFlowResult) GetMinIdOk() (*string, bool) {
	if o == nil || isNil(o.MinId) {
		return nil, false
	}
	return o.MinId, true
}

// HasMinId returns a boolean if a field has been set.
func (o *MarginIsolatedFinFlowResult) HasMinId() bool {
	if o != nil && !isNil(o.MinId) {
		return true
	}

	return false
}

// SetMinId gets a reference to the given string and assigns it to the MinId field.
func (o *MarginIsolatedFinFlowResult) SetMinId(v string) {
	o.MinId = &v
}

// GetResultList returns the ResultList field value if set, zero value otherwise.
func (o *MarginIsolatedFinFlowResult) GetResultList() []MarginIsolatedFinFlowInfo {
	if o == nil || isNil(o.ResultList) {
		var ret []MarginIsolatedFinFlowInfo
		return ret
	}
	return o.ResultList
}

// GetResultListOk returns a tuple with the ResultList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedFinFlowResult) GetResultListOk() ([]MarginIsolatedFinFlowInfo, bool) {
	if o == nil || isNil(o.ResultList) {
		return nil, false
	}
	return o.ResultList, true
}

// HasResultList returns a boolean if a field has been set.
func (o *MarginIsolatedFinFlowResult) HasResultList() bool {
	if o != nil && !isNil(o.ResultList) {
		return true
	}

	return false
}

// SetResultList gets a reference to the given []MarginIsolatedFinFlowInfo and assigns it to the ResultList field.
func (o *MarginIsolatedFinFlowResult) SetResultList(v []MarginIsolatedFinFlowInfo) {
	o.ResultList = v
}

func (o MarginIsolatedFinFlowResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxId) {
		toSerialize["maxId"] = o.MaxId
	}
	if !isNil(o.MinId) {
		toSerialize["minId"] = o.MinId
	}
	if !isNil(o.ResultList) {
		toSerialize["resultList"] = o.ResultList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginIsolatedFinFlowResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginIsolatedFinFlowResult := _MarginIsolatedFinFlowResult{}

	if err = json.Unmarshal(bytes, &varMarginIsolatedFinFlowResult); err == nil {
		*o = MarginIsolatedFinFlowResult(varMarginIsolatedFinFlowResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "maxId")
		delete(additionalProperties, "minId")
		delete(additionalProperties, "resultList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginIsolatedFinFlowResult struct {
	value *MarginIsolatedFinFlowResult
	isSet bool
}

func (v NullableMarginIsolatedFinFlowResult) Get() *MarginIsolatedFinFlowResult {
	return v.value
}

func (v *NullableMarginIsolatedFinFlowResult) Set(val *MarginIsolatedFinFlowResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginIsolatedFinFlowResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginIsolatedFinFlowResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginIsolatedFinFlowResult(val *MarginIsolatedFinFlowResult) *NullableMarginIsolatedFinFlowResult {
	return &NullableMarginIsolatedFinFlowResult{value: val, isSet: true}
}

func (v NullableMarginIsolatedFinFlowResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginIsolatedFinFlowResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
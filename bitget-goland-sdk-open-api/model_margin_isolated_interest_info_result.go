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

// MarginIsolatedInterestInfoResult struct for MarginIsolatedInterestInfoResult
type MarginIsolatedInterestInfoResult struct {
	MaxId                *string                      `json:"maxId,omitempty"`
	MinId                *string                      `json:"minId,omitempty"`
	ResultList           []MarginIsolatedInterestInfo `json:"resultList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginIsolatedInterestInfoResult MarginIsolatedInterestInfoResult

// NewMarginIsolatedInterestInfoResult instantiates a new MarginIsolatedInterestInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginIsolatedInterestInfoResult() *MarginIsolatedInterestInfoResult {
	this := MarginIsolatedInterestInfoResult{}
	return &this
}

// NewMarginIsolatedInterestInfoResultWithDefaults instantiates a new MarginIsolatedInterestInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginIsolatedInterestInfoResultWithDefaults() *MarginIsolatedInterestInfoResult {
	this := MarginIsolatedInterestInfoResult{}
	return &this
}

// GetMaxId returns the MaxId field value if set, zero value otherwise.
func (o *MarginIsolatedInterestInfoResult) GetMaxId() string {
	if o == nil || isNil(o.MaxId) {
		var ret string
		return ret
	}
	return *o.MaxId
}

// GetMaxIdOk returns a tuple with the MaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedInterestInfoResult) GetMaxIdOk() (*string, bool) {
	if o == nil || isNil(o.MaxId) {
		return nil, false
	}
	return o.MaxId, true
}

// HasMaxId returns a boolean if a field has been set.
func (o *MarginIsolatedInterestInfoResult) HasMaxId() bool {
	if o != nil && !isNil(o.MaxId) {
		return true
	}

	return false
}

// SetMaxId gets a reference to the given string and assigns it to the MaxId field.
func (o *MarginIsolatedInterestInfoResult) SetMaxId(v string) {
	o.MaxId = &v
}

// GetMinId returns the MinId field value if set, zero value otherwise.
func (o *MarginIsolatedInterestInfoResult) GetMinId() string {
	if o == nil || isNil(o.MinId) {
		var ret string
		return ret
	}
	return *o.MinId
}

// GetMinIdOk returns a tuple with the MinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedInterestInfoResult) GetMinIdOk() (*string, bool) {
	if o == nil || isNil(o.MinId) {
		return nil, false
	}
	return o.MinId, true
}

// HasMinId returns a boolean if a field has been set.
func (o *MarginIsolatedInterestInfoResult) HasMinId() bool {
	if o != nil && !isNil(o.MinId) {
		return true
	}

	return false
}

// SetMinId gets a reference to the given string and assigns it to the MinId field.
func (o *MarginIsolatedInterestInfoResult) SetMinId(v string) {
	o.MinId = &v
}

// GetResultList returns the ResultList field value if set, zero value otherwise.
func (o *MarginIsolatedInterestInfoResult) GetResultList() []MarginIsolatedInterestInfo {
	if o == nil || isNil(o.ResultList) {
		var ret []MarginIsolatedInterestInfo
		return ret
	}
	return o.ResultList
}

// GetResultListOk returns a tuple with the ResultList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedInterestInfoResult) GetResultListOk() ([]MarginIsolatedInterestInfo, bool) {
	if o == nil || isNil(o.ResultList) {
		return nil, false
	}
	return o.ResultList, true
}

// HasResultList returns a boolean if a field has been set.
func (o *MarginIsolatedInterestInfoResult) HasResultList() bool {
	if o != nil && !isNil(o.ResultList) {
		return true
	}

	return false
}

// SetResultList gets a reference to the given []MarginIsolatedInterestInfo and assigns it to the ResultList field.
func (o *MarginIsolatedInterestInfoResult) SetResultList(v []MarginIsolatedInterestInfo) {
	o.ResultList = v
}

func (o MarginIsolatedInterestInfoResult) MarshalJSON() ([]byte, error) {
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

func (o *MarginIsolatedInterestInfoResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginIsolatedInterestInfoResult := _MarginIsolatedInterestInfoResult{}

	if err = json.Unmarshal(bytes, &varMarginIsolatedInterestInfoResult); err == nil {
		*o = MarginIsolatedInterestInfoResult(varMarginIsolatedInterestInfoResult)
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

type NullableMarginIsolatedInterestInfoResult struct {
	value *MarginIsolatedInterestInfoResult
	isSet bool
}

func (v NullableMarginIsolatedInterestInfoResult) Get() *MarginIsolatedInterestInfoResult {
	return v.value
}

func (v *NullableMarginIsolatedInterestInfoResult) Set(val *MarginIsolatedInterestInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginIsolatedInterestInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginIsolatedInterestInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginIsolatedInterestInfoResult(val *MarginIsolatedInterestInfoResult) *NullableMarginIsolatedInterestInfoResult {
	return &NullableMarginIsolatedInterestInfoResult{value: val, isSet: true}
}

func (v NullableMarginIsolatedInterestInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginIsolatedInterestInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
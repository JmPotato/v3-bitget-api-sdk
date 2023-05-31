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

// MarginCrossRepayResult struct for MarginCrossRepayResult
type MarginCrossRepayResult struct {
	ClientOid            *string `json:"clientOid,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	RemainDebtAmount     *string `json:"remainDebtAmount,omitempty"`
	RepayAmount          *string `json:"repayAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginCrossRepayResult MarginCrossRepayResult

// NewMarginCrossRepayResult instantiates a new MarginCrossRepayResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCrossRepayResult() *MarginCrossRepayResult {
	this := MarginCrossRepayResult{}
	return &this
}

// NewMarginCrossRepayResultWithDefaults instantiates a new MarginCrossRepayResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCrossRepayResultWithDefaults() *MarginCrossRepayResult {
	this := MarginCrossRepayResult{}
	return &this
}

// GetClientOid returns the ClientOid field value if set, zero value otherwise.
func (o *MarginCrossRepayResult) GetClientOid() string {
	if o == nil || isNil(o.ClientOid) {
		var ret string
		return ret
	}
	return *o.ClientOid
}

// GetClientOidOk returns a tuple with the ClientOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRepayResult) GetClientOidOk() (*string, bool) {
	if o == nil || isNil(o.ClientOid) {
		return nil, false
	}
	return o.ClientOid, true
}

// HasClientOid returns a boolean if a field has been set.
func (o *MarginCrossRepayResult) HasClientOid() bool {
	if o != nil && !isNil(o.ClientOid) {
		return true
	}

	return false
}

// SetClientOid gets a reference to the given string and assigns it to the ClientOid field.
func (o *MarginCrossRepayResult) SetClientOid(v string) {
	o.ClientOid = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *MarginCrossRepayResult) GetCoin() string {
	if o == nil || isNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRepayResult) GetCoinOk() (*string, bool) {
	if o == nil || isNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *MarginCrossRepayResult) HasCoin() bool {
	if o != nil && !isNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *MarginCrossRepayResult) SetCoin(v string) {
	o.Coin = &v
}

// GetRemainDebtAmount returns the RemainDebtAmount field value if set, zero value otherwise.
func (o *MarginCrossRepayResult) GetRemainDebtAmount() string {
	if o == nil || isNil(o.RemainDebtAmount) {
		var ret string
		return ret
	}
	return *o.RemainDebtAmount
}

// GetRemainDebtAmountOk returns a tuple with the RemainDebtAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRepayResult) GetRemainDebtAmountOk() (*string, bool) {
	if o == nil || isNil(o.RemainDebtAmount) {
		return nil, false
	}
	return o.RemainDebtAmount, true
}

// HasRemainDebtAmount returns a boolean if a field has been set.
func (o *MarginCrossRepayResult) HasRemainDebtAmount() bool {
	if o != nil && !isNil(o.RemainDebtAmount) {
		return true
	}

	return false
}

// SetRemainDebtAmount gets a reference to the given string and assigns it to the RemainDebtAmount field.
func (o *MarginCrossRepayResult) SetRemainDebtAmount(v string) {
	o.RemainDebtAmount = &v
}

// GetRepayAmount returns the RepayAmount field value if set, zero value otherwise.
func (o *MarginCrossRepayResult) GetRepayAmount() string {
	if o == nil || isNil(o.RepayAmount) {
		var ret string
		return ret
	}
	return *o.RepayAmount
}

// GetRepayAmountOk returns a tuple with the RepayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRepayResult) GetRepayAmountOk() (*string, bool) {
	if o == nil || isNil(o.RepayAmount) {
		return nil, false
	}
	return o.RepayAmount, true
}

// HasRepayAmount returns a boolean if a field has been set.
func (o *MarginCrossRepayResult) HasRepayAmount() bool {
	if o != nil && !isNil(o.RepayAmount) {
		return true
	}

	return false
}

// SetRepayAmount gets a reference to the given string and assigns it to the RepayAmount field.
func (o *MarginCrossRepayResult) SetRepayAmount(v string) {
	o.RepayAmount = &v
}

func (o MarginCrossRepayResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientOid) {
		toSerialize["clientOid"] = o.ClientOid
	}
	if !isNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !isNil(o.RemainDebtAmount) {
		toSerialize["remainDebtAmount"] = o.RemainDebtAmount
	}
	if !isNil(o.RepayAmount) {
		toSerialize["repayAmount"] = o.RepayAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginCrossRepayResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginCrossRepayResult := _MarginCrossRepayResult{}

	if err = json.Unmarshal(bytes, &varMarginCrossRepayResult); err == nil {
		*o = MarginCrossRepayResult(varMarginCrossRepayResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "clientOid")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "remainDebtAmount")
		delete(additionalProperties, "repayAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginCrossRepayResult struct {
	value *MarginCrossRepayResult
	isSet bool
}

func (v NullableMarginCrossRepayResult) Get() *MarginCrossRepayResult {
	return v.value
}

func (v *NullableMarginCrossRepayResult) Set(val *MarginCrossRepayResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCrossRepayResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCrossRepayResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCrossRepayResult(val *MarginCrossRepayResult) *NullableMarginCrossRepayResult {
	return &NullableMarginCrossRepayResult{value: val, isSet: true}
}

func (v NullableMarginCrossRepayResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCrossRepayResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
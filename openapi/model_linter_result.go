/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LinterResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinterResult{}

// LinterResult struct for LinterResult
type LinterResult struct {
	MinimumScore *int32               `json:"minimumScore,omitempty"`
	Passed       *bool                `json:"passed,omitempty"`
	Score        *int32               `json:"score,omitempty"`
	Plugins      []LinterResultPlugin `json:"plugins,omitempty"`
}

// NewLinterResult instantiates a new LinterResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinterResult() *LinterResult {
	this := LinterResult{}
	return &this
}

// NewLinterResultWithDefaults instantiates a new LinterResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinterResultWithDefaults() *LinterResult {
	this := LinterResult{}
	return &this
}

// GetMinimumScore returns the MinimumScore field value if set, zero value otherwise.
func (o *LinterResult) GetMinimumScore() int32 {
	if o == nil || isNil(o.MinimumScore) {
		var ret int32
		return ret
	}
	return *o.MinimumScore
}

// GetMinimumScoreOk returns a tuple with the MinimumScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResult) GetMinimumScoreOk() (*int32, bool) {
	if o == nil || isNil(o.MinimumScore) {
		return nil, false
	}
	return o.MinimumScore, true
}

// HasMinimumScore returns a boolean if a field has been set.
func (o *LinterResult) HasMinimumScore() bool {
	if o != nil && !isNil(o.MinimumScore) {
		return true
	}

	return false
}

// SetMinimumScore gets a reference to the given int32 and assigns it to the MinimumScore field.
func (o *LinterResult) SetMinimumScore(v int32) {
	o.MinimumScore = &v
}

// GetPassed returns the Passed field value if set, zero value otherwise.
func (o *LinterResult) GetPassed() bool {
	if o == nil || isNil(o.Passed) {
		var ret bool
		return ret
	}
	return *o.Passed
}

// GetPassedOk returns a tuple with the Passed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResult) GetPassedOk() (*bool, bool) {
	if o == nil || isNil(o.Passed) {
		return nil, false
	}
	return o.Passed, true
}

// HasPassed returns a boolean if a field has been set.
func (o *LinterResult) HasPassed() bool {
	if o != nil && !isNil(o.Passed) {
		return true
	}

	return false
}

// SetPassed gets a reference to the given bool and assigns it to the Passed field.
func (o *LinterResult) SetPassed(v bool) {
	o.Passed = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *LinterResult) GetScore() int32 {
	if o == nil || isNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResult) GetScoreOk() (*int32, bool) {
	if o == nil || isNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *LinterResult) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *LinterResult) SetScore(v int32) {
	o.Score = &v
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *LinterResult) GetPlugins() []LinterResultPlugin {
	if o == nil || isNil(o.Plugins) {
		var ret []LinterResultPlugin
		return ret
	}
	return o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResult) GetPluginsOk() ([]LinterResultPlugin, bool) {
	if o == nil || isNil(o.Plugins) {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *LinterResult) HasPlugins() bool {
	if o != nil && !isNil(o.Plugins) {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given []LinterResultPlugin and assigns it to the Plugins field.
func (o *LinterResult) SetPlugins(v []LinterResultPlugin) {
	o.Plugins = v
}

func (o LinterResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinterResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinimumScore) {
		toSerialize["minimumScore"] = o.MinimumScore
	}
	if !isNil(o.Passed) {
		toSerialize["passed"] = o.Passed
	}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !isNil(o.Plugins) {
		toSerialize["plugins"] = o.Plugins
	}
	return toSerialize, nil
}

type NullableLinterResult struct {
	value *LinterResult
	isSet bool
}

func (v NullableLinterResult) Get() *LinterResult {
	return v.value
}

func (v *NullableLinterResult) Set(val *LinterResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLinterResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLinterResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinterResult(val *LinterResult) *NullableLinterResult {
	return &NullableLinterResult{value: val, isSet: true}
}

func (v NullableLinterResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinterResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
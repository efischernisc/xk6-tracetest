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

// checks if the VariableSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableSet{}

// VariableSet struct for VariableSet
type VariableSet struct {
	Id          *string            `json:"id,omitempty"`
	Name        *string            `json:"name,omitempty"`
	Description *string            `json:"description,omitempty"`
	Values      []VariableSetValue `json:"values,omitempty"`
}

// NewVariableSet instantiates a new VariableSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableSet() *VariableSet {
	this := VariableSet{}
	return &this
}

// NewVariableSetWithDefaults instantiates a new VariableSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableSetWithDefaults() *VariableSet {
	this := VariableSet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VariableSet) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableSet) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VariableSet) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VariableSet) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VariableSet) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableSet) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VariableSet) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VariableSet) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VariableSet) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableSet) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VariableSet) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VariableSet) SetDescription(v string) {
	o.Description = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *VariableSet) GetValues() []VariableSetValue {
	if o == nil || isNil(o.Values) {
		var ret []VariableSetValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableSet) GetValuesOk() ([]VariableSetValue, bool) {
	if o == nil || isNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *VariableSet) HasValues() bool {
	if o != nil && !isNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []VariableSetValue and assigns it to the Values field.
func (o *VariableSet) SetValues(v []VariableSetValue) {
	o.Values = v
}

func (o VariableSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableVariableSet struct {
	value *VariableSet
	isSet bool
}

func (v NullableVariableSet) Get() *VariableSet {
	return v.value
}

func (v *NullableVariableSet) Set(val *VariableSet) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableSet) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableSet(val *VariableSet) *NullableVariableSet {
	return &NullableVariableSet{value: val, isSet: true}
}

func (v NullableVariableSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
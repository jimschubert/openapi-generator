/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"encoding/json"
)

// MapTest struct for MapTest
type MapTest struct {
	MapMapOfString *map[string]map[string]string `json:"map_map_of_string,omitempty"`
	MapOfEnumString *map[string]MapTestMapOfEnumStringAddlProps `json:"map_of_enum_string,omitempty"`
	DirectMap *map[string]bool `json:"direct_map,omitempty"`
	IndirectMap *map[string]bool `json:"indirect_map,omitempty"`
}

// NewMapTest instantiates a new MapTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapTest() *MapTest {
	this := MapTest{}
	return &this
}

// NewMapTestWithDefaults instantiates a new MapTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapTestWithDefaults() *MapTest {
	this := MapTest{}
	return &this
}

// GetMapMapOfString returns the MapMapOfString field value if set, zero value otherwise.
func (o *MapTest) GetMapMapOfString() map[string]map[string]string {
	if o == nil || o.MapMapOfString == nil {
		var ret map[string]map[string]string
		return ret
	}
	return *o.MapMapOfString
}

// GetMapMapOfStringOk returns a tuple with the MapMapOfString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapTest) GetMapMapOfStringOk() (*map[string]map[string]string, bool) {
	if o == nil || o.MapMapOfString == nil {
		return nil, false
	}
	return o.MapMapOfString, true
}

// HasMapMapOfString returns a boolean if a field has been set.
func (o *MapTest) HasMapMapOfString() bool {
	if o != nil && o.MapMapOfString != nil {
		return true
	}

	return false
}

// SetMapMapOfString gets a reference to the given map[string]map[string]string and assigns it to the MapMapOfString field.
func (o *MapTest) SetMapMapOfString(v map[string]map[string]string) {
	o.MapMapOfString = &v
}

// GetMapOfEnumString returns the MapOfEnumString field value if set, zero value otherwise.
func (o *MapTest) GetMapOfEnumString() map[string]MapTestMapOfEnumStringAddlProps {
	if o == nil || o.MapOfEnumString == nil {
		var ret map[string]MapTestMapOfEnumStringAddlProps
		return ret
	}
	return *o.MapOfEnumString
}

// GetMapOfEnumStringOk returns a tuple with the MapOfEnumString field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *MapTest) GetMapOfEnumStringOk() (map[string]MapTestMapOfEnumStringAddlProps, bool) {
	if o == nil || o.MapOfEnumString == nil {
		var ret map[string]MapTestMapOfEnumStringAddlProps
		return ret, false
=======
func (o *MapTest) GetMapOfEnumStringOk() (*map[string]string, bool) {
	if o == nil || o.MapOfEnumString == nil {
		return nil, false
>>>>>>> origin/master
	}
	return o.MapOfEnumString, true
}

// HasMapOfEnumString returns a boolean if a field has been set.
func (o *MapTest) HasMapOfEnumString() bool {
	if o != nil && o.MapOfEnumString != nil {
		return true
	}

	return false
}

// SetMapOfEnumString gets a reference to the given map[string]MapTestMapOfEnumStringAddlProps and assigns it to the MapOfEnumString field.
func (o *MapTest) SetMapOfEnumString(v map[string]MapTestMapOfEnumStringAddlProps) {
	o.MapOfEnumString = &v
}

// GetDirectMap returns the DirectMap field value if set, zero value otherwise.
func (o *MapTest) GetDirectMap() map[string]bool {
	if o == nil || o.DirectMap == nil {
		var ret map[string]bool
		return ret
	}
	return *o.DirectMap
}

// GetDirectMapOk returns a tuple with the DirectMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapTest) GetDirectMapOk() (*map[string]bool, bool) {
	if o == nil || o.DirectMap == nil {
		return nil, false
	}
	return o.DirectMap, true
}

// HasDirectMap returns a boolean if a field has been set.
func (o *MapTest) HasDirectMap() bool {
	if o != nil && o.DirectMap != nil {
		return true
	}

	return false
}

// SetDirectMap gets a reference to the given map[string]bool and assigns it to the DirectMap field.
func (o *MapTest) SetDirectMap(v map[string]bool) {
	o.DirectMap = &v
}

// GetIndirectMap returns the IndirectMap field value if set, zero value otherwise.
func (o *MapTest) GetIndirectMap() map[string]bool {
	if o == nil || o.IndirectMap == nil {
		var ret map[string]bool
		return ret
	}
	return *o.IndirectMap
}

// GetIndirectMapOk returns a tuple with the IndirectMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapTest) GetIndirectMapOk() (*map[string]bool, bool) {
	if o == nil || o.IndirectMap == nil {
		return nil, false
	}
	return o.IndirectMap, true
}

// HasIndirectMap returns a boolean if a field has been set.
func (o *MapTest) HasIndirectMap() bool {
	if o != nil && o.IndirectMap != nil {
		return true
	}

	return false
}

// SetIndirectMap gets a reference to the given map[string]bool and assigns it to the IndirectMap field.
func (o *MapTest) SetIndirectMap(v map[string]bool) {
	o.IndirectMap = &v
}

func (o MapTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MapMapOfString != nil {
		toSerialize["map_map_of_string"] = o.MapMapOfString
	}
	if o.MapOfEnumString != nil {
		toSerialize["map_of_enum_string"] = o.MapOfEnumString
	}
	if o.DirectMap != nil {
		toSerialize["direct_map"] = o.DirectMap
	}
	if o.IndirectMap != nil {
		toSerialize["indirect_map"] = o.IndirectMap
	}
	return json.Marshal(toSerialize)
}

type NullableMapTest struct {
	value *MapTest
	isSet bool
}

func (v NullableMapTest) Get() *MapTest {
	return v.value
}

func (v *NullableMapTest) Set(val *MapTest) {
	v.value = val
	v.isSet = true
}

func (v NullableMapTest) IsSet() bool {
	return v.isSet
}

func (v *NullableMapTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapTest(val *MapTest) *NullableMapTest {
	return &NullableMapTest{value: val, isSet: true}
}

func (v NullableMapTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

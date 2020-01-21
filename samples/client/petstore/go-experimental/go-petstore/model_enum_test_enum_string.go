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
	"bytes"
	"encoding/json"
)

// EnumTestEnumString the model 'EnumTestEnumString'
type EnumTestEnumString string

// List of Enum_TestEnumString
const (
	ENUM_TEST_ENUM_STRING_UPPER EnumTestEnumString = "UPPER"
	ENUM_TEST_ENUM_STRING_LOWER EnumTestEnumString = "lower"
	ENUM_TEST_ENUM_STRING_EMPTY EnumTestEnumString = ""
)

type NullableEnumTestEnumString struct {
	Value EnumTestEnumString
	ExplicitNull bool
}

func (v NullableEnumTestEnumString) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEnumTestEnumString) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

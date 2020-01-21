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

// OuterEnum the model 'OuterEnum'
type OuterEnum string

// List of OuterEnum
const (
	OUTER_ENUM_PLACED OuterEnum = "placed"
	OUTER_ENUM_APPROVED OuterEnum = "approved"
	OUTER_ENUM_DELIVERED OuterEnum = "delivered"
)

type NullableOuterEnum struct {
	Value OuterEnum
	ExplicitNull bool
}

func (v NullableOuterEnum) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOuterEnum) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

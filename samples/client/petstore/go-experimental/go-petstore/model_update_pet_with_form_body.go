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

// UpdatePetWithFormBody struct for UpdatePetWithFormBody
type UpdatePetWithFormBody struct {
	// Updated name of the pet
	Name *string `json:"name,omitempty"`
	// Updated status of the pet
	Status *string `json:"status,omitempty"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePetWithFormBody) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePetWithFormBody) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePetWithFormBody) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePetWithFormBody) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdatePetWithFormBody) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePetWithFormBody) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdatePetWithFormBody) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdatePetWithFormBody) SetStatus(v string) {
	o.Status = &v
}

type NullableUpdatePetWithFormBody struct {
	Value UpdatePetWithFormBody
	ExplicitNull bool
}

func (v NullableUpdatePetWithFormBody) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdatePetWithFormBody) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

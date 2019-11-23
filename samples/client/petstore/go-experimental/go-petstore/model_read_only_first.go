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

// ReadOnlyFirst struct for ReadOnlyFirst
type ReadOnlyFirst struct {
	Bar *string `json:"bar,omitempty"`
	Baz *string `json:"baz,omitempty"`
}

// GetBar returns the Bar field value if set, zero value otherwise.
func (o *ReadOnlyFirst) GetBar() string {
	if o == nil || o.Bar == nil {
		var ret string
		return ret
	}
	return *o.Bar
}

// GetBarOk returns a tuple with the Bar field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyFirst) GetBarOk() (string, bool) {
	if o == nil || o.Bar == nil {
		var ret string
		return ret, false
	}
	return *o.Bar, true
}

// HasBar returns a boolean if a field has been set.
func (o *ReadOnlyFirst) HasBar() bool {
	if o != nil && o.Bar != nil {
		return true
	}

	return false
}

// SetBar gets a reference to the given string and assigns it to the Bar field.
func (o *ReadOnlyFirst) SetBar(v string) {
	o.Bar = &v
}

// GetBaz returns the Baz field value if set, zero value otherwise.
func (o *ReadOnlyFirst) GetBaz() string {
	if o == nil || o.Baz == nil {
		var ret string
		return ret
	}
	return *o.Baz
}

// GetBazOk returns a tuple with the Baz field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyFirst) GetBazOk() (string, bool) {
	if o == nil || o.Baz == nil {
		var ret string
		return ret, false
	}
	return *o.Baz, true
}

// HasBaz returns a boolean if a field has been set.
func (o *ReadOnlyFirst) HasBaz() bool {
	if o != nil && o.Baz != nil {
		return true
	}

	return false
}

// SetBaz gets a reference to the given string and assigns it to the Baz field.
func (o *ReadOnlyFirst) SetBaz(v string) {
	o.Baz = &v
}

type NullableReadOnlyFirst struct {
	Value ReadOnlyFirst
	ExplicitNull bool
}

func (v NullableReadOnlyFirst) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadOnlyFirst) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}


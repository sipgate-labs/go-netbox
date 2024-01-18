/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.1 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// PatchedWritableRackRequestWidth Rail-to-rail width  * `10` - 10 inches * `19` - 19 inches * `21` - 21 inches * `23` - 23 inches
type PatchedWritableRackRequestWidth int32

// List of PatchedWritableRackRequest_width
const (
	PATCHEDWRITABLERACKREQUESTWIDTH__10 PatchedWritableRackRequestWidth = 10
	PATCHEDWRITABLERACKREQUESTWIDTH__19 PatchedWritableRackRequestWidth = 19
	PATCHEDWRITABLERACKREQUESTWIDTH__21 PatchedWritableRackRequestWidth = 21
	PATCHEDWRITABLERACKREQUESTWIDTH__23 PatchedWritableRackRequestWidth = 23
)

// All allowed values of PatchedWritableRackRequestWidth enum
var AllowedPatchedWritableRackRequestWidthEnumValues = []PatchedWritableRackRequestWidth{
	10,
	19,
	21,
	23,
}

func (v *PatchedWritableRackRequestWidth) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableRackRequestWidth(value)
	for _, existing := range AllowedPatchedWritableRackRequestWidthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableRackRequestWidth", value)
}

// NewPatchedWritableRackRequestWidthFromValue returns a pointer to a valid PatchedWritableRackRequestWidth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableRackRequestWidthFromValue(v int32) (*PatchedWritableRackRequestWidth, error) {
	ev := PatchedWritableRackRequestWidth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableRackRequestWidth: valid values are %v", v, AllowedPatchedWritableRackRequestWidthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableRackRequestWidth) IsValid() bool {
	for _, existing := range AllowedPatchedWritableRackRequestWidthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableRackRequest_width value
func (v PatchedWritableRackRequestWidth) Ptr() *PatchedWritableRackRequestWidth {
	return &v
}

type NullablePatchedWritableRackRequestWidth struct {
	value *PatchedWritableRackRequestWidth
	isSet bool
}

func (v NullablePatchedWritableRackRequestWidth) Get() *PatchedWritableRackRequestWidth {
	return v.value
}

func (v *NullablePatchedWritableRackRequestWidth) Set(val *PatchedWritableRackRequestWidth) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackRequestWidth) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackRequestWidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackRequestWidth(val *PatchedWritableRackRequestWidth) *NullablePatchedWritableRackRequestWidth {
	return &NullablePatchedWritableRackRequestWidth{value: val, isSet: true}
}

func (v NullablePatchedWritableRackRequestWidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackRequestWidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BoolConfigItem bool config item
// swagger:model BoolConfigItem
type BoolConfigItem struct {

	// The configure item can be updated or not
	Editable bool `json:"editable,omitempty"`

	// The boolean value of current config item
	Value bool `json:"value,omitempty"`
}

// Validate validates this bool config item
func (m *BoolConfigItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BoolConfigItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BoolConfigItem) UnmarshalBinary(b []byte) error {
	var res BoolConfigItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

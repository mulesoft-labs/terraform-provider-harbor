// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuotaSwitcher quota switcher
// swagger:model QuotaSwitcher
type QuotaSwitcher struct {

	// The quota is enable or disable
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this quota switcher
func (m *QuotaSwitcher) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaSwitcher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaSwitcher) UnmarshalBinary(b []byte) error {
	var res QuotaSwitcher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Report The harbor native report format
// swagger:model Report
type Report struct {

	// Time of generating this report
	GeneratedAt string `json:"generated_at,omitempty"`

	// scanner
	Scanner *Scanner `json:"scanner,omitempty"`

	// A standard scale for measuring the severity of a vulnerability.
	Severity string `json:"severity,omitempty"`

	// vulnerabilities
	Vulnerabilities []*VulnerabilityItem `json:"vulnerabilities"`
}

// Validate validates this report
func (m *Report) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScanner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Report) validateScanner(formats strfmt.Registry) error {

	if swag.IsZero(m.Scanner) { // not required
		return nil
	}

	if m.Scanner != nil {
		if err := m.Scanner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner")
			}
			return err
		}
	}

	return nil
}

func (m *Report) validateVulnerabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.Vulnerabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Vulnerabilities); i++ {
		if swag.IsZero(m.Vulnerabilities[i]) { // not required
			continue
		}

		if m.Vulnerabilities[i] != nil {
			if err := m.Vulnerabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Report) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Report) UnmarshalBinary(b []byte) error {
	var res Report
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

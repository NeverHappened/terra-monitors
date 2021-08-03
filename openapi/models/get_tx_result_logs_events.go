// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetTxResultLogsEvents get tx result logs events
// swagger:model getTxResult.logs.events
type GetTxResultLogsEvents struct {

	// attributes
	// Required: true
	Attributes []*GetTxResultLogsEventsAttributes `json:"attributes"`

	// types
	// Required: true
	Types *string `json:"types"`
}

// Validate validates this get tx result logs events
func (m *GetTxResultLogsEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxResultLogsEvents) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetTxResultLogsEvents) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("types", "body", m.Types); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTxResultLogsEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxResultLogsEvents) UnmarshalBinary(b []byte) error {
	var res GetTxResultLogsEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

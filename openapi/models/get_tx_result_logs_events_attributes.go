// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetTxResultLogsEventsAttributes get tx result logs events attributes
//
// swagger:model getTxResult.logs.events.attributes
type GetTxResultLogsEventsAttributes struct {

	// key
	// Required: true
	Key *string `json:"key"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this get tx result logs events attributes
func (m *GetTxResultLogsEventsAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxResultLogsEventsAttributes) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *GetTxResultLogsEventsAttributes) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get tx result logs events attributes based on context it is used
func (m *GetTxResultLogsEventsAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetTxResultLogsEventsAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxResultLogsEventsAttributes) UnmarshalBinary(b []byte) error {
	var res GetTxResultLogsEventsAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
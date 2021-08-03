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

// GetTxResultTxValueMsgValue get tx result tx value msg value
// swagger:model getTxResult.tx.value.msg.value
type GetTxResultTxValueMsgValue struct {

	// amount
	// Required: true
	Amount []*GetTxResultTxValueMsgValueAmount `json:"amount"`
}

// Validate validates this get tx result tx value msg value
func (m *GetTxResultTxValueMsgValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxResultTxValueMsgValue) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	for i := 0; i < len(m.Amount); i++ {
		if swag.IsZero(m.Amount[i]) { // not required
			continue
		}

		if m.Amount[i] != nil {
			if err := m.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTxResultTxValueMsgValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxResultTxValueMsgValue) UnmarshalBinary(b []byte) error {
	var res GetTxResultTxValueMsgValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetTxListResultTxsTxValueMsgValue get tx list result txs tx value msg value
//
// swagger:model getTxListResult.txs.tx.value.msg.value
type GetTxListResultTxsTxValueMsgValue struct {

	// coins
	// Required: true
	Coins []*GetTxListResultTxsTxValueMsgValueInputsCoins `json:"coins"`

	// contract
	// Required: true
	Contract *string `json:"contract"`

	// execute msg
	ExecuteMsg interface{} `json:"execute_msg,omitempty"`

	// sender
	// Required: true
	Sender *string `json:"sender"`
}

// Validate validates this get tx list result txs tx value msg value
func (m *GetTxListResultTxsTxValueMsgValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxListResultTxsTxValueMsgValue) validateCoins(formats strfmt.Registry) error {

	if err := validate.Required("coins", "body", m.Coins); err != nil {
		return err
	}

	for i := 0; i < len(m.Coins); i++ {
		if swag.IsZero(m.Coins[i]) { // not required
			continue
		}

		if m.Coins[i] != nil {
			if err := m.Coins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetTxListResultTxsTxValueMsgValue) validateContract(formats strfmt.Registry) error {

	if err := validate.Required("contract", "body", m.Contract); err != nil {
		return err
	}

	return nil
}

func (m *GetTxListResultTxsTxValueMsgValue) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get tx list result txs tx value msg value based on the context it is used
func (m *GetTxListResultTxsTxValueMsgValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCoins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxListResultTxsTxValueMsgValue) contextValidateCoins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Coins); i++ {

		if m.Coins[i] != nil {
			if err := m.Coins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTxListResultTxsTxValueMsgValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxListResultTxsTxValueMsgValue) UnmarshalBinary(b []byte) error {
	var res GetTxListResultTxsTxValueMsgValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

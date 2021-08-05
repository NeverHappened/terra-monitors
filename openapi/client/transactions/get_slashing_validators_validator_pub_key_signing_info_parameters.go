// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetSlashingValidatorsValidatorPubKeySigningInfoParams creates a new GetSlashingValidatorsValidatorPubKeySigningInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSlashingValidatorsValidatorPubKeySigningInfoParams() *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoParamsWithTimeout creates a new GetSlashingValidatorsValidatorPubKeySigningInfoParams object
// with the ability to set a timeout on a request.
func NewGetSlashingValidatorsValidatorPubKeySigningInfoParamsWithTimeout(timeout time.Duration) *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoParams{
		timeout: timeout,
	}
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoParamsWithContext creates a new GetSlashingValidatorsValidatorPubKeySigningInfoParams object
// with the ability to set a context for a request.
func NewGetSlashingValidatorsValidatorPubKeySigningInfoParamsWithContext(ctx context.Context) *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoParams{
		Context: ctx,
	}
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoParamsWithHTTPClient creates a new GetSlashingValidatorsValidatorPubKeySigningInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSlashingValidatorsValidatorPubKeySigningInfoParamsWithHTTPClient(client *http.Client) *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoParams{
		HTTPClient: client,
	}
}

/* GetSlashingValidatorsValidatorPubKeySigningInfoParams contains all the parameters to send to the API endpoint
   for the get slashing validators validator pub key signing info operation.

   Typically these are written to a http.Request.
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoParams struct {

	/* ValidatorPubKey.

	   validator public key
	*/
	ValidatorPubKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get slashing validators validator pub key signing info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) WithDefaults() *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get slashing validators validator pub key signing info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get slashing validators validator pub key signing info params
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) WithTimeout(timeout time.Duration) *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get slashing validators validator pub key signing info params
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get slashing validators validator pub key signing info params
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) WithContext(ctx context.Context) *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get slashing validators validator pub key signing info params
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get slashing validators validator pub key signing info params
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) WithHTTPClient(client *http.Client) *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get slashing validators validator pub key signing info params
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidatorPubKey adds the validatorPubKey to the get slashing validators validator pub key signing info params
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) WithValidatorPubKey(validatorPubKey string) *GetSlashingValidatorsValidatorPubKeySigningInfoParams {
	o.SetValidatorPubKey(validatorPubKey)
	return o
}

// SetValidatorPubKey adds the validatorPubKey to the get slashing validators validator pub key signing info params
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) SetValidatorPubKey(validatorPubKey string) {
	o.ValidatorPubKey = validatorPubKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validatorPubKey
	if err := r.SetPathParam("validatorPubKey", o.ValidatorPubKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

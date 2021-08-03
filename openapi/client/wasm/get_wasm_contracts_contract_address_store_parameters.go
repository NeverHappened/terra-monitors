// Code generated by go-swagger; DO NOT EDIT.

package wasm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWasmContractsContractAddressStoreParams creates a new GetWasmContractsContractAddressStoreParams object
// with the default values initialized.
func NewGetWasmContractsContractAddressStoreParams() *GetWasmContractsContractAddressStoreParams {
	var ()
	return &GetWasmContractsContractAddressStoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWasmContractsContractAddressStoreParamsWithTimeout creates a new GetWasmContractsContractAddressStoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWasmContractsContractAddressStoreParamsWithTimeout(timeout time.Duration) *GetWasmContractsContractAddressStoreParams {
	var ()
	return &GetWasmContractsContractAddressStoreParams{

		timeout: timeout,
	}
}

// NewGetWasmContractsContractAddressStoreParamsWithContext creates a new GetWasmContractsContractAddressStoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWasmContractsContractAddressStoreParamsWithContext(ctx context.Context) *GetWasmContractsContractAddressStoreParams {
	var ()
	return &GetWasmContractsContractAddressStoreParams{

		Context: ctx,
	}
}

// NewGetWasmContractsContractAddressStoreParamsWithHTTPClient creates a new GetWasmContractsContractAddressStoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWasmContractsContractAddressStoreParamsWithHTTPClient(client *http.Client) *GetWasmContractsContractAddressStoreParams {
	var ()
	return &GetWasmContractsContractAddressStoreParams{
		HTTPClient: client,
	}
}

/*GetWasmContractsContractAddressStoreParams contains all the parameters to send to the API endpoint
for the get wasm contracts contract address store operation typically these are written to a http.Request
*/
type GetWasmContractsContractAddressStoreParams struct {

	/*ContractAddress
	  contract address you want to lookup

	*/
	ContractAddress string
	/*QueryMsg
	  json formatted query msg

	*/
	QueryMsg string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) WithTimeout(timeout time.Duration) *GetWasmContractsContractAddressStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) WithContext(ctx context.Context) *GetWasmContractsContractAddressStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) WithHTTPClient(client *http.Client) *GetWasmContractsContractAddressStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContractAddress adds the contractAddress to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) WithContractAddress(contractAddress string) *GetWasmContractsContractAddressStoreParams {
	o.SetContractAddress(contractAddress)
	return o
}

// SetContractAddress adds the contractAddress to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) SetContractAddress(contractAddress string) {
	o.ContractAddress = contractAddress
}

// WithQueryMsg adds the queryMsg to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) WithQueryMsg(queryMsg string) *GetWasmContractsContractAddressStoreParams {
	o.SetQueryMsg(queryMsg)
	return o
}

// SetQueryMsg adds the queryMsg to the get wasm contracts contract address store params
func (o *GetWasmContractsContractAddressStoreParams) SetQueryMsg(queryMsg string) {
	o.QueryMsg = queryMsg
}

// WriteToRequest writes these params to a swagger request
func (o *GetWasmContractsContractAddressStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contractAddress
	if err := r.SetPathParam("contractAddress", o.ContractAddress); err != nil {
		return err
	}

	// query param query_msg
	qrQueryMsg := o.QueryMsg
	qQueryMsg := qrQueryMsg
	if qQueryMsg != "" {
		if err := r.SetQueryParam("query_msg", qQueryMsg); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

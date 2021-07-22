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
	"github.com/go-openapi/swag"
)

// NewGetV1TxsParams creates a new GetV1TxsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TxsParams() *GetV1TxsParams {
	return &GetV1TxsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TxsParamsWithTimeout creates a new GetV1TxsParams object
// with the ability to set a timeout on a request.
func NewGetV1TxsParamsWithTimeout(timeout time.Duration) *GetV1TxsParams {
	return &GetV1TxsParams{
		timeout: timeout,
	}
}

// NewGetV1TxsParamsWithContext creates a new GetV1TxsParams object
// with the ability to set a context for a request.
func NewGetV1TxsParamsWithContext(ctx context.Context) *GetV1TxsParams {
	return &GetV1TxsParams{
		Context: ctx,
	}
}

// NewGetV1TxsParamsWithHTTPClient creates a new GetV1TxsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TxsParamsWithHTTPClient(client *http.Client) *GetV1TxsParams {
	return &GetV1TxsParams{
		HTTPClient: client,
	}
}

/* GetV1TxsParams contains all the parameters to send to the API endpoint
   for the get v1 txs operation.

   Typically these are written to a http.Request.
*/
type GetV1TxsParams struct {

	/* Account.

	   Account address
	*/
	Account *string

	/* Action.

	   Type of tx (account is required)
	*/
	Action *string

	/* Block.

	   Block number
	*/
	Block *string

	/* ChainID.

	   Chain ID of Blockchain (default: chain id of mainnet)
	*/
	ChainID *string

	/* From.

	   Timestamp from
	*/
	From *float64

	/* Limit.

	   Size of page
	*/
	Limit *float64

	/* Memo.

	   Memo filter
	*/
	Memo *string

	/* Offset.

	   Use last id from previous result for pagination
	*/
	Offset *float64

	/* Order.

	   ,'DESC'] Ordering (default: DESC)
	*/
	Order *string

	/* Page.

	   # of page
	*/
	Page *float64

	/* To.

	   Timestamp to
	*/
	To *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 txs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TxsParams) WithDefaults() *GetV1TxsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 txs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TxsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 txs params
func (o *GetV1TxsParams) WithTimeout(timeout time.Duration) *GetV1TxsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 txs params
func (o *GetV1TxsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 txs params
func (o *GetV1TxsParams) WithContext(ctx context.Context) *GetV1TxsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 txs params
func (o *GetV1TxsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 txs params
func (o *GetV1TxsParams) WithHTTPClient(client *http.Client) *GetV1TxsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 txs params
func (o *GetV1TxsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the get v1 txs params
func (o *GetV1TxsParams) WithAccount(account *string) *GetV1TxsParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the get v1 txs params
func (o *GetV1TxsParams) SetAccount(account *string) {
	o.Account = account
}

// WithAction adds the action to the get v1 txs params
func (o *GetV1TxsParams) WithAction(action *string) *GetV1TxsParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the get v1 txs params
func (o *GetV1TxsParams) SetAction(action *string) {
	o.Action = action
}

// WithBlock adds the block to the get v1 txs params
func (o *GetV1TxsParams) WithBlock(block *string) *GetV1TxsParams {
	o.SetBlock(block)
	return o
}

// SetBlock adds the block to the get v1 txs params
func (o *GetV1TxsParams) SetBlock(block *string) {
	o.Block = block
}

// WithChainID adds the chainID to the get v1 txs params
func (o *GetV1TxsParams) WithChainID(chainID *string) *GetV1TxsParams {
	o.SetChainID(chainID)
	return o
}

// SetChainID adds the chainId to the get v1 txs params
func (o *GetV1TxsParams) SetChainID(chainID *string) {
	o.ChainID = chainID
}

// WithFrom adds the from to the get v1 txs params
func (o *GetV1TxsParams) WithFrom(from *float64) *GetV1TxsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get v1 txs params
func (o *GetV1TxsParams) SetFrom(from *float64) {
	o.From = from
}

// WithLimit adds the limit to the get v1 txs params
func (o *GetV1TxsParams) WithLimit(limit *float64) *GetV1TxsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get v1 txs params
func (o *GetV1TxsParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithMemo adds the memo to the get v1 txs params
func (o *GetV1TxsParams) WithMemo(memo *string) *GetV1TxsParams {
	o.SetMemo(memo)
	return o
}

// SetMemo adds the memo to the get v1 txs params
func (o *GetV1TxsParams) SetMemo(memo *string) {
	o.Memo = memo
}

// WithOffset adds the offset to the get v1 txs params
func (o *GetV1TxsParams) WithOffset(offset *float64) *GetV1TxsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get v1 txs params
func (o *GetV1TxsParams) SetOffset(offset *float64) {
	o.Offset = offset
}

// WithOrder adds the order to the get v1 txs params
func (o *GetV1TxsParams) WithOrder(order *string) *GetV1TxsParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the get v1 txs params
func (o *GetV1TxsParams) SetOrder(order *string) {
	o.Order = order
}

// WithPage adds the page to the get v1 txs params
func (o *GetV1TxsParams) WithPage(page *float64) *GetV1TxsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 txs params
func (o *GetV1TxsParams) SetPage(page *float64) {
	o.Page = page
}

// WithTo adds the to to the get v1 txs params
func (o *GetV1TxsParams) WithTo(to *float64) *GetV1TxsParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get v1 txs params
func (o *GetV1TxsParams) SetTo(to *float64) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TxsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Account != nil {

		// query param account
		var qrAccount string

		if o.Account != nil {
			qrAccount = *o.Account
		}
		qAccount := qrAccount
		if qAccount != "" {

			if err := r.SetQueryParam("account", qAccount); err != nil {
				return err
			}
		}
	}

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.Block != nil {

		// query param block
		var qrBlock string

		if o.Block != nil {
			qrBlock = *o.Block
		}
		qBlock := qrBlock
		if qBlock != "" {

			if err := r.SetQueryParam("block", qBlock); err != nil {
				return err
			}
		}
	}

	if o.ChainID != nil {

		// query param chainId
		var qrChainID string

		if o.ChainID != nil {
			qrChainID = *o.ChainID
		}
		qChainID := qrChainID
		if qChainID != "" {

			if err := r.SetQueryParam("chainId", qChainID); err != nil {
				return err
			}
		}
	}

	if o.From != nil {

		// query param from
		var qrFrom float64

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatFloat64(qrFrom)
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit float64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Memo != nil {

		// query param memo
		var qrMemo string

		if o.Memo != nil {
			qrMemo = *o.Memo
		}
		qMemo := qrMemo
		if qMemo != "" {

			if err := r.SetQueryParam("memo", qMemo); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset float64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatFloat64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Order != nil {

		// query param order
		var qrOrder string

		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {

			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage float64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatFloat64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.To != nil {

		// query param to
		var qrTo float64

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := swag.FormatFloat64(qrTo)
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
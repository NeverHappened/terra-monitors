// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/lidofinance/terra-monitors/openapi/models"
)

// GetStakingDelegatorsDelegatorAddrDelegationsReader is a Reader for the GetStakingDelegatorsDelegatorAddrDelegations structure.
type GetStakingDelegatorsDelegatorAddrDelegationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingDelegatorsDelegatorAddrDelegationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsOK creates a GetStakingDelegatorsDelegatorAddrDelegationsOK with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsOK() *GetStakingDelegatorsDelegatorAddrDelegationsOK {
	return &GetStakingDelegatorsDelegatorAddrDelegationsOK{}
}

/* GetStakingDelegatorsDelegatorAddrDelegationsOK describes a response with status code 200, with default header values.

OK
*/
type GetStakingDelegatorsDelegatorAddrDelegationsOK struct {
	Payload *GetStakingDelegatorsDelegatorAddrDelegationsOKBody
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsOK) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations][%d] getStakingDelegatorsDelegatorAddrDelegationsOK  %+v", 200, o.Payload)
}
func (o *GetStakingDelegatorsDelegatorAddrDelegationsOK) GetPayload() *GetStakingDelegatorsDelegatorAddrDelegationsOKBody {
	return o.Payload
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStakingDelegatorsDelegatorAddrDelegationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsBadRequest creates a GetStakingDelegatorsDelegatorAddrDelegationsBadRequest with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsBadRequest() *GetStakingDelegatorsDelegatorAddrDelegationsBadRequest {
	return &GetStakingDelegatorsDelegatorAddrDelegationsBadRequest{}
}

/* GetStakingDelegatorsDelegatorAddrDelegationsBadRequest describes a response with status code 400, with default header values.

Invalid delegator address
*/
type GetStakingDelegatorsDelegatorAddrDelegationsBadRequest struct {
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations][%d] getStakingDelegatorsDelegatorAddrDelegationsBadRequest ", 400)
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsInternalServerError creates a GetStakingDelegatorsDelegatorAddrDelegationsInternalServerError with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsInternalServerError() *GetStakingDelegatorsDelegatorAddrDelegationsInternalServerError {
	return &GetStakingDelegatorsDelegatorAddrDelegationsInternalServerError{}
}

/* GetStakingDelegatorsDelegatorAddrDelegationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStakingDelegatorsDelegatorAddrDelegationsInternalServerError struct {
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations][%d] getStakingDelegatorsDelegatorAddrDelegationsInternalServerError ", 500)
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStakingDelegatorsDelegatorAddrDelegationsOKBody get staking delegators delegator addr delegations o k body
swagger:model GetStakingDelegatorsDelegatorAddrDelegationsOKBody
*/
type GetStakingDelegatorsDelegatorAddrDelegationsOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result []*models.Delegation `json:"result"`
}

// Validate validates this get staking delegators delegator addr delegations o k body
func (o *GetStakingDelegatorsDelegatorAddrDelegationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	for i := 0; i < len(o.Result); i++ {
		if swag.IsZero(o.Result[i]) { // not required
			continue
		}

		if o.Result[i] != nil {
			if err := o.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsOK" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get staking delegators delegator addr delegations o k body based on the context it is used
func (o *GetStakingDelegatorsDelegatorAddrDelegationsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Result); i++ {

		if o.Result[i] != nil {
			if err := o.Result[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsOK" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsOKBody) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrDelegationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

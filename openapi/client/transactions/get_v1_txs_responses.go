// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/lidofinance/terra-monitors/openapi/models"
)

// GetV1TxsReader is a Reader for the GetV1Txs structure.
type GetV1TxsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TxsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TxsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1TxsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1TxsOK creates a GetV1TxsOK with default headers values
func NewGetV1TxsOK() *GetV1TxsOK {
	return &GetV1TxsOK{}
}

/*GetV1TxsOK handles this case with default header values.

Success
*/
type GetV1TxsOK struct {
	Payload *models.GetTxListResult
}

func (o *GetV1TxsOK) Error() string {
	return fmt.Sprintf("[GET /v1/txs][%d] getV1TxsOK  %+v", 200, o.Payload)
}

func (o *GetV1TxsOK) GetPayload() *models.GetTxListResult {
	return o.Payload
}

func (o *GetV1TxsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTxListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1TxsBadRequest creates a GetV1TxsBadRequest with default headers values
func NewGetV1TxsBadRequest() *GetV1TxsBadRequest {
	return &GetV1TxsBadRequest{}
}

/*GetV1TxsBadRequest handles this case with default header values.

Error
*/
type GetV1TxsBadRequest struct {
	Payload *GetV1TxsBadRequestBody
}

func (o *GetV1TxsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/txs][%d] getV1TxsBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1TxsBadRequest) GetPayload() *GetV1TxsBadRequestBody {
	return o.Payload
}

func (o *GetV1TxsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1TxsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetV1TxsBadRequestBody get v1 txs bad request body
swagger:model GetV1TxsBadRequestBody
*/
type GetV1TxsBadRequestBody struct {

	// code
	Code float64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this get v1 txs bad request body
func (o *GetV1TxsBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetV1TxsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1TxsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetV1TxsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

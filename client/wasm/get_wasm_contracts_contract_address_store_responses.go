// Code generated by go-swagger; DO NOT EDIT.

package wasm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetWasmContractsContractAddressStoreReader is a Reader for the GetWasmContractsContractAddressStore structure.
type GetWasmContractsContractAddressStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWasmContractsContractAddressStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWasmContractsContractAddressStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWasmContractsContractAddressStoreOK creates a GetWasmContractsContractAddressStoreOK with default headers values
func NewGetWasmContractsContractAddressStoreOK() *GetWasmContractsContractAddressStoreOK {
	return &GetWasmContractsContractAddressStoreOK{}
}

/* GetWasmContractsContractAddressStoreOK describes a response with status code 200, with default header values.

OK
*/
type GetWasmContractsContractAddressStoreOK struct {
	Payload *GetWasmContractsContractAddressStoreOKBody
}

func (o *GetWasmContractsContractAddressStoreOK) Error() string {
	return fmt.Sprintf("[GET /wasm/contracts/{contractAddress}/store][%d] getWasmContractsContractAddressStoreOK  %+v", 200, o.Payload)
}
func (o *GetWasmContractsContractAddressStoreOK) GetPayload() *GetWasmContractsContractAddressStoreOKBody {
	return o.Payload
}

func (o *GetWasmContractsContractAddressStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWasmContractsContractAddressStoreOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWasmContractsContractAddressStoreOKBody get wasm contracts contract address store o k body
swagger:model GetWasmContractsContractAddressStoreOKBody
*/
type GetWasmContractsContractAddressStoreOKBody struct {

	// error
	Error string `json:"error,omitempty"`

	// height
	Height string `json:"height,omitempty"`

	// result
	Result interface{} `json:"result,omitempty"`
}

// Validate validates this get wasm contracts contract address store o k body
func (o *GetWasmContractsContractAddressStoreOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wasm contracts contract address store o k body based on context it is used
func (o *GetWasmContractsContractAddressStoreOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWasmContractsContractAddressStoreOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWasmContractsContractAddressStoreOKBody) UnmarshalBinary(b []byte) error {
	var res GetWasmContractsContractAddressStoreOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

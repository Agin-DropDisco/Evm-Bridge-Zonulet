// Code generated by go-swagger; DO NOT EDIT.

package erc_1155_swaps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/models"
)

// GetErc1155SwapsOKCode is the HTTP code returned for type GetErc1155SwapsOK
const GetErc1155SwapsOKCode int = 200

/*GetErc1155SwapsOK Success

swagger:response getErc1155SwapsOK
*/
type GetErc1155SwapsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Erc1155Swaps `json:"body,omitempty"`
}

// NewGetErc1155SwapsOK creates GetErc1155SwapsOK with default headers values
func NewGetErc1155SwapsOK() *GetErc1155SwapsOK {

	return &GetErc1155SwapsOK{}
}

// WithPayload adds the payload to the get erc1155 swaps o k response
func (o *GetErc1155SwapsOK) WithPayload(payload *models.Erc1155Swaps) *GetErc1155SwapsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc1155 swaps o k response
func (o *GetErc1155SwapsOK) SetPayload(payload *models.Erc1155Swaps) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc1155SwapsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetErc1155SwapsBadRequestCode is the HTTP code returned for type GetErc1155SwapsBadRequest
const GetErc1155SwapsBadRequestCode int = 400

/*GetErc1155SwapsBadRequest Bad Request

swagger:response getErc1155SwapsBadRequest
*/
type GetErc1155SwapsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetErc1155SwapsBadRequest creates GetErc1155SwapsBadRequest with default headers values
func NewGetErc1155SwapsBadRequest() *GetErc1155SwapsBadRequest {

	return &GetErc1155SwapsBadRequest{}
}

// WithPayload adds the payload to the get erc1155 swaps bad request response
func (o *GetErc1155SwapsBadRequest) WithPayload(payload *models.Error) *GetErc1155SwapsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc1155 swaps bad request response
func (o *GetErc1155SwapsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc1155SwapsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetErc1155SwapsInternalServerErrorCode is the HTTP code returned for type GetErc1155SwapsInternalServerError
const GetErc1155SwapsInternalServerErrorCode int = 500

/*GetErc1155SwapsInternalServerError internal server error

swagger:response getErc1155SwapsInternalServerError
*/
type GetErc1155SwapsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetErc1155SwapsInternalServerError creates GetErc1155SwapsInternalServerError with default headers values
func NewGetErc1155SwapsInternalServerError() *GetErc1155SwapsInternalServerError {

	return &GetErc1155SwapsInternalServerError{}
}

// WithPayload adds the payload to the get erc1155 swaps internal server error response
func (o *GetErc1155SwapsInternalServerError) WithPayload(payload *models.Error) *GetErc1155SwapsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get erc1155 swaps internal server error response
func (o *GetErc1155SwapsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetErc1155SwapsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
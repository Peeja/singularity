// Code generated by go-swagger; DO NOT EDIT.

package data_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// GetItemIDDealsReader is a Reader for the GetItemIDDeals structure.
type GetItemIDDealsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetItemIDDealsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetItemIDDealsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetItemIDDealsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /item/{id}/deals] GetItemIDDeals", response, response.Code())
	}
}

// NewGetItemIDDealsOK creates a GetItemIDDealsOK with default headers values
func NewGetItemIDDealsOK() *GetItemIDDealsOK {
	return &GetItemIDDealsOK{}
}

/*
GetItemIDDealsOK describes a response with status code 200, with default header values.

OK
*/
type GetItemIDDealsOK struct {
	Payload []*models.ModelDeal
}

// IsSuccess returns true when this get item Id deals o k response has a 2xx status code
func (o *GetItemIDDealsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get item Id deals o k response has a 3xx status code
func (o *GetItemIDDealsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item Id deals o k response has a 4xx status code
func (o *GetItemIDDealsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get item Id deals o k response has a 5xx status code
func (o *GetItemIDDealsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get item Id deals o k response a status code equal to that given
func (o *GetItemIDDealsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get item Id deals o k response
func (o *GetItemIDDealsOK) Code() int {
	return 200
}

func (o *GetItemIDDealsOK) Error() string {
	return fmt.Sprintf("[GET /item/{id}/deals][%d] getItemIdDealsOK  %+v", 200, o.Payload)
}

func (o *GetItemIDDealsOK) String() string {
	return fmt.Sprintf("[GET /item/{id}/deals][%d] getItemIdDealsOK  %+v", 200, o.Payload)
}

func (o *GetItemIDDealsOK) GetPayload() []*models.ModelDeal {
	return o.Payload
}

func (o *GetItemIDDealsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemIDDealsInternalServerError creates a GetItemIDDealsInternalServerError with default headers values
func NewGetItemIDDealsInternalServerError() *GetItemIDDealsInternalServerError {
	return &GetItemIDDealsInternalServerError{}
}

/*
GetItemIDDealsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetItemIDDealsInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get item Id deals internal server error response has a 2xx status code
func (o *GetItemIDDealsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item Id deals internal server error response has a 3xx status code
func (o *GetItemIDDealsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item Id deals internal server error response has a 4xx status code
func (o *GetItemIDDealsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get item Id deals internal server error response has a 5xx status code
func (o *GetItemIDDealsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get item Id deals internal server error response a status code equal to that given
func (o *GetItemIDDealsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get item Id deals internal server error response
func (o *GetItemIDDealsInternalServerError) Code() int {
	return 500
}

func (o *GetItemIDDealsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /item/{id}/deals][%d] getItemIdDealsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetItemIDDealsInternalServerError) String() string {
	return fmt.Sprintf("[GET /item/{id}/deals][%d] getItemIdDealsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetItemIDDealsInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetItemIDDealsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
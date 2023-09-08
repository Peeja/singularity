// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// CreateOosInstancePrincipalAuthStorageReader is a Reader for the CreateOosInstancePrincipalAuthStorage structure.
type CreateOosInstancePrincipalAuthStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOosInstancePrincipalAuthStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOosInstancePrincipalAuthStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOosInstancePrincipalAuthStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateOosInstancePrincipalAuthStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/oos/instance_principal_auth] CreateOosInstance_principal_authStorage", response, response.Code())
	}
}

// NewCreateOosInstancePrincipalAuthStorageOK creates a CreateOosInstancePrincipalAuthStorageOK with default headers values
func NewCreateOosInstancePrincipalAuthStorageOK() *CreateOosInstancePrincipalAuthStorageOK {
	return &CreateOosInstancePrincipalAuthStorageOK{}
}

/*
CreateOosInstancePrincipalAuthStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateOosInstancePrincipalAuthStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create oos instance principal auth storage o k response has a 2xx status code
func (o *CreateOosInstancePrincipalAuthStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create oos instance principal auth storage o k response has a 3xx status code
func (o *CreateOosInstancePrincipalAuthStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos instance principal auth storage o k response has a 4xx status code
func (o *CreateOosInstancePrincipalAuthStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create oos instance principal auth storage o k response has a 5xx status code
func (o *CreateOosInstancePrincipalAuthStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create oos instance principal auth storage o k response a status code equal to that given
func (o *CreateOosInstancePrincipalAuthStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create oos instance principal auth storage o k response
func (o *CreateOosInstancePrincipalAuthStorageOK) Code() int {
	return 200
}

func (o *CreateOosInstancePrincipalAuthStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/oos/instance_principal_auth][%d] createOosInstancePrincipalAuthStorageOK  %+v", 200, o.Payload)
}

func (o *CreateOosInstancePrincipalAuthStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/oos/instance_principal_auth][%d] createOosInstancePrincipalAuthStorageOK  %+v", 200, o.Payload)
}

func (o *CreateOosInstancePrincipalAuthStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateOosInstancePrincipalAuthStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOosInstancePrincipalAuthStorageBadRequest creates a CreateOosInstancePrincipalAuthStorageBadRequest with default headers values
func NewCreateOosInstancePrincipalAuthStorageBadRequest() *CreateOosInstancePrincipalAuthStorageBadRequest {
	return &CreateOosInstancePrincipalAuthStorageBadRequest{}
}

/*
CreateOosInstancePrincipalAuthStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateOosInstancePrincipalAuthStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create oos instance principal auth storage bad request response has a 2xx status code
func (o *CreateOosInstancePrincipalAuthStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create oos instance principal auth storage bad request response has a 3xx status code
func (o *CreateOosInstancePrincipalAuthStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos instance principal auth storage bad request response has a 4xx status code
func (o *CreateOosInstancePrincipalAuthStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create oos instance principal auth storage bad request response has a 5xx status code
func (o *CreateOosInstancePrincipalAuthStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create oos instance principal auth storage bad request response a status code equal to that given
func (o *CreateOosInstancePrincipalAuthStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create oos instance principal auth storage bad request response
func (o *CreateOosInstancePrincipalAuthStorageBadRequest) Code() int {
	return 400
}

func (o *CreateOosInstancePrincipalAuthStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/oos/instance_principal_auth][%d] createOosInstancePrincipalAuthStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOosInstancePrincipalAuthStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/oos/instance_principal_auth][%d] createOosInstancePrincipalAuthStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOosInstancePrincipalAuthStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateOosInstancePrincipalAuthStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOosInstancePrincipalAuthStorageInternalServerError creates a CreateOosInstancePrincipalAuthStorageInternalServerError with default headers values
func NewCreateOosInstancePrincipalAuthStorageInternalServerError() *CreateOosInstancePrincipalAuthStorageInternalServerError {
	return &CreateOosInstancePrincipalAuthStorageInternalServerError{}
}

/*
CreateOosInstancePrincipalAuthStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateOosInstancePrincipalAuthStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create oos instance principal auth storage internal server error response has a 2xx status code
func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create oos instance principal auth storage internal server error response has a 3xx status code
func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos instance principal auth storage internal server error response has a 4xx status code
func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create oos instance principal auth storage internal server error response has a 5xx status code
func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create oos instance principal auth storage internal server error response a status code equal to that given
func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create oos instance principal auth storage internal server error response
func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/oos/instance_principal_auth][%d] createOosInstancePrincipalAuthStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/oos/instance_principal_auth][%d] createOosInstancePrincipalAuthStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateOosInstancePrincipalAuthStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
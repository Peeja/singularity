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

// CreateS3DigitalOceanStorageReader is a Reader for the CreateS3DigitalOceanStorage structure.
type CreateS3DigitalOceanStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3DigitalOceanStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3DigitalOceanStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3DigitalOceanStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3DigitalOceanStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/digitalocean] CreateS3DigitalOceanStorage", response, response.Code())
	}
}

// NewCreateS3DigitalOceanStorageOK creates a CreateS3DigitalOceanStorageOK with default headers values
func NewCreateS3DigitalOceanStorageOK() *CreateS3DigitalOceanStorageOK {
	return &CreateS3DigitalOceanStorageOK{}
}

/*
CreateS3DigitalOceanStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3DigitalOceanStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 digital ocean storage o k response has a 2xx status code
func (o *CreateS3DigitalOceanStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 digital ocean storage o k response has a 3xx status code
func (o *CreateS3DigitalOceanStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 digital ocean storage o k response has a 4xx status code
func (o *CreateS3DigitalOceanStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 digital ocean storage o k response has a 5xx status code
func (o *CreateS3DigitalOceanStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 digital ocean storage o k response a status code equal to that given
func (o *CreateS3DigitalOceanStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 digital ocean storage o k response
func (o *CreateS3DigitalOceanStorageOK) Code() int {
	return 200
}

func (o *CreateS3DigitalOceanStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/digitalocean][%d] createS3DigitalOceanStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3DigitalOceanStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/digitalocean][%d] createS3DigitalOceanStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3DigitalOceanStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3DigitalOceanStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3DigitalOceanStorageBadRequest creates a CreateS3DigitalOceanStorageBadRequest with default headers values
func NewCreateS3DigitalOceanStorageBadRequest() *CreateS3DigitalOceanStorageBadRequest {
	return &CreateS3DigitalOceanStorageBadRequest{}
}

/*
CreateS3DigitalOceanStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3DigitalOceanStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 digital ocean storage bad request response has a 2xx status code
func (o *CreateS3DigitalOceanStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 digital ocean storage bad request response has a 3xx status code
func (o *CreateS3DigitalOceanStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 digital ocean storage bad request response has a 4xx status code
func (o *CreateS3DigitalOceanStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 digital ocean storage bad request response has a 5xx status code
func (o *CreateS3DigitalOceanStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 digital ocean storage bad request response a status code equal to that given
func (o *CreateS3DigitalOceanStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 digital ocean storage bad request response
func (o *CreateS3DigitalOceanStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3DigitalOceanStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/digitalocean][%d] createS3DigitalOceanStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3DigitalOceanStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/digitalocean][%d] createS3DigitalOceanStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3DigitalOceanStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3DigitalOceanStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3DigitalOceanStorageInternalServerError creates a CreateS3DigitalOceanStorageInternalServerError with default headers values
func NewCreateS3DigitalOceanStorageInternalServerError() *CreateS3DigitalOceanStorageInternalServerError {
	return &CreateS3DigitalOceanStorageInternalServerError{}
}

/*
CreateS3DigitalOceanStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3DigitalOceanStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 digital ocean storage internal server error response has a 2xx status code
func (o *CreateS3DigitalOceanStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 digital ocean storage internal server error response has a 3xx status code
func (o *CreateS3DigitalOceanStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 digital ocean storage internal server error response has a 4xx status code
func (o *CreateS3DigitalOceanStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 digital ocean storage internal server error response has a 5xx status code
func (o *CreateS3DigitalOceanStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 digital ocean storage internal server error response a status code equal to that given
func (o *CreateS3DigitalOceanStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 digital ocean storage internal server error response
func (o *CreateS3DigitalOceanStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3DigitalOceanStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/digitalocean][%d] createS3DigitalOceanStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3DigitalOceanStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/digitalocean][%d] createS3DigitalOceanStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3DigitalOceanStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3DigitalOceanStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
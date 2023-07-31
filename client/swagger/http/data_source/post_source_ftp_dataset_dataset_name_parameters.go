// Code generated by go-swagger; DO NOT EDIT.

package data_source

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

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewPostSourceFtpDatasetDatasetNameParams creates a new PostSourceFtpDatasetDatasetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceFtpDatasetDatasetNameParams() *PostSourceFtpDatasetDatasetNameParams {
	return &PostSourceFtpDatasetDatasetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceFtpDatasetDatasetNameParamsWithTimeout creates a new PostSourceFtpDatasetDatasetNameParams object
// with the ability to set a timeout on a request.
func NewPostSourceFtpDatasetDatasetNameParamsWithTimeout(timeout time.Duration) *PostSourceFtpDatasetDatasetNameParams {
	return &PostSourceFtpDatasetDatasetNameParams{
		timeout: timeout,
	}
}

// NewPostSourceFtpDatasetDatasetNameParamsWithContext creates a new PostSourceFtpDatasetDatasetNameParams object
// with the ability to set a context for a request.
func NewPostSourceFtpDatasetDatasetNameParamsWithContext(ctx context.Context) *PostSourceFtpDatasetDatasetNameParams {
	return &PostSourceFtpDatasetDatasetNameParams{
		Context: ctx,
	}
}

// NewPostSourceFtpDatasetDatasetNameParamsWithHTTPClient creates a new PostSourceFtpDatasetDatasetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceFtpDatasetDatasetNameParamsWithHTTPClient(client *http.Client) *PostSourceFtpDatasetDatasetNameParams {
	return &PostSourceFtpDatasetDatasetNameParams{
		HTTPClient: client,
	}
}

/*
PostSourceFtpDatasetDatasetNameParams contains all the parameters to send to the API endpoint

	for the post source ftp dataset dataset name operation.

	Typically these are written to a http.Request.
*/
type PostSourceFtpDatasetDatasetNameParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceFtpRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source ftp dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceFtpDatasetDatasetNameParams) WithDefaults() *PostSourceFtpDatasetDatasetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source ftp dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceFtpDatasetDatasetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) WithTimeout(timeout time.Duration) *PostSourceFtpDatasetDatasetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) WithContext(ctx context.Context) *PostSourceFtpDatasetDatasetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) WithHTTPClient(client *http.Client) *PostSourceFtpDatasetDatasetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) WithDatasetName(datasetName string) *PostSourceFtpDatasetDatasetNameParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WithRequest adds the request to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) WithRequest(request *models.DatasourceFtpRequest) *PostSourceFtpDatasetDatasetNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source ftp dataset dataset name params
func (o *PostSourceFtpDatasetDatasetNameParams) SetRequest(request *models.DatasourceFtpRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceFtpDatasetDatasetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasetName
	if err := r.SetPathParam("datasetName", o.DatasetName); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
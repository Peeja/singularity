// Code generated by go-swagger; DO NOT EDIT.

package deal

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

// NewPostDealParams creates a new PostDealParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDealParams() *PostDealParams {
	return &PostDealParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDealParamsWithTimeout creates a new PostDealParams object
// with the ability to set a timeout on a request.
func NewPostDealParamsWithTimeout(timeout time.Duration) *PostDealParams {
	return &PostDealParams{
		timeout: timeout,
	}
}

// NewPostDealParamsWithContext creates a new PostDealParams object
// with the ability to set a context for a request.
func NewPostDealParamsWithContext(ctx context.Context) *PostDealParams {
	return &PostDealParams{
		Context: ctx,
	}
}

// NewPostDealParamsWithHTTPClient creates a new PostDealParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDealParamsWithHTTPClient(client *http.Client) *PostDealParams {
	return &PostDealParams{
		HTTPClient: client,
	}
}

/*
PostDealParams contains all the parameters to send to the API endpoint

	for the post deal operation.

	Typically these are written to a http.Request.
*/
type PostDealParams struct {

	/* Request.

	   ListDealRequest
	*/
	Request *models.DealListDealRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post deal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDealParams) WithDefaults() *PostDealParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post deal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDealParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post deal params
func (o *PostDealParams) WithTimeout(timeout time.Duration) *PostDealParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post deal params
func (o *PostDealParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post deal params
func (o *PostDealParams) WithContext(ctx context.Context) *PostDealParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post deal params
func (o *PostDealParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post deal params
func (o *PostDealParams) WithHTTPClient(client *http.Client) *PostDealParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post deal params
func (o *PostDealParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post deal params
func (o *PostDealParams) WithRequest(request *models.DealListDealRequest) *PostDealParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post deal params
func (o *PostDealParams) SetRequest(request *models.DealListDealRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostDealParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
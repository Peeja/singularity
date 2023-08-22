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
)

// NewGetPackjobIDParams creates a new GetPackjobIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPackjobIDParams() *GetPackjobIDParams {
	return &GetPackjobIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackjobIDParamsWithTimeout creates a new GetPackjobIDParams object
// with the ability to set a timeout on a request.
func NewGetPackjobIDParamsWithTimeout(timeout time.Duration) *GetPackjobIDParams {
	return &GetPackjobIDParams{
		timeout: timeout,
	}
}

// NewGetPackjobIDParamsWithContext creates a new GetPackjobIDParams object
// with the ability to set a context for a request.
func NewGetPackjobIDParamsWithContext(ctx context.Context) *GetPackjobIDParams {
	return &GetPackjobIDParams{
		Context: ctx,
	}
}

// NewGetPackjobIDParamsWithHTTPClient creates a new GetPackjobIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPackjobIDParamsWithHTTPClient(client *http.Client) *GetPackjobIDParams {
	return &GetPackjobIDParams{
		HTTPClient: client,
	}
}

/*
GetPackjobIDParams contains all the parameters to send to the API endpoint

	for the get packjob ID operation.

	Typically these are written to a http.Request.
*/
type GetPackjobIDParams struct {

	/* ID.

	   Pack job ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get packjob ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackjobIDParams) WithDefaults() *GetPackjobIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get packjob ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackjobIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get packjob ID params
func (o *GetPackjobIDParams) WithTimeout(timeout time.Duration) *GetPackjobIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get packjob ID params
func (o *GetPackjobIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get packjob ID params
func (o *GetPackjobIDParams) WithContext(ctx context.Context) *GetPackjobIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get packjob ID params
func (o *GetPackjobIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get packjob ID params
func (o *GetPackjobIDParams) WithHTTPClient(client *http.Client) *GetPackjobIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get packjob ID params
func (o *GetPackjobIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get packjob ID params
func (o *GetPackjobIDParams) WithID(id string) *GetPackjobIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get packjob ID params
func (o *GetPackjobIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackjobIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
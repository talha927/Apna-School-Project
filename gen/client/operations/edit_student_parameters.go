// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEditStudentParams creates a new EditStudentParams object
// with the default values initialized.
func NewEditStudentParams() *EditStudentParams {
	var ()
	return &EditStudentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditStudentParamsWithTimeout creates a new EditStudentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditStudentParamsWithTimeout(timeout time.Duration) *EditStudentParams {
	var ()
	return &EditStudentParams{

		timeout: timeout,
	}
}

// NewEditStudentParamsWithContext creates a new EditStudentParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditStudentParamsWithContext(ctx context.Context) *EditStudentParams {
	var ()
	return &EditStudentParams{

		Context: ctx,
	}
}

// NewEditStudentParamsWithHTTPClient creates a new EditStudentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditStudentParamsWithHTTPClient(client *http.Client) *EditStudentParams {
	var ()
	return &EditStudentParams{
		HTTPClient: client,
	}
}

/*EditStudentParams contains all the parameters to send to the API endpoint
for the edit student operation typically these are written to a http.Request
*/
type EditStudentParams struct {

	/*ID
	  UUID of the student

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit student params
func (o *EditStudentParams) WithTimeout(timeout time.Duration) *EditStudentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit student params
func (o *EditStudentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit student params
func (o *EditStudentParams) WithContext(ctx context.Context) *EditStudentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit student params
func (o *EditStudentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit student params
func (o *EditStudentParams) WithHTTPClient(client *http.Client) *EditStudentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit student params
func (o *EditStudentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edit student params
func (o *EditStudentParams) WithID(id string) *EditStudentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edit student params
func (o *EditStudentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EditStudentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

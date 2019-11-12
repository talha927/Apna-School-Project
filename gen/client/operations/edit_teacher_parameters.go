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

// NewEditTeacherParams creates a new EditTeacherParams object
// with the default values initialized.
func NewEditTeacherParams() *EditTeacherParams {
	var ()
	return &EditTeacherParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditTeacherParamsWithTimeout creates a new EditTeacherParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditTeacherParamsWithTimeout(timeout time.Duration) *EditTeacherParams {
	var ()
	return &EditTeacherParams{

		timeout: timeout,
	}
}

// NewEditTeacherParamsWithContext creates a new EditTeacherParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditTeacherParamsWithContext(ctx context.Context) *EditTeacherParams {
	var ()
	return &EditTeacherParams{

		Context: ctx,
	}
}

// NewEditTeacherParamsWithHTTPClient creates a new EditTeacherParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditTeacherParamsWithHTTPClient(client *http.Client) *EditTeacherParams {
	var ()
	return &EditTeacherParams{
		HTTPClient: client,
	}
}

/*EditTeacherParams contains all the parameters to send to the API endpoint
for the edit teacher operation typically these are written to a http.Request
*/
type EditTeacherParams struct {

	/*ID
	  UUID of the teacher

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit teacher params
func (o *EditTeacherParams) WithTimeout(timeout time.Duration) *EditTeacherParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit teacher params
func (o *EditTeacherParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit teacher params
func (o *EditTeacherParams) WithContext(ctx context.Context) *EditTeacherParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit teacher params
func (o *EditTeacherParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit teacher params
func (o *EditTeacherParams) WithHTTPClient(client *http.Client) *EditTeacherParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit teacher params
func (o *EditTeacherParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edit teacher params
func (o *EditTeacherParams) WithID(id string) *EditTeacherParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edit teacher params
func (o *EditTeacherParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EditTeacherParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

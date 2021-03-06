// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRegistriesIDParams creates a new DeleteRegistriesIDParams object
// with the default values initialized.
func NewDeleteRegistriesIDParams() *DeleteRegistriesIDParams {
	var ()
	return &DeleteRegistriesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRegistriesIDParamsWithTimeout creates a new DeleteRegistriesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRegistriesIDParamsWithTimeout(timeout time.Duration) *DeleteRegistriesIDParams {
	var ()
	return &DeleteRegistriesIDParams{

		timeout: timeout,
	}
}

// NewDeleteRegistriesIDParamsWithContext creates a new DeleteRegistriesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRegistriesIDParamsWithContext(ctx context.Context) *DeleteRegistriesIDParams {
	var ()
	return &DeleteRegistriesIDParams{

		Context: ctx,
	}
}

// NewDeleteRegistriesIDParamsWithHTTPClient creates a new DeleteRegistriesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRegistriesIDParamsWithHTTPClient(client *http.Client) *DeleteRegistriesIDParams {
	var ()
	return &DeleteRegistriesIDParams{
		HTTPClient: client,
	}
}

/*DeleteRegistriesIDParams contains all the parameters to send to the API endpoint
for the delete registries ID operation typically these are written to a http.Request
*/
type DeleteRegistriesIDParams struct {

	/*ID
	  The registry's ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete registries ID params
func (o *DeleteRegistriesIDParams) WithTimeout(timeout time.Duration) *DeleteRegistriesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete registries ID params
func (o *DeleteRegistriesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete registries ID params
func (o *DeleteRegistriesIDParams) WithContext(ctx context.Context) *DeleteRegistriesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete registries ID params
func (o *DeleteRegistriesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete registries ID params
func (o *DeleteRegistriesIDParams) WithHTTPClient(client *http.Client) *DeleteRegistriesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete registries ID params
func (o *DeleteRegistriesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete registries ID params
func (o *DeleteRegistriesIDParams) WithID(id int64) *DeleteRegistriesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete registries ID params
func (o *DeleteRegistriesIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRegistriesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

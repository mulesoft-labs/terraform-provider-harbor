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

// NewGetRetentionsIDParams creates a new GetRetentionsIDParams object
// with the default values initialized.
func NewGetRetentionsIDParams() *GetRetentionsIDParams {
	var ()
	return &GetRetentionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRetentionsIDParamsWithTimeout creates a new GetRetentionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRetentionsIDParamsWithTimeout(timeout time.Duration) *GetRetentionsIDParams {
	var ()
	return &GetRetentionsIDParams{

		timeout: timeout,
	}
}

// NewGetRetentionsIDParamsWithContext creates a new GetRetentionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRetentionsIDParamsWithContext(ctx context.Context) *GetRetentionsIDParams {
	var ()
	return &GetRetentionsIDParams{

		Context: ctx,
	}
}

// NewGetRetentionsIDParamsWithHTTPClient creates a new GetRetentionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRetentionsIDParamsWithHTTPClient(client *http.Client) *GetRetentionsIDParams {
	var ()
	return &GetRetentionsIDParams{
		HTTPClient: client,
	}
}

/*GetRetentionsIDParams contains all the parameters to send to the API endpoint
for the get retentions ID operation typically these are written to a http.Request
*/
type GetRetentionsIDParams struct {

	/*ID
	  Retention ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get retentions ID params
func (o *GetRetentionsIDParams) WithTimeout(timeout time.Duration) *GetRetentionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get retentions ID params
func (o *GetRetentionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get retentions ID params
func (o *GetRetentionsIDParams) WithContext(ctx context.Context) *GetRetentionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get retentions ID params
func (o *GetRetentionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get retentions ID params
func (o *GetRetentionsIDParams) WithHTTPClient(client *http.Client) *GetRetentionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get retentions ID params
func (o *GetRetentionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get retentions ID params
func (o *GetRetentionsIDParams) WithID(id int64) *GetRetentionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get retentions ID params
func (o *GetRetentionsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRetentionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostInternalSyncquotaParams creates a new PostInternalSyncquotaParams object
// with the default values initialized.
func NewPostInternalSyncquotaParams() *PostInternalSyncquotaParams {

	return &PostInternalSyncquotaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostInternalSyncquotaParamsWithTimeout creates a new PostInternalSyncquotaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostInternalSyncquotaParamsWithTimeout(timeout time.Duration) *PostInternalSyncquotaParams {

	return &PostInternalSyncquotaParams{

		timeout: timeout,
	}
}

// NewPostInternalSyncquotaParamsWithContext creates a new PostInternalSyncquotaParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostInternalSyncquotaParamsWithContext(ctx context.Context) *PostInternalSyncquotaParams {

	return &PostInternalSyncquotaParams{

		Context: ctx,
	}
}

// NewPostInternalSyncquotaParamsWithHTTPClient creates a new PostInternalSyncquotaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostInternalSyncquotaParamsWithHTTPClient(client *http.Client) *PostInternalSyncquotaParams {

	return &PostInternalSyncquotaParams{
		HTTPClient: client,
	}
}

/*PostInternalSyncquotaParams contains all the parameters to send to the API endpoint
for the post internal syncquota operation typically these are written to a http.Request
*/
type PostInternalSyncquotaParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post internal syncquota params
func (o *PostInternalSyncquotaParams) WithTimeout(timeout time.Duration) *PostInternalSyncquotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post internal syncquota params
func (o *PostInternalSyncquotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post internal syncquota params
func (o *PostInternalSyncquotaParams) WithContext(ctx context.Context) *PostInternalSyncquotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post internal syncquota params
func (o *PostInternalSyncquotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post internal syncquota params
func (o *PostInternalSyncquotaParams) WithHTTPClient(client *http.Client) *PostInternalSyncquotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post internal syncquota params
func (o *PostInternalSyncquotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostInternalSyncquotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

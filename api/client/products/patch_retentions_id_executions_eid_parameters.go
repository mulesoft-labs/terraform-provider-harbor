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

// NewPatchRetentionsIDExecutionsEidParams creates a new PatchRetentionsIDExecutionsEidParams object
// with the default values initialized.
func NewPatchRetentionsIDExecutionsEidParams() *PatchRetentionsIDExecutionsEidParams {
	var ()
	return &PatchRetentionsIDExecutionsEidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRetentionsIDExecutionsEidParamsWithTimeout creates a new PatchRetentionsIDExecutionsEidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchRetentionsIDExecutionsEidParamsWithTimeout(timeout time.Duration) *PatchRetentionsIDExecutionsEidParams {
	var ()
	return &PatchRetentionsIDExecutionsEidParams{

		timeout: timeout,
	}
}

// NewPatchRetentionsIDExecutionsEidParamsWithContext creates a new PatchRetentionsIDExecutionsEidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchRetentionsIDExecutionsEidParamsWithContext(ctx context.Context) *PatchRetentionsIDExecutionsEidParams {
	var ()
	return &PatchRetentionsIDExecutionsEidParams{

		Context: ctx,
	}
}

// NewPatchRetentionsIDExecutionsEidParamsWithHTTPClient creates a new PatchRetentionsIDExecutionsEidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchRetentionsIDExecutionsEidParamsWithHTTPClient(client *http.Client) *PatchRetentionsIDExecutionsEidParams {
	var ()
	return &PatchRetentionsIDExecutionsEidParams{
		HTTPClient: client,
	}
}

/*PatchRetentionsIDExecutionsEidParams contains all the parameters to send to the API endpoint
for the patch retentions ID executions eid operation typically these are written to a http.Request
*/
type PatchRetentionsIDExecutionsEidParams struct {

	/*Action
	  The action, only support "stop" now.

	*/
	Action PatchRetentionsIDExecutionsEidBody
	/*Eid
	  Retention execution ID.

	*/
	Eid int64
	/*ID
	  Retention ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) WithTimeout(timeout time.Duration) *PatchRetentionsIDExecutionsEidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) WithContext(ctx context.Context) *PatchRetentionsIDExecutionsEidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) WithHTTPClient(client *http.Client) *PatchRetentionsIDExecutionsEidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) WithAction(action PatchRetentionsIDExecutionsEidBody) *PatchRetentionsIDExecutionsEidParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) SetAction(action PatchRetentionsIDExecutionsEidBody) {
	o.Action = action
}

// WithEid adds the eid to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) WithEid(eid int64) *PatchRetentionsIDExecutionsEidParams {
	o.SetEid(eid)
	return o
}

// SetEid adds the eid to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) SetEid(eid int64) {
	o.Eid = eid
}

// WithID adds the id to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) WithID(id int64) *PatchRetentionsIDExecutionsEidParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch retentions ID executions eid params
func (o *PatchRetentionsIDExecutionsEidParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRetentionsIDExecutionsEidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Action); err != nil {
		return err
	}

	// path param eid
	if err := r.SetPathParam("eid", swag.FormatInt64(o.Eid)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
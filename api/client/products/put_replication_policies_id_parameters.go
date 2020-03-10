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

	"github.com/sandhose/terraform-provider-harbor/api/models"
)

// NewPutReplicationPoliciesIDParams creates a new PutReplicationPoliciesIDParams object
// with the default values initialized.
func NewPutReplicationPoliciesIDParams() *PutReplicationPoliciesIDParams {
	var ()
	return &PutReplicationPoliciesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutReplicationPoliciesIDParamsWithTimeout creates a new PutReplicationPoliciesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutReplicationPoliciesIDParamsWithTimeout(timeout time.Duration) *PutReplicationPoliciesIDParams {
	var ()
	return &PutReplicationPoliciesIDParams{

		timeout: timeout,
	}
}

// NewPutReplicationPoliciesIDParamsWithContext creates a new PutReplicationPoliciesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutReplicationPoliciesIDParamsWithContext(ctx context.Context) *PutReplicationPoliciesIDParams {
	var ()
	return &PutReplicationPoliciesIDParams{

		Context: ctx,
	}
}

// NewPutReplicationPoliciesIDParamsWithHTTPClient creates a new PutReplicationPoliciesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutReplicationPoliciesIDParamsWithHTTPClient(client *http.Client) *PutReplicationPoliciesIDParams {
	var ()
	return &PutReplicationPoliciesIDParams{
		HTTPClient: client,
	}
}

/*PutReplicationPoliciesIDParams contains all the parameters to send to the API endpoint
for the put replication policies ID operation typically these are written to a http.Request
*/
type PutReplicationPoliciesIDParams struct {

	/*ID
	  policy ID

	*/
	ID int64
	/*Policy
	  The replication policy model.

	*/
	Policy *models.ReplicationPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) WithTimeout(timeout time.Duration) *PutReplicationPoliciesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) WithContext(ctx context.Context) *PutReplicationPoliciesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) WithHTTPClient(client *http.Client) *PutReplicationPoliciesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) WithID(id int64) *PutReplicationPoliciesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) SetID(id int64) {
	o.ID = id
}

// WithPolicy adds the policy to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) WithPolicy(policy *models.ReplicationPolicy) *PutReplicationPoliciesIDParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the put replication policies ID params
func (o *PutReplicationPoliciesIDParams) SetPolicy(policy *models.ReplicationPolicy) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *PutReplicationPoliciesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

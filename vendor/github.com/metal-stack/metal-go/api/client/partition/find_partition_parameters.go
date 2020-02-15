// Code generated by go-swagger; DO NOT EDIT.

package partition

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

// NewFindPartitionParams creates a new FindPartitionParams object
// with the default values initialized.
func NewFindPartitionParams() *FindPartitionParams {
	var ()
	return &FindPartitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindPartitionParamsWithTimeout creates a new FindPartitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindPartitionParamsWithTimeout(timeout time.Duration) *FindPartitionParams {
	var ()
	return &FindPartitionParams{

		timeout: timeout,
	}
}

// NewFindPartitionParamsWithContext creates a new FindPartitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindPartitionParamsWithContext(ctx context.Context) *FindPartitionParams {
	var ()
	return &FindPartitionParams{

		Context: ctx,
	}
}

// NewFindPartitionParamsWithHTTPClient creates a new FindPartitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindPartitionParamsWithHTTPClient(client *http.Client) *FindPartitionParams {
	var ()
	return &FindPartitionParams{
		HTTPClient: client,
	}
}

/*FindPartitionParams contains all the parameters to send to the API endpoint
for the find partition operation typically these are written to a http.Request
*/
type FindPartitionParams struct {

	/*ID
	  identifier of the Partition

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find partition params
func (o *FindPartitionParams) WithTimeout(timeout time.Duration) *FindPartitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find partition params
func (o *FindPartitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find partition params
func (o *FindPartitionParams) WithContext(ctx context.Context) *FindPartitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find partition params
func (o *FindPartitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find partition params
func (o *FindPartitionParams) WithHTTPClient(client *http.Client) *FindPartitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find partition params
func (o *FindPartitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find partition params
func (o *FindPartitionParams) WithID(id string) *FindPartitionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find partition params
func (o *FindPartitionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindPartitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
// Code generated by go-swagger; DO NOT EDIT.

package namespace_services

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

// NewWatchNamespaceServiceRevisionListParams creates a new WatchNamespaceServiceRevisionListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWatchNamespaceServiceRevisionListParams() *WatchNamespaceServiceRevisionListParams {
	return &WatchNamespaceServiceRevisionListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWatchNamespaceServiceRevisionListParamsWithTimeout creates a new WatchNamespaceServiceRevisionListParams object
// with the ability to set a timeout on a request.
func NewWatchNamespaceServiceRevisionListParamsWithTimeout(timeout time.Duration) *WatchNamespaceServiceRevisionListParams {
	return &WatchNamespaceServiceRevisionListParams{
		timeout: timeout,
	}
}

// NewWatchNamespaceServiceRevisionListParamsWithContext creates a new WatchNamespaceServiceRevisionListParams object
// with the ability to set a context for a request.
func NewWatchNamespaceServiceRevisionListParamsWithContext(ctx context.Context) *WatchNamespaceServiceRevisionListParams {
	return &WatchNamespaceServiceRevisionListParams{
		Context: ctx,
	}
}

// NewWatchNamespaceServiceRevisionListParamsWithHTTPClient creates a new WatchNamespaceServiceRevisionListParams object
// with the ability to set a custom HTTPClient for a request.
func NewWatchNamespaceServiceRevisionListParamsWithHTTPClient(client *http.Client) *WatchNamespaceServiceRevisionListParams {
	return &WatchNamespaceServiceRevisionListParams{
		HTTPClient: client,
	}
}

/* WatchNamespaceServiceRevisionListParams contains all the parameters to send to the API endpoint
   for the watch namespace service revision list operation.

   Typically these are written to a http.Request.
*/
type WatchNamespaceServiceRevisionListParams struct {

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* ServiceName.

	   target service name
	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the watch namespace service revision list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WatchNamespaceServiceRevisionListParams) WithDefaults() *WatchNamespaceServiceRevisionListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the watch namespace service revision list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WatchNamespaceServiceRevisionListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) WithTimeout(timeout time.Duration) *WatchNamespaceServiceRevisionListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) WithContext(ctx context.Context) *WatchNamespaceServiceRevisionListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) WithHTTPClient(client *http.Client) *WatchNamespaceServiceRevisionListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) WithNamespace(namespace string) *WatchNamespaceServiceRevisionListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithServiceName adds the serviceName to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) WithServiceName(serviceName string) *WatchNamespaceServiceRevisionListParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the watch namespace service revision list params
func (o *WatchNamespaceServiceRevisionListParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *WatchNamespaceServiceRevisionListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package payer_authentication

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

// NewCheckPayerAuthEnrollmentParams creates a new CheckPayerAuthEnrollmentParams object
// with the default values initialized.
func NewCheckPayerAuthEnrollmentParams() *CheckPayerAuthEnrollmentParams {
	var ()
	return &CheckPayerAuthEnrollmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckPayerAuthEnrollmentParamsWithTimeout creates a new CheckPayerAuthEnrollmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckPayerAuthEnrollmentParamsWithTimeout(timeout time.Duration) *CheckPayerAuthEnrollmentParams {
	var ()
	return &CheckPayerAuthEnrollmentParams{

		timeout: timeout,
	}
}

// NewCheckPayerAuthEnrollmentParamsWithContext creates a new CheckPayerAuthEnrollmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckPayerAuthEnrollmentParamsWithContext(ctx context.Context) *CheckPayerAuthEnrollmentParams {
	var ()
	return &CheckPayerAuthEnrollmentParams{

		Context: ctx,
	}
}

// NewCheckPayerAuthEnrollmentParamsWithHTTPClient creates a new CheckPayerAuthEnrollmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckPayerAuthEnrollmentParamsWithHTTPClient(client *http.Client) *CheckPayerAuthEnrollmentParams {
	var ()
	return &CheckPayerAuthEnrollmentParams{
		HTTPClient: client,
	}
}

/*CheckPayerAuthEnrollmentParams contains all the parameters to send to the API endpoint
for the check payer auth enrollment operation typically these are written to a http.Request
*/
type CheckPayerAuthEnrollmentParams struct {

	/*CheckPayerAuthEnrollmentRequest*/
	CheckPayerAuthEnrollmentRequest CheckPayerAuthEnrollmentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check payer auth enrollment params
func (o *CheckPayerAuthEnrollmentParams) WithTimeout(timeout time.Duration) *CheckPayerAuthEnrollmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check payer auth enrollment params
func (o *CheckPayerAuthEnrollmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check payer auth enrollment params
func (o *CheckPayerAuthEnrollmentParams) WithContext(ctx context.Context) *CheckPayerAuthEnrollmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check payer auth enrollment params
func (o *CheckPayerAuthEnrollmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check payer auth enrollment params
func (o *CheckPayerAuthEnrollmentParams) WithHTTPClient(client *http.Client) *CheckPayerAuthEnrollmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check payer auth enrollment params
func (o *CheckPayerAuthEnrollmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckPayerAuthEnrollmentRequest adds the checkPayerAuthEnrollmentRequest to the check payer auth enrollment params
func (o *CheckPayerAuthEnrollmentParams) WithCheckPayerAuthEnrollmentRequest(checkPayerAuthEnrollmentRequest CheckPayerAuthEnrollmentBody) *CheckPayerAuthEnrollmentParams {
	o.SetCheckPayerAuthEnrollmentRequest(checkPayerAuthEnrollmentRequest)
	return o
}

// SetCheckPayerAuthEnrollmentRequest adds the checkPayerAuthEnrollmentRequest to the check payer auth enrollment params
func (o *CheckPayerAuthEnrollmentParams) SetCheckPayerAuthEnrollmentRequest(checkPayerAuthEnrollmentRequest CheckPayerAuthEnrollmentBody) {
	o.CheckPayerAuthEnrollmentRequest = checkPayerAuthEnrollmentRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CheckPayerAuthEnrollmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CheckPayerAuthEnrollmentRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

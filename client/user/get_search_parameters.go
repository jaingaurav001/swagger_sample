// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSearchParams creates a new GetSearchParams object
// with the default values initialized.
func NewGetSearchParams() *GetSearchParams {
	var ()
	return &GetSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSearchParamsWithTimeout creates a new GetSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSearchParamsWithTimeout(timeout time.Duration) *GetSearchParams {
	var ()
	return &GetSearchParams{

		timeout: timeout,
	}
}

// NewGetSearchParamsWithContext creates a new GetSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSearchParamsWithContext(ctx context.Context) *GetSearchParams {
	var ()
	return &GetSearchParams{

		Context: ctx,
	}
}

// NewGetSearchParamsWithHTTPClient creates a new GetSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSearchParamsWithHTTPClient(client *http.Client) *GetSearchParams {
	var ()
	return &GetSearchParams{
		HTTPClient: client,
	}
}

/*GetSearchParams contains all the parameters to send to the API endpoint
for the get search operation typically these are written to a http.Request
*/
type GetSearchParams struct {

	/*UserID
	  user id

	*/
	UserID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get search params
func (o *GetSearchParams) WithTimeout(timeout time.Duration) *GetSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get search params
func (o *GetSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get search params
func (o *GetSearchParams) WithContext(ctx context.Context) *GetSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get search params
func (o *GetSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get search params
func (o *GetSearchParams) WithHTTPClient(client *http.Client) *GetSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get search params
func (o *GetSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get search params
func (o *GetSearchParams) WithUserID(userID int32) *GetSearchParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get search params
func (o *GetSearchParams) SetUserID(userID int32) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param user_id
	frUserID := o.UserID
	fUserID := swag.FormatInt32(frUserID)
	if fUserID != "" {
		if err := r.SetFormParam("user_id", fUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSearchParams creates a new GetSearchParams object
// no default values defined in spec.
func NewGetSearchParams() GetSearchParams {

	return GetSearchParams{}
}

// GetSearchParams contains all the bound params for the get search operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetSearch
type GetSearchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*user id
	  Required: true
	  In: query
	*/
	UserID int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSearchParams() beforehand.
func (o *GetSearchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qUserID, qhkUserID, _ := qs.GetOK("user_id")
	if err := o.bindUserID(qUserID, qhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSearchParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("user_id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("user_id", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("user_id", "query", "int32", raw)
	}
	o.UserID = value

	return nil
}

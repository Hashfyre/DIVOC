// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetVaccinatorsParams creates a new GetVaccinatorsParams object
// no default values defined in spec.
func NewGetVaccinatorsParams() GetVaccinatorsParams {

	return GetVaccinatorsParams{}
}

// GetVaccinatorsParams contains all the bound params for the get vaccinators operation
// typically these are obtained from a http.Request
//
// swagger:parameters getVaccinators
type GetVaccinatorsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVaccinatorsParams() beforehand.
func (o *GetVaccinatorsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

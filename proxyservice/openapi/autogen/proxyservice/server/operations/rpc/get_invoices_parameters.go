// Code generated by go-swagger; DO NOT EDIT.

package rpc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInvoicesParams creates a new GetInvoicesParams object
// no default values defined in spec.
func NewGetInvoicesParams() GetInvoicesParams {

	return GetInvoicesParams{}
}

// GetInvoicesParams contains all the bound params for the get invoices operation
// typically these are obtained from a http.Request
//
// swagger:parameters getInvoices
type GetInvoicesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*An identifier of client that invoices will be returned for
	  Required: true
	  In: path
	*/
	ClientID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetInvoicesParams() beforehand.
func (o *GetInvoicesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rClientID, rhkClientID, _ := route.Params.GetOK("client_id")
	if err := o.bindClientID(rClientID, rhkClientID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInvoicesParams) bindClientID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ClientID = raw

	return nil
}

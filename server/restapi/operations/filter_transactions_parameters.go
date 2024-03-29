// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFilterTransactionsParams creates a new FilterTransactionsParams object
// no default values defined in spec.
func NewFilterTransactionsParams() FilterTransactionsParams {

	return FilterTransactionsParams{}
}

// FilterTransactionsParams contains all the bound params for the filter transactions operation
// typically these are obtained from a http.Request
//
// swagger:parameters filterTransactions
type FilterTransactionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Category *string
	/*
	  In: query
	*/
	From *int64
	/*
	  In: query
	*/
	Receiver *string
	/*
	  In: query
	*/
	Sender *string
	/*
	  In: query
	*/
	To *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFilterTransactionsParams() beforehand.
func (o *FilterTransactionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCategory, qhkCategory, _ := qs.GetOK("category")
	if err := o.bindCategory(qCategory, qhkCategory, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	qReceiver, qhkReceiver, _ := qs.GetOK("receiver")
	if err := o.bindReceiver(qReceiver, qhkReceiver, route.Formats); err != nil {
		res = append(res, err)
	}

	qSender, qhkSender, _ := qs.GetOK("sender")
	if err := o.bindSender(qSender, qhkSender, route.Formats); err != nil {
		res = append(res, err)
	}

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCategory binds and validates parameter Category from query.
func (o *FilterTransactionsParams) bindCategory(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Category = &raw

	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *FilterTransactionsParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("from", "query", "int64", raw)
	}
	o.From = &value

	return nil
}

// bindReceiver binds and validates parameter Receiver from query.
func (o *FilterTransactionsParams) bindReceiver(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Receiver = &raw

	return nil
}

// bindSender binds and validates parameter Sender from query.
func (o *FilterTransactionsParams) bindSender(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Sender = &raw

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *FilterTransactionsParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("to", "query", "int64", raw)
	}
	o.To = &value

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// FilterTransactionsURL generates an URL for the filter transactions operation
type FilterTransactionsURL struct {
	Category *string
	From     *int64
	Receiver *string
	Sender   *string
	To       *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FilterTransactionsURL) WithBasePath(bp string) *FilterTransactionsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FilterTransactionsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FilterTransactionsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/transactions/filter"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var categoryQ string
	if o.Category != nil {
		categoryQ = *o.Category
	}
	if categoryQ != "" {
		qs.Set("category", categoryQ)
	}

	var fromQ string
	if o.From != nil {
		fromQ = swag.FormatInt64(*o.From)
	}
	if fromQ != "" {
		qs.Set("from", fromQ)
	}

	var receiverQ string
	if o.Receiver != nil {
		receiverQ = *o.Receiver
	}
	if receiverQ != "" {
		qs.Set("receiver", receiverQ)
	}

	var senderQ string
	if o.Sender != nil {
		senderQ = *o.Sender
	}
	if senderQ != "" {
		qs.Set("sender", senderQ)
	}

	var toQ string
	if o.To != nil {
		toQ = swag.FormatInt64(*o.To)
	}
	if toQ != "" {
		qs.Set("to", toQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FilterTransactionsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FilterTransactionsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FilterTransactionsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FilterTransactionsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FilterTransactionsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *FilterTransactionsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
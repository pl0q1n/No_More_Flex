// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Transaction transaction
// swagger:model transaction
type Transaction struct {

	// category
	Category string `json:"category,omitempty"`

	// receiver
	// Required: true
	Receiver *string `json:"receiver"`

	// sender
	// Required: true
	Sender *string `json:"sender"`

	// time
	// Required: true
	Time *int64 `json:"time"`

	// value
	// Required: true
	Value *int64 `json:"value"`
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReceiver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) validateReceiver(formats strfmt.Registry) error {

	if err := validate.Required("receiver", "body", m.Receiver); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

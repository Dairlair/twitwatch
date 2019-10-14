// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SignupUserInputError signup user input error
// swagger:model SignupUserInputError
type SignupUserInputError struct {

	// code
	// Required: true
	// Enum: [EMAIL_ALREADY_TAKEN]
	Code *string `json:"code"`
}

// Validate validates this signup user input error
func (m *SignupUserInputError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var signupUserInputErrorTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EMAIL_ALREADY_TAKEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		signupUserInputErrorTypeCodePropEnum = append(signupUserInputErrorTypeCodePropEnum, v)
	}
}

const (

	// SignupUserInputErrorCodeEMAILALREADYTAKEN captures enum value "EMAIL_ALREADY_TAKEN"
	SignupUserInputErrorCodeEMAILALREADYTAKEN string = "EMAIL_ALREADY_TAKEN"
)

// prop value enum
func (m *SignupUserInputError) validateCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, signupUserInputErrorTypeCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SignupUserInputError) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SignupUserInputError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignupUserInputError) UnmarshalBinary(b []byte) error {
	var res SignupUserInputError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

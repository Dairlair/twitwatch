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

// CreateTopicRequest create topic request
// swagger:model CreateTopicRequest
type CreateTopicRequest struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// tracks
	// Required: true
	Tracks []string `json:"tracks"`
}

// Validate validates this create topic request
func (m *CreateTopicRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTracks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTopicRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateTopicRequest) validateTracks(formats strfmt.Registry) error {

	if err := validate.Required("tracks", "body", m.Tracks); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTopicRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTopicRequest) UnmarshalBinary(b []byte) error {
	var res CreateTopicRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

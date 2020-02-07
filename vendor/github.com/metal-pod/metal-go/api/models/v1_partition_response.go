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

// V1PartitionResponse v1 partition response
// swagger:model v1.PartitionResponse
type V1PartitionResponse struct {

	// the boot configuration of this partition
	// Required: true
	Bootconfig *V1PartitionBootConfiguration `json:"bootconfig"`

	// the last changed timestamp of this entity
	// Required: true
	// Read Only: true
	// Format: date-time
	Changed strfmt.DateTime `json:"changed"`

	// the creation time of this entity
	// Required: true
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created"`

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the unique ID of this entity
	// Required: true
	// Unique: true
	ID *string `json:"id"`

	// the address to the management service of this partition
	Mgmtserviceaddress string `json:"mgmtserviceaddress,omitempty"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`

	// the length of private networks for the machine's child networks in this partition, default 22
	// Maximum: 30
	// Minimum: 16
	Privatenetworkprefixlength int32 `json:"privatenetworkprefixlength,omitempty"`
}

// Validate validates this v1 partition response
func (m *V1PartitionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootconfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivatenetworkprefixlength(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PartitionResponse) validateBootconfig(formats strfmt.Registry) error {

	if err := validate.Required("bootconfig", "body", m.Bootconfig); err != nil {
		return err
	}

	if m.Bootconfig != nil {
		if err := m.Bootconfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootconfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1PartitionResponse) validateChanged(formats strfmt.Registry) error {

	if err := validate.Required("changed", "body", strfmt.DateTime(m.Changed)); err != nil {
		return err
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1PartitionResponse) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1PartitionResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1PartitionResponse) validatePrivatenetworkprefixlength(formats strfmt.Registry) error {

	if swag.IsZero(m.Privatenetworkprefixlength) { // not required
		return nil
	}

	if err := validate.MinimumInt("privatenetworkprefixlength", "body", int64(m.Privatenetworkprefixlength), 16, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("privatenetworkprefixlength", "body", int64(m.Privatenetworkprefixlength), 30, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PartitionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PartitionResponse) UnmarshalBinary(b []byte) error {
	var res V1PartitionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

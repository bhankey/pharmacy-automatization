// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PharmacyGetAllResponse pharmacy get all response
//
// swagger:model PharmacyGetAllResponse
type PharmacyGetAllResponse struct {

	// pharmacies
	Pharmacies []*Pharmacy `json:"pharmacies"`
}

// Validate validates this pharmacy get all response
func (m *PharmacyGetAllResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePharmacies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PharmacyGetAllResponse) validatePharmacies(formats strfmt.Registry) error {
	if swag.IsZero(m.Pharmacies) { // not required
		return nil
	}

	for i := 0; i < len(m.Pharmacies); i++ {
		if swag.IsZero(m.Pharmacies[i]) { // not required
			continue
		}

		if m.Pharmacies[i] != nil {
			if err := m.Pharmacies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pharmacies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pharmacies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pharmacy get all response based on the context it is used
func (m *PharmacyGetAllResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePharmacies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PharmacyGetAllResponse) contextValidatePharmacies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pharmacies); i++ {

		if m.Pharmacies[i] != nil {
			if err := m.Pharmacies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pharmacies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pharmacies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PharmacyGetAllResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PharmacyGetAllResponse) UnmarshalBinary(b []byte) error {
	var res PharmacyGetAllResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

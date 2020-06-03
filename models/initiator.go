// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Initiator initiator
//
// swagger:model Initiator
type Initiator struct {

	// merchant initiated transaction
	MerchantInitiatedTransaction *InitiatorMerchantInitiatedTransaction `json:"merchantInitiatedTransaction,omitempty"`
}

// Validate validates this initiator
func (m *Initiator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMerchantInitiatedTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Initiator) validateMerchantInitiatedTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.MerchantInitiatedTransaction) { // not required
		return nil
	}

	if m.MerchantInitiatedTransaction != nil {
		if err := m.MerchantInitiatedTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("merchantInitiatedTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Initiator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Initiator) UnmarshalBinary(b []byte) error {
	var res Initiator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InitiatorMerchantInitiatedTransaction initiator merchant initiated transaction
//
// swagger:model InitiatorMerchantInitiatedTransaction
type InitiatorMerchantInitiatedTransaction struct {

	// Previous Consumer Initiated Transaction Id.
	// Max Length: 15
	PreviousTransactionID string `json:"previousTransactionId,omitempty"`
}

// Validate validates this initiator merchant initiated transaction
func (m *InitiatorMerchantInitiatedTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreviousTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitiatorMerchantInitiatedTransaction) validatePreviousTransactionID(formats strfmt.Registry) error {

	if swag.IsZero(m.PreviousTransactionID) { // not required
		return nil
	}

	if err := validate.MaxLength("merchantInitiatedTransaction"+"."+"previousTransactionId", "body", string(m.PreviousTransactionID), 15); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InitiatorMerchantInitiatedTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitiatorMerchantInitiatedTransaction) UnmarshalBinary(b []byte) error {
	var res InitiatorMerchantInitiatedTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

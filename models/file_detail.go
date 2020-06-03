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

// FileDetail file detail
//
// swagger:model FileDetail
type FileDetail struct {

	// Date and time for the file in PST
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"createdTime,omitempty"`

	// Date and time for the file in PST
	// Format: date
	Date strfmt.Date `json:"date,omitempty"`

	// Unique identifier of a file
	FileID string `json:"fileId,omitempty"`

	// Date and time for the file in PST
	// Format: date-time
	LastModifiedTime strfmt.DateTime `json:"lastModifiedTime,omitempty"`

	// 'File extension'
	//
	// Valid values:
	// - 'application/xml'
	// - 'text/csv'
	// - 'application/pdf'
	// - 'application/octet-stream'
	//
	MimeType string `json:"mimeType,omitempty"`

	// Name of the file
	Name string `json:"name,omitempty"`

	// Size of the file in bytes
	Size int64 `json:"size,omitempty"`
}

// Validate validates this file detail
func (m *FileDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileDetail) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createdTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FileDetail) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FileDetail) validateLastModifiedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModifiedTime", "body", "date-time", m.LastModifiedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileDetail) UnmarshalBinary(b []byte) error {
	var res FileDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

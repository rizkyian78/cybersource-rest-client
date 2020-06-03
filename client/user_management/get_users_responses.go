// Code generated by go-swagger; DO NOT EDIT.

package user_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*GetUsersOK handles this case with default header values.

OK
*/
type GetUsersOK struct {
	Payload *GetUsersOKBody
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /ums/v1/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) GetPayload() *GetUsersOKBody {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUsersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersBadRequest creates a GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

/*GetUsersBadRequest handles this case with default header values.

Invalid request.
*/
type GetUsersBadRequest struct {
	Payload *GetUsersBadRequestBody
}

func (o *GetUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /ums/v1/users][%d] getUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersBadRequest) GetPayload() *GetUsersBadRequestBody {
	return o.Payload
}

func (o *GetUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUsersBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersInternalServerError creates a GetUsersInternalServerError with default headers values
func NewGetUsersInternalServerError() *GetUsersInternalServerError {
	return &GetUsersInternalServerError{}
}

/*GetUsersInternalServerError handles this case with default header values.

Unexpected server error.
*/
type GetUsersInternalServerError struct {
}

func (o *GetUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ums/v1/users][%d] getUsersInternalServerError ", 500)
}

func (o *GetUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	Field string `json:"field,omitempty"`

	// Possible reasons for the error.
	//
	// Possible values:
	//  - MISSING_FIELD
	//  - INVALID_DATA
	//
	Reason string `json:"reason,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUsersBadRequestBody umsV1UsersGet400Response
swagger:model GetUsersBadRequestBody
*/
type GetUsersBadRequestBody struct {

	// details
	Details []*DetailsItems0 `json:"details"`

	// The detail message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of the status.
	//
	// Possible values:
	//  - MISSING_FIELD
	//  - INVALID_DATA
	//  - DUPLICATE_REQUEST
	//  - INVALID_CARD
	//  - INVALID_MERCHANT_CONFIGURATION
	//  - INVALID_AMOUNT
	//  - CAPTURE_ALREADY_VOIDED
	//  - ACCOUNT_NOT_ALLOWED_CREDIT
	//
	Reason string `json:"reason,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - INVALID_REQUEST
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this get users bad request body
func (o *GetUsersBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersBadRequestBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUsersBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetUsersBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUsersOKBody umsV1UsersGet200Response
swagger:model GetUsersOKBody
*/
type GetUsersOKBody struct {

	// users
	Users []*UsersItems0 `json:"users"`
}

// Validate validates this get users o k body
func (o *GetUsersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersOKBody) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(o.Users) { // not required
		return nil
	}

	for i := 0; i < len(o.Users); i++ {
		if swag.IsZero(o.Users[i]) { // not required
			continue
		}

		if o.Users[i] != nil {
			if err := o.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUsersOK" + "." + "users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersOKBody) UnmarshalBinary(b []byte) error {
	var res GetUsersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UsersItems0 users items0
swagger:model UsersItems0
*/
type UsersItems0 struct {

	// account information
	AccountInformation *UsersItems0AccountInformation `json:"accountInformation,omitempty"`

	// contact information
	ContactInformation *UsersItems0ContactInformation `json:"contactInformation,omitempty"`

	// custom fields
	CustomFields map[string]string `json:"customFields,omitempty"`

	// organization information
	OrganizationInformation *UsersItems0OrganizationInformation `json:"organizationInformation,omitempty"`
}

// Validate validates this users items0
func (o *UsersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccountInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateContactInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrganizationInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersItems0) validateAccountInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.AccountInformation) { // not required
		return nil
	}

	if o.AccountInformation != nil {
		if err := o.AccountInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accountInformation")
			}
			return err
		}
	}

	return nil
}

func (o *UsersItems0) validateContactInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ContactInformation) { // not required
		return nil
	}

	if o.ContactInformation != nil {
		if err := o.ContactInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactInformation")
			}
			return err
		}
	}

	return nil
}

func (o *UsersItems0) validateOrganizationInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.OrganizationInformation) { // not required
		return nil
	}

	if o.OrganizationInformation != nil {
		if err := o.OrganizationInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizationInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UsersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersItems0) UnmarshalBinary(b []byte) error {
	var res UsersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UsersItems0AccountInformation users items0 account information
swagger:model UsersItems0AccountInformation
*/
type UsersItems0AccountInformation struct {

	// created time
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"createdTime,omitempty"`

	// language preference
	LanguagePreference string `json:"languagePreference,omitempty"`

	// last access time
	// Format: date-time
	LastAccessTime strfmt.DateTime `json:"lastAccessTime,omitempty"`

	// permissions
	Permissions []string `json:"permissions"`

	// role Id
	RoleID string `json:"roleId,omitempty"`

	// Valid values:
	// - active
	// - inactive
	// - locked
	// - disabled
	// - forgotpassword
	// - deleted
	//
	Status string `json:"status,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this users items0 account information
func (o *UsersItems0AccountInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastAccessTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersItems0AccountInformation) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("accountInformation"+"."+"createdTime", "body", "date-time", o.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UsersItems0AccountInformation) validateLastAccessTime(formats strfmt.Registry) error {

	if swag.IsZero(o.LastAccessTime) { // not required
		return nil
	}

	if err := validate.FormatOf("accountInformation"+"."+"lastAccessTime", "body", "date-time", o.LastAccessTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UsersItems0AccountInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersItems0AccountInformation) UnmarshalBinary(b []byte) error {
	var res UsersItems0AccountInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UsersItems0ContactInformation users items0 contact information
swagger:model UsersItems0ContactInformation
*/
type UsersItems0ContactInformation struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// phone number
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// Validate validates this users items0 contact information
func (o *UsersItems0ContactInformation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersItems0ContactInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersItems0ContactInformation) UnmarshalBinary(b []byte) error {
	var res UsersItems0ContactInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UsersItems0OrganizationInformation users items0 organization information
swagger:model UsersItems0OrganizationInformation
*/
type UsersItems0OrganizationInformation struct {

	// organization Id
	OrganizationID string `json:"organizationId,omitempty"`
}

// Validate validates this users items0 organization information
func (o *UsersItems0OrganizationInformation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersItems0OrganizationInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersItems0OrganizationInformation) UnmarshalBinary(b []byte) error {
	var res UsersItems0OrganizationInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"io"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLeasesReader is a Reader for the DeleteLeases structure.
type DeleteLeasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLeasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLeasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLeasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLeasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLeasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLeasesOK creates a DeleteLeasesOK with default headers values
func NewDeleteLeasesOK() *DeleteLeasesOK {
	return &DeleteLeasesOK{}
}

/*DeleteLeasesOK handles this case with default header values.

DeleteLeasesOK delete leases o k
*/
type DeleteLeasesOK struct {
	Payload *DeleteLeasesOKBody
}

func (o *DeleteLeasesOK) Error() string {
	/*
		#nosec CWE-89: false positive. No sql here.
	*/
	return fmt.Sprintf("[DELETE /leases][%d] deleteLeasesOK  %+v", 200, o.Payload)
}

func (o *DeleteLeasesOK) GetPayload() *DeleteLeasesOKBody {
	return o.Payload
}

func (o *DeleteLeasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteLeasesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLeasesBadRequest creates a DeleteLeasesBadRequest with default headers values
func NewDeleteLeasesBadRequest() *DeleteLeasesBadRequest {
	return &DeleteLeasesBadRequest{}
}

/*DeleteLeasesBadRequest handles this case with default header values.

"Failed to Parse Request Body" if the request body is blank or incorrectly formatted. or if there are no account leases found for the specified accountId or if the account specified is not already Active.

*/
type DeleteLeasesBadRequest struct {
}

func (o *DeleteLeasesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /leases][%d] deleteLeasesBadRequest ", 400)
}

func (o *DeleteLeasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLeasesForbidden creates a DeleteLeasesForbidden with default headers values
func NewDeleteLeasesForbidden() *DeleteLeasesForbidden {
	return &DeleteLeasesForbidden{}
}

/*DeleteLeasesForbidden handles this case with default header values.

Failed to authenticate request
*/
type DeleteLeasesForbidden struct {
}

func (o *DeleteLeasesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /leases][%d] deleteLeasesForbidden ", 403)
}

func (o *DeleteLeasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLeasesInternalServerError creates a DeleteLeasesInternalServerError with default headers values
func NewDeleteLeasesInternalServerError() *DeleteLeasesInternalServerError {
	return &DeleteLeasesInternalServerError{}
}

/*DeleteLeasesInternalServerError handles this case with default header values.

Server errors if the database cannot be reached.
*/
type DeleteLeasesInternalServerError struct {
}

func (o *DeleteLeasesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /leases][%d] deleteLeasesInternalServerError ", 500)
}

func (o *DeleteLeasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteLeasesBody delete leases body
swagger:model DeleteLeasesBody
*/
type DeleteLeasesBody struct {

	// account Id
	// Required: true
	AccountID *string `json:"accountId"`

	// principal Id
	// Required: true
	PrincipalID *string `json:"principalId"`
}

// Validate validates this delete leases body
func (o *DeleteLeasesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrincipalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteLeasesBody) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("lease"+"."+"accountId", "body", o.AccountID); err != nil {
		return err
	}

	return nil
}

func (o *DeleteLeasesBody) validatePrincipalID(formats strfmt.Registry) error {

	if err := validate.Required("lease"+"."+"principalId", "body", o.PrincipalID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteLeasesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteLeasesBody) UnmarshalBinary(b []byte) error {
	var res DeleteLeasesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteLeasesOKBody Lease Details
swagger:model DeleteLeasesOKBody
*/
type DeleteLeasesOKBody struct {

	// accountId of the AWS account
	AccountID string `json:"accountId,omitempty"`

	// budget amount
	BudgetAmount float64 `json:"budgetAmount,omitempty"`

	// budget currency
	BudgetCurrency string `json:"budgetCurrency,omitempty"`

	// budget notification emails
	BudgetNotificationEmails []string `json:"budgetNotificationEmails"`

	// creation date in epoch seconds
	CreatedOn float64 `json:"createdOn,omitempty"`

	// date lease should expire in epoch seconds
	ExpiresOn float64 `json:"expiresOn,omitempty"`

	// Lease ID
	ID string `json:"id,omitempty"`

	// date last modified in epoch seconds
	LastModifiedOn float64 `json:"lastModifiedOn,omitempty"`

	// Status of the Lease.
	// "Active": The principal is leased and has access to the account
	// "Inactive": The lease has become inactive, either through expiring, exceeding budget, or by request.
	//
	// Enum: [Active Inactive]
	LeaseStatus string `json:"leaseStatus,omitempty"`

	// date lease status was last modified in epoch seconds
	LeaseStatusModifiedOn float64 `json:"leaseStatusModifiedOn,omitempty"`

	// A reason behind the lease status.
	// "LeaseExpired": The lease exceeded its expiration time ("expiresOn") and
	// the associated account was reset and returned to the account pool.
	// "LeaseOverBudget": The lease exceeded its budgeted amount and the
	// associated account was reset and returned to the account pool.
	// "LeaseDestroyed": The lease was adminstratively ended, which can be done
	// via the leases API.
	// "LeaseActive": The lease is active.
	// "LeaseRolledBack": A system error occurred while provisioning the lease.
	// and it was rolled back.
	//
	// Enum: [LeaseExpired LeaseOverBudget LeaseDestroyed LeaseActive LeaseRolledBack]
	LeaseStatusReason string `json:"leaseStatusReason,omitempty"`

	// principalId of the lease to get
	PrincipalID string `json:"principalId,omitempty"`
}

// Validate validates this delete leases o k body
func (o *DeleteLeasesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLeaseStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLeaseStatusReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deleteLeasesOKBodyTypeLeaseStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteLeasesOKBodyTypeLeaseStatusPropEnum = append(deleteLeasesOKBodyTypeLeaseStatusPropEnum, v)
	}
}

const (

	// DeleteLeasesOKBodyLeaseStatusActive captures enum value "Active"
	DeleteLeasesOKBodyLeaseStatusActive string = "Active"

	// DeleteLeasesOKBodyLeaseStatusInactive captures enum value "Inactive"
	DeleteLeasesOKBodyLeaseStatusInactive string = "Inactive"
)

// prop value enum
func (o *DeleteLeasesOKBody) validateLeaseStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deleteLeasesOKBodyTypeLeaseStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *DeleteLeasesOKBody) validateLeaseStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.LeaseStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateLeaseStatusEnum("deleteLeasesOK"+"."+"leaseStatus", "body", o.LeaseStatus); err != nil {
		return err
	}

	return nil
}

var deleteLeasesOKBodyTypeLeaseStatusReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LeaseExpired","LeaseOverBudget","LeaseDestroyed","LeaseActive","LeaseRolledBack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteLeasesOKBodyTypeLeaseStatusReasonPropEnum = append(deleteLeasesOKBodyTypeLeaseStatusReasonPropEnum, v)
	}
}

const (

	// DeleteLeasesOKBodyLeaseStatusReasonLeaseExpired captures enum value "LeaseExpired"
	DeleteLeasesOKBodyLeaseStatusReasonLeaseExpired string = "LeaseExpired"

	// DeleteLeasesOKBodyLeaseStatusReasonLeaseOverBudget captures enum value "LeaseOverBudget"
	DeleteLeasesOKBodyLeaseStatusReasonLeaseOverBudget string = "LeaseOverBudget"

	// DeleteLeasesOKBodyLeaseStatusReasonLeaseDestroyed captures enum value "LeaseDestroyed"
	DeleteLeasesOKBodyLeaseStatusReasonLeaseDestroyed string = "LeaseDestroyed"

	// DeleteLeasesOKBodyLeaseStatusReasonLeaseActive captures enum value "LeaseActive"
	DeleteLeasesOKBodyLeaseStatusReasonLeaseActive string = "LeaseActive"

	// DeleteLeasesOKBodyLeaseStatusReasonLeaseRolledBack captures enum value "LeaseRolledBack"
	DeleteLeasesOKBodyLeaseStatusReasonLeaseRolledBack string = "LeaseRolledBack"
)

// prop value enum
func (o *DeleteLeasesOKBody) validateLeaseStatusReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deleteLeasesOKBodyTypeLeaseStatusReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *DeleteLeasesOKBody) validateLeaseStatusReason(formats strfmt.Registry) error {

	if swag.IsZero(o.LeaseStatusReason) { // not required
		return nil
	}

	// value enum
	if err := o.validateLeaseStatusReasonEnum("deleteLeasesOK"+"."+"leaseStatusReason", "body", o.LeaseStatusReason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteLeasesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteLeasesOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteLeasesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

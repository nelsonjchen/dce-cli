// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLeasesIDReader is a Reader for the GetLeasesID structure.
type GetLeasesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLeasesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLeasesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLeasesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLeasesIDOK creates a GetLeasesIDOK with default headers values
func NewGetLeasesIDOK() *GetLeasesIDOK {
	return &GetLeasesIDOK{}
}

/*GetLeasesIDOK handles this case with default header values.

GetLeasesIDOK get leases Id o k
*/
type GetLeasesIDOK struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string

	Payload *GetLeasesIDOKBody
}

func (o *GetLeasesIDOK) Error() string {
	return fmt.Sprintf("[GET /leases/{id}][%d] getLeasesIdOK  %+v", 200, o.Payload)
}

func (o *GetLeasesIDOK) GetPayload() *GetLeasesIDOKBody {
	return o.Payload
}

func (o *GetLeasesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Allow-Headers
	o.AccessControlAllowHeaders = response.GetHeader("Access-Control-Allow-Headers")

	// response header Access-Control-Allow-Methods
	o.AccessControlAllowMethods = response.GetHeader("Access-Control-Allow-Methods")

	// response header Access-Control-Allow-Origin
	o.AccessControlAllowOrigin = response.GetHeader("Access-Control-Allow-Origin")

	o.Payload = new(GetLeasesIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLeasesIDForbidden creates a GetLeasesIDForbidden with default headers values
func NewGetLeasesIDForbidden() *GetLeasesIDForbidden {
	return &GetLeasesIDForbidden{}
}

/*GetLeasesIDForbidden handles this case with default header values.

Failed to retrieve lease
*/
type GetLeasesIDForbidden struct {
}

func (o *GetLeasesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /leases/{id}][%d] getLeasesIdForbidden ", 403)
}

func (o *GetLeasesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetLeasesIDOKBody Lease Details
swagger:model GetLeasesIDOKBody
*/
type GetLeasesIDOKBody struct {

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

// Validate validates this get leases ID o k body
func (o *GetLeasesIDOKBody) Validate(formats strfmt.Registry) error {
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

var getLeasesIdOKBodyTypeLeaseStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getLeasesIdOKBodyTypeLeaseStatusPropEnum = append(getLeasesIdOKBodyTypeLeaseStatusPropEnum, v)
	}
}

const (

	// GetLeasesIDOKBodyLeaseStatusActive captures enum value "Active"
	GetLeasesIDOKBodyLeaseStatusActive string = "Active"

	// GetLeasesIDOKBodyLeaseStatusInactive captures enum value "Inactive"
	GetLeasesIDOKBodyLeaseStatusInactive string = "Inactive"
)

// prop value enum
func (o *GetLeasesIDOKBody) validateLeaseStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getLeasesIdOKBodyTypeLeaseStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetLeasesIDOKBody) validateLeaseStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.LeaseStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateLeaseStatusEnum("getLeasesIdOK"+"."+"leaseStatus", "body", o.LeaseStatus); err != nil {
		return err
	}

	return nil
}

var getLeasesIdOKBodyTypeLeaseStatusReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LeaseExpired","LeaseOverBudget","LeaseDestroyed","LeaseActive","LeaseRolledBack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getLeasesIdOKBodyTypeLeaseStatusReasonPropEnum = append(getLeasesIdOKBodyTypeLeaseStatusReasonPropEnum, v)
	}
}

const (

	// GetLeasesIDOKBodyLeaseStatusReasonLeaseExpired captures enum value "LeaseExpired"
	GetLeasesIDOKBodyLeaseStatusReasonLeaseExpired string = "LeaseExpired"

	// GetLeasesIDOKBodyLeaseStatusReasonLeaseOverBudget captures enum value "LeaseOverBudget"
	GetLeasesIDOKBodyLeaseStatusReasonLeaseOverBudget string = "LeaseOverBudget"

	// GetLeasesIDOKBodyLeaseStatusReasonLeaseDestroyed captures enum value "LeaseDestroyed"
	GetLeasesIDOKBodyLeaseStatusReasonLeaseDestroyed string = "LeaseDestroyed"

	// GetLeasesIDOKBodyLeaseStatusReasonLeaseActive captures enum value "LeaseActive"
	GetLeasesIDOKBodyLeaseStatusReasonLeaseActive string = "LeaseActive"

	// GetLeasesIDOKBodyLeaseStatusReasonLeaseRolledBack captures enum value "LeaseRolledBack"
	GetLeasesIDOKBodyLeaseStatusReasonLeaseRolledBack string = "LeaseRolledBack"
)

// prop value enum
func (o *GetLeasesIDOKBody) validateLeaseStatusReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getLeasesIdOKBodyTypeLeaseStatusReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetLeasesIDOKBody) validateLeaseStatusReason(formats strfmt.Registry) error {

	if swag.IsZero(o.LeaseStatusReason) { // not required
		return nil
	}

	// value enum
	if err := o.validateLeaseStatusReasonEnum("getLeasesIdOK"+"."+"leaseStatusReason", "body", o.LeaseStatusReason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLeasesIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLeasesIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetLeasesIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

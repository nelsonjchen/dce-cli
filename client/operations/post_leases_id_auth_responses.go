// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLeasesIDAuthReader is a Reader for the PostLeasesIDAuth structure.
type PostLeasesIDAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLeasesIDAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLeasesIDAuthCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostLeasesIDAuthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLeasesIDAuthForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLeasesIDAuthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLeasesIDAuthCreated creates a PostLeasesIDAuthCreated with default headers values
func NewPostLeasesIDAuthCreated() *PostLeasesIDAuthCreated {
	return &PostLeasesIDAuthCreated{}
}

/*PostLeasesIDAuthCreated handles this case with default header values.

PostLeasesIDAuthCreated post leases Id auth created
*/
type PostLeasesIDAuthCreated struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string

	Payload *PostLeasesIDAuthCreatedBody
}

func (o *PostLeasesIDAuthCreated) Error() string {
	return fmt.Sprintf("[POST /leases/{id}/auth][%d] postLeasesIdAuthCreated  %+v", 201, o.Payload)
}

func (o *PostLeasesIDAuthCreated) GetPayload() *PostLeasesIDAuthCreatedBody {
	return o.Payload
}

func (o *PostLeasesIDAuthCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Allow-Headers
	o.AccessControlAllowHeaders = response.GetHeader("Access-Control-Allow-Headers")

	// response header Access-Control-Allow-Methods
	o.AccessControlAllowMethods = response.GetHeader("Access-Control-Allow-Methods")

	// response header Access-Control-Allow-Origin
	o.AccessControlAllowOrigin = response.GetHeader("Access-Control-Allow-Origin")

	o.Payload = new(PostLeasesIDAuthCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLeasesIDAuthUnauthorized creates a PostLeasesIDAuthUnauthorized with default headers values
func NewPostLeasesIDAuthUnauthorized() *PostLeasesIDAuthUnauthorized {
	return &PostLeasesIDAuthUnauthorized{}
}

/*PostLeasesIDAuthUnauthorized handles this case with default header values.

Unauthorized
*/
type PostLeasesIDAuthUnauthorized struct {
}

func (o *PostLeasesIDAuthUnauthorized) Error() string {
	return fmt.Sprintf("[POST /leases/{id}/auth][%d] postLeasesIdAuthUnauthorized ", 401)
}

func (o *PostLeasesIDAuthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLeasesIDAuthForbidden creates a PostLeasesIDAuthForbidden with default headers values
func NewPostLeasesIDAuthForbidden() *PostLeasesIDAuthForbidden {
	return &PostLeasesIDAuthForbidden{}
}

/*PostLeasesIDAuthForbidden handles this case with default header values.

Failed to retrieve lease authentication
*/
type PostLeasesIDAuthForbidden struct {
}

func (o *PostLeasesIDAuthForbidden) Error() string {
	return fmt.Sprintf("[POST /leases/{id}/auth][%d] postLeasesIdAuthForbidden ", 403)
}

func (o *PostLeasesIDAuthForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLeasesIDAuthInternalServerError creates a PostLeasesIDAuthInternalServerError with default headers values
func NewPostLeasesIDAuthInternalServerError() *PostLeasesIDAuthInternalServerError {
	return &PostLeasesIDAuthInternalServerError{}
}

/*PostLeasesIDAuthInternalServerError handles this case with default header values.

Server failure
*/
type PostLeasesIDAuthInternalServerError struct {
}

func (o *PostLeasesIDAuthInternalServerError) Error() string {
	return fmt.Sprintf("[POST /leases/{id}/auth][%d] postLeasesIdAuthInternalServerError ", 500)
}

func (o *PostLeasesIDAuthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostLeasesIDAuthCreatedBody Lease Authentication
swagger:model PostLeasesIDAuthCreatedBody
*/
type PostLeasesIDAuthCreatedBody struct {

	// Access Key ID for access to the AWS API
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// URL to access the AWS Console
	ConsoleURL string `json:"consoleUrl,omitempty"`

	// Secret Access Key for access to the AWS API
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// Session Token for access to the AWS API
	SessionToken string `json:"sessionToken,omitempty"`
}

// Validate validates this post leases ID auth created body
func (o *PostLeasesIDAuthCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLeasesIDAuthCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLeasesIDAuthCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostLeasesIDAuthCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
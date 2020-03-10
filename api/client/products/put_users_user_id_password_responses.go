// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutUsersUserIDPasswordReader is a Reader for the PutUsersUserIDPassword structure.
type PutUsersUserIDPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUsersUserIDPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUsersUserIDPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUsersUserIDPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUsersUserIDPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUsersUserIDPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUsersUserIDPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutUsersUserIDPasswordOK creates a PutUsersUserIDPasswordOK with default headers values
func NewPutUsersUserIDPasswordOK() *PutUsersUserIDPasswordOK {
	return &PutUsersUserIDPasswordOK{}
}

/*PutUsersUserIDPasswordOK handles this case with default header values.

Updated password successfully.
*/
type PutUsersUserIDPasswordOK struct {
}

func (o *PutUsersUserIDPasswordOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] putUsersUserIdPasswordOK ", 200)
}

func (o *PutUsersUserIDPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDPasswordBadRequest creates a PutUsersUserIDPasswordBadRequest with default headers values
func NewPutUsersUserIDPasswordBadRequest() *PutUsersUserIDPasswordBadRequest {
	return &PutUsersUserIDPasswordBadRequest{}
}

/*PutUsersUserIDPasswordBadRequest handles this case with default header values.

Invalid user ID; Old password is blank; New password is blank.
*/
type PutUsersUserIDPasswordBadRequest struct {
}

func (o *PutUsersUserIDPasswordBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] putUsersUserIdPasswordBadRequest ", 400)
}

func (o *PutUsersUserIDPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDPasswordUnauthorized creates a PutUsersUserIDPasswordUnauthorized with default headers values
func NewPutUsersUserIDPasswordUnauthorized() *PutUsersUserIDPasswordUnauthorized {
	return &PutUsersUserIDPasswordUnauthorized{}
}

/*PutUsersUserIDPasswordUnauthorized handles this case with default header values.

Don't have authority to change password. Please check login status.
*/
type PutUsersUserIDPasswordUnauthorized struct {
}

func (o *PutUsersUserIDPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] putUsersUserIdPasswordUnauthorized ", 401)
}

func (o *PutUsersUserIDPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDPasswordForbidden creates a PutUsersUserIDPasswordForbidden with default headers values
func NewPutUsersUserIDPasswordForbidden() *PutUsersUserIDPasswordForbidden {
	return &PutUsersUserIDPasswordForbidden{}
}

/*PutUsersUserIDPasswordForbidden handles this case with default header values.

The caller does not have permission to update the password of the user with given ID, or the old password in request body is not correct.
*/
type PutUsersUserIDPasswordForbidden struct {
}

func (o *PutUsersUserIDPasswordForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] putUsersUserIdPasswordForbidden ", 403)
}

func (o *PutUsersUserIDPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDPasswordInternalServerError creates a PutUsersUserIDPasswordInternalServerError with default headers values
func NewPutUsersUserIDPasswordInternalServerError() *PutUsersUserIDPasswordInternalServerError {
	return &PutUsersUserIDPasswordInternalServerError{}
}

/*PutUsersUserIDPasswordInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PutUsersUserIDPasswordInternalServerError struct {
}

func (o *PutUsersUserIDPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] putUsersUserIdPasswordInternalServerError ", 500)
}

func (o *PutUsersUserIDPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

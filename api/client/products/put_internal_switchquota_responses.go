// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutInternalSwitchquotaReader is a Reader for the PutInternalSwitchquota structure.
type PutInternalSwitchquotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutInternalSwitchquotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutInternalSwitchquotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutInternalSwitchquotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutInternalSwitchquotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutInternalSwitchquotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutInternalSwitchquotaOK creates a PutInternalSwitchquotaOK with default headers values
func NewPutInternalSwitchquotaOK() *PutInternalSwitchquotaOK {
	return &PutInternalSwitchquotaOK{}
}

/*PutInternalSwitchquotaOK handles this case with default header values.

Enable/Disable quota successfully.
*/
type PutInternalSwitchquotaOK struct {
}

func (o *PutInternalSwitchquotaOK) Error() string {
	return fmt.Sprintf("[PUT /internal/switchquota][%d] putInternalSwitchquotaOK ", 200)
}

func (o *PutInternalSwitchquotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutInternalSwitchquotaUnauthorized creates a PutInternalSwitchquotaUnauthorized with default headers values
func NewPutInternalSwitchquotaUnauthorized() *PutInternalSwitchquotaUnauthorized {
	return &PutInternalSwitchquotaUnauthorized{}
}

/*PutInternalSwitchquotaUnauthorized handles this case with default header values.

User need to log in first.
*/
type PutInternalSwitchquotaUnauthorized struct {
}

func (o *PutInternalSwitchquotaUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /internal/switchquota][%d] putInternalSwitchquotaUnauthorized ", 401)
}

func (o *PutInternalSwitchquotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutInternalSwitchquotaForbidden creates a PutInternalSwitchquotaForbidden with default headers values
func NewPutInternalSwitchquotaForbidden() *PutInternalSwitchquotaForbidden {
	return &PutInternalSwitchquotaForbidden{}
}

/*PutInternalSwitchquotaForbidden handles this case with default header values.

User does not have permission of system admin role.
*/
type PutInternalSwitchquotaForbidden struct {
}

func (o *PutInternalSwitchquotaForbidden) Error() string {
	return fmt.Sprintf("[PUT /internal/switchquota][%d] putInternalSwitchquotaForbidden ", 403)
}

func (o *PutInternalSwitchquotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutInternalSwitchquotaInternalServerError creates a PutInternalSwitchquotaInternalServerError with default headers values
func NewPutInternalSwitchquotaInternalServerError() *PutInternalSwitchquotaInternalServerError {
	return &PutInternalSwitchquotaInternalServerError{}
}

/*PutInternalSwitchquotaInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PutInternalSwitchquotaInternalServerError struct {
}

func (o *PutInternalSwitchquotaInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /internal/switchquota][%d] putInternalSwitchquotaInternalServerError ", 500)
}

func (o *PutInternalSwitchquotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
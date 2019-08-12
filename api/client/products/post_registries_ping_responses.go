// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostRegistriesPingReader is a Reader for the PostRegistriesPing structure.
type PostRegistriesPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRegistriesPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRegistriesPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostRegistriesPingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostRegistriesPingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostRegistriesPingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewPostRegistriesPingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostRegistriesPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRegistriesPingOK creates a PostRegistriesPingOK with default headers values
func NewPostRegistriesPingOK() *PostRegistriesPingOK {
	return &PostRegistriesPingOK{}
}

/*PostRegistriesPingOK handles this case with default header values.

Registry is healthy.
*/
type PostRegistriesPingOK struct {
}

func (o *PostRegistriesPingOK) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] postRegistriesPingOK ", 200)
}

func (o *PostRegistriesPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesPingBadRequest creates a PostRegistriesPingBadRequest with default headers values
func NewPostRegistriesPingBadRequest() *PostRegistriesPingBadRequest {
	return &PostRegistriesPingBadRequest{}
}

/*PostRegistriesPingBadRequest handles this case with default header values.

No proper registry information provided.
*/
type PostRegistriesPingBadRequest struct {
}

func (o *PostRegistriesPingBadRequest) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] postRegistriesPingBadRequest ", 400)
}

func (o *PostRegistriesPingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesPingUnauthorized creates a PostRegistriesPingUnauthorized with default headers values
func NewPostRegistriesPingUnauthorized() *PostRegistriesPingUnauthorized {
	return &PostRegistriesPingUnauthorized{}
}

/*PostRegistriesPingUnauthorized handles this case with default header values.

User need to log in first.
*/
type PostRegistriesPingUnauthorized struct {
}

func (o *PostRegistriesPingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] postRegistriesPingUnauthorized ", 401)
}

func (o *PostRegistriesPingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesPingNotFound creates a PostRegistriesPingNotFound with default headers values
func NewPostRegistriesPingNotFound() *PostRegistriesPingNotFound {
	return &PostRegistriesPingNotFound{}
}

/*PostRegistriesPingNotFound handles this case with default header values.

Registry not found (when registry is provided by ID).
*/
type PostRegistriesPingNotFound struct {
}

func (o *PostRegistriesPingNotFound) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] postRegistriesPingNotFound ", 404)
}

func (o *PostRegistriesPingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesPingUnsupportedMediaType creates a PostRegistriesPingUnsupportedMediaType with default headers values
func NewPostRegistriesPingUnsupportedMediaType() *PostRegistriesPingUnsupportedMediaType {
	return &PostRegistriesPingUnsupportedMediaType{}
}

/*PostRegistriesPingUnsupportedMediaType handles this case with default header values.

The Media Type of the request is not supported, it has to be "application/json"
*/
type PostRegistriesPingUnsupportedMediaType struct {
}

func (o *PostRegistriesPingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] postRegistriesPingUnsupportedMediaType ", 415)
}

func (o *PostRegistriesPingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesPingInternalServerError creates a PostRegistriesPingInternalServerError with default headers values
func NewPostRegistriesPingInternalServerError() *PostRegistriesPingInternalServerError {
	return &PostRegistriesPingInternalServerError{}
}

/*PostRegistriesPingInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PostRegistriesPingInternalServerError struct {
}

func (o *PostRegistriesPingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] postRegistriesPingInternalServerError ", 500)
}

func (o *PostRegistriesPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
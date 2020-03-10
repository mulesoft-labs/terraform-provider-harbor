// Code generated by go-swagger; DO NOT EDIT.

package scanners

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutProjectsProjectIDScannerReader is a Reader for the PutProjectsProjectIDScanner structure.
type PutProjectsProjectIDScannerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectsProjectIDScannerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutProjectsProjectIDScannerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutProjectsProjectIDScannerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutProjectsProjectIDScannerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutProjectsProjectIDScannerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutProjectsProjectIDScannerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutProjectsProjectIDScannerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutProjectsProjectIDScannerOK creates a PutProjectsProjectIDScannerOK with default headers values
func NewPutProjectsProjectIDScannerOK() *PutProjectsProjectIDScannerOK {
	return &PutProjectsProjectIDScannerOK{}
}

/*PutProjectsProjectIDScannerOK handles this case with default header values.

Successfully set the project level scanner
*/
type PutProjectsProjectIDScannerOK struct {
}

func (o *PutProjectsProjectIDScannerOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/scanner][%d] putProjectsProjectIdScannerOK ", 200)
}

func (o *PutProjectsProjectIDScannerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDScannerBadRequest creates a PutProjectsProjectIDScannerBadRequest with default headers values
func NewPutProjectsProjectIDScannerBadRequest() *PutProjectsProjectIDScannerBadRequest {
	return &PutProjectsProjectIDScannerBadRequest{}
}

/*PutProjectsProjectIDScannerBadRequest handles this case with default header values.

Bad project ID
*/
type PutProjectsProjectIDScannerBadRequest struct {
}

func (o *PutProjectsProjectIDScannerBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/scanner][%d] putProjectsProjectIdScannerBadRequest ", 400)
}

func (o *PutProjectsProjectIDScannerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDScannerUnauthorized creates a PutProjectsProjectIDScannerUnauthorized with default headers values
func NewPutProjectsProjectIDScannerUnauthorized() *PutProjectsProjectIDScannerUnauthorized {
	return &PutProjectsProjectIDScannerUnauthorized{}
}

/*PutProjectsProjectIDScannerUnauthorized handles this case with default header values.

Unauthorized request
*/
type PutProjectsProjectIDScannerUnauthorized struct {
}

func (o *PutProjectsProjectIDScannerUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/scanner][%d] putProjectsProjectIdScannerUnauthorized ", 401)
}

func (o *PutProjectsProjectIDScannerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDScannerForbidden creates a PutProjectsProjectIDScannerForbidden with default headers values
func NewPutProjectsProjectIDScannerForbidden() *PutProjectsProjectIDScannerForbidden {
	return &PutProjectsProjectIDScannerForbidden{}
}

/*PutProjectsProjectIDScannerForbidden handles this case with default header values.

Request is not allowed
*/
type PutProjectsProjectIDScannerForbidden struct {
}

func (o *PutProjectsProjectIDScannerForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/scanner][%d] putProjectsProjectIdScannerForbidden ", 403)
}

func (o *PutProjectsProjectIDScannerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDScannerNotFound creates a PutProjectsProjectIDScannerNotFound with default headers values
func NewPutProjectsProjectIDScannerNotFound() *PutProjectsProjectIDScannerNotFound {
	return &PutProjectsProjectIDScannerNotFound{}
}

/*PutProjectsProjectIDScannerNotFound handles this case with default header values.

The requested object is not found
*/
type PutProjectsProjectIDScannerNotFound struct {
}

func (o *PutProjectsProjectIDScannerNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/scanner][%d] putProjectsProjectIdScannerNotFound ", 404)
}

func (o *PutProjectsProjectIDScannerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDScannerInternalServerError creates a PutProjectsProjectIDScannerInternalServerError with default headers values
func NewPutProjectsProjectIDScannerInternalServerError() *PutProjectsProjectIDScannerInternalServerError {
	return &PutProjectsProjectIDScannerInternalServerError{}
}

/*PutProjectsProjectIDScannerInternalServerError handles this case with default header values.

Internal server error happened
*/
type PutProjectsProjectIDScannerInternalServerError struct {
}

func (o *PutProjectsProjectIDScannerInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/scanner][%d] putProjectsProjectIdScannerInternalServerError ", 500)
}

func (o *PutProjectsProjectIDScannerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

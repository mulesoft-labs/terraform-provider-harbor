// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutProjectsProjectIDMetadatasMetaNameReader is a Reader for the PutProjectsProjectIDMetadatasMetaName structure.
type PutProjectsProjectIDMetadatasMetaNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectsProjectIDMetadatasMetaNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutProjectsProjectIDMetadatasMetaNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutProjectsProjectIDMetadatasMetaNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutProjectsProjectIDMetadatasMetaNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutProjectsProjectIDMetadatasMetaNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutProjectsProjectIDMetadatasMetaNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutProjectsProjectIDMetadatasMetaNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutProjectsProjectIDMetadatasMetaNameOK creates a PutProjectsProjectIDMetadatasMetaNameOK with default headers values
func NewPutProjectsProjectIDMetadatasMetaNameOK() *PutProjectsProjectIDMetadatasMetaNameOK {
	return &PutProjectsProjectIDMetadatasMetaNameOK{}
}

/*PutProjectsProjectIDMetadatasMetaNameOK handles this case with default header values.

Updated metadata successfully.
*/
type PutProjectsProjectIDMetadatasMetaNameOK struct {
}

func (o *PutProjectsProjectIDMetadatasMetaNameOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/metadatas/{meta_name}][%d] putProjectsProjectIdMetadatasMetaNameOK ", 200)
}

func (o *PutProjectsProjectIDMetadatasMetaNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDMetadatasMetaNameBadRequest creates a PutProjectsProjectIDMetadatasMetaNameBadRequest with default headers values
func NewPutProjectsProjectIDMetadatasMetaNameBadRequest() *PutProjectsProjectIDMetadatasMetaNameBadRequest {
	return &PutProjectsProjectIDMetadatasMetaNameBadRequest{}
}

/*PutProjectsProjectIDMetadatasMetaNameBadRequest handles this case with default header values.

Invalid request.
*/
type PutProjectsProjectIDMetadatasMetaNameBadRequest struct {
}

func (o *PutProjectsProjectIDMetadatasMetaNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/metadatas/{meta_name}][%d] putProjectsProjectIdMetadatasMetaNameBadRequest ", 400)
}

func (o *PutProjectsProjectIDMetadatasMetaNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDMetadatasMetaNameUnauthorized creates a PutProjectsProjectIDMetadatasMetaNameUnauthorized with default headers values
func NewPutProjectsProjectIDMetadatasMetaNameUnauthorized() *PutProjectsProjectIDMetadatasMetaNameUnauthorized {
	return &PutProjectsProjectIDMetadatasMetaNameUnauthorized{}
}

/*PutProjectsProjectIDMetadatasMetaNameUnauthorized handles this case with default header values.

User need to log in first.
*/
type PutProjectsProjectIDMetadatasMetaNameUnauthorized struct {
}

func (o *PutProjectsProjectIDMetadatasMetaNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/metadatas/{meta_name}][%d] putProjectsProjectIdMetadatasMetaNameUnauthorized ", 401)
}

func (o *PutProjectsProjectIDMetadatasMetaNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDMetadatasMetaNameForbidden creates a PutProjectsProjectIDMetadatasMetaNameForbidden with default headers values
func NewPutProjectsProjectIDMetadatasMetaNameForbidden() *PutProjectsProjectIDMetadatasMetaNameForbidden {
	return &PutProjectsProjectIDMetadatasMetaNameForbidden{}
}

/*PutProjectsProjectIDMetadatasMetaNameForbidden handles this case with default header values.

User does not have permission to the project.
*/
type PutProjectsProjectIDMetadatasMetaNameForbidden struct {
}

func (o *PutProjectsProjectIDMetadatasMetaNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/metadatas/{meta_name}][%d] putProjectsProjectIdMetadatasMetaNameForbidden ", 403)
}

func (o *PutProjectsProjectIDMetadatasMetaNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDMetadatasMetaNameNotFound creates a PutProjectsProjectIDMetadatasMetaNameNotFound with default headers values
func NewPutProjectsProjectIDMetadatasMetaNameNotFound() *PutProjectsProjectIDMetadatasMetaNameNotFound {
	return &PutProjectsProjectIDMetadatasMetaNameNotFound{}
}

/*PutProjectsProjectIDMetadatasMetaNameNotFound handles this case with default header values.

Project or metadata does not exist.
*/
type PutProjectsProjectIDMetadatasMetaNameNotFound struct {
}

func (o *PutProjectsProjectIDMetadatasMetaNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/metadatas/{meta_name}][%d] putProjectsProjectIdMetadatasMetaNameNotFound ", 404)
}

func (o *PutProjectsProjectIDMetadatasMetaNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDMetadatasMetaNameInternalServerError creates a PutProjectsProjectIDMetadatasMetaNameInternalServerError with default headers values
func NewPutProjectsProjectIDMetadatasMetaNameInternalServerError() *PutProjectsProjectIDMetadatasMetaNameInternalServerError {
	return &PutProjectsProjectIDMetadatasMetaNameInternalServerError{}
}

/*PutProjectsProjectIDMetadatasMetaNameInternalServerError handles this case with default header values.

Internal server errors.
*/
type PutProjectsProjectIDMetadatasMetaNameInternalServerError struct {
}

func (o *PutProjectsProjectIDMetadatasMetaNameInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/metadatas/{meta_name}][%d] putProjectsProjectIdMetadatasMetaNameInternalServerError ", 500)
}

func (o *PutProjectsProjectIDMetadatasMetaNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

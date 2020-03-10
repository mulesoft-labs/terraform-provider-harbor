// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sandhose/terraform-provider-harbor/api/models"
)

// GetUsergroupsGroupIDReader is a Reader for the GetUsergroupsGroupID structure.
type GetUsergroupsGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsergroupsGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsergroupsGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsergroupsGroupIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsergroupsGroupIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsergroupsGroupIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsergroupsGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsergroupsGroupIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsergroupsGroupIDOK creates a GetUsergroupsGroupIDOK with default headers values
func NewGetUsergroupsGroupIDOK() *GetUsergroupsGroupIDOK {
	return &GetUsergroupsGroupIDOK{}
}

/*GetUsergroupsGroupIDOK handles this case with default header values.

User group get successfully.
*/
type GetUsergroupsGroupIDOK struct {
	Payload *models.UserGroup
}

func (o *GetUsergroupsGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUsergroupsGroupIdOK  %+v", 200, o.Payload)
}

func (o *GetUsergroupsGroupIDOK) GetPayload() *models.UserGroup {
	return o.Payload
}

func (o *GetUsergroupsGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsergroupsGroupIDBadRequest creates a GetUsergroupsGroupIDBadRequest with default headers values
func NewGetUsergroupsGroupIDBadRequest() *GetUsergroupsGroupIDBadRequest {
	return &GetUsergroupsGroupIDBadRequest{}
}

/*GetUsergroupsGroupIDBadRequest handles this case with default header values.

The user group id is invalid.
*/
type GetUsergroupsGroupIDBadRequest struct {
}

func (o *GetUsergroupsGroupIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUsergroupsGroupIdBadRequest ", 400)
}

func (o *GetUsergroupsGroupIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsergroupsGroupIDUnauthorized creates a GetUsergroupsGroupIDUnauthorized with default headers values
func NewGetUsergroupsGroupIDUnauthorized() *GetUsergroupsGroupIDUnauthorized {
	return &GetUsergroupsGroupIDUnauthorized{}
}

/*GetUsergroupsGroupIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetUsergroupsGroupIDUnauthorized struct {
}

func (o *GetUsergroupsGroupIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUsergroupsGroupIdUnauthorized ", 401)
}

func (o *GetUsergroupsGroupIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsergroupsGroupIDForbidden creates a GetUsergroupsGroupIDForbidden with default headers values
func NewGetUsergroupsGroupIDForbidden() *GetUsergroupsGroupIDForbidden {
	return &GetUsergroupsGroupIDForbidden{}
}

/*GetUsergroupsGroupIDForbidden handles this case with default header values.

User in session does not have permission to the user group.
*/
type GetUsergroupsGroupIDForbidden struct {
}

func (o *GetUsergroupsGroupIDForbidden) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUsergroupsGroupIdForbidden ", 403)
}

func (o *GetUsergroupsGroupIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsergroupsGroupIDNotFound creates a GetUsergroupsGroupIDNotFound with default headers values
func NewGetUsergroupsGroupIDNotFound() *GetUsergroupsGroupIDNotFound {
	return &GetUsergroupsGroupIDNotFound{}
}

/*GetUsergroupsGroupIDNotFound handles this case with default header values.

User group does not exist.
*/
type GetUsergroupsGroupIDNotFound struct {
}

func (o *GetUsergroupsGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUsergroupsGroupIdNotFound ", 404)
}

func (o *GetUsergroupsGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsergroupsGroupIDInternalServerError creates a GetUsergroupsGroupIDInternalServerError with default headers values
func NewGetUsergroupsGroupIDInternalServerError() *GetUsergroupsGroupIDInternalServerError {
	return &GetUsergroupsGroupIDInternalServerError{}
}

/*GetUsergroupsGroupIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetUsergroupsGroupIDInternalServerError struct {
}

func (o *GetUsergroupsGroupIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUsergroupsGroupIdInternalServerError ", 500)
}

func (o *GetUsergroupsGroupIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

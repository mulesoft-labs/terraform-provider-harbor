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

// GetRepositoriesRepoNameLabelsReader is a Reader for the GetRepositoriesRepoNameLabels structure.
type GetRepositoriesRepoNameLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesRepoNameLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesRepoNameLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesRepoNameLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositoriesRepoNameLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesRepoNameLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepositoriesRepoNameLabelsOK creates a GetRepositoriesRepoNameLabelsOK with default headers values
func NewGetRepositoriesRepoNameLabelsOK() *GetRepositoriesRepoNameLabelsOK {
	return &GetRepositoriesRepoNameLabelsOK{}
}

/*GetRepositoriesRepoNameLabelsOK handles this case with default header values.

Successfully.
*/
type GetRepositoriesRepoNameLabelsOK struct {
	Payload []*models.Label
}

func (o *GetRepositoriesRepoNameLabelsOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/labels][%d] getRepositoriesRepoNameLabelsOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesRepoNameLabelsOK) GetPayload() []*models.Label {
	return o.Payload
}

func (o *GetRepositoriesRepoNameLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesRepoNameLabelsUnauthorized creates a GetRepositoriesRepoNameLabelsUnauthorized with default headers values
func NewGetRepositoriesRepoNameLabelsUnauthorized() *GetRepositoriesRepoNameLabelsUnauthorized {
	return &GetRepositoriesRepoNameLabelsUnauthorized{}
}

/*GetRepositoriesRepoNameLabelsUnauthorized handles this case with default header values.

Unauthorized.
*/
type GetRepositoriesRepoNameLabelsUnauthorized struct {
}

func (o *GetRepositoriesRepoNameLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/labels][%d] getRepositoriesRepoNameLabelsUnauthorized ", 401)
}

func (o *GetRepositoriesRepoNameLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesRepoNameLabelsForbidden creates a GetRepositoriesRepoNameLabelsForbidden with default headers values
func NewGetRepositoriesRepoNameLabelsForbidden() *GetRepositoriesRepoNameLabelsForbidden {
	return &GetRepositoriesRepoNameLabelsForbidden{}
}

/*GetRepositoriesRepoNameLabelsForbidden handles this case with default header values.

Forbidden. User should have read permisson for the repository to perform the action.
*/
type GetRepositoriesRepoNameLabelsForbidden struct {
}

func (o *GetRepositoriesRepoNameLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/labels][%d] getRepositoriesRepoNameLabelsForbidden ", 403)
}

func (o *GetRepositoriesRepoNameLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesRepoNameLabelsNotFound creates a GetRepositoriesRepoNameLabelsNotFound with default headers values
func NewGetRepositoriesRepoNameLabelsNotFound() *GetRepositoriesRepoNameLabelsNotFound {
	return &GetRepositoriesRepoNameLabelsNotFound{}
}

/*GetRepositoriesRepoNameLabelsNotFound handles this case with default header values.

Repository not found.
*/
type GetRepositoriesRepoNameLabelsNotFound struct {
}

func (o *GetRepositoriesRepoNameLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/labels][%d] getRepositoriesRepoNameLabelsNotFound ", 404)
}

func (o *GetRepositoriesRepoNameLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

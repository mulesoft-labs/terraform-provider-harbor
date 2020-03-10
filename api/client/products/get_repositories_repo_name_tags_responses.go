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

// GetRepositoriesRepoNameTagsReader is a Reader for the GetRepositoriesRepoNameTags structure.
type GetRepositoriesRepoNameTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesRepoNameTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesRepoNameTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetRepositoriesRepoNameTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepositoriesRepoNameTagsOK creates a GetRepositoriesRepoNameTagsOK with default headers values
func NewGetRepositoriesRepoNameTagsOK() *GetRepositoriesRepoNameTagsOK {
	return &GetRepositoriesRepoNameTagsOK{}
}

/*GetRepositoriesRepoNameTagsOK handles this case with default header values.

Get tags successfully.
*/
type GetRepositoriesRepoNameTagsOK struct {
	Payload []*models.DetailedTag
}

func (o *GetRepositoriesRepoNameTagsOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags][%d] getRepositoriesRepoNameTagsOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesRepoNameTagsOK) GetPayload() []*models.DetailedTag {
	return o.Payload
}

func (o *GetRepositoriesRepoNameTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesRepoNameTagsInternalServerError creates a GetRepositoriesRepoNameTagsInternalServerError with default headers values
func NewGetRepositoriesRepoNameTagsInternalServerError() *GetRepositoriesRepoNameTagsInternalServerError {
	return &GetRepositoriesRepoNameTagsInternalServerError{}
}

/*GetRepositoriesRepoNameTagsInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetRepositoriesRepoNameTagsInternalServerError struct {
}

func (o *GetRepositoriesRepoNameTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags][%d] getRepositoriesRepoNameTagsInternalServerError ", 500)
}

func (o *GetRepositoriesRepoNameTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

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

// GetSysteminfoVolumesReader is a Reader for the GetSysteminfoVolumes structure.
type GetSysteminfoVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSysteminfoVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSysteminfoVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSysteminfoVolumesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSysteminfoVolumesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSysteminfoVolumesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSysteminfoVolumesOK creates a GetSysteminfoVolumesOK with default headers values
func NewGetSysteminfoVolumesOK() *GetSysteminfoVolumesOK {
	return &GetSysteminfoVolumesOK{}
}

/*GetSysteminfoVolumesOK handles this case with default header values.

Get system volumes successfully.
*/
type GetSysteminfoVolumesOK struct {
	Payload *models.SystemInfo
}

func (o *GetSysteminfoVolumesOK) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesOK  %+v", 200, o.Payload)
}

func (o *GetSysteminfoVolumesOK) GetPayload() *models.SystemInfo {
	return o.Payload
}

func (o *GetSysteminfoVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSysteminfoVolumesUnauthorized creates a GetSysteminfoVolumesUnauthorized with default headers values
func NewGetSysteminfoVolumesUnauthorized() *GetSysteminfoVolumesUnauthorized {
	return &GetSysteminfoVolumesUnauthorized{}
}

/*GetSysteminfoVolumesUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetSysteminfoVolumesUnauthorized struct {
}

func (o *GetSysteminfoVolumesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesUnauthorized ", 401)
}

func (o *GetSysteminfoVolumesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSysteminfoVolumesForbidden creates a GetSysteminfoVolumesForbidden with default headers values
func NewGetSysteminfoVolumesForbidden() *GetSysteminfoVolumesForbidden {
	return &GetSysteminfoVolumesForbidden{}
}

/*GetSysteminfoVolumesForbidden handles this case with default header values.

User does not have permission of admin role.
*/
type GetSysteminfoVolumesForbidden struct {
}

func (o *GetSysteminfoVolumesForbidden) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesForbidden ", 403)
}

func (o *GetSysteminfoVolumesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSysteminfoVolumesInternalServerError creates a GetSysteminfoVolumesInternalServerError with default headers values
func NewGetSysteminfoVolumesInternalServerError() *GetSysteminfoVolumesInternalServerError {
	return &GetSysteminfoVolumesInternalServerError{}
}

/*GetSysteminfoVolumesInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetSysteminfoVolumesInternalServerError struct {
}

func (o *GetSysteminfoVolumesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesInternalServerError ", 500)
}

func (o *GetSysteminfoVolumesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

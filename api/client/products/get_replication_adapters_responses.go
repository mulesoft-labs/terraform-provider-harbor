// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetReplicationAdaptersReader is a Reader for the GetReplicationAdapters structure.
type GetReplicationAdaptersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationAdaptersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationAdaptersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReplicationAdaptersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationAdaptersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationAdaptersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReplicationAdaptersOK creates a GetReplicationAdaptersOK with default headers values
func NewGetReplicationAdaptersOK() *GetReplicationAdaptersOK {
	return &GetReplicationAdaptersOK{}
}

/*GetReplicationAdaptersOK handles this case with default header values.

Success.
*/
type GetReplicationAdaptersOK struct {
	Payload []string
}

func (o *GetReplicationAdaptersOK) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] getReplicationAdaptersOK  %+v", 200, o.Payload)
}

func (o *GetReplicationAdaptersOK) GetPayload() []string {
	return o.Payload
}

func (o *GetReplicationAdaptersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationAdaptersUnauthorized creates a GetReplicationAdaptersUnauthorized with default headers values
func NewGetReplicationAdaptersUnauthorized() *GetReplicationAdaptersUnauthorized {
	return &GetReplicationAdaptersUnauthorized{}
}

/*GetReplicationAdaptersUnauthorized handles this case with default header values.

Unauthorized.
*/
type GetReplicationAdaptersUnauthorized struct {
}

func (o *GetReplicationAdaptersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] getReplicationAdaptersUnauthorized ", 401)
}

func (o *GetReplicationAdaptersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationAdaptersForbidden creates a GetReplicationAdaptersForbidden with default headers values
func NewGetReplicationAdaptersForbidden() *GetReplicationAdaptersForbidden {
	return &GetReplicationAdaptersForbidden{}
}

/*GetReplicationAdaptersForbidden handles this case with default header values.

Forbidden.
*/
type GetReplicationAdaptersForbidden struct {
}

func (o *GetReplicationAdaptersForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] getReplicationAdaptersForbidden ", 403)
}

func (o *GetReplicationAdaptersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationAdaptersInternalServerError creates a GetReplicationAdaptersInternalServerError with default headers values
func NewGetReplicationAdaptersInternalServerError() *GetReplicationAdaptersInternalServerError {
	return &GetReplicationAdaptersInternalServerError{}
}

/*GetReplicationAdaptersInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetReplicationAdaptersInternalServerError struct {
}

func (o *GetReplicationAdaptersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] getReplicationAdaptersInternalServerError ", 500)
}

func (o *GetReplicationAdaptersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

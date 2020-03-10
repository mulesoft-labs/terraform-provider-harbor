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

// GetScansAllMetricsReader is a Reader for the GetScansAllMetrics structure.
type GetScansAllMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScansAllMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScansAllMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScansAllMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScansAllMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScansAllMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScansAllMetricsOK creates a GetScansAllMetricsOK with default headers values
func NewGetScansAllMetricsOK() *GetScansAllMetricsOK {
	return &GetScansAllMetricsOK{}
}

/*GetScansAllMetricsOK handles this case with default header values.

OK
*/
type GetScansAllMetricsOK struct {
	Payload *models.Stats
}

func (o *GetScansAllMetricsOK) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getScansAllMetricsOK  %+v", 200, o.Payload)
}

func (o *GetScansAllMetricsOK) GetPayload() *models.Stats {
	return o.Payload
}

func (o *GetScansAllMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Stats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScansAllMetricsUnauthorized creates a GetScansAllMetricsUnauthorized with default headers values
func NewGetScansAllMetricsUnauthorized() *GetScansAllMetricsUnauthorized {
	return &GetScansAllMetricsUnauthorized{}
}

/*GetScansAllMetricsUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetScansAllMetricsUnauthorized struct {
}

func (o *GetScansAllMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getScansAllMetricsUnauthorized ", 401)
}

func (o *GetScansAllMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScansAllMetricsForbidden creates a GetScansAllMetricsForbidden with default headers values
func NewGetScansAllMetricsForbidden() *GetScansAllMetricsForbidden {
	return &GetScansAllMetricsForbidden{}
}

/*GetScansAllMetricsForbidden handles this case with default header values.

Request is not allowed
*/
type GetScansAllMetricsForbidden struct {
}

func (o *GetScansAllMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getScansAllMetricsForbidden ", 403)
}

func (o *GetScansAllMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScansAllMetricsInternalServerError creates a GetScansAllMetricsInternalServerError with default headers values
func NewGetScansAllMetricsInternalServerError() *GetScansAllMetricsInternalServerError {
	return &GetScansAllMetricsInternalServerError{}
}

/*GetScansAllMetricsInternalServerError handles this case with default header values.

Internal server error happened
*/
type GetScansAllMetricsInternalServerError struct {
}

func (o *GetScansAllMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getScansAllMetricsInternalServerError ", 500)
}

func (o *GetScansAllMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// MachineOffReader is a Reader for the MachineOff structure.
type MachineOffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachineOffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMachineOffOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewMachineOffDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMachineOffOK creates a MachineOffOK with default headers values
func NewMachineOffOK() *MachineOffOK {
	return &MachineOffOK{}
}

/*MachineOffOK handles this case with default header values.

OK
*/
type MachineOffOK struct {
	Payload *models.V1MachineResponse
}

func (o *MachineOffOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/off][%d] machineOffOK  %+v", 200, o.Payload)
}

func (o *MachineOffOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMachineOffDefault creates a MachineOffDefault with default headers values
func NewMachineOffDefault(code int) *MachineOffDefault {
	return &MachineOffDefault{
		_statusCode: code,
	}
}

/*MachineOffDefault handles this case with default header values.

Error
*/
type MachineOffDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the machine off default response
func (o *MachineOffDefault) Code() int {
	return o._statusCode
}

func (o *MachineOffDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/off][%d] machineOff default  %+v", o._statusCode, o.Payload)
}

func (o *MachineOffDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

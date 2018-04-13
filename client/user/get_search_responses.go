// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jaingaurav001/swagger_sample/models"
)

// GetSearchReader is a Reader for the GetSearch structure.
type GetSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSearchOK creates a GetSearchOK with default headers values
func NewGetSearchOK() *GetSearchOK {
	return &GetSearchOK{}
}

/*GetSearchOK handles this case with default header values.

OK
*/
type GetSearchOK struct {
	Payload *models.User
}

func (o *GetSearchOK) Error() string {
	return fmt.Sprintf("[GET /search][%d] getSearchOK  %+v", 200, o.Payload)
}

func (o *GetSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchDefault creates a GetSearchDefault with default headers values
func NewGetSearchDefault(code int) *GetSearchDefault {
	return &GetSearchDefault{
		_statusCode: code,
	}
}

/*GetSearchDefault handles this case with default header values.

error
*/
type GetSearchDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get search default response
func (o *GetSearchDefault) Code() int {
	return o._statusCode
}

func (o *GetSearchDefault) Error() string {
	return fmt.Sprintf("[GET /search][%d] GetSearch default  %+v", o._statusCode, o.Payload)
}

func (o *GetSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteTeacherReader is a Reader for the DeleteTeacher structure.
type DeleteTeacherReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeacherReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTeacherNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTeacherNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTeacherNoContent creates a DeleteTeacherNoContent with default headers values
func NewDeleteTeacherNoContent() *DeleteTeacherNoContent {
	return &DeleteTeacherNoContent{}
}

/*DeleteTeacherNoContent handles this case with default header values.

teacher deleted
*/
type DeleteTeacherNoContent struct {
}

func (o *DeleteTeacherNoContent) Error() string {
	return fmt.Sprintf("[DELETE /teacher/{ID}][%d] deleteTeacherNoContent ", 204)
}

func (o *DeleteTeacherNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTeacherNotFound creates a DeleteTeacherNotFound with default headers values
func NewDeleteTeacherNotFound() *DeleteTeacherNotFound {
	return &DeleteTeacherNotFound{}
}

/*DeleteTeacherNotFound handles this case with default header values.

teacher not found
*/
type DeleteTeacherNotFound struct {
}

func (o *DeleteTeacherNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teacher/{ID}][%d] deleteTeacherNotFound ", 404)
}

func (o *DeleteTeacherNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

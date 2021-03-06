// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/divoc/portal-api/swagger_gen/models"
)

// GetVaccinatorsOKCode is the HTTP code returned for type GetVaccinatorsOK
const GetVaccinatorsOKCode int = 200

/*GetVaccinatorsOK OK

swagger:response getVaccinatorsOK
*/
type GetVaccinatorsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Vaccinator `json:"body,omitempty"`
}

// NewGetVaccinatorsOK creates GetVaccinatorsOK with default headers values
func NewGetVaccinatorsOK() *GetVaccinatorsOK {

	return &GetVaccinatorsOK{}
}

// WithPayload adds the payload to the get vaccinators o k response
func (o *GetVaccinatorsOK) WithPayload(payload []*models.Vaccinator) *GetVaccinatorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vaccinators o k response
func (o *GetVaccinatorsOK) SetPayload(payload []*models.Vaccinator) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVaccinatorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Vaccinator, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

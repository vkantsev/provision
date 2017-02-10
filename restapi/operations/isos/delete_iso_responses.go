package isos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*DeleteIsoNoContent delete iso no content

swagger:response deleteIsoNoContent
*/
type DeleteIsoNoContent struct {
}

// NewDeleteIsoNoContent creates DeleteIsoNoContent with default headers values
func NewDeleteIsoNoContent() *DeleteIsoNoContent {
	return &DeleteIsoNoContent{}
}

// WriteResponse to the client
func (o *DeleteIsoNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteIsoUnauthorized delete iso unauthorized

swagger:response deleteIsoUnauthorized
*/
type DeleteIsoUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteIsoUnauthorized creates DeleteIsoUnauthorized with default headers values
func NewDeleteIsoUnauthorized() *DeleteIsoUnauthorized {
	return &DeleteIsoUnauthorized{}
}

// WithPayload adds the payload to the delete iso unauthorized response
func (o *DeleteIsoUnauthorized) WithPayload(payload *models.Error) *DeleteIsoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete iso unauthorized response
func (o *DeleteIsoUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIsoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteIsoNotFound delete iso not found

swagger:response deleteIsoNotFound
*/
type DeleteIsoNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteIsoNotFound creates DeleteIsoNotFound with default headers values
func NewDeleteIsoNotFound() *DeleteIsoNotFound {
	return &DeleteIsoNotFound{}
}

// WithPayload adds the payload to the delete iso not found response
func (o *DeleteIsoNotFound) WithPayload(payload *models.Error) *DeleteIsoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete iso not found response
func (o *DeleteIsoNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIsoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteIsoConflict delete iso conflict

swagger:response deleteIsoConflict
*/
type DeleteIsoConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteIsoConflict creates DeleteIsoConflict with default headers values
func NewDeleteIsoConflict() *DeleteIsoConflict {
	return &DeleteIsoConflict{}
}

// WithPayload adds the payload to the delete iso conflict response
func (o *DeleteIsoConflict) WithPayload(payload *models.Error) *DeleteIsoConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete iso conflict response
func (o *DeleteIsoConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIsoConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteIsoInternalServerError delete iso internal server error

swagger:response deleteIsoInternalServerError
*/
type DeleteIsoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteIsoInternalServerError creates DeleteIsoInternalServerError with default headers values
func NewDeleteIsoInternalServerError() *DeleteIsoInternalServerError {
	return &DeleteIsoInternalServerError{}
}

// WithPayload adds the payload to the delete iso internal server error response
func (o *DeleteIsoInternalServerError) WithPayload(payload *models.Error) *DeleteIsoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete iso internal server error response
func (o *DeleteIsoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIsoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

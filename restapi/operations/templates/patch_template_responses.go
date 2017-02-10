package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*PatchTemplateOK patch template o k

swagger:response patchTemplateOK
*/
type PatchTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.TemplateOutput `json:"body,omitempty"`
}

// NewPatchTemplateOK creates PatchTemplateOK with default headers values
func NewPatchTemplateOK() *PatchTemplateOK {
	return &PatchTemplateOK{}
}

// WithPayload adds the payload to the patch template o k response
func (o *PatchTemplateOK) WithPayload(payload *models.TemplateOutput) *PatchTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch template o k response
func (o *PatchTemplateOK) SetPayload(payload *models.TemplateOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchTemplateUnauthorized patch template unauthorized

swagger:response patchTemplateUnauthorized
*/
type PatchTemplateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchTemplateUnauthorized creates PatchTemplateUnauthorized with default headers values
func NewPatchTemplateUnauthorized() *PatchTemplateUnauthorized {
	return &PatchTemplateUnauthorized{}
}

// WithPayload adds the payload to the patch template unauthorized response
func (o *PatchTemplateUnauthorized) WithPayload(payload *models.Error) *PatchTemplateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch template unauthorized response
func (o *PatchTemplateUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTemplateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchTemplateNotFound patch template not found

swagger:response patchTemplateNotFound
*/
type PatchTemplateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchTemplateNotFound creates PatchTemplateNotFound with default headers values
func NewPatchTemplateNotFound() *PatchTemplateNotFound {
	return &PatchTemplateNotFound{}
}

// WithPayload adds the payload to the patch template not found response
func (o *PatchTemplateNotFound) WithPayload(payload *models.Error) *PatchTemplateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch template not found response
func (o *PatchTemplateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTemplateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchTemplateConflict patch template conflict

swagger:response patchTemplateConflict
*/
type PatchTemplateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchTemplateConflict creates PatchTemplateConflict with default headers values
func NewPatchTemplateConflict() *PatchTemplateConflict {
	return &PatchTemplateConflict{}
}

// WithPayload adds the payload to the patch template conflict response
func (o *PatchTemplateConflict) WithPayload(payload *models.Error) *PatchTemplateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch template conflict response
func (o *PatchTemplateConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTemplateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchTemplateExpectationFailed patch template expectation failed

swagger:response patchTemplateExpectationFailed
*/
type PatchTemplateExpectationFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchTemplateExpectationFailed creates PatchTemplateExpectationFailed with default headers values
func NewPatchTemplateExpectationFailed() *PatchTemplateExpectationFailed {
	return &PatchTemplateExpectationFailed{}
}

// WithPayload adds the payload to the patch template expectation failed response
func (o *PatchTemplateExpectationFailed) WithPayload(payload *models.Error) *PatchTemplateExpectationFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch template expectation failed response
func (o *PatchTemplateExpectationFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTemplateExpectationFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(417)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchTemplateInternalServerError patch template internal server error

swagger:response patchTemplateInternalServerError
*/
type PatchTemplateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchTemplateInternalServerError creates PatchTemplateInternalServerError with default headers values
func NewPatchTemplateInternalServerError() *PatchTemplateInternalServerError {
	return &PatchTemplateInternalServerError{}
}

// WithPayload adds the payload to the patch template internal server error response
func (o *PatchTemplateInternalServerError) WithPayload(payload *models.Error) *PatchTemplateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch template internal server error response
func (o *PatchTemplateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTemplateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

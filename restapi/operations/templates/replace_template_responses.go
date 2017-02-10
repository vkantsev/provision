package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*ReplaceTemplateOK replace template o k

swagger:response replaceTemplateOK
*/
type ReplaceTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.TemplateOutput `json:"body,omitempty"`
}

// NewReplaceTemplateOK creates ReplaceTemplateOK with default headers values
func NewReplaceTemplateOK() *ReplaceTemplateOK {
	return &ReplaceTemplateOK{}
}

// WithPayload adds the payload to the replace template o k response
func (o *ReplaceTemplateOK) WithPayload(payload *models.TemplateOutput) *ReplaceTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace template o k response
func (o *ReplaceTemplateOK) SetPayload(payload *models.TemplateOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceTemplateCreated replace template created

swagger:response replaceTemplateCreated
*/
type ReplaceTemplateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.TemplateOutput `json:"body,omitempty"`
}

// NewReplaceTemplateCreated creates ReplaceTemplateCreated with default headers values
func NewReplaceTemplateCreated() *ReplaceTemplateCreated {
	return &ReplaceTemplateCreated{}
}

// WithPayload adds the payload to the replace template created response
func (o *ReplaceTemplateCreated) WithPayload(payload *models.TemplateOutput) *ReplaceTemplateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace template created response
func (o *ReplaceTemplateCreated) SetPayload(payload *models.TemplateOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTemplateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceTemplateUnauthorized replace template unauthorized

swagger:response replaceTemplateUnauthorized
*/
type ReplaceTemplateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTemplateUnauthorized creates ReplaceTemplateUnauthorized with default headers values
func NewReplaceTemplateUnauthorized() *ReplaceTemplateUnauthorized {
	return &ReplaceTemplateUnauthorized{}
}

// WithPayload adds the payload to the replace template unauthorized response
func (o *ReplaceTemplateUnauthorized) WithPayload(payload *models.Error) *ReplaceTemplateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace template unauthorized response
func (o *ReplaceTemplateUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTemplateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceTemplateConflict replace template conflict

swagger:response replaceTemplateConflict
*/
type ReplaceTemplateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTemplateConflict creates ReplaceTemplateConflict with default headers values
func NewReplaceTemplateConflict() *ReplaceTemplateConflict {
	return &ReplaceTemplateConflict{}
}

// WithPayload adds the payload to the replace template conflict response
func (o *ReplaceTemplateConflict) WithPayload(payload *models.Error) *ReplaceTemplateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace template conflict response
func (o *ReplaceTemplateConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTemplateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceTemplateExpectationFailed replace template expectation failed

swagger:response replaceTemplateExpectationFailed
*/
type ReplaceTemplateExpectationFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTemplateExpectationFailed creates ReplaceTemplateExpectationFailed with default headers values
func NewReplaceTemplateExpectationFailed() *ReplaceTemplateExpectationFailed {
	return &ReplaceTemplateExpectationFailed{}
}

// WithPayload adds the payload to the replace template expectation failed response
func (o *ReplaceTemplateExpectationFailed) WithPayload(payload *models.Error) *ReplaceTemplateExpectationFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace template expectation failed response
func (o *ReplaceTemplateExpectationFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTemplateExpectationFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(417)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceTemplateInternalServerError replace template internal server error

swagger:response replaceTemplateInternalServerError
*/
type ReplaceTemplateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTemplateInternalServerError creates ReplaceTemplateInternalServerError with default headers values
func NewReplaceTemplateInternalServerError() *ReplaceTemplateInternalServerError {
	return &ReplaceTemplateInternalServerError{}
}

// WithPayload adds the payload to the replace template internal server error response
func (o *ReplaceTemplateInternalServerError) WithPayload(payload *models.Error) *ReplaceTemplateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace template internal server error response
func (o *ReplaceTemplateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTemplateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

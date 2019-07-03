// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetBindsOKCode is the HTTP code returned for type GetBindsOK
const GetBindsOKCode int = 200

/*GetBindsOK Successful operation

swagger:response getBindsOK
*/
type GetBindsOK struct {

	/*
	  In: Body
	*/
	Payload *GetBindsOKBody `json:"body,omitempty"`
}

// NewGetBindsOK creates GetBindsOK with default headers values
func NewGetBindsOK() *GetBindsOK {

	return &GetBindsOK{}
}

// WithPayload adds the payload to the get binds o k response
func (o *GetBindsOK) WithPayload(payload *GetBindsOKBody) *GetBindsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get binds o k response
func (o *GetBindsOK) SetPayload(payload *GetBindsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBindsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBindsDefault General Error

swagger:response getBindsDefault
*/
type GetBindsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBindsDefault creates GetBindsDefault with default headers values
func NewGetBindsDefault(code int) *GetBindsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBindsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get binds default response
func (o *GetBindsDefault) WithStatusCode(code int) *GetBindsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get binds default response
func (o *GetBindsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get binds default response
func (o *GetBindsDefault) WithPayload(payload *models.Error) *GetBindsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get binds default response
func (o *GetBindsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBindsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

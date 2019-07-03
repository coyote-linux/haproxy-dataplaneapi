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

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetLogTargetsOKCode is the HTTP code returned for type GetLogTargetsOK
const GetLogTargetsOKCode int = 200

/*GetLogTargetsOK Successful operation

swagger:response getLogTargetsOK
*/
type GetLogTargetsOK struct {

	/*
	  In: Body
	*/
	Payload *GetLogTargetsOKBody `json:"body,omitempty"`
}

// NewGetLogTargetsOK creates GetLogTargetsOK with default headers values
func NewGetLogTargetsOK() *GetLogTargetsOK {

	return &GetLogTargetsOK{}
}

// WithPayload adds the payload to the get log targets o k response
func (o *GetLogTargetsOK) WithPayload(payload *GetLogTargetsOKBody) *GetLogTargetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log targets o k response
func (o *GetLogTargetsOK) SetPayload(payload *GetLogTargetsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogTargetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetLogTargetsDefault General Error

swagger:response getLogTargetsDefault
*/
type GetLogTargetsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLogTargetsDefault creates GetLogTargetsDefault with default headers values
func NewGetLogTargetsDefault(code int) *GetLogTargetsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLogTargetsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get log targets default response
func (o *GetLogTargetsDefault) WithStatusCode(code int) *GetLogTargetsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get log targets default response
func (o *GetLogTargetsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get log targets default response
func (o *GetLogTargetsDefault) WithPayload(payload *models.Error) *GetLogTargetsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log targets default response
func (o *GetLogTargetsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogTargetsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

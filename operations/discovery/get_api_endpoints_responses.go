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

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetAPIEndpointsOKCode is the HTTP code returned for type GetAPIEndpointsOK
const GetAPIEndpointsOKCode int = 200

/*GetAPIEndpointsOK Success

swagger:response getApiEndpointsOK
*/
type GetAPIEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload models.Endpoints `json:"body,omitempty"`
}

// NewGetAPIEndpointsOK creates GetAPIEndpointsOK with default headers values
func NewGetAPIEndpointsOK() *GetAPIEndpointsOK {

	return &GetAPIEndpointsOK{}
}

// WithPayload adds the payload to the get Api endpoints o k response
func (o *GetAPIEndpointsOK) WithPayload(payload models.Endpoints) *GetAPIEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api endpoints o k response
func (o *GetAPIEndpointsOK) SetPayload(payload models.Endpoints) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Endpoints{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetAPIEndpointsDefault General Error

swagger:response getApiEndpointsDefault
*/
type GetAPIEndpointsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAPIEndpointsDefault creates GetAPIEndpointsDefault with default headers values
func NewGetAPIEndpointsDefault(code int) *GetAPIEndpointsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAPIEndpointsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get API endpoints default response
func (o *GetAPIEndpointsDefault) WithStatusCode(code int) *GetAPIEndpointsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get API endpoints default response
func (o *GetAPIEndpointsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get API endpoints default response
func (o *GetAPIEndpointsDefault) WithPayload(payload *models.Error) *GetAPIEndpointsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get API endpoints default response
func (o *GetAPIEndpointsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIEndpointsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

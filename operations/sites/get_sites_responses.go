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

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetSitesOKCode is the HTTP code returned for type GetSitesOK
const GetSitesOKCode int = 200

/*GetSitesOK Successful operation

swagger:response getSitesOK
*/
type GetSitesOK struct {

	/*
	  In: Body
	*/
	Payload *GetSitesOKBody `json:"body,omitempty"`
}

// NewGetSitesOK creates GetSitesOK with default headers values
func NewGetSitesOK() *GetSitesOK {

	return &GetSitesOK{}
}

// WithPayload adds the payload to the get sites o k response
func (o *GetSitesOK) WithPayload(payload *GetSitesOKBody) *GetSitesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sites o k response
func (o *GetSitesOK) SetPayload(payload *GetSitesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSitesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSitesDefault General Error

swagger:response getSitesDefault
*/
type GetSitesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSitesDefault creates GetSitesDefault with default headers values
func NewGetSitesDefault(code int) *GetSitesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSitesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sites default response
func (o *GetSitesDefault) WithStatusCode(code int) *GetSitesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sites default response
func (o *GetSitesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get sites default response
func (o *GetSitesDefault) WithPayload(payload *models.Error) *GetSitesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sites default response
func (o *GetSitesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSitesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetAclsOKCode is the HTTP code returned for type GetAclsOK
const GetAclsOKCode int = 200

/*GetAclsOK Successful operation

swagger:response getAclsOK
*/
type GetAclsOK struct {

	/*
	  In: Body
	*/
	Payload *GetAclsOKBody `json:"body,omitempty"`
}

// NewGetAclsOK creates GetAclsOK with default headers values
func NewGetAclsOK() *GetAclsOK {

	return &GetAclsOK{}
}

// WithPayload adds the payload to the get acls o k response
func (o *GetAclsOK) WithPayload(payload *GetAclsOKBody) *GetAclsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get acls o k response
func (o *GetAclsOK) SetPayload(payload *GetAclsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAclsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAclsDefault General Error

swagger:response getAclsDefault
*/
type GetAclsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAclsDefault creates GetAclsDefault with default headers values
func NewGetAclsDefault(code int) *GetAclsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAclsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get acls default response
func (o *GetAclsDefault) WithStatusCode(code int) *GetAclsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get acls default response
func (o *GetAclsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get acls default response
func (o *GetAclsDefault) WithPayload(payload *models.Error) *GetAclsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get acls default response
func (o *GetAclsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAclsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

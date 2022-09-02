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

	"github.com/haproxytech/client-native/v4/models"
)

// GetHaproxyEndpointsOKCode is the HTTP code returned for type GetHaproxyEndpointsOK
const GetHaproxyEndpointsOKCode int = 200

/*
GetHaproxyEndpointsOK Success

swagger:response getHaproxyEndpointsOK
*/
type GetHaproxyEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload models.Endpoints `json:"body,omitempty"`
}

// NewGetHaproxyEndpointsOK creates GetHaproxyEndpointsOK with default headers values
func NewGetHaproxyEndpointsOK() *GetHaproxyEndpointsOK {

	return &GetHaproxyEndpointsOK{}
}

// WithPayload adds the payload to the get haproxy endpoints o k response
func (o *GetHaproxyEndpointsOK) WithPayload(payload models.Endpoints) *GetHaproxyEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get haproxy endpoints o k response
func (o *GetHaproxyEndpointsOK) SetPayload(payload models.Endpoints) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHaproxyEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*
GetHaproxyEndpointsDefault General Error

swagger:response getHaproxyEndpointsDefault
*/
type GetHaproxyEndpointsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHaproxyEndpointsDefault creates GetHaproxyEndpointsDefault with default headers values
func NewGetHaproxyEndpointsDefault(code int) *GetHaproxyEndpointsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHaproxyEndpointsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get haproxy endpoints default response
func (o *GetHaproxyEndpointsDefault) WithStatusCode(code int) *GetHaproxyEndpointsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get haproxy endpoints default response
func (o *GetHaproxyEndpointsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get haproxy endpoints default response
func (o *GetHaproxyEndpointsDefault) WithConfigurationVersion(configurationVersion string) *GetHaproxyEndpointsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get haproxy endpoints default response
func (o *GetHaproxyEndpointsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get haproxy endpoints default response
func (o *GetHaproxyEndpointsDefault) WithPayload(payload *models.Error) *GetHaproxyEndpointsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get haproxy endpoints default response
func (o *GetHaproxyEndpointsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHaproxyEndpointsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

package dgram_bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetDgramBindsOKCode is the HTTP code returned for type GetDgramBindsOK
const GetDgramBindsOKCode int = 200

/*
GetDgramBindsOK Successful operation

swagger:response getDgramBindsOK
*/
type GetDgramBindsOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetDgramBindsOKBody `json:"body,omitempty"`
}

// NewGetDgramBindsOK creates GetDgramBindsOK with default headers values
func NewGetDgramBindsOK() *GetDgramBindsOK {

	return &GetDgramBindsOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get dgram binds o k response
func (o *GetDgramBindsOK) WithConfigurationVersion(configurationVersion string) *GetDgramBindsOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get dgram binds o k response
func (o *GetDgramBindsOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get dgram binds o k response
func (o *GetDgramBindsOK) WithPayload(payload *GetDgramBindsOKBody) *GetDgramBindsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dgram binds o k response
func (o *GetDgramBindsOK) SetPayload(payload *GetDgramBindsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDgramBindsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetDgramBindsDefault General Error

swagger:response getDgramBindsDefault
*/
type GetDgramBindsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDgramBindsDefault creates GetDgramBindsDefault with default headers values
func NewGetDgramBindsDefault(code int) *GetDgramBindsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDgramBindsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get dgram binds default response
func (o *GetDgramBindsDefault) WithStatusCode(code int) *GetDgramBindsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get dgram binds default response
func (o *GetDgramBindsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get dgram binds default response
func (o *GetDgramBindsDefault) WithConfigurationVersion(configurationVersion string) *GetDgramBindsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get dgram binds default response
func (o *GetDgramBindsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get dgram binds default response
func (o *GetDgramBindsDefault) WithPayload(payload *models.Error) *GetDgramBindsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dgram binds default response
func (o *GetDgramBindsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDgramBindsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

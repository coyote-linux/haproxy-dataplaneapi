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

package userlist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetUserlistsOKCode is the HTTP code returned for type GetUserlistsOK
const GetUserlistsOKCode int = 200

/*
GetUserlistsOK Successful operation

swagger:response getUserlistsOK
*/
type GetUserlistsOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetUserlistsOKBody `json:"body,omitempty"`
}

// NewGetUserlistsOK creates GetUserlistsOK with default headers values
func NewGetUserlistsOK() *GetUserlistsOK {

	return &GetUserlistsOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get userlists o k response
func (o *GetUserlistsOK) WithConfigurationVersion(configurationVersion string) *GetUserlistsOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get userlists o k response
func (o *GetUserlistsOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get userlists o k response
func (o *GetUserlistsOK) WithPayload(payload *GetUserlistsOKBody) *GetUserlistsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get userlists o k response
func (o *GetUserlistsOK) SetPayload(payload *GetUserlistsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserlistsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetUserlistsDefault General Error

swagger:response getUserlistsDefault
*/
type GetUserlistsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserlistsDefault creates GetUserlistsDefault with default headers values
func NewGetUserlistsDefault(code int) *GetUserlistsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserlistsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get userlists default response
func (o *GetUserlistsDefault) WithStatusCode(code int) *GetUserlistsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get userlists default response
func (o *GetUserlistsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get userlists default response
func (o *GetUserlistsDefault) WithConfigurationVersion(configurationVersion string) *GetUserlistsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get userlists default response
func (o *GetUserlistsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get userlists default response
func (o *GetUserlistsDefault) WithPayload(payload *models.Error) *GetUserlistsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get userlists default response
func (o *GetUserlistsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserlistsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

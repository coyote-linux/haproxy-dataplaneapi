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

	"github.com/haproxytech/client-native/v4/models"
)

// GetLogTargetOKCode is the HTTP code returned for type GetLogTargetOK
const GetLogTargetOKCode int = 200

/*
GetLogTargetOK Successful operation

swagger:response getLogTargetOK
*/
type GetLogTargetOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetLogTargetOKBody `json:"body,omitempty"`
}

// NewGetLogTargetOK creates GetLogTargetOK with default headers values
func NewGetLogTargetOK() *GetLogTargetOK {

	return &GetLogTargetOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get log target o k response
func (o *GetLogTargetOK) WithConfigurationVersion(configurationVersion string) *GetLogTargetOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get log target o k response
func (o *GetLogTargetOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get log target o k response
func (o *GetLogTargetOK) WithPayload(payload *GetLogTargetOKBody) *GetLogTargetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log target o k response
func (o *GetLogTargetOK) SetPayload(payload *GetLogTargetOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogTargetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetLogTargetNotFoundCode is the HTTP code returned for type GetLogTargetNotFound
const GetLogTargetNotFoundCode int = 404

/*
GetLogTargetNotFound The specified resource was not found

swagger:response getLogTargetNotFound
*/
type GetLogTargetNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLogTargetNotFound creates GetLogTargetNotFound with default headers values
func NewGetLogTargetNotFound() *GetLogTargetNotFound {

	return &GetLogTargetNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get log target not found response
func (o *GetLogTargetNotFound) WithConfigurationVersion(configurationVersion string) *GetLogTargetNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get log target not found response
func (o *GetLogTargetNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get log target not found response
func (o *GetLogTargetNotFound) WithPayload(payload *models.Error) *GetLogTargetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log target not found response
func (o *GetLogTargetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogTargetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetLogTargetDefault General Error

swagger:response getLogTargetDefault
*/
type GetLogTargetDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLogTargetDefault creates GetLogTargetDefault with default headers values
func NewGetLogTargetDefault(code int) *GetLogTargetDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLogTargetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get log target default response
func (o *GetLogTargetDefault) WithStatusCode(code int) *GetLogTargetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get log target default response
func (o *GetLogTargetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get log target default response
func (o *GetLogTargetDefault) WithConfigurationVersion(configurationVersion string) *GetLogTargetDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get log target default response
func (o *GetLogTargetDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get log target default response
func (o *GetLogTargetDefault) WithPayload(payload *models.Error) *GetLogTargetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log target default response
func (o *GetLogTargetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogTargetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

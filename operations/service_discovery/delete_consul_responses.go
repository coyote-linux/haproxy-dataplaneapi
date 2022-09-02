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

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteConsulNoContentCode is the HTTP code returned for type DeleteConsulNoContent
const DeleteConsulNoContentCode int = 204

/*
DeleteConsulNoContent Consul server deleted

swagger:response deleteConsulNoContent
*/
type DeleteConsulNoContent struct {
}

// NewDeleteConsulNoContent creates DeleteConsulNoContent with default headers values
func NewDeleteConsulNoContent() *DeleteConsulNoContent {

	return &DeleteConsulNoContent{}
}

// WriteResponse to the client
func (o *DeleteConsulNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConsulNotFoundCode is the HTTP code returned for type DeleteConsulNotFound
const DeleteConsulNotFoundCode int = 404

/*
DeleteConsulNotFound The specified resource was not found

swagger:response deleteConsulNotFound
*/
type DeleteConsulNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConsulNotFound creates DeleteConsulNotFound with default headers values
func NewDeleteConsulNotFound() *DeleteConsulNotFound {

	return &DeleteConsulNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete consul not found response
func (o *DeleteConsulNotFound) WithConfigurationVersion(configurationVersion string) *DeleteConsulNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete consul not found response
func (o *DeleteConsulNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete consul not found response
func (o *DeleteConsulNotFound) WithPayload(payload *models.Error) *DeleteConsulNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete consul not found response
func (o *DeleteConsulNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConsulNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteConsulDefault General Error

swagger:response deleteConsulDefault
*/
type DeleteConsulDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConsulDefault creates DeleteConsulDefault with default headers values
func NewDeleteConsulDefault(code int) *DeleteConsulDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteConsulDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete consul default response
func (o *DeleteConsulDefault) WithStatusCode(code int) *DeleteConsulDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete consul default response
func (o *DeleteConsulDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete consul default response
func (o *DeleteConsulDefault) WithConfigurationVersion(configurationVersion string) *DeleteConsulDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete consul default response
func (o *DeleteConsulDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete consul default response
func (o *DeleteConsulDefault) WithPayload(payload *models.Error) *DeleteConsulDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete consul default response
func (o *DeleteConsulDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConsulDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

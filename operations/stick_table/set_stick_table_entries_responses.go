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

package stick_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// SetStickTableEntriesNoContentCode is the HTTP code returned for type SetStickTableEntriesNoContent
const SetStickTableEntriesNoContentCode int = 204

/*
SetStickTableEntriesNoContent Successful operation

swagger:response setStickTableEntriesNoContent
*/
type SetStickTableEntriesNoContent struct {
}

// NewSetStickTableEntriesNoContent creates SetStickTableEntriesNoContent with default headers values
func NewSetStickTableEntriesNoContent() *SetStickTableEntriesNoContent {

	return &SetStickTableEntriesNoContent{}
}

// WriteResponse to the client
func (o *SetStickTableEntriesNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
SetStickTableEntriesDefault General Error

swagger:response setStickTableEntriesDefault
*/
type SetStickTableEntriesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetStickTableEntriesDefault creates SetStickTableEntriesDefault with default headers values
func NewSetStickTableEntriesDefault(code int) *SetStickTableEntriesDefault {
	if code <= 0 {
		code = 500
	}

	return &SetStickTableEntriesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the set stick table entries default response
func (o *SetStickTableEntriesDefault) WithStatusCode(code int) *SetStickTableEntriesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the set stick table entries default response
func (o *SetStickTableEntriesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the set stick table entries default response
func (o *SetStickTableEntriesDefault) WithConfigurationVersion(configurationVersion string) *SetStickTableEntriesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the set stick table entries default response
func (o *SetStickTableEntriesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the set stick table entries default response
func (o *SetStickTableEntriesDefault) WithPayload(payload *models.Error) *SetStickTableEntriesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set stick table entries default response
func (o *SetStickTableEntriesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetStickTableEntriesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

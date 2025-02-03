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

package log_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// DeleteLogProfileAcceptedCode is the HTTP code returned for type DeleteLogProfileAccepted
const DeleteLogProfileAcceptedCode int = 202

/*
DeleteLogProfileAccepted Configuration change accepted and reload requested

swagger:response deleteLogProfileAccepted
*/
type DeleteLogProfileAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteLogProfileAccepted creates DeleteLogProfileAccepted with default headers values
func NewDeleteLogProfileAccepted() *DeleteLogProfileAccepted {

	return &DeleteLogProfileAccepted{}
}

// WithReloadID adds the reloadId to the delete log profile accepted response
func (o *DeleteLogProfileAccepted) WithReloadID(reloadID string) *DeleteLogProfileAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete log profile accepted response
func (o *DeleteLogProfileAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteLogProfileAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteLogProfileNoContentCode is the HTTP code returned for type DeleteLogProfileNoContent
const DeleteLogProfileNoContentCode int = 204

/*
DeleteLogProfileNoContent log_profile deleted

swagger:response deleteLogProfileNoContent
*/
type DeleteLogProfileNoContent struct {
}

// NewDeleteLogProfileNoContent creates DeleteLogProfileNoContent with default headers values
func NewDeleteLogProfileNoContent() *DeleteLogProfileNoContent {

	return &DeleteLogProfileNoContent{}
}

// WriteResponse to the client
func (o *DeleteLogProfileNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteLogProfileNotFoundCode is the HTTP code returned for type DeleteLogProfileNotFound
const DeleteLogProfileNotFoundCode int = 404

/*
DeleteLogProfileNotFound The specified resource was not found

swagger:response deleteLogProfileNotFound
*/
type DeleteLogProfileNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteLogProfileNotFound creates DeleteLogProfileNotFound with default headers values
func NewDeleteLogProfileNotFound() *DeleteLogProfileNotFound {

	return &DeleteLogProfileNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete log profile not found response
func (o *DeleteLogProfileNotFound) WithConfigurationVersion(configurationVersion string) *DeleteLogProfileNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete log profile not found response
func (o *DeleteLogProfileNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete log profile not found response
func (o *DeleteLogProfileNotFound) WithPayload(payload *models.Error) *DeleteLogProfileNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete log profile not found response
func (o *DeleteLogProfileNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLogProfileNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteLogProfileDefault General Error

swagger:response deleteLogProfileDefault
*/
type DeleteLogProfileDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteLogProfileDefault creates DeleteLogProfileDefault with default headers values
func NewDeleteLogProfileDefault(code int) *DeleteLogProfileDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteLogProfileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete log profile default response
func (o *DeleteLogProfileDefault) WithStatusCode(code int) *DeleteLogProfileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete log profile default response
func (o *DeleteLogProfileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete log profile default response
func (o *DeleteLogProfileDefault) WithConfigurationVersion(configurationVersion string) *DeleteLogProfileDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete log profile default response
func (o *DeleteLogProfileDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete log profile default response
func (o *DeleteLogProfileDefault) WithPayload(payload *models.Error) *DeleteLogProfileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete log profile default response
func (o *DeleteLogProfileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLogProfileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

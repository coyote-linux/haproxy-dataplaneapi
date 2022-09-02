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

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteDeclareCaptureAcceptedCode is the HTTP code returned for type DeleteDeclareCaptureAccepted
const DeleteDeclareCaptureAcceptedCode int = 202

/*
DeleteDeclareCaptureAccepted Configuration change accepted and reload requested

swagger:response deleteDeclareCaptureAccepted
*/
type DeleteDeclareCaptureAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteDeclareCaptureAccepted creates DeleteDeclareCaptureAccepted with default headers values
func NewDeleteDeclareCaptureAccepted() *DeleteDeclareCaptureAccepted {

	return &DeleteDeclareCaptureAccepted{}
}

// WithReloadID adds the reloadId to the delete declare capture accepted response
func (o *DeleteDeclareCaptureAccepted) WithReloadID(reloadID string) *DeleteDeclareCaptureAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete declare capture accepted response
func (o *DeleteDeclareCaptureAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteDeclareCaptureAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteDeclareCaptureNoContentCode is the HTTP code returned for type DeleteDeclareCaptureNoContent
const DeleteDeclareCaptureNoContentCode int = 204

/*
DeleteDeclareCaptureNoContent Declare Capture deleted

swagger:response deleteDeclareCaptureNoContent
*/
type DeleteDeclareCaptureNoContent struct {
}

// NewDeleteDeclareCaptureNoContent creates DeleteDeclareCaptureNoContent with default headers values
func NewDeleteDeclareCaptureNoContent() *DeleteDeclareCaptureNoContent {

	return &DeleteDeclareCaptureNoContent{}
}

// WriteResponse to the client
func (o *DeleteDeclareCaptureNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteDeclareCaptureNotFoundCode is the HTTP code returned for type DeleteDeclareCaptureNotFound
const DeleteDeclareCaptureNotFoundCode int = 404

/*
DeleteDeclareCaptureNotFound The specified resource was not found

swagger:response deleteDeclareCaptureNotFound
*/
type DeleteDeclareCaptureNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDeclareCaptureNotFound creates DeleteDeclareCaptureNotFound with default headers values
func NewDeleteDeclareCaptureNotFound() *DeleteDeclareCaptureNotFound {

	return &DeleteDeclareCaptureNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete declare capture not found response
func (o *DeleteDeclareCaptureNotFound) WithConfigurationVersion(configurationVersion string) *DeleteDeclareCaptureNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete declare capture not found response
func (o *DeleteDeclareCaptureNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete declare capture not found response
func (o *DeleteDeclareCaptureNotFound) WithPayload(payload *models.Error) *DeleteDeclareCaptureNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete declare capture not found response
func (o *DeleteDeclareCaptureNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeclareCaptureNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteDeclareCaptureDefault General Error

swagger:response deleteDeclareCaptureDefault
*/
type DeleteDeclareCaptureDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDeclareCaptureDefault creates DeleteDeclareCaptureDefault with default headers values
func NewDeleteDeclareCaptureDefault(code int) *DeleteDeclareCaptureDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteDeclareCaptureDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete declare capture default response
func (o *DeleteDeclareCaptureDefault) WithStatusCode(code int) *DeleteDeclareCaptureDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete declare capture default response
func (o *DeleteDeclareCaptureDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete declare capture default response
func (o *DeleteDeclareCaptureDefault) WithConfigurationVersion(configurationVersion string) *DeleteDeclareCaptureDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete declare capture default response
func (o *DeleteDeclareCaptureDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete declare capture default response
func (o *DeleteDeclareCaptureDefault) WithPayload(payload *models.Error) *DeleteDeclareCaptureDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete declare capture default response
func (o *DeleteDeclareCaptureDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeclareCaptureDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

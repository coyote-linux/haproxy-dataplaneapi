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

// CreateLogProfileCreatedCode is the HTTP code returned for type CreateLogProfileCreated
const CreateLogProfileCreatedCode int = 201

/*
CreateLogProfileCreated Log Profile created

swagger:response createLogProfileCreated
*/
type CreateLogProfileCreated struct {

	/*
	  In: Body
	*/
	Payload *models.LogProfile `json:"body,omitempty"`
}

// NewCreateLogProfileCreated creates CreateLogProfileCreated with default headers values
func NewCreateLogProfileCreated() *CreateLogProfileCreated {

	return &CreateLogProfileCreated{}
}

// WithPayload adds the payload to the create log profile created response
func (o *CreateLogProfileCreated) WithPayload(payload *models.LogProfile) *CreateLogProfileCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log profile created response
func (o *CreateLogProfileCreated) SetPayload(payload *models.LogProfile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogProfileCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateLogProfileAcceptedCode is the HTTP code returned for type CreateLogProfileAccepted
const CreateLogProfileAcceptedCode int = 202

/*
CreateLogProfileAccepted Configuration change accepted and reload requested

swagger:response createLogProfileAccepted
*/
type CreateLogProfileAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.LogProfile `json:"body,omitempty"`
}

// NewCreateLogProfileAccepted creates CreateLogProfileAccepted with default headers values
func NewCreateLogProfileAccepted() *CreateLogProfileAccepted {

	return &CreateLogProfileAccepted{}
}

// WithReloadID adds the reloadId to the create log profile accepted response
func (o *CreateLogProfileAccepted) WithReloadID(reloadID string) *CreateLogProfileAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create log profile accepted response
func (o *CreateLogProfileAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create log profile accepted response
func (o *CreateLogProfileAccepted) WithPayload(payload *models.LogProfile) *CreateLogProfileAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log profile accepted response
func (o *CreateLogProfileAccepted) SetPayload(payload *models.LogProfile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogProfileAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateLogProfileBadRequestCode is the HTTP code returned for type CreateLogProfileBadRequest
const CreateLogProfileBadRequestCode int = 400

/*
CreateLogProfileBadRequest Bad request

swagger:response createLogProfileBadRequest
*/
type CreateLogProfileBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogProfileBadRequest creates CreateLogProfileBadRequest with default headers values
func NewCreateLogProfileBadRequest() *CreateLogProfileBadRequest {

	return &CreateLogProfileBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create log profile bad request response
func (o *CreateLogProfileBadRequest) WithConfigurationVersion(configurationVersion string) *CreateLogProfileBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log profile bad request response
func (o *CreateLogProfileBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log profile bad request response
func (o *CreateLogProfileBadRequest) WithPayload(payload *models.Error) *CreateLogProfileBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log profile bad request response
func (o *CreateLogProfileBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogProfileBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateLogProfileConflictCode is the HTTP code returned for type CreateLogProfileConflict
const CreateLogProfileConflictCode int = 409

/*
CreateLogProfileConflict The specified resource already exists

swagger:response createLogProfileConflict
*/
type CreateLogProfileConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogProfileConflict creates CreateLogProfileConflict with default headers values
func NewCreateLogProfileConflict() *CreateLogProfileConflict {

	return &CreateLogProfileConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create log profile conflict response
func (o *CreateLogProfileConflict) WithConfigurationVersion(configurationVersion string) *CreateLogProfileConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log profile conflict response
func (o *CreateLogProfileConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log profile conflict response
func (o *CreateLogProfileConflict) WithPayload(payload *models.Error) *CreateLogProfileConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log profile conflict response
func (o *CreateLogProfileConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogProfileConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateLogProfileDefault General Error

swagger:response createLogProfileDefault
*/
type CreateLogProfileDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogProfileDefault creates CreateLogProfileDefault with default headers values
func NewCreateLogProfileDefault(code int) *CreateLogProfileDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateLogProfileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create log profile default response
func (o *CreateLogProfileDefault) WithStatusCode(code int) *CreateLogProfileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create log profile default response
func (o *CreateLogProfileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create log profile default response
func (o *CreateLogProfileDefault) WithConfigurationVersion(configurationVersion string) *CreateLogProfileDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log profile default response
func (o *CreateLogProfileDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log profile default response
func (o *CreateLogProfileDefault) WithPayload(payload *models.Error) *CreateLogProfileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log profile default response
func (o *CreateLogProfileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogProfileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

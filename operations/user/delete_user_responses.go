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

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteUserAcceptedCode is the HTTP code returned for type DeleteUserAccepted
const DeleteUserAcceptedCode int = 202

/*
DeleteUserAccepted Configuration change accepted and reload requested

swagger:response deleteUserAccepted
*/
type DeleteUserAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteUserAccepted creates DeleteUserAccepted with default headers values
func NewDeleteUserAccepted() *DeleteUserAccepted {

	return &DeleteUserAccepted{}
}

// WithReloadID adds the reloadId to the delete user accepted response
func (o *DeleteUserAccepted) WithReloadID(reloadID string) *DeleteUserAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete user accepted response
func (o *DeleteUserAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteUserAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteUserNoContentCode is the HTTP code returned for type DeleteUserNoContent
const DeleteUserNoContentCode int = 204

/*
DeleteUserNoContent User deleted

swagger:response deleteUserNoContent
*/
type DeleteUserNoContent struct {
}

// NewDeleteUserNoContent creates DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {

	return &DeleteUserNoContent{}
}

// WriteResponse to the client
func (o *DeleteUserNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteUserNotFoundCode is the HTTP code returned for type DeleteUserNotFound
const DeleteUserNotFoundCode int = 404

/*
DeleteUserNotFound The specified resource was not found

swagger:response deleteUserNotFound
*/
type DeleteUserNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserNotFound creates DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {

	return &DeleteUserNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete user not found response
func (o *DeleteUserNotFound) WithConfigurationVersion(configurationVersion string) *DeleteUserNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete user not found response
func (o *DeleteUserNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete user not found response
func (o *DeleteUserNotFound) WithPayload(payload *models.Error) *DeleteUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user not found response
func (o *DeleteUserNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteUserDefault General Error

swagger:response deleteUserDefault
*/
type DeleteUserDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserDefault creates DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete user default response
func (o *DeleteUserDefault) WithStatusCode(code int) *DeleteUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete user default response
func (o *DeleteUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete user default response
func (o *DeleteUserDefault) WithConfigurationVersion(configurationVersion string) *DeleteUserDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete user default response
func (o *DeleteUserDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete user default response
func (o *DeleteUserDefault) WithPayload(payload *models.Error) *DeleteUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user default response
func (o *DeleteUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

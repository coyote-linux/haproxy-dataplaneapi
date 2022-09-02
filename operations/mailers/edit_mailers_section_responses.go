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

package mailers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// EditMailersSectionOKCode is the HTTP code returned for type EditMailersSectionOK
const EditMailersSectionOKCode int = 200

/*
EditMailersSectionOK Mailers configuration updated

swagger:response editMailersSectionOK
*/
type EditMailersSectionOK struct {

	/*
	  In: Body
	*/
	Payload *models.MailersSection `json:"body,omitempty"`
}

// NewEditMailersSectionOK creates EditMailersSectionOK with default headers values
func NewEditMailersSectionOK() *EditMailersSectionOK {

	return &EditMailersSectionOK{}
}

// WithPayload adds the payload to the edit mailers section o k response
func (o *EditMailersSectionOK) WithPayload(payload *models.MailersSection) *EditMailersSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit mailers section o k response
func (o *EditMailersSectionOK) SetPayload(payload *models.MailersSection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditMailersSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EditMailersSectionAcceptedCode is the HTTP code returned for type EditMailersSectionAccepted
const EditMailersSectionAcceptedCode int = 202

/*
EditMailersSectionAccepted Configuration change accepted and reload requested

swagger:response editMailersSectionAccepted
*/
type EditMailersSectionAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.MailersSection `json:"body,omitempty"`
}

// NewEditMailersSectionAccepted creates EditMailersSectionAccepted with default headers values
func NewEditMailersSectionAccepted() *EditMailersSectionAccepted {

	return &EditMailersSectionAccepted{}
}

// WithReloadID adds the reloadId to the edit mailers section accepted response
func (o *EditMailersSectionAccepted) WithReloadID(reloadID string) *EditMailersSectionAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the edit mailers section accepted response
func (o *EditMailersSectionAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the edit mailers section accepted response
func (o *EditMailersSectionAccepted) WithPayload(payload *models.MailersSection) *EditMailersSectionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit mailers section accepted response
func (o *EditMailersSectionAccepted) SetPayload(payload *models.MailersSection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditMailersSectionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// EditMailersSectionBadRequestCode is the HTTP code returned for type EditMailersSectionBadRequest
const EditMailersSectionBadRequestCode int = 400

/*
EditMailersSectionBadRequest Bad request

swagger:response editMailersSectionBadRequest
*/
type EditMailersSectionBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEditMailersSectionBadRequest creates EditMailersSectionBadRequest with default headers values
func NewEditMailersSectionBadRequest() *EditMailersSectionBadRequest {

	return &EditMailersSectionBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the edit mailers section bad request response
func (o *EditMailersSectionBadRequest) WithConfigurationVersion(configurationVersion string) *EditMailersSectionBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the edit mailers section bad request response
func (o *EditMailersSectionBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the edit mailers section bad request response
func (o *EditMailersSectionBadRequest) WithPayload(payload *models.Error) *EditMailersSectionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit mailers section bad request response
func (o *EditMailersSectionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditMailersSectionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// EditMailersSectionNotFoundCode is the HTTP code returned for type EditMailersSectionNotFound
const EditMailersSectionNotFoundCode int = 404

/*
EditMailersSectionNotFound The specified resource was not found

swagger:response editMailersSectionNotFound
*/
type EditMailersSectionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEditMailersSectionNotFound creates EditMailersSectionNotFound with default headers values
func NewEditMailersSectionNotFound() *EditMailersSectionNotFound {

	return &EditMailersSectionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the edit mailers section not found response
func (o *EditMailersSectionNotFound) WithConfigurationVersion(configurationVersion string) *EditMailersSectionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the edit mailers section not found response
func (o *EditMailersSectionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the edit mailers section not found response
func (o *EditMailersSectionNotFound) WithPayload(payload *models.Error) *EditMailersSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit mailers section not found response
func (o *EditMailersSectionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditMailersSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
EditMailersSectionDefault General Error

swagger:response editMailersSectionDefault
*/
type EditMailersSectionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEditMailersSectionDefault creates EditMailersSectionDefault with default headers values
func NewEditMailersSectionDefault(code int) *EditMailersSectionDefault {
	if code <= 0 {
		code = 500
	}

	return &EditMailersSectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the edit mailers section default response
func (o *EditMailersSectionDefault) WithStatusCode(code int) *EditMailersSectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the edit mailers section default response
func (o *EditMailersSectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the edit mailers section default response
func (o *EditMailersSectionDefault) WithConfigurationVersion(configurationVersion string) *EditMailersSectionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the edit mailers section default response
func (o *EditMailersSectionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the edit mailers section default response
func (o *EditMailersSectionDefault) WithPayload(payload *models.Error) *EditMailersSectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit mailers section default response
func (o *EditMailersSectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditMailersSectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

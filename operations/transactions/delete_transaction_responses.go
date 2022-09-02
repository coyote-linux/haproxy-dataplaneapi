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

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteTransactionNoContentCode is the HTTP code returned for type DeleteTransactionNoContent
const DeleteTransactionNoContentCode int = 204

/*
DeleteTransactionNoContent Transaction deleted

swagger:response deleteTransactionNoContent
*/
type DeleteTransactionNoContent struct {
}

// NewDeleteTransactionNoContent creates DeleteTransactionNoContent with default headers values
func NewDeleteTransactionNoContent() *DeleteTransactionNoContent {

	return &DeleteTransactionNoContent{}
}

// WriteResponse to the client
func (o *DeleteTransactionNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTransactionNotFoundCode is the HTTP code returned for type DeleteTransactionNotFound
const DeleteTransactionNotFoundCode int = 404

/*
DeleteTransactionNotFound The specified resource was not found

swagger:response deleteTransactionNotFound
*/
type DeleteTransactionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTransactionNotFound creates DeleteTransactionNotFound with default headers values
func NewDeleteTransactionNotFound() *DeleteTransactionNotFound {

	return &DeleteTransactionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete transaction not found response
func (o *DeleteTransactionNotFound) WithConfigurationVersion(configurationVersion string) *DeleteTransactionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete transaction not found response
func (o *DeleteTransactionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete transaction not found response
func (o *DeleteTransactionNotFound) WithPayload(payload *models.Error) *DeleteTransactionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete transaction not found response
func (o *DeleteTransactionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTransactionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteTransactionDefault General Error

swagger:response deleteTransactionDefault
*/
type DeleteTransactionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTransactionDefault creates DeleteTransactionDefault with default headers values
func NewDeleteTransactionDefault(code int) *DeleteTransactionDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTransactionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete transaction default response
func (o *DeleteTransactionDefault) WithStatusCode(code int) *DeleteTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete transaction default response
func (o *DeleteTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete transaction default response
func (o *DeleteTransactionDefault) WithConfigurationVersion(configurationVersion string) *DeleteTransactionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete transaction default response
func (o *DeleteTransactionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete transaction default response
func (o *DeleteTransactionDefault) WithPayload(payload *models.Error) *DeleteTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete transaction default response
func (o *DeleteTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

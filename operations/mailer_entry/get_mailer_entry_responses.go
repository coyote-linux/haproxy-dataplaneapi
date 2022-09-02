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

package mailer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetMailerEntryOKCode is the HTTP code returned for type GetMailerEntryOK
const GetMailerEntryOKCode int = 200

/*
GetMailerEntryOK Successful operation

swagger:response getMailerEntryOK
*/
type GetMailerEntryOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetMailerEntryOKBody `json:"body,omitempty"`
}

// NewGetMailerEntryOK creates GetMailerEntryOK with default headers values
func NewGetMailerEntryOK() *GetMailerEntryOK {

	return &GetMailerEntryOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get mailer entry o k response
func (o *GetMailerEntryOK) WithConfigurationVersion(configurationVersion string) *GetMailerEntryOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get mailer entry o k response
func (o *GetMailerEntryOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get mailer entry o k response
func (o *GetMailerEntryOK) WithPayload(payload *GetMailerEntryOKBody) *GetMailerEntryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get mailer entry o k response
func (o *GetMailerEntryOK) SetPayload(payload *GetMailerEntryOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMailerEntryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetMailerEntryNotFoundCode is the HTTP code returned for type GetMailerEntryNotFound
const GetMailerEntryNotFoundCode int = 404

/*
GetMailerEntryNotFound The specified resource already exists

swagger:response getMailerEntryNotFound
*/
type GetMailerEntryNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMailerEntryNotFound creates GetMailerEntryNotFound with default headers values
func NewGetMailerEntryNotFound() *GetMailerEntryNotFound {

	return &GetMailerEntryNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get mailer entry not found response
func (o *GetMailerEntryNotFound) WithConfigurationVersion(configurationVersion string) *GetMailerEntryNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get mailer entry not found response
func (o *GetMailerEntryNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get mailer entry not found response
func (o *GetMailerEntryNotFound) WithPayload(payload *models.Error) *GetMailerEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get mailer entry not found response
func (o *GetMailerEntryNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMailerEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetMailerEntryDefault General Error

swagger:response getMailerEntryDefault
*/
type GetMailerEntryDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMailerEntryDefault creates GetMailerEntryDefault with default headers values
func NewGetMailerEntryDefault(code int) *GetMailerEntryDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMailerEntryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get mailer entry default response
func (o *GetMailerEntryDefault) WithStatusCode(code int) *GetMailerEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get mailer entry default response
func (o *GetMailerEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get mailer entry default response
func (o *GetMailerEntryDefault) WithConfigurationVersion(configurationVersion string) *GetMailerEntryDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get mailer entry default response
func (o *GetMailerEntryDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get mailer entry default response
func (o *GetMailerEntryDefault) WithPayload(payload *models.Error) *GetMailerEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get mailer entry default response
func (o *GetMailerEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMailerEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

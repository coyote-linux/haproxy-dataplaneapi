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

package http_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetHTTPRequestRuleOKCode is the HTTP code returned for type GetHTTPRequestRuleOK
const GetHTTPRequestRuleOKCode int = 200

/*
GetHTTPRequestRuleOK Successful operation

swagger:response getHttpRequestRuleOK
*/
type GetHTTPRequestRuleOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetHTTPRequestRuleOKBody `json:"body,omitempty"`
}

// NewGetHTTPRequestRuleOK creates GetHTTPRequestRuleOK with default headers values
func NewGetHTTPRequestRuleOK() *GetHTTPRequestRuleOK {

	return &GetHTTPRequestRuleOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http request rule o k response
func (o *GetHTTPRequestRuleOK) WithConfigurationVersion(configurationVersion string) *GetHTTPRequestRuleOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http request rule o k response
func (o *GetHTTPRequestRuleOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http request rule o k response
func (o *GetHTTPRequestRuleOK) WithPayload(payload *GetHTTPRequestRuleOKBody) *GetHTTPRequestRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http request rule o k response
func (o *GetHTTPRequestRuleOK) SetPayload(payload *GetHTTPRequestRuleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPRequestRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetHTTPRequestRuleNotFoundCode is the HTTP code returned for type GetHTTPRequestRuleNotFound
const GetHTTPRequestRuleNotFoundCode int = 404

/*
GetHTTPRequestRuleNotFound The specified resource was not found

swagger:response getHttpRequestRuleNotFound
*/
type GetHTTPRequestRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPRequestRuleNotFound creates GetHTTPRequestRuleNotFound with default headers values
func NewGetHTTPRequestRuleNotFound() *GetHTTPRequestRuleNotFound {

	return &GetHTTPRequestRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http request rule not found response
func (o *GetHTTPRequestRuleNotFound) WithConfigurationVersion(configurationVersion string) *GetHTTPRequestRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http request rule not found response
func (o *GetHTTPRequestRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http request rule not found response
func (o *GetHTTPRequestRuleNotFound) WithPayload(payload *models.Error) *GetHTTPRequestRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http request rule not found response
func (o *GetHTTPRequestRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPRequestRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetHTTPRequestRuleDefault General Error

swagger:response getHttpRequestRuleDefault
*/
type GetHTTPRequestRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPRequestRuleDefault creates GetHTTPRequestRuleDefault with default headers values
func NewGetHTTPRequestRuleDefault(code int) *GetHTTPRequestRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHTTPRequestRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get HTTP request rule default response
func (o *GetHTTPRequestRuleDefault) WithStatusCode(code int) *GetHTTPRequestRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get HTTP request rule default response
func (o *GetHTTPRequestRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get HTTP request rule default response
func (o *GetHTTPRequestRuleDefault) WithConfigurationVersion(configurationVersion string) *GetHTTPRequestRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get HTTP request rule default response
func (o *GetHTTPRequestRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get HTTP request rule default response
func (o *GetHTTPRequestRuleDefault) WithPayload(payload *models.Error) *GetHTTPRequestRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get HTTP request rule default response
func (o *GetHTTPRequestRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPRequestRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

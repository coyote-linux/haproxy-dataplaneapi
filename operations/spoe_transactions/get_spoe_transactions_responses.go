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

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetSpoeTransactionsOKCode is the HTTP code returned for type GetSpoeTransactionsOK
const GetSpoeTransactionsOKCode int = 200

/*
GetSpoeTransactionsOK Success

swagger:response getSpoeTransactionsOK
*/
type GetSpoeTransactionsOK struct {

	/*
	  In: Body
	*/
	Payload models.SpoeTransactions `json:"body,omitempty"`
}

// NewGetSpoeTransactionsOK creates GetSpoeTransactionsOK with default headers values
func NewGetSpoeTransactionsOK() *GetSpoeTransactionsOK {

	return &GetSpoeTransactionsOK{}
}

// WithPayload adds the payload to the get spoe transactions o k response
func (o *GetSpoeTransactionsOK) WithPayload(payload models.SpoeTransactions) *GetSpoeTransactionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe transactions o k response
func (o *GetSpoeTransactionsOK) SetPayload(payload models.SpoeTransactions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeTransactionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.SpoeTransactions{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetSpoeTransactionsDefault General Error

swagger:response getSpoeTransactionsDefault
*/
type GetSpoeTransactionsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeTransactionsDefault creates GetSpoeTransactionsDefault with default headers values
func NewGetSpoeTransactionsDefault(code int) *GetSpoeTransactionsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSpoeTransactionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get spoe transactions default response
func (o *GetSpoeTransactionsDefault) WithStatusCode(code int) *GetSpoeTransactionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoe transactions default response
func (o *GetSpoeTransactionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get spoe transactions default response
func (o *GetSpoeTransactionsDefault) WithConfigurationVersion(configurationVersion string) *GetSpoeTransactionsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe transactions default response
func (o *GetSpoeTransactionsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe transactions default response
func (o *GetSpoeTransactionsDefault) WithPayload(payload *models.Error) *GetSpoeTransactionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe transactions default response
func (o *GetSpoeTransactionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeTransactionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

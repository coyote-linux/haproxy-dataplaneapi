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

package cache

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceCacheOKCode is the HTTP code returned for type ReplaceCacheOK
const ReplaceCacheOKCode int = 200

/*ReplaceCacheOK Cache replaced

swagger:response replaceCacheOK
*/
type ReplaceCacheOK struct {

	/*
	  In: Body
	*/
	Payload *models.Cache `json:"body,omitempty"`
}

// NewReplaceCacheOK creates ReplaceCacheOK with default headers values
func NewReplaceCacheOK() *ReplaceCacheOK {

	return &ReplaceCacheOK{}
}

// WithPayload adds the payload to the replace cache o k response
func (o *ReplaceCacheOK) WithPayload(payload *models.Cache) *ReplaceCacheOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace cache o k response
func (o *ReplaceCacheOK) SetPayload(payload *models.Cache) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCacheOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCacheAcceptedCode is the HTTP code returned for type ReplaceCacheAccepted
const ReplaceCacheAcceptedCode int = 202

/*ReplaceCacheAccepted Configuration change accepted and reload requested

swagger:response replaceCacheAccepted
*/
type ReplaceCacheAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Cache `json:"body,omitempty"`
}

// NewReplaceCacheAccepted creates ReplaceCacheAccepted with default headers values
func NewReplaceCacheAccepted() *ReplaceCacheAccepted {

	return &ReplaceCacheAccepted{}
}

// WithReloadID adds the reloadId to the replace cache accepted response
func (o *ReplaceCacheAccepted) WithReloadID(reloadID string) *ReplaceCacheAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace cache accepted response
func (o *ReplaceCacheAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace cache accepted response
func (o *ReplaceCacheAccepted) WithPayload(payload *models.Cache) *ReplaceCacheAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace cache accepted response
func (o *ReplaceCacheAccepted) SetPayload(payload *models.Cache) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCacheAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceCacheBadRequestCode is the HTTP code returned for type ReplaceCacheBadRequest
const ReplaceCacheBadRequestCode int = 400

/*ReplaceCacheBadRequest Bad request

swagger:response replaceCacheBadRequest
*/
type ReplaceCacheBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceCacheBadRequest creates ReplaceCacheBadRequest with default headers values
func NewReplaceCacheBadRequest() *ReplaceCacheBadRequest {

	return &ReplaceCacheBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace cache bad request response
func (o *ReplaceCacheBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceCacheBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace cache bad request response
func (o *ReplaceCacheBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace cache bad request response
func (o *ReplaceCacheBadRequest) WithPayload(payload *models.Error) *ReplaceCacheBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace cache bad request response
func (o *ReplaceCacheBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCacheBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceCacheNotFoundCode is the HTTP code returned for type ReplaceCacheNotFound
const ReplaceCacheNotFoundCode int = 404

/*ReplaceCacheNotFound The specified resource was not found

swagger:response replaceCacheNotFound
*/
type ReplaceCacheNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceCacheNotFound creates ReplaceCacheNotFound with default headers values
func NewReplaceCacheNotFound() *ReplaceCacheNotFound {

	return &ReplaceCacheNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace cache not found response
func (o *ReplaceCacheNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceCacheNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace cache not found response
func (o *ReplaceCacheNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace cache not found response
func (o *ReplaceCacheNotFound) WithPayload(payload *models.Error) *ReplaceCacheNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace cache not found response
func (o *ReplaceCacheNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCacheNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceCacheDefault General Error

swagger:response replaceCacheDefault
*/
type ReplaceCacheDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceCacheDefault creates ReplaceCacheDefault with default headers values
func NewReplaceCacheDefault(code int) *ReplaceCacheDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceCacheDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace cache default response
func (o *ReplaceCacheDefault) WithStatusCode(code int) *ReplaceCacheDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace cache default response
func (o *ReplaceCacheDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace cache default response
func (o *ReplaceCacheDefault) WithConfigurationVersion(configurationVersion string) *ReplaceCacheDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace cache default response
func (o *ReplaceCacheDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace cache default response
func (o *ReplaceCacheDefault) WithPayload(payload *models.Error) *ReplaceCacheDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace cache default response
func (o *ReplaceCacheDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCacheDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

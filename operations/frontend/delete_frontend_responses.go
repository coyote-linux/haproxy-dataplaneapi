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

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// DeleteFrontendAcceptedCode is the HTTP code returned for type DeleteFrontendAccepted
const DeleteFrontendAcceptedCode int = 202

/*DeleteFrontendAccepted Configuration change accepted and reload requested

swagger:response deleteFrontendAccepted
*/
type DeleteFrontendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteFrontendAccepted creates DeleteFrontendAccepted with default headers values
func NewDeleteFrontendAccepted() *DeleteFrontendAccepted {

	return &DeleteFrontendAccepted{}
}

// WithReloadID adds the reloadId to the delete frontend accepted response
func (o *DeleteFrontendAccepted) WithReloadID(reloadID string) *DeleteFrontendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete frontend accepted response
func (o *DeleteFrontendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteFrontendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteFrontendNoContentCode is the HTTP code returned for type DeleteFrontendNoContent
const DeleteFrontendNoContentCode int = 204

/*DeleteFrontendNoContent Frontend deleted

swagger:response deleteFrontendNoContent
*/
type DeleteFrontendNoContent struct {
}

// NewDeleteFrontendNoContent creates DeleteFrontendNoContent with default headers values
func NewDeleteFrontendNoContent() *DeleteFrontendNoContent {

	return &DeleteFrontendNoContent{}
}

// WriteResponse to the client
func (o *DeleteFrontendNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteFrontendNotFoundCode is the HTTP code returned for type DeleteFrontendNotFound
const DeleteFrontendNotFoundCode int = 404

/*DeleteFrontendNotFound The specified resource was not found

swagger:response deleteFrontendNotFound
*/
type DeleteFrontendNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteFrontendNotFound creates DeleteFrontendNotFound with default headers values
func NewDeleteFrontendNotFound() *DeleteFrontendNotFound {

	return &DeleteFrontendNotFound{}
}

// WithPayload adds the payload to the delete frontend not found response
func (o *DeleteFrontendNotFound) WithPayload(payload *models.Error) *DeleteFrontendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete frontend not found response
func (o *DeleteFrontendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFrontendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteFrontendDefault General Error

swagger:response deleteFrontendDefault
*/
type DeleteFrontendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteFrontendDefault creates DeleteFrontendDefault with default headers values
func NewDeleteFrontendDefault(code int) *DeleteFrontendDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteFrontendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete frontend default response
func (o *DeleteFrontendDefault) WithStatusCode(code int) *DeleteFrontendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete frontend default response
func (o *DeleteFrontendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete frontend default response
func (o *DeleteFrontendDefault) WithPayload(payload *models.Error) *DeleteFrontendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete frontend default response
func (o *DeleteFrontendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFrontendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

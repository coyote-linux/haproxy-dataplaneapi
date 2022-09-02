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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReplaceRuntimeMapEntryHandlerFunc turns a function with the right signature into a replace runtime map entry handler
type ReplaceRuntimeMapEntryHandlerFunc func(ReplaceRuntimeMapEntryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceRuntimeMapEntryHandlerFunc) Handle(params ReplaceRuntimeMapEntryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceRuntimeMapEntryHandler interface for that can handle valid replace runtime map entry params
type ReplaceRuntimeMapEntryHandler interface {
	Handle(ReplaceRuntimeMapEntryParams, interface{}) middleware.Responder
}

// NewReplaceRuntimeMapEntry creates a new http.Handler for the replace runtime map entry operation
func NewReplaceRuntimeMapEntry(ctx *middleware.Context, handler ReplaceRuntimeMapEntryHandler) *ReplaceRuntimeMapEntry {
	return &ReplaceRuntimeMapEntry{Context: ctx, Handler: handler}
}

/*
	ReplaceRuntimeMapEntry swagger:route PUT /services/haproxy/runtime/maps_entries/{id} Maps replaceRuntimeMapEntry

# Replace the value corresponding to each id in a map

Replaces the value corresponding to each id in a map.
*/
type ReplaceRuntimeMapEntry struct {
	Context *middleware.Context
	Handler ReplaceRuntimeMapEntryHandler
}

func (o *ReplaceRuntimeMapEntry) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewReplaceRuntimeMapEntryParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ReplaceRuntimeMapEntryBody replace runtime map entry body
//
// swagger:model ReplaceRuntimeMapEntryBody
type ReplaceRuntimeMapEntryBody struct {

	// Map value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this replace runtime map entry body
func (o *ReplaceRuntimeMapEntryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplaceRuntimeMapEntryBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this replace runtime map entry body based on context it is used
func (o *ReplaceRuntimeMapEntryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ReplaceRuntimeMapEntryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplaceRuntimeMapEntryBody) UnmarshalBinary(b []byte) error {
	var res ReplaceRuntimeMapEntryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

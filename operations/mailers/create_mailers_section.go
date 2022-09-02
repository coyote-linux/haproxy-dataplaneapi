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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateMailersSectionHandlerFunc turns a function with the right signature into a create mailers section handler
type CreateMailersSectionHandlerFunc func(CreateMailersSectionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateMailersSectionHandlerFunc) Handle(params CreateMailersSectionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateMailersSectionHandler interface for that can handle valid create mailers section params
type CreateMailersSectionHandler interface {
	Handle(CreateMailersSectionParams, interface{}) middleware.Responder
}

// NewCreateMailersSection creates a new http.Handler for the create mailers section operation
func NewCreateMailersSection(ctx *middleware.Context, handler CreateMailersSectionHandler) *CreateMailersSection {
	return &CreateMailersSection{Context: ctx, Handler: handler}
}

/*
	CreateMailersSection swagger:route POST /services/haproxy/configuration/mailers_section Mailers createMailersSection

# Add a new Mailers section

Creates a new empty Mailers section
*/
type CreateMailersSection struct {
	Context *middleware.Context
	Handler CreateMailersSectionHandler
}

func (o *CreateMailersSection) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateMailersSectionParams()
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

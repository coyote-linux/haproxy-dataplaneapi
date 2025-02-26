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

package traces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateTraceEntryHandlerFunc turns a function with the right signature into a create trace entry handler
type CreateTraceEntryHandlerFunc func(CreateTraceEntryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTraceEntryHandlerFunc) Handle(params CreateTraceEntryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateTraceEntryHandler interface for that can handle valid create trace entry params
type CreateTraceEntryHandler interface {
	Handle(CreateTraceEntryParams, interface{}) middleware.Responder
}

// NewCreateTraceEntry creates a new http.Handler for the create trace entry operation
func NewCreateTraceEntry(ctx *middleware.Context, handler CreateTraceEntryHandler) *CreateTraceEntry {
	return &CreateTraceEntry{Context: ctx, Handler: handler}
}

/*
	CreateTraceEntry swagger:route POST /services/haproxy/configuration/traces/entries Traces createTraceEntry

# Add a new trace entry

Adds a new trace entry into the traces section. The traces section will be created if needed.
*/
type CreateTraceEntry struct {
	Context *middleware.Context
	Handler CreateTraceEntryHandler
}

func (o *CreateTraceEntry) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateTraceEntryParams()
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

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

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReplaceTCPRequestRuleHandlerFunc turns a function with the right signature into a replace TCP request rule handler
type ReplaceTCPRequestRuleHandlerFunc func(ReplaceTCPRequestRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceTCPRequestRuleHandlerFunc) Handle(params ReplaceTCPRequestRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceTCPRequestRuleHandler interface for that can handle valid replace TCP request rule params
type ReplaceTCPRequestRuleHandler interface {
	Handle(ReplaceTCPRequestRuleParams, interface{}) middleware.Responder
}

// NewReplaceTCPRequestRule creates a new http.Handler for the replace TCP request rule operation
func NewReplaceTCPRequestRule(ctx *middleware.Context, handler ReplaceTCPRequestRuleHandler) *ReplaceTCPRequestRule {
	return &ReplaceTCPRequestRule{Context: ctx, Handler: handler}
}

/*ReplaceTCPRequestRule swagger:route PUT /services/haproxy/configuration/tcp_request_rules/{id} TCPRequestRule replaceTcpRequestRule

Replace a TCP Request Rule

Replaces a TCP Request Rule configuration by it's ID in the specified parent.

*/
type ReplaceTCPRequestRule struct {
	Context *middleware.Context
	Handler ReplaceTCPRequestRuleHandler
}

func (o *ReplaceTCPRequestRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceTCPRequestRuleParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

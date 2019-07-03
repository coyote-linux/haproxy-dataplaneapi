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

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetStatsHandlerFunc turns a function with the right signature into a get stats handler
type GetStatsHandlerFunc func(GetStatsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStatsHandlerFunc) Handle(params GetStatsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetStatsHandler interface for that can handle valid get stats params
type GetStatsHandler interface {
	Handle(GetStatsParams, interface{}) middleware.Responder
}

// NewGetStats creates a new http.Handler for the get stats operation
func NewGetStats(ctx *middleware.Context, handler GetStatsHandler) *GetStats {
	return &GetStats{Context: ctx, Handler: handler}
}

/*GetStats swagger:route GET /services/haproxy/stats/native Stats getStats

Gets stats

Getting stats from the HAProxy.

*/
type GetStats struct {
	Context *middleware.Context
	Handler GetStatsHandler
}

func (o *GetStats) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStatsParams()

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

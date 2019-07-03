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

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetTCPResponseRuleHandlerFunc turns a function with the right signature into a get TCP response rule handler
type GetTCPResponseRuleHandlerFunc func(GetTCPResponseRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTCPResponseRuleHandlerFunc) Handle(params GetTCPResponseRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTCPResponseRuleHandler interface for that can handle valid get TCP response rule params
type GetTCPResponseRuleHandler interface {
	Handle(GetTCPResponseRuleParams, interface{}) middleware.Responder
}

// NewGetTCPResponseRule creates a new http.Handler for the get TCP response rule operation
func NewGetTCPResponseRule(ctx *middleware.Context, handler GetTCPResponseRuleHandler) *GetTCPResponseRule {
	return &GetTCPResponseRule{Context: ctx, Handler: handler}
}

/*GetTCPResponseRule swagger:route GET /services/haproxy/configuration/tcp_response_rules/{id} TCPResponseRule getTcpResponseRule

Return one TCP Response Rule

Returns one TCP Response Rule configuration by it's ID in the specified backend.

*/
type GetTCPResponseRule struct {
	Context *middleware.Context
	Handler GetTCPResponseRuleHandler
}

func (o *GetTCPResponseRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTCPResponseRuleParams()

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

// GetTCPResponseRuleOKBody get TCP response rule o k body
// swagger:model GetTCPResponseRuleOKBody
type GetTCPResponseRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.TCPResponseRule `json:"data,omitempty"`
}

// Validate validates this get TCP response rule o k body
func (o *GetTCPResponseRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPResponseRuleOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTcpResponseRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTCPResponseRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTCPResponseRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetTCPResponseRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

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

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLogTargetParams creates a new GetLogTargetParams object
// no default values defined in spec.
func NewGetLogTargetParams() GetLogTargetParams {

	return GetLogTargetParams{}
}

// GetLogTargetParams contains all the bound params for the get log target operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLogTarget
type GetLogTargetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Log Target ID
	  Required: true
	  In: path
	*/
	ID int64
	/*Parent name
	  Required: true
	  In: query
	*/
	ParentName string
	/*Parent type
	  Required: true
	  In: query
	*/
	ParentType string
	/*ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	  In: query
	*/
	TransactionID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLogTargetParams() beforehand.
func (o *GetLogTargetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qParentName, qhkParentName, _ := qs.GetOK("parent_name")
	if err := o.bindParentName(qParentName, qhkParentName, route.Formats); err != nil {
		res = append(res, err)
	}

	qParentType, qhkParentType, _ := qs.GetOK("parent_type")
	if err := o.bindParentType(qParentType, qhkParentType, route.Formats); err != nil {
		res = append(res, err)
	}

	qTransactionID, qhkTransactionID, _ := qs.GetOK("transaction_id")
	if err := o.bindTransactionID(qTransactionID, qhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetLogTargetParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}

// bindParentName binds and validates parameter ParentName from query.
func (o *GetLogTargetParams) bindParentName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("parent_name", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("parent_name", "query", raw); err != nil {
		return err
	}

	o.ParentName = raw

	return nil
}

// bindParentType binds and validates parameter ParentType from query.
func (o *GetLogTargetParams) bindParentType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("parent_type", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("parent_type", "query", raw); err != nil {
		return err
	}

	o.ParentType = raw

	if err := o.validateParentType(formats); err != nil {
		return err
	}

	return nil
}

// validateParentType carries on validations for parameter ParentType
func (o *GetLogTargetParams) validateParentType(formats strfmt.Registry) error {

	if err := validate.Enum("parent_type", "query", o.ParentType, []interface{}{"frontend", "backend"}); err != nil {
		return err
	}

	return nil
}

// bindTransactionID binds and validates parameter TransactionID from query.
func (o *GetLogTargetParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TransactionID = &raw

	return nil
}

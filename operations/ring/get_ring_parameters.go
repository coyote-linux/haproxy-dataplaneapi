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

package ring

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRingParams creates a new GetRingParams object
// with the default values initialized.
func NewGetRingParams() GetRingParams {

	var (
		// initialize parameters with default values

		fullSectionDefault = bool(false)
	)

	return GetRingParams{
		FullSection: &fullSectionDefault,
	}
}

// GetRingParams contains all the bound params for the get ring operation
// typically these are obtained from a http.Request
//
// swagger:parameters getRing
type GetRingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Indicates if the action affects the specified child resources as well
	  In: query
	  Default: false
	*/
	FullSection *bool
	/*Ring name
	  Required: true
	  In: path
	*/
	Name string
	/*ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	  In: query
	*/
	TransactionID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRingParams() beforehand.
func (o *GetRingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFullSection, qhkFullSection, _ := qs.GetOK("full_section")
	if err := o.bindFullSection(qFullSection, qhkFullSection, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
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

// bindFullSection binds and validates parameter FullSection from query.
func (o *GetRingParams) bindFullSection(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRingParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("full_section", "query", "bool", raw)
	}
	o.FullSection = &value

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *GetRingParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Name = raw

	return nil
}

// bindTransactionID binds and validates parameter TransactionID from query.
func (o *GetRingParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

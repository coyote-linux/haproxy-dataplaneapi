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

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// NewCreateHTTPResponseRuleParams creates a new CreateHTTPResponseRuleParams object
// with the default values initialized.
func NewCreateHTTPResponseRuleParams() CreateHTTPResponseRuleParams {

	var (
		// initialize parameters with default values

		forceReloadDefault = bool(false)
	)

	return CreateHTTPResponseRuleParams{
		ForceReload: &forceReloadDefault,
	}
}

// CreateHTTPResponseRuleParams contains all the bound params for the create HTTP response rule operation
// typically these are obtained from a http.Request
//
// swagger:parameters createHTTPResponseRule
type CreateHTTPResponseRuleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Data *models.HTTPResponseRule
	/*If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	  In: query
	  Default: false
	*/
	ForceReload *bool
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
	/*Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	  In: query
	*/
	Version *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateHTTPResponseRuleParams() beforehand.
func (o *CreateHTTPResponseRuleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.HTTPResponseRule
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body"))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body"))
	}
	qForceReload, qhkForceReload, _ := qs.GetOK("force_reload")
	if err := o.bindForceReload(qForceReload, qhkForceReload, route.Formats); err != nil {
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

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindForceReload binds and validates parameter ForceReload from query.
func (o *CreateHTTPResponseRuleParams) bindForceReload(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewCreateHTTPResponseRuleParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("force_reload", "query", "bool", raw)
	}
	o.ForceReload = &value

	return nil
}

// bindParentName binds and validates parameter ParentName from query.
func (o *CreateHTTPResponseRuleParams) bindParentName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *CreateHTTPResponseRuleParams) bindParentType(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *CreateHTTPResponseRuleParams) validateParentType(formats strfmt.Registry) error {

	if err := validate.Enum("parent_type", "query", o.ParentType, []interface{}{"frontend", "backend"}); err != nil {
		return err
	}

	return nil
}

// bindTransactionID binds and validates parameter TransactionID from query.
func (o *CreateHTTPResponseRuleParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindVersion binds and validates parameter Version from query.
func (o *CreateHTTPResponseRuleParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("version", "query", "int64", raw)
	}
	o.Version = &value

	return nil
}

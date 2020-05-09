// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/go_http_client/run_model"
)

// NewCreateRunParams creates a new CreateRunParams object
// with the default values initialized.
func NewCreateRunParams() *CreateRunParams {
	var ()
	return &CreateRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRunParamsWithTimeout creates a new CreateRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRunParamsWithTimeout(timeout time.Duration) *CreateRunParams {
	var ()
	return &CreateRunParams{

		timeout: timeout,
	}
}

// NewCreateRunParamsWithContext creates a new CreateRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRunParamsWithContext(ctx context.Context) *CreateRunParams {
	var ()
	return &CreateRunParams{

		Context: ctx,
	}
}

// NewCreateRunParamsWithHTTPClient creates a new CreateRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRunParamsWithHTTPClient(client *http.Client) *CreateRunParams {
	var ()
	return &CreateRunParams{
		HTTPClient: client,
	}
}

/*CreateRunParams contains all the parameters to send to the API endpoint
for the create run operation typically these are written to a http.Request
*/
type CreateRunParams struct {

	/*Body*/
	Body *run_model.APIRun

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create run params
func (o *CreateRunParams) WithTimeout(timeout time.Duration) *CreateRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create run params
func (o *CreateRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create run params
func (o *CreateRunParams) WithContext(ctx context.Context) *CreateRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create run params
func (o *CreateRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create run params
func (o *CreateRunParams) WithHTTPClient(client *http.Client) *CreateRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create run params
func (o *CreateRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create run params
func (o *CreateRunParams) WithBody(body *run_model.APIRun) *CreateRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create run params
func (o *CreateRunParams) SetBody(body *run_model.APIRun) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

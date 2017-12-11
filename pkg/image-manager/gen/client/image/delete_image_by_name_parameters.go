///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteImageByNameParams creates a new DeleteImageByNameParams object
// with the default values initialized.
func NewDeleteImageByNameParams() *DeleteImageByNameParams {
	var ()
	return &DeleteImageByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteImageByNameParamsWithTimeout creates a new DeleteImageByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteImageByNameParamsWithTimeout(timeout time.Duration) *DeleteImageByNameParams {
	var ()
	return &DeleteImageByNameParams{

		timeout: timeout,
	}
}

// NewDeleteImageByNameParamsWithContext creates a new DeleteImageByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteImageByNameParamsWithContext(ctx context.Context) *DeleteImageByNameParams {
	var ()
	return &DeleteImageByNameParams{

		Context: ctx,
	}
}

// NewDeleteImageByNameParamsWithHTTPClient creates a new DeleteImageByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteImageByNameParamsWithHTTPClient(client *http.Client) *DeleteImageByNameParams {
	var ()
	return &DeleteImageByNameParams{
		HTTPClient: client,
	}
}

/*DeleteImageByNameParams contains all the parameters to send to the API endpoint
for the delete image by name operation typically these are written to a http.Request
*/
type DeleteImageByNameParams struct {

	/*ImageName
	  Name of image to return

	*/
	ImageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete image by name params
func (o *DeleteImageByNameParams) WithTimeout(timeout time.Duration) *DeleteImageByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete image by name params
func (o *DeleteImageByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete image by name params
func (o *DeleteImageByNameParams) WithContext(ctx context.Context) *DeleteImageByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete image by name params
func (o *DeleteImageByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete image by name params
func (o *DeleteImageByNameParams) WithHTTPClient(client *http.Client) *DeleteImageByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete image by name params
func (o *DeleteImageByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageName adds the imageName to the delete image by name params
func (o *DeleteImageByNameParams) WithImageName(imageName string) *DeleteImageByNameParams {
	o.SetImageName(imageName)
	return o
}

// SetImageName adds the imageName to the delete image by name params
func (o *DeleteImageByNameParams) SetImageName(imageName string) {
	o.ImageName = imageName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteImageByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageName
	if err := r.SetPathParam("imageName", o.ImageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
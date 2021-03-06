///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// AddFunctionURL generates an URL for the add function operation
type AddFunctionURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AddFunctionURL) WithBasePath(bp string) *AddFunctionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AddFunctionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *AddFunctionURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1/function"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *AddFunctionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *AddFunctionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *AddFunctionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on AddFunctionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on AddFunctionURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *AddFunctionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

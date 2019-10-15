// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// DeleteStreamURL generates an URL for the delete stream operation
type DeleteStreamURL struct {
	StreamID int64
	TopicID  int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteStreamURL) WithBasePath(bp string) *DeleteStreamURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteStreamURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteStreamURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/topics/{topicId}/streams/{streamId}"

	streamID := swag.FormatInt64(o.StreamID)
	if streamID != "" {
		_path = strings.Replace(_path, "{streamId}", streamID, -1)
	} else {
		return nil, errors.New("streamId is required on DeleteStreamURL")
	}

	topicID := swag.FormatInt64(o.TopicID)
	if topicID != "" {
		_path = strings.Replace(_path, "{topicId}", topicID, -1)
	} else {
		return nil, errors.New("topicId is required on DeleteStreamURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteStreamURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteStreamURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteStreamURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteStreamURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteStreamURL")
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
func (o *DeleteStreamURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

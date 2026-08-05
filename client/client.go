/*
Copyright 2024 ByteDance Ltd. and/or its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package client

import (
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/bytedance/douyin-openapi-credential-go/client"
	util "github.com/bytedance/douyin-openapi-util-go/client"

	"github.com/bytedance/douyin-openapi-sdk-go/internal/transport"
)

// Default runtime settings applied by Init.
const (
	defaultReadTimeout    = 5000
	defaultConnectTimeout = 1000
	defaultMaxAttempts    = 3
)

// Client calls the Douyin OpenAPI. Create one with NewClient and then call the
// generated method for the endpoint you need.
type Client struct {
	// Endpoint overrides the default host of every endpoint that does not
	// pin its own host. Leave it empty to use each endpoint's default.
	Endpoint       *string
	Credential     credential.Credential
	Autoretry      *bool
	IgnoreSSL      *bool
	MaxAttempts    *int
	ReadTimeout    *int
	ConnectTimeout *int
}

func NewClient(config *credential.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *credential.Config) (_err error) {
	client.Credential, _err = credential.NewCredential(config)
	if _err != nil {
		return _err
	}

	client.ReadTimeout = tea.Int(defaultReadTimeout)
	client.ConnectTimeout = tea.Int(defaultConnectTimeout)
	client.Autoretry = tea.Bool(true)
	client.IgnoreSSL = tea.Bool(true)
	client.MaxAttempts = tea.Int(defaultMaxAttempts)
	return nil
}

// SetCredential replaces the credential the client authenticates with.
func (client *Client) SetCredential(credential credential.Credential) {
	client.Credential = credential
}

// SetRuntimeOptions replaces the client's timeout and retry settings.
func (client *Client) SetRuntimeOptions(options *RuntimeOptions) {
	client.ReadTimeout = options.ReadTimeout
	client.ConnectTimeout = options.ConnectTimeout
	client.Autoretry = options.Autoretry
	client.IgnoreSSL = options.IgnoreSSL
	client.MaxAttempts = options.MaxAttempts
}

// GetHost returns the host to call: Endpoint when it is set, defaultHost
// otherwise.
func (client *Client) GetHost(defaultHost string) *string {
	if client.Endpoint != nil {
		return client.Endpoint
	}
	return tea.String(defaultHost)
}

// transportOptions snapshots the client's runtime settings for one call.
func (client *Client) transportOptions() *transport.Options {
	return &transport.Options{
		ReadTimeout:    client.ReadTimeout,
		ConnectTimeout: client.ConnectTimeout,
		Autoretry:      client.Autoretry,
		IgnoreSSL:      client.IgnoreSSL,
		MaxAttempts:    client.MaxAttempts,
	}
}

// invoke validates request, performs the call described by req, and decodes the
// response into a freshly allocated T. It is the single entry point shared by
// every generated API method.
func invoke[T any](client *Client, request interface{}, req *transport.Request) (*T, error) {
	if err := tea.Validate(request); err != nil {
		return nil, err
	}

	var result *T
	err := transport.Do(client.transportOptions(), req, func(body map[string]interface{}) error {
		result = new(T)
		return tea.Convert(tea.ToMap(body), &result)
	})
	return result, err
}

// CommonOpenAPI calls an arbitrary endpoint, for APIs the SDK does not yet
// expose as a generated method. The response body is returned undecoded and is
// not checked for a business error code.
func (client *Client) CommonOpenAPI(request *CommonRequest) (_result map[string]interface{}, _err error) {
	if _err = tea.Validate(request); _err != nil {
		return _result, _err
	}

	req := &transport.Request{
		Method:            tea.StringValue(request.Method),
		Path:              tea.StringValue(request.Path),
		Host:              request.Host,
		Headers:           request.Header,
		RawQuery:          request.Query,
		SkipBusinessError: true,
	}
	// The caller picks the encoding through the content-type they supply.
	if tea.BoolValue(util.EqualString(request.Header["content-type"], tea.String(transport.ContentTypeJSON))) {
		req.Body = request.Body
		req.Encoding = transport.EncodingJSON
	}

	_err = transport.Do(client.transportOptions(), req, func(body map[string]interface{}) error {
		_result = make(map[string]interface{})
		return tea.Convert(tea.ToMap(body), &_result)
	})
	return _result, _err
}

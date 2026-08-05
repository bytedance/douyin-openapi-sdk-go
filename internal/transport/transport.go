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

// Package transport carries out a single OpenAPI call: it builds the HTTP
// request, retries it according to the caller's runtime options, checks the
// transport-level and business-level error codes, and hands the decoded JSON
// body back to the caller.
//
// Every generated API method in the SDK funnels through Do; the method itself
// only supplies the values that actually differ between endpoints.
package transport

import (
	"github.com/alibabacloud-go/tea/tea"
	util "github.com/bytedance/douyin-openapi-util-go/client"
)

// Encoding selects how Request.Body is serialised onto the wire.
type Encoding int

const (
	// EncodingNone sends no request body.
	EncodingNone Encoding = iota
	// EncodingJSON sends the body as a JSON document.
	EncodingJSON
	// EncodingForm sends the body as application/x-www-form-urlencoded.
	EncodingForm
	// EncodingFileForm sends the body as a multipart/form-data upload. The
	// multipart boundary is generated per attempt and written into the
	// content-type header, overriding Request.ContentType.
	EncodingFileForm
)

// Common content types used by the generated operations.
const (
	ContentTypeJSON = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"
)

// Options are the per-client runtime settings that govern timeouts and retries.
type Options struct {
	ReadTimeout    *int
	ConnectTimeout *int
	Autoretry      *bool
	IgnoreSSL      *bool
	MaxAttempts    *int
}

// runtime renders the options into the map shape the tea runtime expects.
func (o *Options) runtime() map[string]interface{} {
	return map[string]interface{}{
		"timeouted":      "retry",
		"readTimeout":    tea.IntValue(o.ReadTimeout),
		"connectTimeout": tea.IntValue(o.ConnectTimeout),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(o.Autoretry),
			"maxAttempts": tea.IntValue(o.MaxAttempts),
		},
		"ignoreSSL": tea.BoolValue(o.IgnoreSSL),
	}
}

// Request describes one OpenAPI call.
type Request struct {
	// Method is the HTTP verb, e.g. "GET" or "POST".
	Method string
	// Path is the request path, e.g. "/js/getticket/".
	Path string
	// Host is the resolved hostname to send to.
	Host *string

	// Headers seeds the request headers. It is used directly rather than
	// copied, matching the behaviour of the generated code: the SDK writes
	// its own headers into the caller's map.
	Headers map[string]*string
	// ContentType, when non-empty, sets the content-type header. Callers that
	// leave it empty keep whatever the caller passed in Headers.
	ContentType string

	// TokenHeader names the header carrying Token, e.g. "access-token".
	// The header is only set when both TokenHeader and Token are non-empty.
	TokenHeader string
	Token       *string

	// Query is stringified into query parameters. RawQuery, when set, is used
	// verbatim instead and takes precedence.
	Query    map[string]interface{}
	RawQuery map[string]*string

	// Body is serialised according to Encoding. EncodingForm and
	// EncodingFileForm require a map[string]interface{}; anything else is
	// treated as an empty body by those encodings.
	Body     interface{}
	Encoding Encoding

	// SkipBusinessError skips the err_no/err_msg check on the response body,
	// for endpoints that do not follow that convention.
	SkipBusinessError bool
}

// build renders the Request into a tea request. A fresh multipart boundary is
// generated on each call, so retries do not reuse a consumed body.
func (r *Request) build() *tea.Request {
	out := tea.NewRequest()
	if !tea.BoolValue(util.IsUnset(r.Headers)) {
		out.Headers = r.Headers
	}

	out.Protocol = tea.String("HTTPS")
	out.Method = tea.String(r.Method)
	out.Pathname = tea.String(r.Path)
	out.Headers["host"] = r.Host
	if r.ContentType != "" {
		out.Headers["content-type"] = tea.String(r.ContentType)
	}

	switch r.Encoding {
	case EncodingJSON:
		out.Body = tea.ToReader(util.ToJSONString(r.Body))
	case EncodingForm:
		out.Body = tea.ToReader(util.ToFormString(bodyMap(r.Body)))
	case EncodingFileForm:
		boundary := util.GetBoundary()
		out.Headers["content-type"] = tea.String("multipart/form-data; boundary=" + tea.StringValue(boundary))
		out.Body = util.ToFileForm(bodyMap(r.Body), boundary)
	}

	if r.RawQuery != nil {
		out.Query = r.RawQuery
	} else if r.Query != nil {
		out.Query = util.StringifyMapValue(r.Query)
	}

	out.Headers["x-from"] = tea.String("openapi-sdk")
	if r.TokenHeader != "" && !tea.BoolValue(util.Empty(r.Token)) {
		out.Headers[r.TokenHeader] = r.Token
	}
	return out
}

// bodyMap adapts Body for the encodings that need keyed fields.
func bodyMap(body interface{}) map[string]interface{} {
	m, _ := body.(map[string]interface{})
	return m
}

// Do sends req, retrying as the options allow, and passes the decoded response
// body to decode. decode is called at most once, only after the response has
// passed both the HTTP status check and the business error check.
func Do(opts *Options, req *Request, decode func(body map[string]interface{}) error) error {
	runtime := opts.runtime()

	var err error
	for attempt := 0; tea.BoolValue(tea.AllowRetry(runtime["retry"], tea.Int(attempt))); attempt++ {
		if attempt > 0 {
			backoff := tea.GetBackoffTime(runtime["backoff"], tea.Int(attempt))
			if tea.IntValue(backoff) > 0 {
				tea.Sleep(backoff)
			}
		}

		err = send(req, runtime, decode)
		if !tea.BoolValue(tea.Retryable(err)) {
			break
		}
	}
	return err
}

// send performs a single attempt.
func send(req *Request, runtime map[string]interface{}, decode func(map[string]interface{}) error) error {
	response, err := tea.DoRequest(req.build(), runtime)
	if err != nil {
		return err
	}

	if !tea.BoolValue(util.EqualNumber(response.StatusCode, tea.Int(200))) {
		return tea.NewSDKError(map[string]interface{}{
			// Spelling preserved from the generated SDK for compatibility with
			// callers that match on the message.
			"message": "Reqeust Failed!",
			"code":    tea.ToString(tea.IntValue(response.StatusCode)),
		})
	}

	bodyStr, err := util.ReadAsString(response.Body)
	if err != nil {
		return err
	}

	if !req.SkipBusinessError {
		if err := checkBusinessError(bodyStr); err != nil {
			return err
		}
	}

	body, err := util.AssertAsMap(util.ParseJSON(bodyStr))
	if err != nil {
		return err
	}
	return decode(body)
}

// checkBusinessError reports the err_no/err_msg pair carried in the response
// body as an SDK error when err_no is non-zero.
func checkBusinessError(bodyStr *string) error {
	resErr := util.GetErrMessage(bodyStr)
	code, err := util.AssertAsNumber(resErr["err_no"])
	if err != nil {
		return err
	}
	if !tea.BoolValue(util.EqualNumber(code, tea.Int(0))) {
		return tea.NewSDKError(map[string]interface{}{
			"message": resErr["err_msg"],
			"code":    tea.ToString(tea.IntValue(code)),
		})
	}
	return nil
}

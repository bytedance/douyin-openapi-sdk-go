package transport

import (
	"io"
	"strings"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	util "github.com/bytedance/douyin-openapi-util-go/client"
)

func testOptions() *Options {
	return &Options{
		ReadTimeout:    tea.Int(5000),
		ConnectTimeout: tea.Int(1000),
		Autoretry:      tea.Bool(false),
		IgnoreSSL:      tea.Bool(true),
		MaxAttempts:    tea.Int(3),
	}
}

func readAll(t *testing.T, r io.Reader) string {
	t.Helper()
	if r == nil {
		return ""
	}
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	return string(b)
}

func header(t *testing.T, req *tea.Request, key string) string {
	t.Helper()
	v, ok := req.Headers[key]
	if !ok {
		t.Fatalf("header %q not set", key)
	}
	return tea.StringValue(v)
}

func TestOptionsRuntime(t *testing.T) {
	runtime := testOptions().runtime()

	if runtime["readTimeout"] != 5000 {
		t.Errorf("readTimeout = %v, want 5000", runtime["readTimeout"])
	}
	if runtime["connectTimeout"] != 1000 {
		t.Errorf("connectTimeout = %v, want 1000", runtime["connectTimeout"])
	}
	if runtime["ignoreSSL"] != true {
		t.Errorf("ignoreSSL = %v, want true", runtime["ignoreSSL"])
	}
	if runtime["timeouted"] != "retry" {
		t.Errorf("timeouted = %v, want retry", runtime["timeouted"])
	}

	retry, ok := runtime["retry"].(map[string]interface{})
	if !ok {
		t.Fatalf("retry = %T, want map", runtime["retry"])
	}
	if retry["retryable"] != false {
		t.Errorf("retryable = %v, want false", retry["retryable"])
	}
	if retry["maxAttempts"] != 3 {
		t.Errorf("maxAttempts = %v, want 3", retry["maxAttempts"])
	}

	// The runtime must not carry a backoff policy: the SDK relies on
	// GetBackoffTime returning 0 for an absent one.
	if _, ok := runtime["backoff"]; ok {
		t.Error("runtime unexpectedly carries a backoff policy")
	}
}

func TestBuildSetsBaseFields(t *testing.T) {
	req := &Request{
		Method:      "GET",
		Path:        "/js/getticket/",
		Host:        tea.String("open.douyin.com"),
		ContentType: ContentTypeJSON,
	}

	out := req.build()

	if tea.StringValue(out.Protocol) != "HTTPS" {
		t.Errorf("protocol = %q, want HTTPS", tea.StringValue(out.Protocol))
	}
	if tea.StringValue(out.Method) != "GET" {
		t.Errorf("method = %q, want GET", tea.StringValue(out.Method))
	}
	if tea.StringValue(out.Pathname) != "/js/getticket/" {
		t.Errorf("pathname = %q", tea.StringValue(out.Pathname))
	}
	if got := header(t, out, "host"); got != "open.douyin.com" {
		t.Errorf("host = %q", got)
	}
	if got := header(t, out, "content-type"); got != ContentTypeJSON {
		t.Errorf("content-type = %q", got)
	}
	// Every SDK request identifies itself to the gateway.
	if got := header(t, out, "x-from"); got != "openapi-sdk" {
		t.Errorf("x-from = %q, want openapi-sdk", got)
	}
}

func TestBuildOmitsContentTypeWhenUnset(t *testing.T) {
	req := &Request{
		Method:  "GET",
		Path:    "/p/",
		Host:    tea.String("h"),
		Headers: map[string]*string{"content-type": tea.String("text/plain")},
	}

	out := req.build()

	// Endpoints that declare no content type must leave the caller's alone.
	if got := header(t, out, "content-type"); got != "text/plain" {
		t.Errorf("content-type = %q, want the caller's text/plain", got)
	}
}

func TestBuildUsesCallerHeaders(t *testing.T) {
	headers := map[string]*string{"x-custom": tea.String("v")}
	req := &Request{Method: "GET", Path: "/p/", Host: tea.String("h"), Headers: headers}

	out := req.build()

	if got := header(t, out, "x-custom"); got != "v" {
		t.Errorf("x-custom = %q, want v", got)
	}
	// The generated SDK writes its headers into the caller's map rather than a
	// copy; callers may observe that, so it is pinned here deliberately.
	if _, ok := headers["x-from"]; !ok {
		t.Error("expected build to write into the caller's header map")
	}
}

func TestBuildTokenHeader(t *testing.T) {
	tests := []struct {
		name        string
		tokenHeader string
		token       *string
		wantKey     string
		wantSet     bool
	}{
		{"access token", "access-token", tea.String("tok"), "access-token", true},
		{"webcast token", "x-token", tea.String("tok"), "x-token", true},
		{"empty token", "access-token", tea.String(""), "access-token", false},
		{"nil token", "access-token", nil, "access-token", false},
		{"endpoint takes no token", "", tea.String("tok"), "access-token", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &Request{
				Method:      "GET",
				Path:        "/p/",
				Host:        tea.String("h"),
				TokenHeader: tt.tokenHeader,
				Token:       tt.token,
			}

			out := req.build()

			got, ok := out.Headers[tt.wantKey]
			if ok != tt.wantSet {
				t.Fatalf("header %q present = %v, want %v", tt.wantKey, ok, tt.wantSet)
			}
			if tt.wantSet && tea.StringValue(got) != tea.StringValue(tt.token) {
				t.Errorf("token = %q, want %q", tea.StringValue(got), tea.StringValue(tt.token))
			}
		})
	}
}

func TestBuildQuery(t *testing.T) {
	req := &Request{
		Method: "GET",
		Path:   "/p/",
		Host:   tea.String("h"),
		Query: map[string]interface{}{
			"open_id": "abc",
			"count":   int64(10),
		},
	}

	out := req.build()

	if got := tea.StringValue(out.Query["open_id"]); got != "abc" {
		t.Errorf("open_id = %q, want abc", got)
	}
	// Non-string query values are stringified for the wire.
	if got := tea.StringValue(out.Query["count"]); got != "10" {
		t.Errorf("count = %q, want 10", got)
	}
}

func TestBuildRawQueryWins(t *testing.T) {
	req := &Request{
		Method:   "GET",
		Path:     "/p/",
		Host:     tea.String("h"),
		Query:    map[string]interface{}{"a": "1"},
		RawQuery: map[string]*string{"b": tea.String("2")},
	}

	out := req.build()

	if _, ok := out.Query["a"]; ok {
		t.Error("RawQuery should replace Query entirely")
	}
	if got := tea.StringValue(out.Query["b"]); got != "2" {
		t.Errorf("b = %q, want 2", got)
	}
}

func TestBuildBodyEncodings(t *testing.T) {
	body := map[string]interface{}{"code": "abc"}

	t.Run("none", func(t *testing.T) {
		out := (&Request{Method: "GET", Path: "/p/", Host: tea.String("h"), Body: body}).build()
		if out.Body != nil {
			t.Errorf("body = %q, want no body", readAll(t, out.Body))
		}
	})

	t.Run("json", func(t *testing.T) {
		out := (&Request{Method: "POST", Path: "/p/", Host: tea.String("h"),
			Body: body, Encoding: EncodingJSON}).build()
		if got := readAll(t, out.Body); got != `{"code":"abc"}` {
			t.Errorf("body = %q", got)
		}
	})

	t.Run("form", func(t *testing.T) {
		out := (&Request{Method: "POST", Path: "/p/", Host: tea.String("h"),
			Body: body, Encoding: EncodingForm}).build()
		if got := readAll(t, out.Body); got != "code=abc" {
			t.Errorf("body = %q, want code=abc", got)
		}
	})
}

func TestBuildFileFormSetsBoundary(t *testing.T) {
	newReq := func() *Request {
		return &Request{
			Method:      "POST",
			Path:        "/tool/imagex/client_upload/",
			Host:        tea.String("h"),
			ContentType: ContentTypeJSON,
			Body: map[string]interface{}{
				"image": &util.FileField{
					Filename:    tea.String("a.png"),
					ContentType: tea.String("image/png"),
					Content:     strings.NewReader("PNGDATA"),
				},
			},
			Encoding: EncodingFileForm,
		}
	}

	req := newReq()
	first := req.build()

	ct := header(t, first, "content-type")
	if !strings.HasPrefix(ct, "multipart/form-data; boundary=") {
		t.Fatalf("content-type = %q, want a multipart type overriding ContentType", ct)
	}

	// Each attempt needs its own boundary, so a retry cannot reuse a boundary
	// that the previous attempt already committed to the wire.
	second := newReq().build()
	if ct2 := header(t, second, "content-type"); ct2 == ct {
		t.Error("expected a fresh boundary per build")
	}

	boundary := strings.TrimPrefix(ct, "multipart/form-data; boundary=")
	got := readAll(t, first.Body)
	if !strings.Contains(got, boundary) {
		t.Errorf("body does not carry its own boundary %q: %q", boundary, got)
	}
	if !strings.Contains(got, `filename="a.png"`) || !strings.Contains(got, "PNGDATA") {
		t.Errorf("body is missing the uploaded file: %q", got)
	}
}

func TestCheckBusinessError(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantErr  bool
		wantCode string
		wantMsg  string
	}{
		{name: "success", body: `{"err_no":0,"err_msg":"success"}`},
		{name: "no error field", body: `{"data":{"ticket":"t"}}`},
		{
			name:     "err_no",
			body:     `{"err_no":40001,"err_msg":"invalid token","log_id":"lg1"}`,
			wantErr:  true,
			wantCode: "40001",
			wantMsg:  "msg=invalid token, logid=lg1",
		},
		{
			// Douyin reports failures under several different key names; the
			// SDK must surface all of them.
			name:     "extra.error_code",
			body:     `{"extra":{"error_code":2190008,"description":"no permission","logid":"lg2"}}`,
			wantErr:  true,
			wantCode: "2190008",
			wantMsg:  "msg=no permission, logid=lg2",
		},
		{
			name:     "nested err",
			body:     `{"err":{"err_code":50001,"err_msg":"internal"}}`,
			wantErr:  true,
			wantCode: "50001",
			wantMsg:  "msg=internal, logid=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkBusinessError(tea.String(tt.body))
			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			sdkErr, ok := err.(*tea.SDKError)
			if !ok {
				t.Fatalf("error = %T, want *tea.SDKError", err)
			}
			if got := tea.StringValue(sdkErr.Code); got != tt.wantCode {
				t.Errorf("code = %q, want %q", got, tt.wantCode)
			}
			if got := tea.StringValue(sdkErr.Message); got != tt.wantMsg {
				t.Errorf("message = %q, want %q", got, tt.wantMsg)
			}
		})
	}
}

func TestCheckBusinessErrorRejectsUnparsableBody(t *testing.T) {
	// A body the error decoder cannot read yields no err_no at all. The SDK
	// reports that as an error rather than treating it as a success.
	err := checkBusinessError(tea.String("not json"))
	if err == nil {
		t.Fatal("expected an error for an unparsable body")
	}
	if _, ok := err.(*tea.SDKError); ok {
		t.Errorf("error = %v, want a plain assertion error, not an SDKError", err)
	}
}

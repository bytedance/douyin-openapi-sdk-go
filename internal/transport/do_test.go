package transport

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
)

func TestMain(m *testing.M) {
	// tea honours the proxy environment variables but not NO_PROXY, so a
	// developer's local proxy would swallow these loopback requests.
	for _, key := range []string{"HTTP_PROXY", "http_proxy", "HTTPS_PROXY", "https_proxy"} {
		os.Unsetenv(key)
	}
	os.Exit(m.Run())
}

// newServer starts a TLS test server and returns the host:port to point a
// Request at. Requests reach it because the SDK's IgnoreSSL option is on.
func newServer(t *testing.T, handler http.HandlerFunc) *string {
	t.Helper()
	srv := httptest.NewTLSServer(handler)
	t.Cleanup(srv.Close)
	return tea.String(strings.TrimPrefix(srv.URL, "https://"))
}

// closeConnection drops the connection without answering, which surfaces to the
// caller as a network error rather than an HTTP status.
func closeConnection(t *testing.T, w http.ResponseWriter) {
	t.Helper()
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		t.Errorf("hijack: %v", err)
		return
	}
	conn.Close()
}

func collect(body *map[string]interface{}) func(map[string]interface{}) error {
	return func(b map[string]interface{}) error {
		*body = b
		return nil
	}
}

func TestDoSuccess(t *testing.T) {
	var gotPath, gotMethod, gotQuery, gotToken, gotFrom, gotBody string
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath, gotMethod = r.URL.Path, r.Method
		gotQuery = r.URL.RawQuery
		gotToken = r.Header.Get("access-token")
		gotFrom = r.Header.Get("x-from")
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		io.WriteString(w, `{"data":{"ticket":"tk","err_no":0}}`)
	})

	var body map[string]interface{}
	err := Do(testOptions(), &Request{
		Method:      "POST",
		Path:        "/js/getticket/",
		Host:        host,
		ContentType: ContentTypeJSON,
		Query:       map[string]interface{}{"open_id": "oid"},
		Body:        map[string]interface{}{"k": "v"},
		Encoding:    EncodingJSON,
		TokenHeader: "access-token",
		Token:       tea.String("tok"),
	}, collect(&body))
	if err != nil {
		t.Fatalf("Do: %v", err)
	}

	if gotMethod != "POST" || gotPath != "/js/getticket/" {
		t.Errorf("got %s %s, want POST /js/getticket/", gotMethod, gotPath)
	}
	if gotQuery != "open_id=oid" {
		t.Errorf("query = %q, want open_id=oid", gotQuery)
	}
	if gotToken != "tok" {
		t.Errorf("access-token = %q, want tok", gotToken)
	}
	if gotFrom != "openapi-sdk" {
		t.Errorf("x-from = %q, want openapi-sdk", gotFrom)
	}
	if gotBody != `{"k":"v"}` {
		t.Errorf("body = %q", gotBody)
	}

	data, ok := body["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("data = %T, want map", body["data"])
	}
	if data["ticket"] != "tk" {
		t.Errorf("ticket = %v, want tk", data["ticket"])
	}
}

func TestDoHTTPStatusError(t *testing.T) {
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, `{"err_no":0}`)
	})

	called := false
	err := Do(testOptions(), &Request{Method: "GET", Path: "/p/", Host: host},
		func(map[string]interface{}) error { called = true; return nil })

	sdkErr, ok := err.(*tea.SDKError)
	if !ok {
		t.Fatalf("error = %v (%T), want *tea.SDKError", err, err)
	}
	if got := tea.StringValue(sdkErr.Code); got != "403" {
		t.Errorf("code = %q, want 403", got)
	}
	if called {
		t.Error("decode must not run for a non-200 response")
	}
}

func TestDoBusinessError(t *testing.T) {
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"err_no":40001,"err_msg":"bad token","log_id":"lg"}`)
	})

	called := false
	err := Do(testOptions(), &Request{Method: "GET", Path: "/p/", Host: host},
		func(map[string]interface{}) error { called = true; return nil })

	sdkErr, ok := err.(*tea.SDKError)
	if !ok {
		t.Fatalf("error = %v (%T), want *tea.SDKError", err, err)
	}
	if got := tea.StringValue(sdkErr.Code); got != "40001" {
		t.Errorf("code = %q, want 40001", got)
	}
	if called {
		t.Error("decode must not run when the response carries a business error")
	}
}

func TestDoSkipBusinessError(t *testing.T) {
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"err_no":40001,"err_msg":"bad token"}`)
	})

	var body map[string]interface{}
	err := Do(testOptions(), &Request{
		Method: "GET", Path: "/p/", Host: host, SkipBusinessError: true,
	}, collect(&body))
	if err != nil {
		t.Fatalf("Do: %v", err)
	}

	// CommonOpenAPI hands the raw body back and lets the caller judge it.
	// Numbers arrive as json.Number, not float64.
	if body["err_no"] != json.Number("40001") {
		t.Errorf("err_no = %v (%T), want the raw body passed through", body["err_no"], body["err_no"])
	}
}

func TestDoRawQueryReachesServer(t *testing.T) {
	var gotQuery string
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		io.WriteString(w, `{"err_no":0}`)
	})

	err := Do(testOptions(), &Request{
		Method:   "GET",
		Path:     "/p/",
		Host:     host,
		RawQuery: map[string]*string{"a": tea.String("1")},
	}, func(map[string]interface{}) error { return nil })
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if gotQuery != "a=1" {
		t.Errorf("query = %q, want a=1", gotQuery)
	}
}

func TestDoNoRetryWhenDisabled(t *testing.T) {
	var attempts int32
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		closeConnection(t, w)
	})

	opts := testOptions() // Autoretry is off.
	err := Do(opts, &Request{Method: "POST", Path: "/p/", Host: host,
		Body: map[string]interface{}{"k": "v"}, Encoding: EncodingJSON},
		func(map[string]interface{}) error { return nil })

	if err == nil {
		t.Fatal("expected a network error")
	}
	if got := atomic.LoadInt32(&attempts); got != 1 {
		t.Errorf("attempts = %d, want 1 with retries disabled", got)
	}
}

func TestDoRetriesUntilSuccess(t *testing.T) {
	var attempts int32
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&attempts, 1) < 3 {
			closeConnection(t, w)
			return
		}
		io.WriteString(w, `{"data":{"ticket":"tk"},"err_no":0}`)
	})

	opts := testOptions()
	opts.Autoretry = tea.Bool(true)
	opts.MaxAttempts = tea.Int(3)

	var body map[string]interface{}
	err := Do(opts, &Request{Method: "POST", Path: "/p/", Host: host,
		Body: map[string]interface{}{"k": "v"}, Encoding: EncodingJSON}, collect(&body))
	if err != nil {
		t.Fatalf("Do: %v", err)
	}

	if got := atomic.LoadInt32(&attempts); got != 3 {
		t.Errorf("attempts = %d, want 3", got)
	}
	if body["data"] == nil {
		t.Error("expected the successful attempt's body to be decoded")
	}
}

func TestDoStopsRetryingAtMaxAttempts(t *testing.T) {
	var attempts int32
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		closeConnection(t, w)
	})

	opts := testOptions()
	opts.Autoretry = tea.Bool(true)
	opts.MaxAttempts = tea.Int(2)

	err := Do(opts, &Request{Method: "POST", Path: "/p/", Host: host,
		Body: map[string]interface{}{"k": "v"}, Encoding: EncodingJSON},
		func(map[string]interface{}) error { return nil })

	if err == nil {
		t.Fatal("expected the final network error to be returned")
	}
	// tea allows the initial attempt plus MaxAttempts retries.
	if got := atomic.LoadInt32(&attempts); got != 3 {
		t.Errorf("attempts = %d, want 3 (1 initial + 2 retries)", got)
	}
}

func TestDoDoesNotRetryStatusErrors(t *testing.T) {
	var attempts int32
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusInternalServerError)
	})

	opts := testOptions()
	opts.Autoretry = tea.Bool(true)

	err := Do(opts, &Request{Method: "GET", Path: "/p/", Host: host},
		func(map[string]interface{}) error { return nil })

	if err == nil {
		t.Fatal("expected an error")
	}
	// The SDK's own status error carries no StatusCode, so tea treats it as
	// final. Preserved from the generated code.
	if got := atomic.LoadInt32(&attempts); got != 1 {
		t.Errorf("attempts = %d, want 1", got)
	}
}

func TestDoPropagatesDecodeError(t *testing.T) {
	host := newServer(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"err_no":0}`)
	})

	want := io.ErrUnexpectedEOF
	err := Do(testOptions(), &Request{Method: "GET", Path: "/p/", Host: host},
		func(map[string]interface{}) error { return want })
	if err != want {
		t.Errorf("error = %v, want %v", err, want)
	}
}

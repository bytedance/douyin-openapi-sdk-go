package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/bytedance/douyin-openapi-credential-go/client"
)

func TestMain(m *testing.M) {
	// tea honours the proxy environment variables but not NO_PROXY, so a
	// developer's local proxy would swallow these loopback requests.
	for _, key := range []string{"HTTP_PROXY", "http_proxy", "HTTPS_PROXY", "https_proxy"} {
		os.Unsetenv(key)
	}
	os.Exit(m.Run())
}

// newTestClient returns a client whose requests are routed to a local TLS test
// server via the Endpoint override.
func newTestClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewTLSServer(handler)
	t.Cleanup(srv.Close)

	client, err := NewClient(&credential.Config{
		ClientKey:    tea.String("test-client-key"),
		ClientSecret: tea.String("test-client-secret"),
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	client.Endpoint = tea.String(strings.TrimPrefix(srv.URL, "https://"))
	return client
}

func TestNewClient(t *testing.T) {
	config := &credential.Config{
		ClientKey:    tea.String("test-client-key"),
		ClientSecret: tea.String("test-client-secret"),
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Fatalf("Expected client to be created, got nil")
	}

	if client.Credential == nil {
		t.Error("Expected credential to be initialized")
	}

	if tea.IntValue(client.ReadTimeout) != 5000 {
		t.Errorf("Expected ReadTimeout 5000, got %d", tea.IntValue(client.ReadTimeout))
	}
	if tea.IntValue(client.ConnectTimeout) != 1000 {
		t.Errorf("Expected ConnectTimeout 1000, got %d", tea.IntValue(client.ConnectTimeout))
	}
	if tea.IntValue(client.MaxAttempts) != 3 {
		t.Errorf("Expected MaxAttempts 3, got %d", tea.IntValue(client.MaxAttempts))
	}
	if !tea.BoolValue(client.Autoretry) {
		t.Error("Expected Autoretry to default to true")
	}
	if !tea.BoolValue(client.IgnoreSSL) {
		t.Error("Expected IgnoreSSL to default to true")
	}
}

func TestGetHost(t *testing.T) {
	client := &Client{}

	// Test default host fallback
	host := client.GetHost("open.douyin.com")
	if host == nil || *host != "open.douyin.com" {
		t.Errorf("Expected 'open.douyin.com', got %v", host)
	}

	// Test overridden host
	client.Endpoint = tea.String("api.example.com")
	host2 := client.GetHost("open.douyin.com")
	if host2 == nil || *host2 != "api.example.com" {
		t.Errorf("Expected 'api.example.com', got %v", host2)
	}
}

func TestSetRuntimeOptions(t *testing.T) {
	client := &Client{}
	client.SetRuntimeOptions(&RuntimeOptions{
		Autoretry:      tea.Bool(false),
		IgnoreSSL:      tea.Bool(false),
		MaxAttempts:    tea.Int(7),
		ReadTimeout:    tea.Int(1234),
		ConnectTimeout: tea.Int(567),
	})

	if tea.BoolValue(client.Autoretry) {
		t.Error("Autoretry not applied")
	}
	if tea.BoolValue(client.IgnoreSSL) {
		t.Error("IgnoreSSL not applied")
	}
	if tea.IntValue(client.MaxAttempts) != 7 {
		t.Errorf("MaxAttempts = %d, want 7", tea.IntValue(client.MaxAttempts))
	}
	if tea.IntValue(client.ReadTimeout) != 1234 {
		t.Errorf("ReadTimeout = %d, want 1234", tea.IntValue(client.ReadTimeout))
	}
	if tea.IntValue(client.ConnectTimeout) != 567 {
		t.Errorf("ConnectTimeout = %d, want 567", tea.IntValue(client.ConnectTimeout))
	}
}

func TestSetCredential(t *testing.T) {
	cred, err := credential.NewCredential(&credential.Config{
		ClientKey:    tea.String("k"),
		ClientSecret: tea.String("s"),
	})
	if err != nil {
		t.Fatalf("NewCredential: %v", err)
	}

	client := &Client{}
	client.SetCredential(cred)
	if client.Credential != cred {
		t.Error("SetCredential did not store the credential")
	}
}

// TestGeneratedGetMethod exercises a generated GET endpoint end to end:
// query-less request, access-token header, typed response decoding.
func TestGeneratedGetMethod(t *testing.T) {
	var gotPath, gotToken string
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotToken = r.Header.Get("access-token")
		io.WriteString(w, `{"data":{"ticket":"tk","expires_in":7200}}`)
	})

	resp, err := client.JsGetticket(&JsGetticketRequest{
		AccessToken: tea.String("tok"),
	})
	if err != nil {
		t.Fatalf("JsGetticket: %v", err)
	}

	if gotPath != "/js/getticket/" {
		t.Errorf("path = %q, want /js/getticket/", gotPath)
	}
	if gotToken != "tok" {
		t.Errorf("access-token = %q, want tok", gotToken)
	}
	if resp.Data == nil {
		t.Fatal("response data not decoded")
	}
	if tea.StringValue(resp.Data.Ticket) != "tk" {
		t.Errorf("ticket = %q, want tk", tea.StringValue(resp.Data.Ticket))
	}
	if tea.Int64Value(resp.Data.ExpiresIn) != 7200 {
		t.Errorf("expires_in = %d, want 7200", tea.Int64Value(resp.Data.ExpiresIn))
	}
}

// TestGeneratedFormMethod exercises a generated POST endpoint that sends a
// form-encoded body and takes no access token.
func TestGeneratedFormMethod(t *testing.T) {
	var gotBody, gotContentType string
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotContentType = r.Header.Get("content-type")
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		io.WriteString(w, `{"data":{"access_token":"at","expires_in":86400}}`)
	})

	resp, err := client.OauthAccessToken(&OauthAccessTokenRequest{
		Code:         tea.String("c"),
		GrantType:    tea.String("authorization_code"),
		ClientKey:    tea.String("k"),
		ClientSecret: tea.String("s"),
	})
	if err != nil {
		t.Fatalf("OauthAccessToken: %v", err)
	}

	if gotContentType != "application/x-www-form-urlencoded" {
		t.Errorf("content-type = %q", gotContentType)
	}
	for _, want := range []string{"code=c", "grant_type=authorization_code", "client_key=k", "client_secret=s"} {
		if !strings.Contains(gotBody, want) {
			t.Errorf("body %q missing %q", gotBody, want)
		}
	}
	if resp.Data == nil || tea.StringValue(resp.Data.AccessToken) != "at" {
		t.Errorf("access_token not decoded from %+v", resp.Data)
	}
}

// TestGeneratedMethodQueryAndBody covers an endpoint that sends both a query
// string and a JSON body.
func TestGeneratedMethodQueryAndBody(t *testing.T) {
	var gotQuery, gotBody string
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		io.WriteString(w, `{"data":{"error_code":0}}`)
	})

	_, err := client.AuditSet(&AuditSetRequest{
		OpenId:      tea.String("oid"),
		ApplyId:     tea.String("aid"),
		Status:      tea.Int64(1),
		AccessToken: tea.String("tok"),
	})
	if err != nil {
		t.Fatalf("AuditSet: %v", err)
	}

	if gotQuery != "open_id=oid" {
		t.Errorf("query = %q, want open_id=oid", gotQuery)
	}
	if !strings.Contains(gotBody, `"apply_id":"aid"`) || !strings.Contains(gotBody, `"status":1`) {
		t.Errorf("body = %q", gotBody)
	}
}

func TestGeneratedMethodValidationError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("no request should be sent when validation fails")
	})

	// Code, GrantType, ClientKey and ClientSecret are all required.
	resp, err := client.OauthAccessToken(&OauthAccessTokenRequest{})
	if err == nil {
		t.Fatal("expected a validation error")
	}
	if resp != nil {
		t.Errorf("expected a nil response on error, got %+v", resp)
	}
}

func TestGeneratedMethodBusinessError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"err_no":40001,"err_msg":"invalid token"}`)
	})

	resp, err := client.JsGetticket(&JsGetticketRequest{AccessToken: tea.String("bad")})
	if err == nil {
		t.Fatal("expected a business error")
	}
	if resp != nil {
		t.Errorf("expected a nil response on error, got %+v", resp)
	}

	sdkErr, ok := err.(*tea.SDKError)
	if !ok {
		t.Fatalf("error = %T, want *tea.SDKError", err)
	}
	if tea.StringValue(sdkErr.Code) != "40001" {
		t.Errorf("code = %q, want 40001", tea.StringValue(sdkErr.Code))
	}
}

func TestGeneratedMethodHTTPError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	})

	resp, err := client.JsGetticket(&JsGetticketRequest{AccessToken: tea.String("tok")})
	if err == nil {
		t.Fatal("expected an HTTP error")
	}
	if resp != nil {
		t.Errorf("expected a nil response on error, got %+v", resp)
	}
}

// TestPinnedHostIgnoresEndpoint documents that endpoints which hard-code their
// host are not redirected by Client.Endpoint.
func TestPinnedHostIgnoresEndpoint(t *testing.T) {
	client := &Client{Endpoint: tea.String("127.0.0.1:1")}

	// QuizGet pins webcast.bytedance.com; GetHost is only consulted by the
	// endpoints that do not pin a host.
	if got := tea.StringValue(client.GetHost("open.douyin.com")); got != "127.0.0.1:1" {
		t.Errorf("GetHost = %q, want the endpoint override", got)
	}
}

func TestCommonOpenAPI(t *testing.T) {
	var gotPath, gotMethod, gotQuery, gotBody, gotCustom string
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath, gotMethod = r.URL.Path, r.Method
		gotQuery = r.URL.RawQuery
		gotCustom = r.Header.Get("x-custom")
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		io.WriteString(w, `{"err_no":0,"data":{"value":"v"}}`)
	})

	result, err := client.CommonOpenAPI(&CommonRequest{
		Host:   client.Endpoint,
		Path:   tea.String("/some/new/api/"),
		Method: tea.String("POST"),
		Header: map[string]*string{
			"content-type": tea.String("application/json"),
			"x-custom":     tea.String("cv"),
		},
		Query: map[string]*string{"a": tea.String("1")},
		Body:  map[string]interface{}{"k": "v"},
	})
	if err != nil {
		t.Fatalf("CommonOpenAPI: %v", err)
	}

	if gotMethod != "POST" || gotPath != "/some/new/api/" {
		t.Errorf("got %s %s", gotMethod, gotPath)
	}
	if gotQuery != "a=1" {
		t.Errorf("query = %q, want a=1", gotQuery)
	}
	if gotCustom != "cv" {
		t.Errorf("x-custom = %q, want cv", gotCustom)
	}
	if gotBody != `{"k":"v"}` {
		t.Errorf("body = %q", gotBody)
	}
	if result["data"] == nil {
		t.Errorf("result not decoded: %+v", result)
	}
}

// TestCommonOpenAPISkipsBusinessError pins that CommonOpenAPI returns the raw
// body rather than converting err_no into an error.
func TestCommonOpenAPISkipsBusinessError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"err_no":40001,"err_msg":"nope"}`)
	})

	result, err := client.CommonOpenAPI(&CommonRequest{
		Host:   client.Endpoint,
		Path:   tea.String("/p/"),
		Method: tea.String("GET"),
		Header: map[string]*string{"content-type": tea.String("application/json")},
		Query:  map[string]*string{},
		Body:   map[string]interface{}{},
	})
	if err != nil {
		t.Fatalf("CommonOpenAPI: %v", err)
	}
	if result["err_no"] == nil {
		t.Errorf("expected the raw body, got %+v", result)
	}
}

func TestCommonOpenAPIValidationError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("no request should be sent when validation fails")
	})

	// Every CommonRequest field is required.
	result, err := client.CommonOpenAPI(&CommonRequest{})
	if err == nil {
		t.Fatal("expected a validation error")
	}
	if result != nil {
		t.Errorf("expected a nil result on error, got %+v", result)
	}
}

package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

const (
	apiKeyEnvKey  string = "WARREN_API_KEY"
	baseURLEnvKey string = "WARREN_API_BASE_URL"
)

// Default creates API where both BaseURL and APIKey comes from environment variables.
var Default *API = New(os.Getenv(baseURLEnvKey), os.Getenv(apiKeyEnvKey))

// ClientResponse is a data structured returned by `doRequest()`.
// To make the client compatible even when the server changed their response format.
// User of this library is responsible to handle the Body which is a slice of byte.
type ClientResponse struct {
	Error error
	Body  []byte
}

// API used to holds objects that are needed to make a HTTP call.
type API struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// FormRequest make a call with form-encoded payload
func (a *API) FormRequest(ctx context.Context, cfg RequestConfig) *ClientResponse {
	req, err := a.buildRequest(ctx, cfg)
	if err != nil {
		return &ClientResponse{Error: err}
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return a.doRequest(req)
}

// JsonRequest make a call with json-encoded payload
func (a *API) JSONRequest(ctx context.Context, cfg RequestConfig) *ClientResponse {
	req, err := a.buildRequest(ctx, cfg)
	if err != nil {
		return &ClientResponse{Error: err}
	}
	req.Header.Set("Content-Type", "application/json")
	return a.doRequest(req)
}

// buildRequest wraps `http.NewRequestWithContext` and set necessary header for authentication.
func (a *API) buildRequest(ctx context.Context, cfg RequestConfig) (*http.Request, error) {
	body, err := cfg.body()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(cfg.Method), cfg.url(a.BaseURL), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("apikey", a.APIKey)
	return req, nil
}

// doRequest doing the actual request
func (a *API) doRequest(req *http.Request) *ClientResponse {
	res, err := a.HTTPClient.Do(req)
	if err != nil {
		return &ClientResponse{Error: err}
	}
	defer res.Body.Close()

	// we'll only accept 2xx and 3xx as success
	if res.StatusCode >= 400 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return &ClientResponse{
				Error: fmt.Errorf("api call failed with status code=%d", res.StatusCode),
			}
		}
		return &ClientResponse{
			Body:  b,
			Error: fmt.Errorf("api call failed with status code=%d: %s", res.StatusCode, b),
		}
	}
	b, err := io.ReadAll(res.Body)
	return &ClientResponse{Body: b, Error: err}
}

// New create an instance of API
func New(baseURL, apiKey string) *API {
	return &API{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		HTTPClient: http.DefaultClient,
	}
}

// MockClientServer returns API client and test server to simplify API call testing
func MockClientServer(fn func(w http.ResponseWriter, r *http.Request)) (*API, *httptest.Server) {
	s := httptest.NewServer(http.HandlerFunc(fn))
	a := &API{
		APIKey:     "secret",
		BaseURL:    s.URL,
		HTTPClient: s.Client(),
	}
	return a, s
}

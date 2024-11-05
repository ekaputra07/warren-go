package http

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
	// ApiKeyEnvKey is the key used to get API Key from environment variable
	ApiKeyEnvKey string = "IDCLOUDHOST_API_KEY"

	// BaseUrl is the base URL of the API
	BaseUrl string = "https://api.idcloudhost.com"
)

// DefaultClient create Client with default configuration
var DefaultClient *Client = NewClient()

// ClientResponse is a data structured returned by `doRequest()`.
// To make the client compatible even when the server changed their response format.
// User of this library is responsible to handle the Body which is a slice of byte.
type ClientResponse struct {
	Error error
	Body  []byte
}

// Client used to holds objects that are needed to make a HTTP call.
type Client struct {
	ApiKey     string
	BaseUrl    string
	HTTPClient *http.Client
}

// SetApiKey manually set API Key field
func (c *Client) SetApiKey(key string) *Client {
	c.ApiKey = key
	return c
}

// FormRequest make a call with form-encoded payload
func (c *Client) FormRequest(ctx context.Context, cfg RequestConfig) *ClientResponse {
	req, err := c.buildRequest(ctx, cfg)
	if err != nil {
		return &ClientResponse{Error: err}
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.doRequest(req)
}

// JsonRequest make a call with json-encoded payload
func (c *Client) JsonRequest(ctx context.Context, cfg RequestConfig) *ClientResponse {
	req, err := c.buildRequest(ctx, cfg)
	if err != nil {
		return &ClientResponse{Error: err}
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

// buildRequest wraps `http.NewRequestWithContext` and set necessary header for authentication.
func (c *Client) buildRequest(ctx context.Context, cfg RequestConfig) (*http.Request, error) {
	body, err := cfg.body()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(cfg.Method), cfg.url(c.BaseUrl), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("apikey", c.ApiKey)
	return req, nil
}

// doRequest doing the actual request
func (c *Client) doRequest(req *http.Request) *ClientResponse {
	res, err := c.HTTPClient.Do(req)
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

// NewClient create an instance of Client
func NewClient() *Client {
	return &Client{
		ApiKey:     os.Getenv(ApiKeyEnvKey),
		BaseUrl:    BaseUrl,
		HTTPClient: http.DefaultClient,
	}
}

// MockClientServer returns client and test server to simplify API call testing
func MockClientServer(fn func(w http.ResponseWriter, r *http.Request)) (*Client, *httptest.Server) {
	s := httptest.NewServer(http.HandlerFunc(fn))
	c := &Client{
		ApiKey:     "secret",
		BaseUrl:    s.URL,
		HTTPClient: s.Client(),
	}
	return c, s
}

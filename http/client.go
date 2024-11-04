package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	// ApiKeyEnvKey is the key used to get API Key from environment variable
	ApiKeyEnvKey string = "IDCLOUDHOST_API_KEY"

	// BaseUrl is the base URL of the API
	BaseUrl string = "https://api.idcloudhost.com"
)

type RequestConfig struct {
	Method string
	Path   string
	Query  url.Values
	Data   url.Values
	Json   map[string]interface{}
}

// URL returns full request URL composed from baseURL, Path and Query field.
func (r RequestConfig) url(baseURL string) string {
	url := fmt.Sprintf("%s/%s", baseURL, strings.TrimLeft(r.Path, "/"))
	if r.Query == nil {
		return url
	}
	qs := r.Query.Encode()
	return fmt.Sprintf("%s?%s", url, qs)
}

// body returns io.Reader either from Data or Json field
func (r RequestConfig) body() (io.Reader, error) {
	if r.Data != nil && r.Json != nil {
		return nil, errors.New("data and json can not be set at the same time")
	}
	if r.Data != nil {
		return strings.NewReader(r.Data.Encode()), nil
	}
	if r.Json != nil {
		jsonStr, err := json.Marshal(r.Json)
		if err != nil {
			return nil, err
		}
		return strings.NewReader(string(jsonStr)), nil
	}

	// for request that don't have body
	return nil, nil
}

// ClientResponse is a data structured returned by the Client.Call() method.
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

// JsonRequest make a call json-encoded payload
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
	req, err := http.NewRequestWithContext(ctx, cfg.Method, cfg.url(c.BaseUrl), body)
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

// DefaultClient create Client with default configuration
var DefaultClient *Client = NewClient()

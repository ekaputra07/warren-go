package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	apiKeyEnvKey string = "IDCLOUDHOST_API_KEY"
	baseUrl      string = "https://api.idcloudhost.com"
)

// ApiClient used to holds objects that are needed to make an HTTP call.
type ApiClient struct {
	apiKey     string
	baseUrl    string
	httpClient *http.Client
}

// SetApiKey manually set API Key field
func (ac *ApiClient) SetApiKey(key string) *ApiClient {
	ac.apiKey = key
	return ac
}

// request wraps `http.NewRequestWithContext` and set necesarry header for authentication.
func (ac *ApiClient) request(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	url := buildUrl(ac.baseUrl, path)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("apikey", ac.apiKey)
	return req, nil
}

// maybeFormRequest returns normal or form request object depends on `data` value.
func (ac *ApiClient) maybeFormRequest(ctx context.Context, method string, path string, data url.Values) (*http.Request, error) {
	if data != nil {
		req, err := ac.request(ctx, method, path, strings.NewReader(data.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		return req, nil
	}
	return ac.request(ctx, method, path, nil)
}

// DoRequest builds Request object from given parameters and then make the actual request via `http.Client.Do()`.
// DoRequest is a generic method to make a HTTP request regardless of its method.
// For all method except GET: If data is provided (not nil), we'll do a form request
// For GET request: `data` is converted into query params and appended to the path
func (ac *ApiClient) DoRequest(ctx context.Context, method string, path string, data url.Values) (*http.Response, error) {
	switch method {
	case "GET":
		// if data is set, convert them to query params
		if data != nil {
			path = fmt.Sprintf("%s?%s", path, data.Encode())
		}
		req, err := ac.request(ctx, "GET", path, nil)
		if err != nil {
			return nil, err
		}
		return ac.httpClient.Do(req)
	default:
		req, err := ac.maybeFormRequest(ctx, method, path, data)
		if err != nil {
			return nil, err
		}
		return ac.httpClient.Do(req)
	}
}

func buildUrl(baseUrl, path string) string {
	return fmt.Sprintf("%s/%s", baseUrl, strings.TrimLeft(path, "/"))
}

// NewApiClient create an instance of ApiClient
func NewApiClient() *ApiClient {
	return &ApiClient{
		apiKey:     os.Getenv(apiKeyEnvKey),
		baseUrl:    baseUrl,
		httpClient: http.DefaultClient,
	}
}

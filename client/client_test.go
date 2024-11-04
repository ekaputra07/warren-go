package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUrl(t *testing.T) {
	url := buildUrl("https://example.com", "/some/path")
	assert.Equal(t, "https://example.com/some/path", url)

	url = buildUrl("https://example.com", "some/path")
	assert.Equal(t, "https://example.com/some/path", url)

	url = buildUrl("https://example.com", "///some/path")
	assert.Equal(t, "https://example.com/some/path", url)
}

func TestNewApiClient(t *testing.T) {
	// no api key in Env variable
	c := NewApiClient()
	assert.Equal(t, "", c.apiKey)
	assert.Equal(t, baseUrl, c.baseUrl)

	// api key set in Env variable
	os.Setenv(apiKeyEnvKey, "secret")
	c = NewApiClient()
	assert.Equal(t, "secret", c.apiKey)

	// api key manually set
	c = NewApiClient().SetApiKey("secret-new")
	assert.Equal(t, "secret-new", c.apiKey)
}

func TestDoRequestNilContext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	c := &ApiClient{
		apiKey:     "secret",
		baseUrl:    server.URL,
		httpClient: server.Client(),
	}
	_, err := c.DoRequest(nil, "GET", "/test", nil)
	assert.Error(t, err)
}

func TestDoRequestMethodInvalid(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	c := &ApiClient{
		apiKey:     "secret",
		baseUrl:    server.URL,
		httpClient: server.Client(),
	}
	_, err := c.DoRequest(context.Background(), "**BAD😀METHOD**", "/test", nil)
	assert.Error(t, err)
}

func TestDoRequestGET(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/test", r.RequestURI)
		assert.Equal(t, "secret", r.Header.Get("apikey"))
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	c := &ApiClient{
		apiKey:     "secret",
		baseUrl:    server.URL,
		httpClient: server.Client(),
	}
	resp, _ := c.DoRequest(context.Background(), "GET", "/test", nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDoRequestGETQueryParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/test?name=test", r.RequestURI)
		assert.Equal(t, "secret", r.Header.Get("apikey"))
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	c := &ApiClient{
		apiKey:     "secret",
		baseUrl:    server.URL,
		httpClient: server.Client(),
	}
	query := url.Values{}
	query.Add("name", "test")

	resp, _ := c.DoRequest(context.Background(), "GET", "/test", query)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDoRequestPOST(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/test", r.RequestURI)
		assert.Equal(t, "secret", r.Header.Get("apikey"))
		assert.Equal(t, "", r.Header.Get("Content-Type"))
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	c := &ApiClient{
		apiKey:     "secret",
		baseUrl:    server.URL,
		httpClient: server.Client(),
	}
	resp, _ := c.DoRequest(context.Background(), "POST", "/test", nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDoRequestPOSTForm(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/test", r.RequestURI)
		assert.Equal(t, "secret", r.Header.Get("apikey"))
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))

		// parse the posted form data
		_ = r.ParseForm()

		assert.Equal(t, "test", r.Form.Get("name"))
		assert.Equal(t, "20", r.Form.Get("age"))

		w.Write([]byte("OK"))
	}))
	defer server.Close()

	c := &ApiClient{
		apiKey:     "secret",
		baseUrl:    server.URL,
		httpClient: server.Client(),
	}
	body := url.Values{}
	body.Add("name", "test")
	body.Add("age", "20")
	resp, _ := c.DoRequest(context.Background(), "POST", "/test", body)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

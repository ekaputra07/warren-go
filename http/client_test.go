package http

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	os.Setenv(ApiKeyEnvKey, "")
	defer os.Setenv(ApiKeyEnvKey, "")

	// no api key in Env variable
	c := NewClient()
	assert.Equal(t, "", c.ApiKey)
	assert.Equal(t, BaseURL, c.BaseURL)

	// api key set in Env variable
	os.Setenv(ApiKeyEnvKey, "secret")
	c = NewClient()
	assert.Equal(t, "secret", c.ApiKey)

	// api key manually set
	c = NewClient().SetApiKey("secret-new")
	assert.Equal(t, "secret-new", c.ApiKey)
}

func TestFormRequest_NoContext(t *testing.T) {
	c, s := MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	defer s.Close()

	cfg := RequestConfig{
		Method: "GET",
		Path:   "/test",
	}
	resp := c.FormRequest(nil, cfg)
	assert.Error(t, resp.Error)
}

func TestFormRequest_MethodInvalid(t *testing.T) {
	c, s := MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	defer s.Close()

	cfg := RequestConfig{
		Method: "**BAD METHOD**",
		Path:   "/test",
	}
	resp := c.FormRequest(context.Background(), cfg)
	assert.Error(t, resp.Error)
}

func TestFormRequest_GET(t *testing.T) {
	c, s := MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/test", r.RequestURI)
		assert.Equal(t, "secret", r.Header.Get("apikey"))
		w.Write([]byte("OK"))
	})
	defer s.Close()

	cfg := RequestConfig{
		Method: "GET",
		Path:   "/test",
	}
	resp := c.FormRequest(context.Background(), cfg)
	assert.Equal(t, []byte("OK"), resp.Body)
}

func TestFormRequest_GET_QueryParams(t *testing.T) {
	c, s := MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/test?name=test", r.RequestURI)
		assert.Equal(t, "secret", r.Header.Get("apikey"))
		w.Write([]byte("OK"))
	})
	defer s.Close()

	q := url.Values{}
	q.Add("name", "test")

	cfg := RequestConfig{
		Method: "GET",
		Path:   "/test",
		Query:  q,
	}
	resp := c.FormRequest(context.Background(), cfg)
	assert.Equal(t, []byte("OK"), resp.Body)
}

func TestFormRequest(t *testing.T) {
	c, s := MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/test", r.RequestURI)
		assert.Equal(t, "secret", r.Header.Get("apikey"))
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))

		// parse the posted form data
		_ = r.ParseForm()

		assert.Equal(t, "test", r.Form.Get("name"))
		assert.Equal(t, "20", r.Form.Get("age"))

		w.Write([]byte("OK"))
	})
	defer s.Close()

	d := url.Values{}
	d.Add("name", "test")
	d.Add("age", "20")

	cfg := RequestConfig{
		Method: "POST",
		Path:   "/test",
		Data:   d,
	}

	resp := c.FormRequest(context.Background(), cfg)
	assert.Equal(t, []byte("OK"), resp.Body)
}

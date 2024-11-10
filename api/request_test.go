package api

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestConfig_url(t *testing.T) {
	q := url.Values{}
	q.Add("name", "test")

	// with slash prefix
	cfg := RequestConfig{
		Path:  "/some/path",
		Query: q,
	}
	assert.Equal(t, "https://example.com/some/path?name=test", cfg.url("https://example.com"))

	// with many slash prefix
	cfg = RequestConfig{
		Path:  "////some/path",
		Query: q,
	}
	assert.Equal(t, "https://example.com/some/path?name=test", cfg.url("https://example.com"))

	// no slash prefix
	cfg = RequestConfig{
		Path:  "some/path",
		Query: q,
	}
	assert.Equal(t, "https://example.com/some/path?name=test", cfg.url("https://example.com"))
}

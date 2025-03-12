package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
)

type RequestConfig struct {
	Method string
	Path   string
	Query  url.Values
	Data   url.Values
	JSON   map[string]any
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
	if r.Data != nil && r.JSON != nil {
		return nil, errors.New("data and json can not be set at the same time")
	}
	if r.Data != nil {
		return strings.NewReader(r.Data.Encode()), nil
	}
	if r.JSON != nil {
		jsonStr, err := json.Marshal(r.JSON)
		if err != nil {
			return nil, err
		}
		return strings.NewReader(string(jsonStr)), nil
	}

	// for request that don't have body
	return nil, nil
}

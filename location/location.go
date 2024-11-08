package location

import (
	"context"

	"github.com/ekaputra07/idcloudhost-go/http"
)

func NewClient() *Client {
	return &Client{
		H: http.DefaultClient,
	}
}

type Client struct {
	H *http.Client
}

func (c *Client) ListLocations(ctx context.Context) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/config/locations",
	}
	return c.H.FormRequest(ctx, rc)
}

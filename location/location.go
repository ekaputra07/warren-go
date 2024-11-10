package location

import (
	"context"

	"github.com/ekaputra07/warren-go/api"
)

func NewClient(client *api.API) *Client {
	return &Client{
		API: client,
	}
}

type Client struct {
	API *api.API
}

func (c *Client) ListLocations(ctx context.Context) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/config/locations",
	}
	return c.API.FormRequest(ctx, rc)
}

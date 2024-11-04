package location

import (
	"context"

	"github.com/ekaputra07/idcloudhost-go/http"
)

type LocationClient struct {
	H *http.Client
}

func (c *LocationClient) ListLocations(ctx context.Context) *http.ClientResponse {
	cfg := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/config/locations",
	}
	return c.H.FormRequest(ctx, cfg)
}

func NewClient() *LocationClient {
	return &LocationClient{
		H: http.DefaultClient,
	}
}

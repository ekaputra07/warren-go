package location

import (
	"context"
	"encoding/json"

	"github.com/ekaputra07/warren-go/api"
)

// Location represents data center location
type Location struct {
	DisplayName string `json:"display_name"`
	IsDefault   bool   `json:"is_default"`
	IsPreferred bool   `json:"is_preferred"`
	Description string `json:"description"`
	OrderNr     int    `json:"order_nr"`
	Slug        string `json:"slug"`
	CountryCode string `json:"country_code"`
}

func NewClient(client *api.API) *Client {
	return &Client{
		API: client,
	}
}

type Client struct {
	API *api.API
}

// ListLocations https://api.warren.io/#list-locations
func (c *Client) ListLocations(ctx context.Context) ([]Location, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/config/locations",
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}

	var locations []Location
	if err := json.Unmarshal(resp.Body, &locations); err != nil {
		return nil, err
	}
	return locations, nil
}

package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ekaputra07/warren-go/api"
	"github.com/google/uuid"
)

func NewClient(client *api.API, location string) *Client {
	return &Client{
		API:      client,
		Location: location,
	}
}

// ListNetworks https://api.warren.io/#list-networks
func (c *Client) ListNetworks(ctx context.Context) (*[]NetworkInfo, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/%s/network/networks", c.Location),
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return nil, res.Error
	}
	var i []NetworkInfo
	if err := json.Unmarshal(res.Body, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

// GetNetwork https://api.warren.io/#get-network-data
func (c *Client) GetNetwork(ctx context.Context, id uuid.UUID) (*NetworkInfo, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/%s/network/network/%s", c.Location, id),
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return nil, res.Error
	}
	var i NetworkInfo
	if err := json.Unmarshal(res.Body, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

// DeleteNetwork https://api.warren.io/#delete-network
func (c *Client) DeleteNetwork(ctx context.Context, id uuid.UUID) error {
	rc := api.RequestConfig{
		Method: "DELETE",
		Path:   fmt.Sprintf("/v1/%s/network/network/%s", c.Location, id),
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// RenameNetwork https://api.warren.io/#change-network-name
func (c *Client) RenameNetwork(ctx context.Context, id uuid.UUID, newName string) error {
	rc := api.RequestConfig{
		Method: "PATCH",
		Path:   fmt.Sprintf("/v1/%s/network/network/%s", c.Location, id),
		JSON:   map[string]interface{}{"name": newName},
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// GetOrCreateDefaultNetwork https://api.warren.io/#create-or-get-default-network
func (c *Client) GetOrCreateDefaultNetwork(ctx context.Context, name string) (*NetworkInfo, error) {
	rc := api.RequestConfig{
		Method: "POST",
		Path:   fmt.Sprintf("/v1/%s/network/network", c.Location),
		Query:  url.Values{"name": []string{name}},
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return nil, res.Error
	}
	var i NetworkInfo
	if err := json.Unmarshal(res.Body, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

// SetDefaultNetwork https://api.warren.io/#change-network-to-default
func (c *Client) SetDefaultNetwork(ctx context.Context, id uuid.UUID) error {
	rc := api.RequestConfig{
		Method: "PUT",
		Path:   fmt.Sprintf("/v1/%s/network/network/%s/default", c.Location, id),
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

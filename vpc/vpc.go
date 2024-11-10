package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ekaputra07/idcloudhost-go/http"
	"github.com/google/uuid"
)

func NewClient() *Client {
	return &Client{
		H: http.DefaultClient,
	}
}

// ListNetworks https://api.idcloudhost.com/#list-networks
func (c *Client) ListNetworks(ctx context.Context, location string) (*[]NetworkInfo, error) {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/%s/network/networks", location),
	}
	res := c.H.JSONRequest(ctx, rc)
	if res.Error != nil {
		return nil, res.Error
	}
	var i []NetworkInfo
	if err := json.Unmarshal(res.Body, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

// GetNetwork https://api.idcloudhost.com/#get-network-data
func (c *Client) GetNetwork(ctx context.Context, location string, id uuid.UUID) (*NetworkInfo, error) {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/%s/network/network/%s", location, id),
	}
	res := c.H.JSONRequest(ctx, rc)
	if res.Error != nil {
		return nil, res.Error
	}
	var i NetworkInfo
	if err := json.Unmarshal(res.Body, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

// DeleteNetwork https://api.idcloudhost.com/#delete-network
func (c *Client) DeleteNetwork(ctx context.Context, location string, id uuid.UUID) error {
	rc := http.RequestConfig{
		Method: "DELETE",
		Path:   fmt.Sprintf("/v1/%s/network/network/%s", location, id),
	}
	res := c.H.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// RenameNetwork https://api.idcloudhost.com/#change-network-name
func (c *Client) RenameNetwork(ctx context.Context, location string, id uuid.UUID, newName string) error {
	rc := http.RequestConfig{
		Method: "PATCH",
		Path:   fmt.Sprintf("/v1/%s/network/network/%s", location, id),
		JSON:   map[string]interface{}{"name": newName},
	}
	res := c.H.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// GetOrCreateDefaultNetwork https://api.idcloudhost.com/#create-or-get-default-network
func (c *Client) GetOrCreateDefaultNetwork(ctx context.Context, location string, name string) (*NetworkInfo, error) {
	rc := http.RequestConfig{
		Method: "POST",
		Path:   fmt.Sprintf("/v1/%s/network/network", location),
		Query:  url.Values{"name": []string{name}},
	}
	res := c.H.JSONRequest(ctx, rc)
	if res.Error != nil {
		return nil, res.Error
	}
	var i NetworkInfo
	if err := json.Unmarshal(res.Body, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

// SetDefaultNetwork https://api.idcloudhost.com/#change-network-to-default
func (c *Client) SetDefaultNetwork(ctx context.Context, location string, id uuid.UUID) error {
	rc := http.RequestConfig{
		Method: "PUT",
		Path:   fmt.Sprintf("/v1/%s/network/network/%s/default", location, id),
	}
	res := c.H.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

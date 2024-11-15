package ip

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ekaputra07/warren-go/api"
	"github.com/google/uuid"
)

func NewClient(client *api.API, location string) *Client {
	return &Client{
		API:      client,
		Location: location,
	}
}

// ListFloatingIPs https://api.warren.io/#list-floating-ips
func (c *Client) ListFloatingIPs(ctx context.Context) (*[]IPAddressInfo, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/%s/network/ip_addresses", c.Location),
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return nil, res.Error
	}
	var ips []IPAddressInfo
	if err := json.Unmarshal(res.Body, &ips); err != nil {
		return nil, err
	}
	return &ips, nil
}

// CreateFloatingIP https://api.warren.io/#create-floating-ip
func (c *Client) CreateFloatingIP(ctx context.Context, info *IPAddressInfo) error {
	if info.BillingAccountID == 0 {
		return fmt.Errorf("BillingAccountID with value of %v is invalid", info.BillingAccountID)
	}

	rc := api.RequestConfig{
		Method: "POST",
		Path:   fmt.Sprintf("/v1/%s/network/ip_addresses", c.Location),
		JSON: map[string]interface{}{
			"name":               info.Name,
			"billing_account_id": info.BillingAccountID,
		},
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	if err := json.Unmarshal(res.Body, info); err != nil {
		return err
	}
	return nil
}

// GetFloatingIP https://api.warren.io/#get-floating-ip
func (c *Client) GetFloatingIP(ctx context.Context, address string) (*IPAddressInfo, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/%s/network/ip_addresses/%s", c.Location, address),
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return nil, res.Error
	}
	var ip IPAddressInfo
	if err := json.Unmarshal(res.Body, &ip); err != nil {
		return nil, err
	}
	return &ip, nil
}

// UpdateFloatingIP https://api.warren.io/#update-floating-ip
func (c *Client) UpdateFloatingIP(ctx context.Context, info *IPAddressInfo) error {
	if info.BillingAccountID == 0 {
		return fmt.Errorf("BillingAccountID with value of %v is invalid", info.BillingAccountID)
	}

	rc := api.RequestConfig{
		Method: "PATCH",
		Path:   fmt.Sprintf("/v1/%s/network/ip_addresses", c.Location),
		JSON: map[string]interface{}{
			"name":               info.Name,
			"billing_account_id": info.BillingAccountID,
		},
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// DeleteFloatingIP https://api.warren.io/#delete-floating-ip
func (c *Client) DeleteFloatingIP(ctx context.Context, address string) error {
	rc := api.RequestConfig{
		Method: "DELETE",
		Path:   fmt.Sprintf("/v1/%s/network/ip_addresses/%s", c.Location, address),
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// AssignFloatingIPToVM https://api.warren.io/#assign-floating-ip
func (c *Client) AssignFloatingIPToVM(ctx context.Context, address string, vmUUID uuid.UUID) error {
	rc := api.RequestConfig{
		Method: "POST",
		Path:   fmt.Sprintf("/v1/%s/network/ip_addresses/%s/assign", c.Location, address),
		JSON:   map[string]interface{}{"vm_uuid": vmUUID},
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// UnassignFloatingIPFromVM https://api.warren.io/#un-assign-floating-ip
func (c *Client) UnassignFloatingIPFromVM(ctx context.Context, address string, vmUUID uuid.UUID) error {
	rc := api.RequestConfig{
		Method: "POST",
		Path:   fmt.Sprintf("/v1/%s/network/ip_addresses/%s/unassign", c.Location, address),
		JSON:   map[string]interface{}{"vm_uuid": vmUUID},
	}
	res := c.API.JSONRequest(ctx, rc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

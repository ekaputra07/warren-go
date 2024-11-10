package blockstorage

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ekaputra07/warren-go/api"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

func NewClient(client *api.API) *Client {
	return &Client{
		API: client,
	}
}

// ListDisks https://api.warren.io/#list-disks
func (c *Client) LisDisks(ctx context.Context) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/disks",
	}
	return c.API.FormRequest(ctx, rc)
}

// CreateDisk https://api.warren.io/#create-disk
func (c *Client) CreateDisk(ctx context.Context, cfg CreateDiskConfig) *api.ClientResponse {
	enc := schema.NewEncoder()
	d := url.Values{}
	if err := enc.Encode(cfg, d); err != nil {
		return &api.ClientResponse{Error: err}
	}

	rc := api.RequestConfig{
		Method: "POST",
		Path:   "/v1/storage/disks",
		Data:   d,
	}
	return c.API.FormRequest(ctx, rc)
}

// GetDisk https://api.warren.io/#get-disk
func (c *Client) GetDisk(ctx context.Context, diskID uuid.UUID) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
	}
	return c.API.FormRequest(ctx, rc)
}

// DeleteDisk https://api.warren.io/#delete-disk
func (c *Client) DeleteDisk(ctx context.Context, diskID uuid.UUID) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "DELETE",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
	}
	return c.API.FormRequest(ctx, rc)
}

// AttachDiskToVM https://api.warren.io/#attach-disk
func (c *Client) AttachDiskToVM(ctx context.Context, diskID, vmID uuid.UUID) *api.ClientResponse {
	d := url.Values{
		"uuid":         []string{vmID.String()},
		"storage_uuid": []string{diskID.String()},
	}
	rc := api.RequestConfig{
		Method: "POST",
		Path:   "/v1/user-resource/vm/storage/attach",
		Data:   d,
	}
	return c.API.FormRequest(ctx, rc)
}

// DetachDiskFromVM https://api.warren.io/#detach-disk
func (c *Client) DetachDiskFromVM(ctx context.Context, diskID, vmID uuid.UUID) *api.ClientResponse {
	d := url.Values{
		"uuid":         []string{vmID.String()},
		"storage_uuid": []string{diskID.String()},
	}
	rc := api.RequestConfig{
		Method: "POST",
		Path:   "/v1/user-resource/vm/storage/detach",
		Data:   d,
	}
	return c.API.FormRequest(ctx, rc)
}

// UpdateDiskBillingAccount https://api.warren.io/#modify-disk-info
func (c *Client) UpdateDiskBillingAccount(ctx context.Context, diskID uuid.UUID, billingAccountID int) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "PATCH",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
		Data:   url.Values{"billing_account_id": []string{strconv.Itoa(billingAccountID)}},
	}
	return c.API.FormRequest(ctx, rc)
}

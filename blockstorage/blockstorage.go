package blockstorage

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ekaputra07/idcloudhost-go/http"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

func NewClient() *Client {
	return &Client{
		H: http.DefaultClient,
	}
}

// ListDisks https://api.idcloudhost.com/#list-disks
func (c *Client) LisDisks(ctx context.Context) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/disks",
	}
	return c.H.FormRequest(ctx, rc)
}

// CreateDisk https://api.idcloudhost.com/#create-disk
func (c *Client) CreateDisk(ctx context.Context, cfg CreateDiskConfig) *http.ClientResponse {
	enc := schema.NewEncoder()
	d := url.Values{}
	if err := enc.Encode(cfg, d); err != nil {
		return &http.ClientResponse{Error: err}
	}

	rc := http.RequestConfig{
		Method: "POST",
		Path:   "/v1/storage/disks",
		Data:   d,
	}
	return c.H.FormRequest(ctx, rc)
}

// GetDisk https://api.idcloudhost.com/#get-disk
func (c *Client) GetDisk(ctx context.Context, diskID uuid.UUID) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
	}
	return c.H.FormRequest(ctx, rc)
}

// DeleteDisk https://api.idcloudhost.com/#delete-disk
func (c *Client) DeleteDisk(ctx context.Context, diskID uuid.UUID) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "DELETE",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
	}
	return c.H.FormRequest(ctx, rc)
}

// AttachDiskToVM https://api.idcloudhost.com/#attach-disk
func (c *Client) AttachDiskToVM(ctx context.Context, diskID, vmID uuid.UUID) *http.ClientResponse {
	d := url.Values{
		"uuid":         []string{vmID.String()},
		"storage_uuid": []string{diskID.String()},
	}
	rc := http.RequestConfig{
		Method: "POST",
		Path:   "/v1/user-resource/vm/storage/attach",
		Data:   d,
	}
	return c.H.FormRequest(ctx, rc)
}

// DetachDiskFromVM https://api.idcloudhost.com/#detach-disk
func (c *Client) DetachDiskFromVM(ctx context.Context, diskID, vmID uuid.UUID) *http.ClientResponse {
	d := url.Values{
		"uuid":         []string{vmID.String()},
		"storage_uuid": []string{diskID.String()},
	}
	rc := http.RequestConfig{
		Method: "POST",
		Path:   "/v1/user-resource/vm/storage/detach",
		Data:   d,
	}
	return c.H.FormRequest(ctx, rc)
}

// UpdateDiskBillingAccount https://api.idcloudhost.com/#modify-disk-info
func (c *Client) UpdateDiskBillingAccount(ctx context.Context, diskID uuid.UUID, billingAccountID int) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "PATCH",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
		Data:   url.Values{"billing_account_id": []string{strconv.Itoa(billingAccountID)}},
	}
	return c.H.FormRequest(ctx, rc)
}

package blockstorage

import (
	"context"
	"encoding/json"
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
func (c *Client) LisDisks(ctx context.Context) (*[]Disk, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/disks",
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}
	var disks []Disk
	if err := json.Unmarshal(resp.Body, &disks); err != nil {
		return nil, err
	}
	return &disks, nil
}

// CreateDisk https://api.warren.io/#create-disk
func (c *Client) CreateDisk(ctx context.Context, disk *Disk) error {
	enc := schema.NewEncoder()
	d := url.Values{}
	if err := enc.Encode(disk, d); err != nil {
		return err
	}

	rc := api.RequestConfig{
		Method: "POST",
		Path:   "/v1/storage/disks",
		Data:   d,
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return resp.Error
	}
	return json.Unmarshal(resp.Body, disk)
}

// GetDisk https://api.warren.io/#get-disk
func (c *Client) GetDisk(ctx context.Context, diskID uuid.UUID) (*Disk, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}
	var disk Disk
	if err := json.Unmarshal(resp.Body, &disk); err != nil {
		return nil, err
	}
	return &disk, nil
}

// DeleteDisk https://api.warren.io/#delete-disk
func (c *Client) DeleteDisk(ctx context.Context, diskID uuid.UUID) error {
	rc := api.RequestConfig{
		Method: "DELETE",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
	}
	return c.API.FormRequest(ctx, rc).Error
}

// AttachDiskToVM https://api.warren.io/#attach-disk
func (c *Client) AttachDiskToVM(ctx context.Context, diskID, vmID uuid.UUID) error {
	d := url.Values{
		"uuid":         []string{vmID.String()},
		"storage_uuid": []string{diskID.String()},
	}
	rc := api.RequestConfig{
		Method: "POST",
		Path:   "/v1/user-resource/vm/storage/attach",
		Data:   d,
	}
	return c.API.FormRequest(ctx, rc).Error
}

// DetachDiskFromVM https://api.warren.io/#detach-disk
func (c *Client) DetachDiskFromVM(ctx context.Context, diskID, vmID uuid.UUID) error {
	d := url.Values{
		"uuid":         []string{vmID.String()},
		"storage_uuid": []string{diskID.String()},
	}
	rc := api.RequestConfig{
		Method: "POST",
		Path:   "/v1/user-resource/vm/storage/detach",
		Data:   d,
	}
	return c.API.FormRequest(ctx, rc).Error
}

// UpdateDiskBillingAccount https://api.warren.io/#modify-disk-info
func (c *Client) UpdateDiskBillingAccount(ctx context.Context, diskID uuid.UUID, billingAccountID int) error {
	rc := api.RequestConfig{
		Method: "PATCH",
		Path:   fmt.Sprintf("/v1/storage/disks/%s", diskID),
		Data:   url.Values{"billing_account_id": []string{strconv.Itoa(billingAccountID)}},
	}
	return c.API.FormRequest(ctx, rc).Error
}

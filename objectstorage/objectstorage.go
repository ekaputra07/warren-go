package objectstorage

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/ekaputra07/warren-go/api"
)

func NewClient(client *api.API) *Client {
	return &Client{
		API: client,
	}
}

// ForBillingAccount set the value of BillingAccountID
func (c *Client) ForBillingAccount(id int) *Client {
	c.BillingAccountID = id
	return c
}

// GetS3ApiURL https://api.warren.io/#s3-api-info
func (c *Client) GetS3ApiURL(ctx context.Context) (*map[string]string, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/api/s3",
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}
	var data map[string]string
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetS3UserInfo https://api.warren.io/#get-s3-user
func (c *Client) GetS3UserInfo(ctx context.Context) (*S3UserInfo, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/user",
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}
	var info S3UserInfo
	if err := json.Unmarshal(resp.Body, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

// GetS3UserKeys https://api.warren.io/#get-keys
func (c *Client) GetS3UserKeys(ctx context.Context) (*[]S3Credential, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/user/keys",
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}
	var credentials []S3Credential
	if err := json.Unmarshal(resp.Body, &credentials); err != nil {
		return nil, err
	}
	return &credentials, nil
}

// GenerateS3UserKey https://api.warren.io/#generate-key
func (c *Client) GenerateS3UserKey(ctx context.Context) (*[]S3Credential, error) {
	rc := api.RequestConfig{
		Method: "POST",
		Path:   "/v1/storage/user/keys",
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}
	var credentials []S3Credential
	if err := json.Unmarshal(resp.Body, &credentials); err != nil {
		return nil, err
	}
	return &credentials, nil
}

// DeleteS3UserKey https://api.warren.io/#generate-key
func (c *Client) DeleteS3UserKey(ctx context.Context, accessKey string) error {
	rc := api.RequestConfig{
		Method: "DELETE",
		Path:   "/v1/storage/user/keys",
		Query:  url.Values{"access_key": []string{accessKey}},
	}
	return c.API.FormRequest(ctx, rc).Error
}

// ListBuckets https://api.warren.io/#list-buckets
func (c *Client) ListBuckets(ctx context.Context) (*[]S3Bucket, error) {
	var resp *api.ClientResponse

	if c.BillingAccountID == 0 {
		rc := api.RequestConfig{
			Method: "GET",
			Path:   "/v1/storage/bucket/list",
		}
		resp = c.API.FormRequest(ctx, rc)
	} else {
		rc := api.RequestConfig{
			Method: "GET",
			Path:   "/v1/storage/bucket/list",
			Query:  url.Values{"billing_account_id": []string{strconv.Itoa(c.BillingAccountID)}},
		}
		resp = c.API.FormRequest(ctx, rc)
	}

	if resp.Error != nil {
		return nil, resp.Error
	}
	var buckets []S3Bucket
	if err := json.Unmarshal(resp.Body, &buckets); err != nil {
		return nil, err
	}
	return &buckets, nil
}

// GetBucket https://api.warren.io/#get-bucket
func (c *Client) GetBucket(ctx context.Context, bucketName string) (*S3Bucket, error) {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/bucket",
		Query:  url.Values{"name": []string{bucketName}},
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}
	var bucket S3Bucket
	if err := json.Unmarshal(resp.Body, &bucket); err != nil {
		return nil, err
	}
	return &bucket, nil
}

// CreateBucket https://api.warren.io/#create-bucket
func (c *Client) CreateBucket(ctx context.Context, bucketName string) (*S3Bucket, error) {
	d := url.Values{"name": []string{bucketName}}
	if c.BillingAccountID != 0 {
		d.Add("billing_account_id", strconv.Itoa(c.BillingAccountID))
	}

	rc := api.RequestConfig{
		Method: "PUT",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	resp := c.API.FormRequest(ctx, rc)
	if resp.Error != nil {
		return nil, resp.Error
	}
	var bucket S3Bucket
	if err := json.Unmarshal(resp.Body, &bucket); err != nil {
		return nil, err
	}
	return &bucket, nil
}

// DeleteBucket https://api.warren.io/#delete-bucket
func (c *Client) DeleteBucket(ctx context.Context, bucketName string) error {
	rc := api.RequestConfig{
		Method: "DELETE",
		Path:   "/v1/storage/bucket",
		Query:  url.Values{"name": []string{bucketName}},
	}
	return c.API.FormRequest(ctx, rc).Error
}

// UpdateBucketBillingAccount https://api.warren.io/#modify-bucket
func (c *Client) UpdateBucketBillingAccount(ctx context.Context, bucketName string, billingAccountID int) error {
	d := url.Values{
		"name":               []string{bucketName},
		"billing_account_id": []string{strconv.Itoa(billingAccountID)},
	}

	rc := api.RequestConfig{
		Method: "PATCH",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	return c.API.FormRequest(ctx, rc).Error
}

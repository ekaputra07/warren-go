package objectstorage

import (
	"context"
	"net/url"
	"strconv"

	"github.com/ekaputra07/warren-go/api"
)

func NewClient(client *api.API) *Client {
	return &Client{
		API: client,
	}
}

type Client struct {
	BillingAccountID int
	API              *api.API
}

// ForBillingAccount set the value of BillingAccountID
func (c *Client) ForBillingAccount(id int) *Client {
	c.BillingAccountID = id
	return c
}

// GetS3ApiURL https://api.warren.io/#s3-api-info
func (c *Client) GetS3ApiURL(ctx context.Context) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/api/s3",
	}
	return c.API.FormRequest(ctx, rc)
}

// GetS3UserInfo https://api.warren.io/#get-s3-user
func (c *Client) GetS3UserInfo(ctx context.Context) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/user",
	}
	return c.API.FormRequest(ctx, rc)
}

// GetS3UserKeys https://api.warren.io/#get-keys
func (c *Client) GetS3UserKeys(ctx context.Context) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/user/keys",
	}
	return c.API.FormRequest(ctx, rc)
}

// GenerateS3UserKey https://api.warren.io/#generate-key
func (c *Client) GenerateS3UserKey(ctx context.Context) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "POST",
		Path:   "/v1/storage/user/keys",
	}
	return c.API.FormRequest(ctx, rc)
}

// DeleteS3UserKey https://api.warren.io/#generate-key
func (c *Client) DeleteS3UserKey(ctx context.Context, accessKey string) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "DELETE",
		Path:   "/v1/storage/user/keys",
		Query:  url.Values{"access_key": []string{accessKey}},
	}
	return c.API.FormRequest(ctx, rc)
}

// ListBuckets https://api.warren.io/#list-buckets
func (c *Client) ListBuckets(ctx context.Context) *api.ClientResponse {
	if c.BillingAccountID == 0 {
		rc := api.RequestConfig{
			Method: "GET",
			Path:   "/v1/storage/bucket/list",
		}
		return c.API.FormRequest(ctx, rc)
	}

	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/bucket/list",
		Query:  url.Values{"billing_account_id": []string{strconv.Itoa(c.BillingAccountID)}},
	}
	return c.API.FormRequest(ctx, rc)
}

// GetBucket https://api.warren.io/#get-bucket
func (c *Client) GetBucket(ctx context.Context, bucketName string) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/bucket",
		Query:  url.Values{"name": []string{bucketName}},
	}
	return c.API.FormRequest(ctx, rc)
}

// CreateBucket https://api.warren.io/#create-bucket
func (c *Client) CreateBucket(ctx context.Context, bucketName string) *api.ClientResponse {
	d := url.Values{"name": []string{bucketName}}
	if c.BillingAccountID != 0 {
		d.Add("billing_account_id", strconv.Itoa(c.BillingAccountID))
	}

	rc := api.RequestConfig{
		Method: "PUT",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	return c.API.FormRequest(ctx, rc)
}

// DeleteBucket https://api.warren.io/#delete-bucket
func (c *Client) DeleteBucket(ctx context.Context, bucketName string) *api.ClientResponse {
	rc := api.RequestConfig{
		Method: "DELETE",
		Path:   "/v1/storage/bucket",
		Query:  url.Values{"name": []string{bucketName}},
	}
	return c.API.FormRequest(ctx, rc)
}

// UpdateBucketBillingAccount https://api.warren.io/#modify-bucket
func (c *Client) UpdateBucketBillingAccount(ctx context.Context, bucketName string, billingAccountID int) *api.ClientResponse {
	d := url.Values{
		"name":               []string{bucketName},
		"billing_account_id": []string{strconv.Itoa(billingAccountID)},
	}

	rc := api.RequestConfig{
		Method: "PATCH",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	return c.API.FormRequest(ctx, rc)
}

package objectstorage

import (
	"context"
	"net/url"
	"strconv"

	"github.com/ekaputra07/idcloudhost-go/http"
)

func NewClient() *Client {
	return &Client{
		H: http.DefaultClient,
	}
}

type Client struct {
	BillingAccountID int
	H                *http.Client
}

// ForBillingAccount set the value of BillingAccountID
func (c *Client) ForBillingAccount(id int) *Client {
	c.BillingAccountID = id
	return c
}

// GetS3ApiURL https://api.idcloudhost.com/#s3-api-info
func (c *Client) GetS3ApiURL(ctx context.Context) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/api/s3",
	}
	return c.H.FormRequest(ctx, rc)
}

// GetS3UserInfo https://api.idcloudhost.com/#get-s3-user
func (c *Client) GetS3UserInfo(ctx context.Context) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/user",
	}
	return c.H.FormRequest(ctx, rc)
}

// GetS3UserKeys https://api.idcloudhost.com/#get-keys
func (c *Client) GetS3UserKeys(ctx context.Context) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/user/keys",
	}
	return c.H.FormRequest(ctx, rc)
}

// GenerateS3UserKey https://api.idcloudhost.com/#generate-key
func (c *Client) GenerateS3UserKey(ctx context.Context) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "POST",
		Path:   "/v1/storage/user/keys",
	}
	return c.H.FormRequest(ctx, rc)
}

// DeleteS3UserKey https://api.idcloudhost.com/#generate-key
func (c *Client) DeleteS3UserKey(ctx context.Context, accessKey string) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "DELETE",
		Path:   "/v1/storage/user/keys",
		Query:  url.Values{"access_key": []string{accessKey}},
	}
	return c.H.FormRequest(ctx, rc)
}

// ListBuckets https://api.idcloudhost.com/#list-buckets
func (c *Client) ListBuckets(ctx context.Context) *http.ClientResponse {
	if c.BillingAccountID == 0 {
		rc := http.RequestConfig{
			Method: "GET",
			Path:   "/v1/storage/bucket/list",
		}
		return c.H.FormRequest(ctx, rc)
	}

	rc := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/bucket/list",
		Query:  url.Values{"billing_account_id": []string{strconv.Itoa(c.BillingAccountID)}},
	}
	return c.H.FormRequest(ctx, rc)
}

// GetBucket https://api.idcloudhost.com/#get-bucket
func (c *Client) GetBucket(ctx context.Context, bucketName string) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/bucket",
		Query:  url.Values{"name": []string{bucketName}},
	}
	return c.H.FormRequest(ctx, rc)
}

// CreateBucket https://api.idcloudhost.com/#create-bucket
func (c *Client) CreateBucket(ctx context.Context, bucketName string) *http.ClientResponse {
	d := url.Values{"name": []string{bucketName}}
	if c.BillingAccountID != 0 {
		d.Add("billing_account_id", strconv.Itoa(c.BillingAccountID))
	}

	rc := http.RequestConfig{
		Method: "PUT",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	return c.H.FormRequest(ctx, rc)
}

// DeleteBucket https://api.idcloudhost.com/#delete-bucket
func (c *Client) DeleteBucket(ctx context.Context, bucketName string) *http.ClientResponse {
	rc := http.RequestConfig{
		Method: "DELETE",
		Path:   "/v1/storage/bucket",
		Query:  url.Values{"name": []string{bucketName}},
	}
	return c.H.FormRequest(ctx, rc)
}

// UpdateBucketBillingAccount https://api.idcloudhost.com/#modify-bucket
func (c *Client) UpdateBucketBillingAccount(ctx context.Context, bucketName string, billingAccountID int) *http.ClientResponse {
	d := url.Values{
		"name":               []string{bucketName},
		"billing_account_id": []string{strconv.Itoa(billingAccountID)},
	}

	rc := http.RequestConfig{
		Method: "PATCH",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	return c.H.FormRequest(ctx, rc)
}

package objectstorage

import (
	"context"
	"net/url"

	"github.com/ekaputra07/idcloudhost-go/http"
)

type ObjectStorageClient struct {
	BillingAccountID string
	H                *http.Client
}

// ForBillingAccount set the value of BillingAccountID
func (c *ObjectStorageClient) ForBillingAccount(id string) *ObjectStorageClient {
	c.BillingAccountID = id
	return c
}

// GetS3ApiURL https://api.idcloudhost.com/#s3-api-info
func (c *ObjectStorageClient) GetS3ApiURL(ctx context.Context) *http.ClientResponse {
	cfg := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/api/s3",
	}
	return c.H.FormRequest(ctx, cfg)
}

// GetS3UserInfo https://api.idcloudhost.com/#get-s3-user
func (c *ObjectStorageClient) GetS3UserInfo(ctx context.Context) *http.ClientResponse {
	cfg := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/user",
	}
	return c.H.FormRequest(ctx, cfg)
}

// GetS3UserKeys https://api.idcloudhost.com/#get-keys
func (c *ObjectStorageClient) GetS3UserKeys(ctx context.Context) *http.ClientResponse {
	cfg := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/user/keys",
	}
	return c.H.FormRequest(ctx, cfg)
}

// GenerateS3UserKey https://api.idcloudhost.com/#generate-key
func (c *ObjectStorageClient) GenerateS3UserKey(ctx context.Context) *http.ClientResponse {
	cfg := http.RequestConfig{
		Method: "POST",
		Path:   "/v1/storage/user/keys",
	}
	return c.H.FormRequest(ctx, cfg)
}

// DeleteS3UserKey https://api.idcloudhost.com/#generate-key
func (c *ObjectStorageClient) DeleteS3UserKey(ctx context.Context, accessKey string) *http.ClientResponse {
	d := url.Values{}
	d.Add("access_key", accessKey)

	cfg := http.RequestConfig{
		Method: "DELETE",
		Path:   "/v1/storage/user/keys",
		Data:   d,
	}
	return c.H.FormRequest(ctx, cfg)
}

// ListBuckets https://api.idcloudhost.com/#list-buckets
func (c *ObjectStorageClient) ListBuckets(ctx context.Context) *http.ClientResponse {
	if c.BillingAccountID == "" {
		cfg := http.RequestConfig{
			Method: "GET",
			Path:   "/v1/storage/bucket/list",
		}
		return c.H.FormRequest(ctx, cfg)
	}
	q := url.Values{}
	q.Add("billing_account_id", c.BillingAccountID)

	cfg := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/bucket/list",
		Query:  q,
	}
	return c.H.FormRequest(ctx, cfg)
}

// GetBucket https://api.idcloudhost.com/#get-bucket
func (c *ObjectStorageClient) GetBucket(ctx context.Context, name string) *http.ClientResponse {
	q := url.Values{}
	q.Add("name", name)

	cfg := http.RequestConfig{
		Method: "GET",
		Path:   "/v1/storage/bucket",
		Query:  q,
	}
	return c.H.FormRequest(ctx, cfg)
}

// CreateBucket https://api.idcloudhost.com/#create-bucket
func (c *ObjectStorageClient) CreateBucket(ctx context.Context, bucketName string) *http.ClientResponse {
	d := url.Values{}
	d.Add("name", bucketName)
	if c.BillingAccountID != "" {
		d.Add("billing_account_id", c.BillingAccountID)
	}

	cfg := http.RequestConfig{
		Method: "PUT",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	return c.H.FormRequest(ctx, cfg)
}

// DeleteBucket https://api.idcloudhost.com/#delete-bucket
func (c *ObjectStorageClient) DeleteBucket(ctx context.Context, bucketName string) *http.ClientResponse {
	d := url.Values{}
	d.Add("name", bucketName)

	cfg := http.RequestConfig{
		Method: "DELETE",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	return c.H.FormRequest(ctx, cfg)
}

// UpdateBucketBillingAccount https://api.idcloudhost.com/#modify-bucket
func (c *ObjectStorageClient) UpdateBucketBillingAccount(ctx context.Context, bucketName, billingAccountId string) *http.ClientResponse {
	d := url.Values{}
	d.Add("name", bucketName)
	d.Add("billing_account_id", billingAccountId)

	cfg := http.RequestConfig{
		Method: "PATCH",
		Path:   "/v1/storage/bucket",
		Data:   d,
	}
	return c.H.FormRequest(ctx, cfg)
}

func NewClient() *ObjectStorageClient {
	return &ObjectStorageClient{
		H: http.DefaultClient,
	}
}

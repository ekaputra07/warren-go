package objectstorage

import (
	"context"
	"net/http"
	"testing"

	h "github.com/ekaputra07/idcloudhost-go/http"
	"github.com/stretchr/testify/assert"
)

func ForBillingAccount(t *testing.T) {
	c := NewClient().ForBillingAccount("test")
	assert.Equal(t, "test", c.BillingAccountID)
}

func TestGetS3ApiURL(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/api/s3", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.GetS3ApiURL(context.Background())
}

func TestGetS3UserInfo(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/user", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.GetS3UserInfo(context.Background())
}

func TestGetS3UserKeys(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/user/keys", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.GetS3UserKeys(context.Background())
}

func TestGenerateS3UserKey(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/storage/user/keys", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.GenerateS3UserKey(context.Background())
}

func TestDeleteS3UserKey(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/v1/storage/user/keys?access_key=testKey", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.DeleteS3UserKey(context.Background(), "testKey")
}

func TestListBuckets(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/bucket/list", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.ListBuckets(context.Background())
}

func TestListBucketsWithBillingAccount(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/bucket/list?billing_account_id=testId", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.BillingAccountID = "testId"
	osc.ListBuckets(context.Background())
}

func TestGetBucket(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/bucket?name=testBucket", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.GetBucket(context.Background(), "testBucket")
}

func TestCreateBucket(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, "/v1/storage/bucket", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "testBucket", r.Form.Get("name"))
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.CreateBucket(context.Background(), "testBucket")
}

func TestCreateBucketWithBillingAccount(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, "/v1/storage/bucket", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "testBucket", r.Form.Get("name"))
		assert.Equal(t, "testId", r.Form.Get("billing_account_id"))
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.BillingAccountID = "testId"
	osc.CreateBucket(context.Background(), "testBucket")
}

func TestDeleteBucket(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/v1/storage/bucket?name=testBucket", r.RequestURI)
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.DeleteBucket(context.Background(), "testBucket")
}

func TestUpdateBucketBillingAccount(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.Equal(t, "/v1/storage/bucket", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "testBucket", r.Form.Get("name"))
		assert.Equal(t, "testId", r.Form.Get("billing_account_id"))
	})
	defer s.Close()

	osc := ObjectStorageClient{H: c}
	osc.UpdateBucketBillingAccount(context.Background(), "testBucket", "testId")
}

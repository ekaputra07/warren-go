package objectstorage

import (
	"context"
	"net/http"
	"testing"

	"github.com/ekaputra07/warren-go/api"
	"github.com/stretchr/testify/assert"
)

func ForBillingAccount(t *testing.T) {
	c := NewClient(api.Default).ForBillingAccount(123)
	assert.Equal(t, 123, c.BillingAccountID)
}

func TestGetS3ApiURL(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/api/s3", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.GetS3ApiURL(context.Background())
}

func TestGetS3UserInfo(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/user", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.GetS3UserInfo(context.Background())
}

func TestGetS3UserKeys(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/user/keys", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.GetS3UserKeys(context.Background())
}

func TestGenerateS3UserKey(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/storage/user/keys", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.GenerateS3UserKey(context.Background())
}

func TestDeleteS3UserKey(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/v1/storage/user/keys?access_key=testKey", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.DeleteS3UserKey(context.Background(), "testKey")
}

func TestListBuckets(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/bucket/list", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.ListBuckets(context.Background())
}

func TestListBucketsWithBillingAccount(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/bucket/list?billing_account_id=123", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.BillingAccountID = 123
	os.ListBuckets(context.Background())
}

func TestGetBucket(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/bucket?name=testBucket", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.GetBucket(context.Background(), "testBucket")
}

func TestCreateBucket(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, "/v1/storage/bucket", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "testBucket", r.Form.Get("name"))
	})
	defer s.Close()

	os := Client{API: a}
	os.CreateBucket(context.Background(), "testBucket")
}

func TestCreateBucketWithBillingAccount(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, "/v1/storage/bucket", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "testBucket", r.Form.Get("name"))
		assert.Equal(t, "123", r.Form.Get("billing_account_id"))
	})
	defer s.Close()

	os := Client{API: a}
	os.BillingAccountID = 123
	os.CreateBucket(context.Background(), "testBucket")
}

func TestDeleteBucket(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/v1/storage/bucket?name=testBucket", r.RequestURI)
	})
	defer s.Close()

	os := Client{API: a}
	os.DeleteBucket(context.Background(), "testBucket")
}

func TestUpdateBucketBillingAccount(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.Equal(t, "/v1/storage/bucket", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "testBucket", r.Form.Get("name"))
		assert.Equal(t, "123", r.Form.Get("billing_account_id"))
	})
	defer s.Close()

	os := Client{API: a}
	os.UpdateBucketBillingAccount(context.Background(), "testBucket", 123)
}

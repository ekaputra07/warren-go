package blockstorage

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	h "github.com/ekaputra07/idcloudhost-go/http"
	"github.com/stretchr/testify/assert"
)

func TestListDisks(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/disks", r.RequestURI)
	})
	defer s.Close()

	bs := Client{H: c}
	bs.LisDisks(context.Background())
}

func TestCreateDisk(t *testing.T) {
	config := CreateDiskConfig{
		SizeGb:           10,
		BillingAccountId: 123,
		SourceImageType:  ImageTypeOSBase,
		SourceImage:      "ubuntu_20.04",
	}
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/storage/disks", r.RequestURI)

		_ = r.ParseForm()

		assert.Equal(t, strconv.Itoa(config.SizeGb), r.Form.Get("size_gb"))
		assert.Equal(t, strconv.Itoa(config.BillingAccountId), r.Form.Get("billing_account_id"))
		assert.Equal(t, string(ImageTypeOSBase), r.Form.Get("source_image_type"))
		assert.Equal(t, config.SourceImage, r.Form.Get("source_image"))
	})
	defer s.Close()

	bs := Client{H: c}
	bs.CreateDisk(context.Background(), config)
}

func TestGetDisk(t *testing.T) {
	id := "testId"
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/disks/testId", r.RequestURI)
	})
	defer s.Close()

	bs := Client{H: c}
	bs.GetDisk(context.Background(), id)
}

func TestDeleteDisk(t *testing.T) {
	id := "testId"
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/v1/storage/disks/testId", r.RequestURI)
	})
	defer s.Close()

	bs := Client{H: c}
	bs.DeleteDisk(context.Background(), id)
}

func TestAttachDiskToVM(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/user-resource/vm/storage/attach", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "vmId", r.Form.Get("uuid"))
		assert.Equal(t, "diskId", r.Form.Get("storage_uuid"))
	})
	defer s.Close()

	bs := Client{H: c}
	bs.AttachDiskToVM(context.Background(), "diskId", "vmId")
}

func TestDetachDiskFromVM(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/user-resource/vm/storage/detach", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "vmId", r.Form.Get("uuid"))
		assert.Equal(t, "diskId", r.Form.Get("storage_uuid"))
	})
	defer s.Close()

	bs := Client{H: c}
	bs.DetachDiskFromVM(context.Background(), "diskId", "vmId")
}

func TestUpdateBucketBillingAccount(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.Equal(t, "/v1/storage/disks/testId", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "123", r.Form.Get("billing_account_id"))
	})
	defer s.Close()

	bs := Client{H: c}
	bs.UpdateDiskBillingAccount(context.Background(), "testId", 123)
}

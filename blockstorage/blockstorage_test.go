package blockstorage

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/ekaputra07/warren-go/api"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListDisks(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/storage/disks", r.RequestURI)
	})
	defer s.Close()

	bs := Client{API: a}
	bs.LisDisks(context.Background())
}

func TestCreateDisk(t *testing.T) {
	config := CreateDiskConfig{
		SizeGB:           10,
		BillingAccountID: 123,
		SourceImageType:  ImageTypeOSBase,
		SourceImage:      "ubuntu_20.04",
	}
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/storage/disks", r.RequestURI)

		_ = r.ParseForm()

		assert.Equal(t, strconv.Itoa(config.SizeGB), r.Form.Get("size_gb"))
		assert.Equal(t, strconv.Itoa(config.BillingAccountID), r.Form.Get("billing_account_id"))
		assert.Equal(t, string(ImageTypeOSBase), r.Form.Get("source_image_type"))
		assert.Equal(t, config.SourceImage, r.Form.Get("source_image"))
	})
	defer s.Close()

	bs := Client{API: a}
	bs.CreateDisk(context.Background(), config)
}

func TestGetDisk(t *testing.T) {
	id := uuid.New()
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/storage/disks/%s", id), r.RequestURI)
	})
	defer s.Close()

	bs := Client{API: a}
	bs.GetDisk(context.Background(), id)
}

func TestDeleteDisk(t *testing.T) {
	id := uuid.New()
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/storage/disks/%s", id), r.RequestURI)
	})
	defer s.Close()

	bs := Client{API: a}
	bs.DeleteDisk(context.Background(), id)
}

func TestAttachDiskToVM(t *testing.T) {
	diskId := uuid.New()
	vmId := uuid.New()
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/user-resource/vm/storage/attach", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, vmId.String(), r.Form.Get("uuid"))
		assert.Equal(t, diskId.String(), r.Form.Get("storage_uuid"))
	})
	defer s.Close()

	bs := Client{API: a}
	bs.AttachDiskToVM(context.Background(), diskId, vmId)
}

func TestDetachDiskFromVM(t *testing.T) {
	diskId := uuid.New()
	vmId := uuid.New()
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/v1/user-resource/vm/storage/detach", r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, vmId.String(), r.Form.Get("uuid"))
		assert.Equal(t, diskId.String(), r.Form.Get("storage_uuid"))
	})
	defer s.Close()

	bs := Client{API: a}
	bs.DetachDiskFromVM(context.Background(), diskId, vmId)
}

func TestUpdateBucketBillingAccount(t *testing.T) {
	id := uuid.New()
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/storage/disks/%s", id), r.RequestURI)

		_ = r.ParseForm()
		assert.Equal(t, "123", r.Form.Get("billing_account_id"))
	})
	defer s.Close()

	bs := Client{API: a}
	bs.UpdateDiskBillingAccount(context.Background(), id, 123)
}

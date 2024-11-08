package blockstorage

import (
	"github.com/ekaputra07/idcloudhost-go/http"
)

type Client struct {
	H *http.Client
}

type SourceImageType string

const (
	ImageTypeOSBase   SourceImageType = "OS_BASE"
	ImageTypeDisk     SourceImageType = "DISK"
	ImageTypeSnapshot SourceImageType = "SNAPSHOT"
	ImageTypeEmpty    SourceImageType = "EMPTY"
)

type CreateDiskConfig struct {
	SizeGb           int             `schema:"size_gb"`
	BillingAccountId int             `schema:"billing_account_id"`
	SourceImageType  SourceImageType `schema:"source_image_type,default:EMPTY"`
	SourceImage      string          `schema:"source_image"`
}

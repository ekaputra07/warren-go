package blockstorage

import "github.com/ekaputra07/warren-go/api"

type Client struct {
	API *api.API
}

type SourceImageType string

const (
	ImageTypeOSBase   SourceImageType = "OS_BASE"
	ImageTypeDisk     SourceImageType = "DISK"
	ImageTypeSnapshot SourceImageType = "SNAPSHOT"
	ImageTypeEmpty    SourceImageType = "EMPTY"
)

type CreateDiskConfig struct {
	SizeGB           int             `schema:"size_gb"`
	BillingAccountID int             `schema:"billing_account_id"`
	SourceImageType  SourceImageType `schema:"source_image_type,default:EMPTY"`
	SourceImage      string          `schema:"source_image"`
}

package blockstorage

import (
	"github.com/ekaputra07/warren-go/api"
	"github.com/google/uuid"
)

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

type Snapshot struct {
	UUID      uuid.UUID `schema:"uuid"`
	SizeGB    int       `schema:"sizeGb"`
	CreatedAt string    `schema:"created_at"`
	DiskUUID  uuid.UUID `schema:"disk_uuid"`
}

type Disk struct {
	UUID             uuid.UUID       `schema:"uuid"`
	Status           string          `schema:"status"`
	Snapshots        []Snapshot      `schema:"snapshots"`
	UserID           int             `schema:"user_id"`
	BillingAccountID int             `schema:"billing_account_id"`
	SizeGB           int             `schema:"size_gb"`
	SourceImageType  SourceImageType `schema:"source_image_type"`
	SourceImage      string          `schema:"source_image"`
	CreatedAt        string          `schema:"created_at"`
	UpdatedAt        string          `schema:"updated_at"`
}

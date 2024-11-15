package ip

import (
	"github.com/ekaputra07/warren-go/api"
	"github.com/google/uuid"
)

type Client struct {
	API      *api.API
	Location string
}

type IPAddressInfo struct {
	ID                     int           `json:"id"`
	Address                string        `json:"address"`
	UserID                 int           `json:"user_id"`
	BillingAccountID       int           `json:"billing_account_id"`
	Type                   string        `json:"type"`
	NetworkID              uuid.NullUUID `json:"network_id"`
	Name                   string        `json:"name"`
	Enabled                bool          `json:"enabled"`
	CreatedAt              string        `json:"created_at"`
	UpdatedAt              string        `json:"updated_at"`
	IsDeleted              bool          `json:"is_deleted"`
	IsVirtual              bool          `json:"is_virtual"`
	AssignedTo             uuid.NullUUID `json:"assigned_to"`
	AssignedToResourceType string        `json:"assigned_to_resource_type"`
	AssignedToPrivateIP    string        `json:"assigned_to_private_ip"`
}

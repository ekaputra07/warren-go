package vpc

import (
	"github.com/ekaputra07/warren-go/api"
	"github.com/google/uuid"
)

type Client struct {
	API      *api.API
	Location string
}

type NetworkInfo struct {
	VLANID        int        `json:"vlan_id"`
	UUID          uuid.UUID  `json:"uuid"`
	Name          string     `json:"name"`
	Subnet        string     `json:"subnet"`
	SubnetIPV6    string     `json:"subnet_ipv6"`
	Type          string     `json:"type"`
	IsDefault     bool       `json:"is_default"`
	ResourceCount int        `json:"resources_count"`
	VMUUIDs       uuid.UUIDs `json:"vm_uuids"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
}

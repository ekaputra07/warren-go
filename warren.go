package warren

import (
	"github.com/ekaputra07/warren-go/api"
	"github.com/ekaputra07/warren-go/blockstorage"
	"github.com/ekaputra07/warren-go/ip"
	"github.com/ekaputra07/warren-go/location"
	"github.com/ekaputra07/warren-go/objectstorage"
	"github.com/ekaputra07/warren-go/vpc"
)

// Warren a single object to access all APIs
type Warren struct {
	Location      *location.Client
	ObjectStorage *objectstorage.Client
	BlockStorage  *blockstorage.Client
	VPC           *vpc.Client
	IP            *ip.Client
}

// Init initialize Warren with given API client
func Init(api *api.API, loc string) *Warren {
	return &Warren{
		Location:      location.NewClient(api),
		ObjectStorage: objectstorage.NewClient(api),
		BlockStorage:  blockstorage.NewClient(api),
		VPC:           vpc.NewClient(api, loc),
		IP:            ip.NewClient(api, loc),
	}
}

// New returns Warren that initialized with Default API client.
// Use this if you want to manage resources that doesn't require datacenter location such as:
// location, objectstorage, blockstorage
func New() *Warren {
	return Init(api.Default, "")
}

// New returns Warren that initialized with Default API client and specified location.
// Use this if you want to manage resources that require datacenter location such as:
// vpc, ip
func NewWithLocation(location string) *Warren {
	return Init(api.Default, location)
}

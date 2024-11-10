package warren

import (
	"github.com/ekaputra07/warren-go/api"
	"github.com/ekaputra07/warren-go/blockstorage"
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
}

// Init initialize Warren with given API client
func Init(api *api.API) *Warren {
	return &Warren{
		Location:      location.NewClient(api),
		ObjectStorage: objectstorage.NewClient(api),
		BlockStorage:  blockstorage.NewClient(api),
		VPC:           vpc.NewClient(api),
	}
}

// New returns Warren that initialized with Default API client
func New() *Warren {
	return Init(api.Default)
}

package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ekaputra07/warren-go/api"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	loc string    = "jkt01"
	id  uuid.UUID = uuid.MustParse("4e5eadd3-8b11-4c34-812a-2cf97120b628")
)

func TestListNetworks(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/networks", loc), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{API: a, Location: loc}
	vpc.ListNetworks(context.Background())
}

func TestGetNetwork(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network/%s", loc, id), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{API: a, Location: loc}
	vpc.GetNetwork(context.Background(), id)
}

func TestDeleteNetwork(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network/%s", loc, id), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{API: a, Location: loc}
	vpc.DeleteNetwork(context.Background(), id)
}

func TestRenameNetwork(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network/%s", loc, id), r.RequestURI)

		var data map[string]interface{}
		_ = json.NewDecoder(r.Body).Decode(&data)
		assert.Equal(t, "Test", data["name"])
	})
	defer s.Close()

	vpc := Client{API: a, Location: loc}
	vpc.RenameNetwork(context.Background(), id, "Test")
}

func TestGetOrCreateDefaultNetwork(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network?name=Default", loc), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{API: a, Location: loc}
	vpc.GetOrCreateDefaultNetwork(context.Background(), "Default")
}

func TestSetDefaultNetwork(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network/%s/default", loc, id), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{API: a, Location: loc}
	vpc.SetDefaultNetwork(context.Background(), id)
}

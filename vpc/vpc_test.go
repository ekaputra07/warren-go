package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	h "github.com/ekaputra07/idcloudhost-go/http"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	loc string    = "jkt01"
	id  uuid.UUID = uuid.MustParse("4e5eadd3-8b11-4c34-812a-2cf97120b628")
)

func TestListNetworks(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/networks", loc), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{H: c}
	vpc.ListNetworks(context.Background(), loc)
}

func TestGetNetwork(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network/%s", loc, id), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{H: c}
	vpc.GetNetwork(context.Background(), loc, id)
}

func TestDeleteNetwork(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network/%s", loc, id), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{H: c}
	vpc.DeleteNetwork(context.Background(), loc, id)
}

func TestRenameNetwork(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network/%s", loc, id), r.RequestURI)

		var data map[string]interface{}
		_ = json.NewDecoder(r.Body).Decode(&data)
		assert.Equal(t, "Test", data["name"])
	})
	defer s.Close()

	vpc := Client{H: c}
	vpc.RenameNetwork(context.Background(), loc, id, "Test")
}

func TestGetOrCreateDefaultNetwork(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network?name=Default", loc), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{H: c}
	vpc.GetOrCreateDefaultNetwork(context.Background(), loc, "Default")
}

func TestSetDefaultNetwork(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/network/%s/default", loc, id), r.RequestURI)
	})
	defer s.Close()

	vpc := Client{H: c}
	vpc.SetDefaultNetwork(context.Background(), loc, id)
}

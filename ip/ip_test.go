package ip

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
	loc     string    = "jkt01"
	address string    = "1.2.3.4"
	vmUUID  uuid.UUID = uuid.MustParse("4e5eadd3-8b11-4c34-812a-2cf97120b628")
)

func TestListFloatingIPs(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/ip_addresses", loc), r.RequestURI)
	})
	defer s.Close()

	ip := Client{API: a, Location: loc}
	ip.ListFloatingIPs(context.Background())
}

func TestCreateFloatingIP(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/ip_addresses", loc), r.RequestURI)

		var data map[string]any
		_ = json.NewDecoder(r.Body).Decode(&data)
		assert.Equal(t, "Test", data["name"])
		assert.Equal(t, float64(123), data["billing_account_id"])
	})
	defer s.Close()

	ip := Client{API: a, Location: loc}
	info := IPAddressInfo{Name: "Test"}

	// BillingAccountID not set
	err := ip.CreateFloatingIP(context.Background(), &info)
	assert.Error(t, err)

	// Success
	info.BillingAccountID = 123
	ip.CreateFloatingIP(context.Background(), &info)
}

func TestGetFloatingIP(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/ip_addresses/%s", loc, address), r.RequestURI)
	})
	defer s.Close()

	ip := Client{API: a, Location: loc}
	ip.GetFloatingIP(context.Background(), address)
}

func TestUpdateFloatingIP(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/ip_addresses/%s", loc, address), r.RequestURI)

		var data map[string]any
		_ = json.NewDecoder(r.Body).Decode(&data)
		assert.Equal(t, "Test", data["name"])
		assert.Equal(t, float64(123), data["billing_account_id"])
	})
	defer s.Close()

	ip := Client{API: a, Location: loc}
	info := IPAddressInfo{Name: "Test", Address: address}

	// BillingAccountID not set
	err := ip.UpdateFloatingIP(context.Background(), info)
	assert.Error(t, err)

	// Success
	info.BillingAccountID = 123
	ip.UpdateFloatingIP(context.Background(), info)
}

func TestDeleteFloatingIP(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/ip_addresses/%s", loc, address), r.RequestURI)
	})
	defer s.Close()

	ip := Client{API: a, Location: loc}
	ip.DeleteFloatingIP(context.Background(), address)
}

func TestAssignFloatingIPToVM(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/ip_addresses/%s/assign", loc, address), r.RequestURI)

		var data map[string]any
		_ = json.NewDecoder(r.Body).Decode(&data)
		assert.Equal(t, vmUUID.String(), data["vm_uuid"])
	})
	defer s.Close()

	ip := Client{API: a, Location: loc}
	ip.AssignFloatingIPToVM(context.Background(), address, vmUUID)
}

func TestUnassignFloatingIPFromVM(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, fmt.Sprintf("/v1/%s/network/ip_addresses/%s/unassign", loc, address), r.RequestURI)

		var data map[string]any
		_ = json.NewDecoder(r.Body).Decode(&data)
		assert.Equal(t, vmUUID.String(), data["vm_uuid"])
	})
	defer s.Close()

	ip := Client{API: a, Location: loc}
	ip.UnassignFloatingIPFromVM(context.Background(), address, vmUUID)
}

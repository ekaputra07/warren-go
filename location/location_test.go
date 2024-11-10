package location

import (
	"context"
	"net/http"
	"testing"

	"github.com/ekaputra07/warren-go/api"
	"github.com/stretchr/testify/assert"
)

func TestListLocations(t *testing.T) {
	a, s := api.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/config/locations", r.RequestURI)
	})
	defer s.Close()

	lc := Client{API: a}
	lc.ListLocations(context.Background())
}

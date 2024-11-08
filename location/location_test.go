package location

import (
	"context"
	"net/http"
	"testing"

	h "github.com/ekaputra07/idcloudhost-go/http"
	"github.com/stretchr/testify/assert"
)

func TestListLocations(t *testing.T) {
	c, s := h.MockClientServer(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/config/locations", r.RequestURI)
	})
	defer s.Close()

	lc := Client{H: c}
	lc.ListLocations(context.Background())
}

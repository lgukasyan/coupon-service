package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerStart(t *testing.T) {
	// Start server
	r, err := Start()
	assert.NoError(t, err, "error starting the server")

	// Create server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Send test
	resp, err := http.Get(ts.URL + "/api/coupon")
	assert.NoError(t, err, "error while sending a GET request")

	// Close at the end
	defer resp.Body.Close()

	// Check if success
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
}

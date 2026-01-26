package spaces

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCDNResponse(t *testing.T) {
	f, err := os.ReadFile("do_endpoints.json")
	require.NoError(t, err)

	var resp CDNResponse
	err = json.Unmarshal(f, &resp)
	require.NoError(t, err)

	assert.Equal(t, 4, len(resp.Endpoints))
}

func TestCDNPurge(t *testing.T) {
	t.Skip("requires api call")

	c := Config{
		CDNEndpointID: "385158b2-750d-4f71-9fa6-3eddd1b81dd0",
		Token:         "dop...",
	}

	// https://static.free-entertainment-guide.com/discover.json
	err := c.Purge("discover.json")
	require.NoError(t, err)
	// this just logs a message. If successful body is empty
	// assert.Equal(t, 1, 2)
}

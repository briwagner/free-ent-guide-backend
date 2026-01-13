package moviedb

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

var c Cred

type Cred struct {
	Tms     string
	Moviedb string
}

func (c *Cred) Setup() {
	yamlFile, err := os.ReadFile("../../creds.yaml")
	if err != nil {
		log.Fatalf("File not loading: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("No credentials %v", err)
	}
}

func TestGetTrending(t *testing.T) {
	c.Setup()
	mdb := MovieDb{Key: c.Moviedb}
	mdb.GetTrending(t.Context())

	if mdb.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", mdb.Status, http.StatusOK)
	}
}

func TestGetDiscover(t *testing.T) {
	mdb := MovieDb{Key: c.Moviedb}
	dateParam := time.Now().Format("2006-01-02")
	mdb.GetDiscover(t.Context(), dateParam, "1")

	if mdb.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", mdb.Status, http.StatusOK)
	}
}

func TestGetToken(t *testing.T) {
	mdb := MovieDb{Key: c.Moviedb}
	mdb.getToken(t.Context())

	if mdb.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", mdb.Status, http.StatusOK)
	}
	if mdb.Token.Success != true {
		t.Errorf("Token returned wrong status. Got %v, Wanted %v", mdb.Token.Success, true)
	}
}

func TestDiscover(t *testing.T) {
	mdb := NewMovieDB("secret")

	data, err := os.ReadFile("./testdata/discover.json")
	require.NoError(t, err)

	var results DiscoverResults
	err = json.Unmarshal(data, &results)
	require.NoError(t, err)

	assert.Len(t, results.Results, 20)
	assert.Equal(t, "The Mouse Trap", results.Results[0].Title)
	assert.Equal(t, "Aug. 23, 2024", results.Results[0].ReleaseDate.String())
	assert.Equal(t, int(1225377), results.Results[0].ID)
	assert.Equal(t, false, results.Results[0].Video)
	assert.Equal(t, "It's Alex's 21st Birthday, but she's stuck at the amusement arcade on a late shift so her friends decide to surprise her, but a masked killer dressed as Mickey Mouse decides to play a game of his own with them which she must survive.", results.Results[0].Description)
	assert.Equal(t, "/3ovFaFeojLFIl5ClqhtgYMDS8sE.jpg", results.Results[0].Poster)

	ti, err := time.Parse("2006-01-02", "2024-08-16")
	require.NoError(t, err)
	filtered := mdb.filterReleases(ti, results.Results)
	assert.Len(t, filtered, 8)

	limit := ti.Add(time.Hour * 24 * 30)
	SortByDate(filtered)
	assert.Equal(t, int(5492), filtered[0].ID)
	assert.Equal(t, "Aug. 16, 2024", filtered[0].ReleaseDate.String())
	last := filtered[len(filtered)-1]
	assert.True(t, last.ReleaseDate.Before(limit))
}

func TestDiscoverPaging(t *testing.T) {
	t.Skip("makes api call")

	mdb := MovieDb{Key: c.Moviedb}
	res, err := mdb.GetDiscoverPaged(t.Context(), "2024-08-19")
	require.NoError(t, err)
	assert.Equal(t, 313, len(res))
}

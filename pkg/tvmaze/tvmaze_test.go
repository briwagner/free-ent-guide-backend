package tvmaze

import (
	"net/http"
	"testing"
)

func TestGetSearch(t *testing.T) {
	tmz := TvMaze{}
	tmz.GetSearch("Riverdale")

	if tmz.Status != http.StatusOK {
		t.Errorf("Fetch return wrong status. Got %v, Wanted %v", tmz.Status, http.StatusOK)
	}
}

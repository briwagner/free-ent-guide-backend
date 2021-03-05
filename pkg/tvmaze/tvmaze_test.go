package tvmaze

import (
	"net/http"
	"testing"
)

func TestGetSearch(t *testing.T) {
	tmz := TvMaze{}
	s, _ := tmz.GetSearch("Riverdale")

	if s != http.StatusOK {
		t.Errorf("Fetch return wrong status. Got %v, Wanted %v", s, http.StatusOK)
	}
}

package tvmaze

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTvMaze_Get(t *testing.T) {
	tvm := TvMaze{Client: &http.Client{}}
	tvm.GetSearch(t.Context(), "Riverdale")
	assert.Equal(t, http.StatusOK, tvm.Status)
	assert.Contains(t, string(tvm.Response), `"url":"https://www.tvmaze.com/shows/5262/riverdale"`)

	tvm2 := TvMaze{Client: &http.Client{}}
	tvm2.GetShow(t.Context(), int64(60398))
	assert.Equal(t, http.StatusOK, tvm2.Status)
	assert.Contains(t, string(tvm2.Response), `"url":"https://www.tvmaze.com/shows/60398/sunny"`)
}

package tvmaze

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTvMaze_Get(t *testing.T) {
	tvm := TvMaze{}
	tvm.GetSearch("Riverdale")
	assert.Equal(t, http.StatusOK, tvm.Status)
	assert.Contains(t, string(tvm.Response), `"url":"https://www.tvmaze.com/shows/5262/riverdale"`)

	tvm2 := TvMaze{}
	tvm2.GetShow(int64(60398))
	assert.Equal(t, http.StatusOK, tvm2.Status)
	assert.Contains(t, string(tvm2.Response), `"url":"https://www.tvmaze.com/shows/60398/sunny"`)
}

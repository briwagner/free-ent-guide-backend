package tvmaze

import (
	"context"
	"fmt"
	"free-ent-guide-backend/pkg/bri_otel"
	"io"
	"log"
	"net/http"
)

// TvMaze is a wrapper for requests to Tv Maze api.
type TvMaze struct {
	Status   int
	Response []byte
	Client   *http.Client
}

const base = "http://api.tvmaze.com/"

func NewTVMaze() TvMaze {
	return TvMaze{Client: bri_otel.NewOtelClient(5)}
}

// GetSearch makes the api call to find a show by title.
func (t *TvMaze) GetSearch(ctx context.Context, title string) error {
	url := base + "search/shows"
	params := make(map[string]string, 0)
	params["q"] = title
	return t.fetch(ctx, url, params)
}

func (t *TvMaze) GetShow(ctx context.Context, showID int64) error {
	url := fmt.Sprintf("%sshows/%d", base, showID)
	params := make(map[string]string, 0)
	return t.fetch(ctx, url, params)
}

func (t *TvMaze) GetEpisode(ctx context.Context, id int64) error {
	url := fmt.Sprintf("%sepisodes/%d", base, id)
	params := make(map[string]string, 0)
	return t.fetch(ctx, url, params)
}

// fetch is a wrapper for api calls.
func (t *TvMaze) fetch(ctx context.Context, url string, params map[string]string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("Cannot build URL for TV Maze. %v", err)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	resp, doErr := t.Client.Do(req)
	if doErr != nil {
		log.Printf("Failed to make request to TV Maze %v", doErr.Error())
		t.Status = 500
		t.Response = []byte("Server error")
		return doErr
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Status = 500
		t.Response = []byte("Cannot parse body.")
		return err
	}

	t.Status = resp.StatusCode
	t.Response = body
	return nil
}

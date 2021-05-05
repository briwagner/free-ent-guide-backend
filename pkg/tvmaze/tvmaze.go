package tvmaze

import (
	"io/ioutil"
	"log"
	"net/http"
)

// TvMaze is a wrapper for requests to Tv Maze api.
type TvMaze struct {
	Status   int
	Response []byte
}

const base = "http://api.tvmaze.com/"

// GetSearch makes the api call to find a show by title.
func (t *TvMaze) GetSearch(title string) {
	url := base + "search/shows"
	params := make(map[string]string)
	params["q"] = title
	// Todo: should this be returned or pointer is good enough?
	_ = t.fetch(url, params)
}

// fetch is a wrapper for api calls.
func (t *TvMaze) fetch(url string, params map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Cannot build URL for TV Maze. %v", err)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, doErr := client.Do(req)

	if doErr != nil {
		log.Printf("Failed to make request to TV Maze %v", doErr.Error())
		t.Status = 500
		t.Response = []byte("Server error")
		return doErr
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		t.Status = 500
		t.Response = []byte("Cannot parse body.")
		return err
	}

	t.Status = resp.StatusCode
	t.Response = body
	return nil
}

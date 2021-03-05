package tvmaze

import (
	"io/ioutil"
	"log"
	"net/http"
)

type TvMaze struct{}

var base = "http://api.tvmaze.com/"

func (t TvMaze) GetSearch(title string) (int, []byte) {
	url := base + "search/shows"
	params := make(map[string]string)
	params["q"] = title
	return t.Fetch(url, params)
}

func (t TvMaze) Fetch(url string, params map[string]string) (int, []byte) {
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
		log.Fatalf("Failed to make request to TV Maze %v", doErr)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return 0, []byte("Cannot parse body.")
	}

	return resp.StatusCode, body
}

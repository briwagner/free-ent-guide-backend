package tmsapi

import (
	"io/ioutil"
	"log"
	"net/http"
)

const base = "http://data.tmsapi.com/v1.1/"
const lineup = "USA-TX42500-X"

type TmsApi struct {
	Key      string
	Status   int
	Response []byte
}

// GetCinema returns listings for movies showing by date, zip.
func (t *TmsApi) GetCinema(zip string, date string) {
	url := base + "movies/showings"
	params := make(map[string]string)
	params["zip"] = zip
	params["startDate"] = date

	_ = t.fetch(url, params)
}

// GetTvMovies returns listings for tv movies by date, lineup.
func (t *TmsApi) GetTvMovies(date string, lineup string) {
	url := base + "movies/airings"
	params := make(map[string]string)
	params["startDateTime"] = date
	params["lineupId"] = lineup

	_ = t.fetch(url, params)
}

// GetTvSports returns listings for tv sports by date, lineup.
func (t *TmsApi) GetTvSports(date string, lineup string) {
	url := base + "sports/all/events/airings"
	params := make(map[string]string)
	params["startDateTime"] = date
	params["lineupId"] = lineup

	_ = t.fetch(url, params)
}

// fetch is a wrapper for api calls.
func (t *TmsApi) fetch(url string, p map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Cannot build URL for TmsApi. %v", err)
		t.Status = 500
		return err
	}

	q := req.URL.Query()
	q.Add("api_key", t.Key)
	for k, v := range p {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, doErr := client.Do(req)

	if doErr != nil {
		log.Printf("Failed to make request to TmsApi %v", doErr.Error())
		t.Status = 500
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

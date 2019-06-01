package tmsapi

import (
	"io/ioutil"
	"log"
	"net/http"
)

type TmsApi struct {
	Key string
}

func (t TmsApi) GetCinema(zip string, date string) (int, []byte) {
	url := "http://data.tmsapi.com/v1.1/movies/showings"
	params := make(map[string]string)
	params["zip"] = zip
	params["startDate"] = date

	return t.Fetch(url, params)
}

func (t TmsApi) GetTvMovies(date string) (int, []byte) {
	url := "http://data.tmsapi.com/v1.1/movies/airings"
	params := make(map[string]string)
	params["startDate"] = date
	params["lineupId"] = "USA-TX42500-X"

	return t.Fetch(url, params)
}

func (t TmsApi) GetTvSports(date string) (int, []byte) {
	url := "http://data.tmsapi.com/v1.1/sports/all/events/airings"
	params := make(map[string]string)
	params["startDate"] = date
	params["lineupId"] = "USA-TX42500-X"

	return t.Fetch(url, params)
}

func (t TmsApi) Fetch(url string, p map[string]string) (int, []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Cannot build URL for TmsApi. %v", err)
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
		log.Fatalf("Failed to make request to TmsApi %v", doErr)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return 0, []byte("Cannot parse body.")
	}

	return resp.StatusCode, body
}

package tmsapi

import (
	"context"
	"free-ent-guide-backend/pkg/bri_otel"
	"io"
	"log"
	"net/http"
)

const base = "http://data.tmsapi.com/v1.1/"
const lineup = "USA-TX42500-X"

type TmsApi struct {
	Key      string
	Status   int
	Response []byte
	Client   *http.Client
}

func NewTMSApi(key string) TmsApi {
	return TmsApi{
		Key:    key,
		Client: bri_otel.NewOtelClient(5),
	}
}

// GetCinema returns listings for movies showing by date, zip.
func (t *TmsApi) GetCinema(ctx context.Context, zip string, date string) error {
	url := base + "movies/showings"
	params := make(map[string]string)
	params["zip"] = zip
	params["startDate"] = date

	return t.fetch(ctx, url, params)
}

// GetTvMovies returns listings for tv movies by date, lineup.
func (t *TmsApi) GetTvMovies(ctx context.Context, date string, lineup string) error {
	url := base + "movies/airings"
	params := make(map[string]string)
	params["startDateTime"] = date
	params["lineupId"] = lineup

	return t.fetch(ctx, url, params)
}

// GetTvSports returns listings for tv sports by date, lineup.
func (t *TmsApi) GetTvSports(ctx context.Context, date string, lineup string) error {
	url := base + "sports/all/events/airings"
	params := make(map[string]string)
	params["startDateTime"] = date
	params["lineupId"] = lineup

	return t.fetch(ctx, url, params)
}

// fetch is a wrapper for api calls.
func (t *TmsApi) fetch(ctx context.Context, url string, p map[string]string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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
	resp, doErr := t.Client.Do(req)

	if doErr != nil {
		log.Printf("Failed to make request to TmsApi %v", doErr.Error())
		t.Status = 500
		return doErr
	}

	body, err := io.ReadAll(resp.Body)
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

package moviedb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"
)

const base = "https://api.themoviedb.org/3/"

type MovieDb struct {
	Key      string
	Status   int
	Response []byte
	Token    Token
}

type Token struct {
	Success bool   `json:"success"`
	Expires string `json:"expires_at"`
	Token   string `json:"request_token"`
}

// GetTrending retrieves the trending TV listings.
func (m *MovieDb) GetTrending() {
	url := base + "trending/tv/week"
	p := map[string]string{
		"api_key": m.Key,
	}
	_ = m.fetch(url, p)
}

func (m *MovieDb) GetDiscoverPaged(date string) ([]MovieRelease, error) {
	var mr []MovieRelease
	p := 1
	totalP := 2

	for p <= totalP {
		m.GetDiscover(date, strconv.Itoa(p))
		if m.Status != 200 {
			return mr, fmt.Errorf("error getting discover results; status %d", m.Status)
		}

		var results DiscoverResults
		err := json.Unmarshal(m.Response, &results)
		if err != nil {
			return mr, err
		}

		if p != results.Page {
			return mr, fmt.Errorf("wrong page, asked for %d, got %d", p, results.Page)
		}

		if results.Page == 0 || results.TotalPages == 0 {
			return mr, errors.New("something wrong. No page")
		}

		// Update values
		p++
		totalP = results.TotalPages
		mr = append(mr, results.Results...)
	}

	ti, err := time.Parse("2006-01-02", date)
	if err != nil {
		return mr, err
	}

	filtered := FilterReleases(ti, mr)
	SortByDate(filtered)

	return filtered, nil
}

// GetDiscover retrieves the coming movie listings.
func (m *MovieDb) GetDiscover(date string, p string) {
	url := base + "discover/movie"
	// from discoverMoviesReq(), must include date and other filters
	// in query param
	params := map[string]string{
		"api_key":                  m.Key,
		"primary_release_date.gte": date,
		"adult":                    "false",
		"page":                     p,
		// todo: language filter was added because Asian
		// languages cannot be entered into DB as utf-8.
		// Try to convert text before pushing to DB?
		"with_original_language": "en",
	}
	_ = m.fetch(url, params)
}

// GetToken retrieves a token.
// TODO unused except in test.
func (m *MovieDb) GetToken() error {
	url := base + "authentication/token/new"
	p := map[string]string{
		"api_key": m.Key,
	}
	err := m.fetch(url, p)

	if err != nil {
		return err
	}

	var token Token
	json.Unmarshal([]byte(m.Response), &token)
	m.Token = token
	return nil
}

// fetch is a wrapper for api calls.
func (m *MovieDb) fetch(url string, params map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Cannot build URL for MovieDB. %v", err)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	// q.Add("api_key", m.Key)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, doErr := client.Do(req)

	if doErr != nil {
		log.Printf("Failed to make request to MovieDB %v", doErr.Error())
		m.Status = 500
		m.Response = []byte("Server error")
		return doErr
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		m.Status = 500
		m.Response = []byte("Cannot parse body.")
		return err
	}

	m.Status = resp.StatusCode
	m.Response = body
	return nil
}

type MovieRelease struct {
	ID               int         `json:"id"`
	Title            string      `json:"title"`
	ReleaseDate      ReleaseDate `json:"release_date"`
	Description      string      `json:"overview"`
	Poster           string      `json:"poster_path"`
	Video            bool        `json:"video"`
	OriginalLanguage string      `json:"original_language"`
	Popularity       float64     `json:"popularity"`
}

func SortByDate(rel []MovieRelease) {
	slices.SortFunc(rel, func(a, b MovieRelease) int {
		return a.ReleaseDate.Time.Compare(b.ReleaseDate.Time)
	})
}

func FilterReleases(d time.Time, rel []MovieRelease) []MovieRelease {
	var filtered []MovieRelease

	limit := d.Add(time.Hour * 24 * 30)
	for _, r := range rel {
		// This list is really big. How to control?
		if r.ReleaseDate.Time.Before(limit) && r.Popularity > float64(50) {
			filtered = append(filtered, r)
		}
	}

	return filtered
}

type DiscoverResults struct {
	Page         int            `json:"page"`
	Results      []MovieRelease `json:"results"`
	TotalPages   int            `json:"total_pages"`
	TotalResults int            `json:"total_results"`
}

type ReleaseDate struct {
	time.Time
}

func (rd *ReleaseDate) UnmarshalJSON(b []byte) error {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	rd.Time = date
	return nil
}

func (rd ReleaseDate) String() string {
	return rd.Time.Format("Jan. 2, 2006")
}

package moviedb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/pkg/bri_otel"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"
)

const Base = "https://api.themoviedb.org/3/"

type MovieDb struct {
	Key        string
	popularity float64 // filter movies in Discover
	Client     *http.Client
	// TODO add a logger here!!

	Status   int
	Response []byte
	Token    Token
}

type Token struct {
	Success bool   `json:"success"`
	Expires string `json:"expires_at"`
	Token   string `json:"request_token"`
}

func NewMovieDB(key string) *MovieDb {
	return &MovieDb{
		Key:        key,
		popularity: float64(10),
		Client:     bri_otel.NewOtelClient(5),
	}
}

// GetTrending retrieves the trending TV listings.
func (m *MovieDb) GetTrending(ctx context.Context) {
	url := Base + "trending/tv/week"
	p := map[string]string{
		"api_key": m.Key,
	}
	_ = m.fetch(ctx, url, p)
}

func (m *MovieDb) GetDiscoverPaged(ctx context.Context, date string) ([]MovieRelease, error) {
	var mr []MovieRelease
	p := 1
	totalP := 2

	for p <= totalP {
		m.GetDiscover(ctx, date, strconv.Itoa(p))
		if m.Status != 200 {
			return mr, fmt.Errorf("error getting DiscoverResults; status %d", m.Status)
		}

		var results DiscoverResults
		err := json.Unmarshal(m.Response, &results)
		if err != nil {
			// TODO write this to file in /tmp ? to view what's bad in here?
			return mr, fmt.Errorf("error unmarshal DiscoverResults: %w", err)
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
		return mr, fmt.Errorf("error parsing time for %s: %w", date, err)
	}

	filtered := m.filterReleases(ti, mr)
	SortByDate(filtered)

	return filtered, nil
}

// GetDiscover retrieves the coming movie listings.
func (m *MovieDb) GetDiscover(ctx context.Context, date string, p string) {
	url := Base + "discover/movie"
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
	_ = m.fetch(ctx, url, params)
}

// getToken retrieves a token.
// TODO unused except in test.
func (m *MovieDb) getToken(ctx context.Context) error {
	url := Base + "authentication/token/new"
	p := map[string]string{
		"api_key": m.Key,
	}
	err := m.fetch(ctx, url, p)

	if err != nil {
		return err
	}

	var token Token
	json.Unmarshal([]byte(m.Response), &token)
	m.Token = token
	return nil
}

// fetch is a wrapper for api calls.
func (m *MovieDb) fetch(ctx context.Context, url string, params map[string]string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("Cannot build URL for MovieDB. %v", err)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	// q.Add("api_key", m.Key)
	req.URL.RawQuery = q.Encode()
	resp, doErr := m.Client.Do(req)

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

func (m *MovieDb) filterReleases(d time.Time, rel []MovieRelease) []MovieRelease {
	var filtered []MovieRelease

	limit := d.Add(time.Hour * 24 * 30)
	for _, r := range rel {
		// This list is really big. How to control?
		if r.ReleaseDate.Time.Before(limit) && r.Popularity > m.popularity {
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
		rd.Time = time.Now()
		log.Printf("error parsing date DiscoverMovies: %s", err.Error())
	} else {
		rd.Time = date
	}
	return nil
}

func (rd ReleaseDate) String() string {
	return rd.Time.Format("Jan. 2, 2006")
}

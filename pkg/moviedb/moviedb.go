package moviedb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

// GetDiscover retrieves the coming movie listings.
func (m *MovieDb) GetDiscover(date string) {
	url := base + "discover/movie"
	// from discoverMoviesReq(), must include date and other filters
	// in query param
	params := map[string]string{
		"api_key":                  m.Key,
		"primary_release_date.gte": date,
		"adult":                    "false",
		// todo: language filter was added because Asian
		// languages cannot be entered into DB as utf-8.
		// Try to convert text before pushing to DB?
		"with_original_language": "en",
	}
	_ = m.fetch(url, params)
}

// GetToken retrieves a token.
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

	body, err := ioutil.ReadAll(resp.Body)
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

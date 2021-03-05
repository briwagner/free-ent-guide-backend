package moviedb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type MovieDb struct {
	Key string
}

type Token struct {
	Success bool   `json:"success"`
	Expires string `json:"expires_at"`
	Token   string `json:"request_token"`
}

func (m MovieDb) GetTrending() (int, []byte) {
	url := "https://api.themoviedb.org/3/trending/tv/week"
	return m.Fetch(url)
}

func (m MovieDb) GetDiscover() (int, []byte) {
	url := "https://api.themoviedb.org/3/discover/movie"
	return m.Fetch(url)
}

func (m MovieDb) GetToken() (int, Token) {
	url := "https://api.themoviedb.org/3/authentication/token/new"
	s, b := m.Fetch(url)

	var token Token
	json.Unmarshal([]byte(b), &token)
	return s, token
}

func (m MovieDb) Fetch(url string) (int, []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Cannot build URL for MovieDB. %v", err)
	}

	q := req.URL.Query()
	q.Add("api_key", m.Key)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, doErr := client.Do(req)

	if doErr != nil {
		log.Fatalf("Failed to make request to MovieDB %v", doErr)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return 0, []byte("Cannot parse body.")
	}

	return resp.StatusCode, body
}

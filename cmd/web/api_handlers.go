package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetMovies handler for movies by zip.
// Responds to /v1/movies?zip={ZIP}.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	zip := r.URL.Query().Get("zip")
	if zip == "" {
		w.Write([]byte("Must pass a valid zip code"))
		return
	}

	params := make(map[string]string)
	params["zip"] = zip

	// Set timezone to avoid using UTC on server.
	zone, err := time.LoadLocation(c.Timezone)
	if err != nil {
		log.Printf("Cannot load timezone %e", err)
	}
	tt := time.Now().In(zone).Format("2006-01-02")
	log.Printf("Cinema request zip: %s, on: %v", zip, tt)

	params["startDate"] = tt
	w.Write([]byte(GetTMSReq(params, "movies/showings")))
}

func discoverMoviesReq(date string) (string, error) {
	url := "https://api.themoviedb.org/3/discover/movie"

	req, _ := http.NewRequest("GET", url, nil)

	var dateParam string

	q := req.URL.Query()
	q.Add("api_key", c.Moviedb)
	if date != "" {
		dateParam = date
	} else {
		dateParam = time.Now().Format("2006-01-02")
	}
	q.Add("primary_release_date.gte", dateParam)
	q.Add("adult", "false")

	req.URL.RawQuery = q.Encode()

	client := http.Client{Timeout: time.Second * 2}
	resp, doErr := client.Do(req)

	if doErr != nil {
		fmt.Println(doErr)
		return "Failed to make the discover request from TMS.", doErr
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "Failed to parse body", err
	}
	return string(body), nil
}

// DiscoverMovies handler for coming-soon movies.
// Responds to /v1/discover?date={YYYY-MM-DD}.
func DiscoverMovies(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	cacheKey := "discover"
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(406)
		w.Write([]byte("Must pass a date"))
		return
	}
	// Check cache for stored response.
	cache, err := getCache(cacheKey)

	// No cache found.
	if err != nil {
		cacheStatus(&w, false)

		req, getErr := discoverMoviesReq(date)
		if getErr != nil {
			w.WriteHeader(500)
		} else {
			_ = setCache(cacheKey, req, time.Hour)
		}
		w.Write([]byte(req))
	} else {
		// Cache found.
		cacheStatus(&w, true)
		w.Write([]byte(cache))
	}

}

func getTvSearchReq(query string) string {
	url := "http://api.tvmaze.com/search/shows"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return "Could not build http request"
	}

	q := req.URL.Query()
	q.Add("q", query)

	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, doErr := client.Do(req)

	if doErr != nil {
		fmt.Println(doErr)
		return "Failed to make the request"
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return "Failed to parse body"
	}

	return string(body)
}

// GetTMSReq is general getter for API calls to TMS service.
func GetTMSReq(params map[string]string, loc string) string {
	url := "http://data.tmsapi.com/v1.1/" + loc

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return "Could not build http request"
	}

	q := req.URL.Query()
	q.Add("api_key", c.Tms)
	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, doErr := client.Do(req)

	if doErr != nil {
		fmt.Println(doErr)
		return "Failed to make the discover request from TMS."
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return "Failed to parse body"
	}

	if len(body) == 0 {
		// During Covid, some requests return no showings.
		log.Printf("No programs found: %s", loc)
		return ""
	}
	return string(body)
}

// GetTvMovies responds to /v1/tv-movies?date={YYYY-MM-DD}.
func GetTvMovies(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	cacheKey := "tvmovies"
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(406)
		w.Write([]byte("Must pass a date"))
		return
	}
	cache, err := getCache(cacheKey)
	if err != nil {
		cacheStatus(&w, false)

		params := make(map[string]string)
		params["startDateTime"] = date
		params["lineupId"] = "USA-TX42500-X"
		req := GetTMSReq(params, "movies/airings")
		w.Write([]byte(req))
		_ = setCache(cacheKey, req, time.Hour)
	} else {
		cacheStatus(&w, true)
		w.Write([]byte(cache))
	}
}

// GetTvSports responds to /v1/tv-sports?date={YYYY-MM-DD}.
func GetTvSports(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(406)
		w.Write([]byte("Must pass a date"))
		return
	}

	cacheKey := "tvsports?" + date
	cache, err := getCache(cacheKey)
	if err != nil {
		cacheStatus(&w, false)

		params := make(map[string]string)
		params["startDateTime"] = date
		params["lineupId"] = "USA-TX42500-X"
		req := GetTMSReq(params, "sports/all/events/airings")
		w.Write([]byte(req))
		_ = setCache(cacheKey, req, time.Hour)
	} else {
		cacheStatus(&w, true)
		w.Write([]byte(cache))
	}
}

// GetTvSearch responds to /v1/tv-search?title={TITLE}.
func GetTvSearch(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	title := r.URL.Query().Get("title")
	if title == "" {
		w.WriteHeader(406)
		w.Write([]byte("Must pass a title to search"))
		return
	}
	w.Write([]byte(getTvSearchReq(title)))
}

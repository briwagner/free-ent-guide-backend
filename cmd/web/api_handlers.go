package main

import (
	"fmt"
	"free-ent-guide-backend/pkg/moviedb"
	"free-ent-guide-backend/pkg/tvmaze"
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

		mdb := moviedb.MovieDb{}
		mdb.GetDiscover(date)

		if mdb.Status == http.StatusOK {
			_ = setCache(cacheKey, string(mdb.Response), time.Hour)
		} else {
			w.WriteHeader(500)
		}
		w.Write(mdb.Response)
	} else {
		// Cache found.
		cacheStatus(&w, true)
		w.Write([]byte(cache))
	}
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

	tmz := tvmaze.TvMaze{}
	tmz.GetSearch(title)

	w.WriteHeader(tmz.Status)
	w.Write(tmz.Response)
}

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	yaml "gopkg.in/yaml.v2"
)

var c Cred
var tmsApi string
var movieDb string
var port string
var env string
var today = time.Now()

func main() {
	c.getCreds()
	tmsApi = c.Tms
	movieDb = c.Moviedb
	port = fmt.Sprintf(":%v", c.Port)

	fmt.Printf("ENT API is live. Listening on port %v ...\n", port)
	http.HandleFunc("/v1/movies", GetMovies)
	http.HandleFunc("/v1/discover", DiscoverMovies)
	http.HandleFunc("/v1/tv-movies", GetTvMovies)
	http.HandleFunc("/v1/tv-sports", GetTvSports)
	http.HandleFunc("/v1/tv-search", GetTvSearch)
	http.ListenAndServe(port, nil)
}

type Cred struct {
	Tms     string
	Moviedb string
	Port    int
	Env     string
}

func (c *Cred) getCreds() *Cred {
	yamlFile, err := ioutil.ReadFile("creds.yaml")
	if err != nil {
		log.Printf("File not loading: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("No credentials %v", err)
	}
	return c
}

// Response to /v1/movies?zip={ZIP}.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		enableCors(&w)
		if zip, ok := r.URL.Query()["zip"]; ok {
			params := make(map[string]string)
			params["zip"] = zip[0]
			params["startDate"] = time.Now().Format("2006-01-02")
			w.Write([]byte(GetTMSReq(params, "movies/showings")))
		} else {
			w.Write([]byte("Must pass a valid zip code"))
		}
	default:
		w.Write([]byte("Only GET requests are accepted."))
	}
}

func discoverMoviesReq(date string) (string, error) {
	url := "https://api.themoviedb.org/3/discover/movie"

	req, _ := http.NewRequest("GET", url, nil)

	var dateParam string

	q := req.URL.Query()
	q.Add("api_key", movieDb)
	if date != "" {
		dateParam = date
	} else {
		dateParam = today.Format("2006-01-02")
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

// Response to /v1/discover?date={YYYY-MM-DD}.
func DiscoverMovies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		enableCors(&w)
		cacheKey := "discover"
		if date, ok := r.URL.Query()["date"]; !ok {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a date"))
		} else {
			cache, err := getCache(cacheKey)
			if err != nil {
				cacheStatus(&w, false)

				req, getErr := discoverMoviesReq(date[0])
				if getErr != nil {
					w.WriteHeader(500)
				} else {
					setCache(cacheKey, req)
				}
				w.Write([]byte(req))
			} else {
				cacheStatus(&w, true)
				w.Write([]byte(cache))
			}

		}
	default:
		w.WriteHeader(403)
		w.Write([]byte("Only GET requests are accepted"))
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

func GetTMSReq(params map[string]string, loc string) string {
	url := "http://data.tmsapi.com/v1.1/" + loc

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return "Could not build http request"
	}

	q := req.URL.Query()
	q.Add("api_key", tmsApi)
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
		return "No programs found"
	}
	return string(body)
}

// Response to /v1/tv-movies?date={YYYY-MM-DD}.
func GetTvMovies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		enableCors(&w)
		cacheKey := "tvmovies"
		if date, ok := r.URL.Query()["date"]; ok {
			cache, err := getCache(cacheKey)
			if err != nil {
				cacheStatus(&w, false)

				params := make(map[string]string)
				params["startDate"] = date[0]
				params["lineupId"] = "USA-TX42500-X"
				req := GetTMSReq(params, "movies/airings")
				w.Write([]byte(req))
				setCache(cacheKey, req)
			} else {
				cacheStatus(&w, true)
				w.Write([]byte(cache))
			}
		} else {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a date"))
		}
	default:
		w.Write([]byte("Only GET requests are accepted"))
	}
}

// Response to /v1/tv-sports?date={YYYY-MM-DD}.
func GetTvSports(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		enableCors(&w)
		if date, ok := r.URL.Query()["date"]; ok {
			cacheKey := "tvsports?" + date[0]
			cache, err := getCache(cacheKey)
			if err != nil {
				cacheStatus(&w, false)

				params := make(map[string]string)
				params["startDate"] = date[0]
				params["lineupId"] = "USA-TX42500-X"
				req := GetTMSReq(params, "sports/all/events/airings")
				w.Write([]byte(req))
				setCache(cacheKey, req)
			} else {
				cacheStatus(&w, true)
				w.Write([]byte(cache))
			}
		} else {
			w.Write([]byte("Must pass a date"))
		}
	default:
		w.Write([]byte("Only GET requests are accepted"))
	}
}

// Response to /v1/tv-search?title={TITLE}.
func GetTvSearch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		enableCors(&w)
		if title, ok := r.URL.Query()["title"]; ok {
			w.Write([]byte(getTvSearchReq(title[0])))
		} else {
			w.Write([]byte("Must pass a title to search"))
		}
	default:
		w.Write([]byte("Only GET requests are accepted"))
	}
}

func getCache(key string) (string, error) {
	if env != "prod" {
		return "", errors.New("Caching is not enabled")
	}
	addr := "localhost:6379"
	pw := ""
	db := 0

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})
	val, err := client.Get(key).Result()
	if err != nil {
		log.Printf("Key not found for: %v", key)
	}
	return val, err
}

func setCache(key string, val string) {
	if env != "prod" {
		return
	}
	addr := "localhost:6379"
	pw := ""
	db := 0

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})

	// Default cache timing to one hour.
	err := client.Set(key, val, time.Hour).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	if env != "prod" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		return
	}
	(*w).Header().Set("Access-Control-Allow-Origin", "http://www.free-entertainment-guide.com")
}

func cacheStatus(w *http.ResponseWriter, status bool) {
	if status {
		(*w).Header().Set("Cache", "HIT")
		fmt.Println("Cache HIT")
	} else {
		(*w).Header().Set("Cache", "MISS")
		fmt.Println("Cache MISS")
	}
}

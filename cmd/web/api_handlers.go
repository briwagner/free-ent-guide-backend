package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/moviedb"
	"free-ent-guide-backend/pkg/tmsapi"
	"free-ent-guide-backend/pkg/tvmaze"
	"log"
	"net/http"
	"time"
)

// TODO: this should be passed per user.
const lineup = "USA-TX42500-X"

// GetMovies handler for movies by zip.
// Responds to /v1/movies?zip={ZIP}.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	zip := r.URL.Query().Get("zip")
	if zip == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Must pass a valid zip code"))
		return
	}

	// Check cache for stored response.
	Sqlc := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	cache := &models.Cache{}
	cacheKey := r.URL.String()
	err := cache.GetByName(cacheKey, Sqlc)

	if err != nil {
		cacheStatus(&w, false)

		// Date is optional to look at future days.
		date := r.URL.Query().Get("date")
		if date == "" {
			// Set timezone to avoid using UTC on server.
			zone, err := time.LoadLocation(c.Timezone)
			if err != nil {
				log.Printf("Cannot load timezone %e", err)
			}
			date = time.Now().In(zone).Format("2006-01-02")
		}

		tms := tmsapi.TmsApi{Key: c.Tms}
		tms.GetCinema(zip, date)

		// Save to cache.
		if tms.Status == http.StatusOK {
			newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
			err = newC.Insert(Sqlc)
			if err != nil {
				log.Printf("error setting cache: %v", err)
			}
		} else {
			w.WriteHeader(500)
		}
		w.Write(tms.Response)
		return
	}

	// Cache found.
	cacheStatus(&w, true)
	w.Write([]byte(cache.Value))
}

// DiscoverMovies handler for coming-soon movies.
// Responds to /v1/discover?date={YYYY-MM-DD}.
func DiscoverMovies(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	date := r.URL.Query().Get("date")
	// TODO parse this date to time (2006-01-02) and catch errors here.
	if date == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	// Check cache for stored response.
	Sqlc := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	ca := &models.Cache{}
	cacheKey := fmt.Sprintf("discover?date=%s", date)
	err := ca.GetByName(cacheKey, Sqlc)

	// No cache found.
	if err != nil {
		cacheStatus(&w, false)

		// Pass credential with struct.
		mdb := moviedb.MovieDb{Key: c.Moviedb}
		mdb.GetDiscover(date)

		if mdb.Status == http.StatusOK {
			newC := &models.Cache{Name: cacheKey, Value: string(mdb.Response)}
			err = newC.Insert(Sqlc)
			if err != nil {
				log.Printf("error setting cache: %v", err)
			}
		} else {
			w.WriteHeader(500)
		}
		w.Write(mdb.Response)
		return
	}
	// Cache found.
	cacheStatus(&w, true)
	w.Write([]byte(ca.Value))
}

// GetTvMovies responds to /v1/tv-movies?date={YYYY-MM-DD}.
func GetTvMovies(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	// Cache
	Sqlc := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	cacheKey := "tvmovies"
	ca := &models.Cache{}
	err := ca.GetByName(cacheKey, Sqlc)
	if err != nil {
		cacheStatus(&w, false)

		tms := tmsapi.TmsApi{Key: c.Tms}
		tms.GetTvMovies(date, lineup)
		w.WriteHeader(tms.Status)
		w.Write(tms.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
		err = newC.Insert(Sqlc)
		if err != nil {
			log.Printf("error setting cache: %v", err)
		}
	} else {
		cacheStatus(&w, true)
		w.Write([]byte(ca.Value))
	}
}

// GetTvSports responds to /v1/tv-sports?date={YYYY-MM-DD}.
func GetTvSports(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	Sqlc := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	cacheKey := "tvsports?" + date
	ca := &models.Cache{}
	err := ca.GetByName(cacheKey, Sqlc)
	if err != nil {
		cacheStatus(&w, false)

		tms := tmsapi.TmsApi{Key: c.Tms}
		tms.GetTvSports(date, lineup)
		w.WriteHeader(tms.Status)
		w.Write(tms.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
		err = newC.Insert(Sqlc)
		if err != nil {
			log.Printf("error setting cache: %v", err)
		}
	} else {
		cacheStatus(&w, true)
		w.Write([]byte(ca.Value))
	}
}

// GetTvSearch responds to /v1/tv-search?title={TITLE}.
func GetTvSearch(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	title := r.URL.Query().Get("title")
	if title == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a title to search"))
		return
	}

	tmz := tvmaze.TvMaze{}
	tmz.GetSearch(title)

	w.WriteHeader(tmz.Status)
	w.Write(tmz.Response)
}

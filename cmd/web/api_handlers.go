package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/tmsapi"
	"free-ent-guide-backend/pkg/tvmaze"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// TODO: this should be passed per user.
const lineup = "USA-TX42500-X"

// GetMovies handler for movies by zip.
// Responds to /v1/movies?zip={ZIP}.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryZip == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Must pass a valid zip code"))
		return
	}

	// Check cache for stored response.
	cache := &models.Cache{}
	cacheKey := r.URL.String()
	err := cache.GetByName(cacheKey, common.queries)
	if err != nil {
		cacheStatus(&w, false)

		// Date is optional to look at future days.
		if common.queryDate == "" {
			// Set timezone to avoid using UTC on server.
			zone, err := time.LoadLocation(c.Timezone)
			if err != nil {
				log.Printf("Cannot load timezone %e", err)
			}
			common.queryDate = common.now.In(zone).Format("2006-01-02")
		}

		tms := tmsapi.TmsApi{Key: c.Tms}
		err = tms.GetCinema(common.queryZip, common.queryDate)
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			return
		}

		// Save to cache.
		if tms.Status == http.StatusOK {
			newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
			err = newC.Insert(common.queries)
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
	common := prepareResponse(w, r)

	_, err := time.Parse("2006-01-02", common.queryDate)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	data, err := os.ReadFile("./public/discover.json")
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}

// GetTvMovies responds to /v1/tv-movies?date={YYYY-MM-DD}.
func GetTvMovies(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryDate == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	// Cache
	cacheKey := "tvmovies"
	ca := &models.Cache{}
	err := ca.GetByName(cacheKey, common.queries)
	if err != nil {
		cacheStatus(&w, false)

		tms := tmsapi.TmsApi{Key: c.Tms}
		tms.GetTvMovies(common.queryDate, lineup)
		w.WriteHeader(tms.Status)
		w.Write(tms.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
		err = newC.Insert(common.queries)
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
	common := prepareResponse(w, r)

	if common.queryDate == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	cacheKey := "tvsports?" + common.queryDate
	ca := &models.Cache{}
	err := ca.GetByName(cacheKey, common.queries)
	if err != nil {
		cacheStatus(&w, false)

		tms := tmsapi.TmsApi{Key: c.Tms}
		tms.GetTvSports(common.queryDate, lineup) // TODO this should return a possible error. else we store no data.
		w.WriteHeader(tms.Status)
		w.Write(tms.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
		err = newC.Insert(common.queries)
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

// GetTvShow responds to /v1/tv-show/{show_id}.
func GetTvShow(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	tvm := tvmaze.TvMaze{}
	showID, err := strconv.Atoi(common.vars["show_id"])
	if err != nil {
		log.Printf("error reading show id: %s", err)
		w.WriteHeader(400)
		return
	}

	// Check cache.
	cacheKey := fmt.Sprintf("tvshow=%d", showID)
	ca := &models.Cache{}
	err = ca.GetByName(cacheKey, common.queries)

	if err != nil {
		cacheStatus(&w, false)
		tvm.GetShow(int64(showID))

		w.WriteHeader(tvm.Status)
		w.Write(tvm.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tvm.Response)}
		err = newC.Insert(common.queries)
		if err != nil {
			log.Printf("error setting cache: %v", err)
		}
		return
	}

	cacheStatus(&w, true)
	w.Write([]byte(ca.Value))
}

// GetTvEpisode responds to /v1/tv-show/episode/{id}.
func GetTvEpisode(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	tvm := tvmaze.TvMaze{}
	episodeID, err := strconv.Atoi(common.vars["id"])
	if err != nil {
		log.Printf("error reading episode id: %s", err)
		w.WriteHeader(400)
		return
	}

	// Check cache.
	cacheKey := fmt.Sprintf("tvepisode=%d", episodeID)
	ca := &models.Cache{}
	err = ca.GetByName(cacheKey, common.queries)

	if err != nil {
		cacheStatus(&w, false)
		tvm.GetEpisode(int64(episodeID))

		w.WriteHeader(tvm.Status)
		w.Write(tvm.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tvm.Response)}
		err = newC.Insert(common.queries)
		if err != nil {
			log.Printf("error setting cache: %v", err)
		}
		return
	}

	cacheStatus(&w, true)
	w.Write([]byte(ca.Value))
}

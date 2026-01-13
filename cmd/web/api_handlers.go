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

	"go.opentelemetry.io/otel/attribute"
)

// TODO: this should be passed per user.
const lineup = "USA-TX42500-X"

// GetMovies handler for movies by zip.
// Responds to /v1/movies?zip={ZIP}.
func (app *App) GetMovies(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryZip == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Must pass a valid zip code"))
		return
	}

	// Check cache for stored response.
	cache := &models.Cache{}
	cacheKey := r.URL.String()
	err := cache.GetByName(r.Context(), cacheKey, common.queries)
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

		ctx, span := app.Tracer().Start(r.Context(), "getMovies")
		span.SetAttributes(
			attribute.KeyValue{Key: "zip", Value: attribute.StringValue(common.queryZip)},
		)
		defer span.End()

		tms := tmsapi.NewTMSApi(c.Tms)
		err = tms.GetCinema(ctx, common.queryZip, common.queryDate)
		if err != nil {
			app.l.Error("error getMovies", "error", err)
			w.WriteHeader(500)
			return
		}

		// Save to cache.
		if tms.Status == http.StatusOK {
			newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
			err = newC.Insert(r.Context(), common.queries)
			if err != nil {
				app.l.Error("error setting cache GetMovies", "error", err)
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
func (app *App) DiscoverMovies(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	_, err := time.Parse("2006-01-02", common.queryDate)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	data, err := os.ReadFile("./public/discover.json")
	if err != nil {
		app.l.Error("error discoverMovies", "error", err)
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}

// GetTvMovies responds to /v1/tv-movies?date={YYYY-MM-DD}.
func (app *App) GetTvMovies(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryDate == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	// Cache
	cacheKey := "tvmovies"
	ca := &models.Cache{}
	err := ca.GetByName(r.Context(), cacheKey, common.queries)
	if err != nil {
		cacheStatus(&w, false)

		ctx, span := app.Tracer().Start(r.Context(), "getTvMovies")
		span.SetAttributes(
			attribute.KeyValue{Key: "date", Value: attribute.StringValue(common.queryDate)},
		)
		defer span.End()

		tms := tmsapi.NewTMSApi(c.Tms)
		doErr := tms.GetTvMovies(ctx, common.queryDate, lineup)
		if doErr != nil {
			app.l.Error("error getTvMovies", "error", doErr)
			w.WriteHeader(tms.Status)
			return
		}

		w.WriteHeader(tms.Status)
		w.Write(tms.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
		err = newC.Insert(r.Context(), common.queries)
		if err != nil {
			app.l.Error("error setting cache GetTvMovies", "error", err)
		}
	} else {
		cacheStatus(&w, true)
		w.Write([]byte(ca.Value))
	}
}

// GetTvSports responds to /v1/tv-sports?date={YYYY-MM-DD}.
func (app *App) GetTvSports(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryDate == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	cacheKey := "tvsports?" + common.queryDate
	ca := &models.Cache{}
	err := ca.GetByName(r.Context(), cacheKey, common.queries)
	if err != nil {
		cacheStatus(&w, false)

		ctx, span := app.Tracer().Start(r.Context(), "getTvSports")
		span.SetAttributes(
			attribute.KeyValue{Key: "date", Value: attribute.StringValue(common.queryDate)},
		)
		defer span.End()

		tms := tmsapi.NewTMSApi(c.Tms)
		doErr := tms.GetTvSports(ctx, common.queryDate, lineup)
		if doErr != nil {
			app.l.Error("error getTvSports", "error", doErr)
			w.WriteHeader(tms.Status)
			return
		}

		w.WriteHeader(tms.Status)
		w.Write(tms.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tms.Response)}
		err = newC.Insert(ctx, common.queries)
		if err != nil {
			app.l.Error("error setting cache GetTvSports", "error", err)
		}
	} else {
		cacheStatus(&w, true)
		w.Write([]byte(ca.Value))
	}
}

// GetTvSearch responds to /v1/tv-search?title={TITLE}.
func (app *App) GetTvSearch(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	title := r.URL.Query().Get("title")
	if title == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a title to search"))
		return
	}

	ctx, span := app.Tracer().Start(r.Context(), "getTvSearch")
	span.SetAttributes(
		attribute.KeyValue{Key: "title", Value: attribute.StringValue(title)},
	)
	defer span.End()

	tvm := tvmaze.NewTVMaze()
	doErr := tvm.GetSearch(ctx, title)
	if doErr != nil {
		app.l.Error("error getTvSearch", "error", doErr)
		w.WriteHeader(tvm.Status)
		return
	}

	w.WriteHeader(tvm.Status)
	w.Write(tvm.Response)
}

// GetTvShow responds to /v1/tv-show/{show_id}.
func (app *App) GetTvShow(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	tvm := tvmaze.NewTVMaze()
	showID, err := strconv.Atoi(common.vars["show_id"])
	if err != nil {
		app.l.Error("error reading show id", "error", err)
		w.WriteHeader(400)
		return
	}

	// Check cache.
	cacheKey := fmt.Sprintf("tvshow=%d", showID)
	ca := &models.Cache{}
	err = ca.GetByName(r.Context(), cacheKey, common.queries)

	if err != nil {
		cacheStatus(&w, false)

		ctx, span := app.Tracer().Start(r.Context(), "getTvShow")
		span.SetAttributes(
			attribute.KeyValue{Key: "showID", Value: attribute.IntValue(showID)},
		)
		defer span.End()

		doErr := tvm.GetShow(ctx, int64(showID))
		if doErr != nil {
			app.l.Error("error getTvShow", "error", doErr)
			w.WriteHeader(tvm.Status)
			return
		}

		w.WriteHeader(tvm.Status)
		w.Write(tvm.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tvm.Response)}
		err = newC.Insert(ctx, common.queries)
		if err != nil {
			app.l.Error("error setting cache GetTvShow", "error", err)
		}
		return
	}

	cacheStatus(&w, true)
	w.Write([]byte(ca.Value))
}

// GetTvEpisode responds to /v1/tv-show/episode/{id}.
func (app *App) GetTvEpisode(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	tvm := tvmaze.NewTVMaze()
	episodeID, err := strconv.Atoi(common.vars["id"])
	if err != nil {
		app.l.Error("error reading episode id: %s", "error", err)
		w.WriteHeader(400)
		return
	}

	// Check cache.
	cacheKey := fmt.Sprintf("tvepisode=%d", episodeID)
	ca := &models.Cache{}
	err = ca.GetByName(r.Context(), cacheKey, common.queries)

	if err != nil {
		cacheStatus(&w, false)

		ctx, span := app.Tracer().Start(r.Context(), "getTvEpisode")
		span.SetAttributes(
			attribute.KeyValue{Key: "episode", Value: attribute.IntValue(episodeID)},
		)
		defer span.End()

		doErr := tvm.GetEpisode(ctx, int64(episodeID))
		if doErr != nil {
			app.l.Error("error getTvEpisode", "error", doErr)
			w.WriteHeader(tvm.Status)
			return
		}

		w.WriteHeader(tvm.Status)
		w.Write(tvm.Response)

		newC := &models.Cache{Name: cacheKey, Value: string(tvm.Response)}
		err = newC.Insert(ctx, common.queries)
		if err != nil {
			app.l.Error("error setting cache: %v", "error", err)
		}
		return
	}

	cacheStatus(&w, true)
	w.Write([]byte(ca.Value))
}

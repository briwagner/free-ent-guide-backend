package main

import (
	"encoding/json"
	"errors"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/nhlapi"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// NHLGameHandler returns a single scheduled game by GameID,
// and includes live scores and timing from the official api.
func NHLGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queries := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)

	// Fetch game from database.
	var g models.NHLGame
	gid, err := strconv.Atoi(vars["game_id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = g.FindByGameID(queries, gid)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Fetch scores, timing from the NHL api.
	gu, err := nhlapi.GetUpdate(g.GameID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TODO streamline this so we don't need a custom struct to return?
	g.HomeScore = int(gu.HomeScore)
	g.VisitorScore = int(gu.VisitorScore)
	g.Period = int(gu.Period)
	g.Status = gu.Status
	ngu := &models.NHLGameUpdate{}
	ngu.FromGame(g)

	ret, err := json.Marshal(ngu)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(ret))
}

// NHLGamesHandler responds to GET and returns the
// schedule of NHL games for the specified date.
func NHLGamesHandler(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query().Get("date")
	if d == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Must pass a date"))
		return
	}

	queries := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	gs := &models.NHLGames{}
	err := gs.LoadByDate(queries, d)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(*gs) == 0 {
		enableCors(&w)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	gsJSON, err := json.Marshal(gs)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gsJSON))
}

func NHLGamesLatest(w http.ResponseWriter, r *http.Request) {
	queries := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	gs, err := models.NHLGetLatestGames(queries)
	if err != nil || len(gs) == 0 {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var games models.NHLGames
	err = games.LoadByDate(queries, gs[0].Gametime.Format("2006-01-02"))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	gsJSON, err := json.Marshal(games)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gsJSON))
}

func NHLGamesNext(w http.ResponseWriter, r *http.Request) {
	queries := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	d := time.Now()
	gs, err := models.NHLGetNextGameday(queries, d)
	if err != nil {
		log.Printf("error getting next NHL games for %s: %s", d.Format("2006-01-02"), err)
	}

	gsJSON, err := json.Marshal(gs)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gsJSON))
}

// MLBGameHandler responds to GET and returns a single
// game matching the specified GameID.
func MLBGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)

	enableCors(&w)

	// Fetch game from database.
	g := models.MLBGame{}
	gameID, err := strconv.Atoi(vars["game_id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = g.FindByGameID(q, gameID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Fetch scores, timing from the MLB api.
	g.GetUpdate()
	if err != nil {
		log.Print(err)
		// Special error case for canceled game that should be marked deleted in DB.
		if errors.Is(err, models.ErrorGameCanceled) {
			w.WriteHeader(http.StatusGone)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	gJSON, err := json.Marshal(g)
	if err != nil {
		log.Printf("game update failed: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO we are returning the Game, not a GameUpdate anymore.
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gJSON))
}

// MLBGamesHandler responds to GET and returns the
// schedule of MLB games for the specified date.
func MLBGamesHandler(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query().Get("date")
	if d == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Must pass a date"))
		return
	}

	q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	gs := &models.MLBGames{}
	err := gs.LoadByDate(q, d)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	gsJSON, err := json.Marshal(gs)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gsJSON))
}

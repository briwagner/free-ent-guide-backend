package main

import (
	"encoding/json"
	"errors"
	"free-ent-guide-backend/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// NHLGameHandler returns a single scheduled game by GameID,
// and includes live scores and timing from the official api.
func NHLGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DB := r.Context().Value(models.StorageContextKey).(models.Store)

	// Fetch game from database.
	var g models.NHLGame
	err := g.FindByID(vars["game_id"], DB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Fetch scores, timing from the NHL api.
	gu, err := g.GetUpdate()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	guJSON, err := json.Marshal(gu)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(guJSON))
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
	DB := r.Context().Value(models.StorageContextKey).(models.Store)
	gs := &models.NHLGames{}
	err := gs.LoadByDate(d, DB)
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

// MLBGameHandler responds to GET and returns a single
// game matching the specified GameID.
func MLBGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DB := r.Context().Value(models.StorageContextKey).(models.Store)

	enableCors(&w)

	// Fetch game from database.
	var g models.MLBGame
	err := g.FindByID(vars["game_id"], DB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Fetch scores, timing from the MLB api.
	gu, err := g.GetUpdate()
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

	guJSON, err := json.Marshal(gu)
	if err != nil {
		log.Printf("game update failed: %s", err)
		// why is this 500? can we just return the game?
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(guJSON))
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

	DB := r.Context().Value(models.StorageContextKey).(models.Store)
	gs := &models.MLBGames{}
	err := gs.LoadByDate(d, DB)
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

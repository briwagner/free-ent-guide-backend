package main

import (
	"encoding/json"
	"free-ent-guide-backend/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// NHLGameHandler responds to GET and returns a single
// game matching the specified GameID.
func NHLGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DB := r.Context().Value(models.StorageContextKey).(models.Store)

	var g models.NHLGame
	err := g.FindByID(vars["game_id"], DB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	gJSON, err := json.Marshal(g)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gJSON))
}

// NHLGamesHandler responds to GET and returns the
// schedule of NHL games for the specified date.
func NHLGamesHandler(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query().Get("date")
	if d == "" {
		w.WriteHeader(400)
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

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gsJSON))
}

// MLBGameHandler responds to GET and returns a single
// game matching the specified GameID.
func MLBGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DB := r.Context().Value(models.StorageContextKey).(models.Store)

	var g models.MLBGame
	err := g.FindByID(vars["game_id"], DB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	gJSON, err := json.Marshal(g)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gJSON))
}

// MLBGamesHandler responds to GET and returns the
// schedule of MLB games for the specified date.
func MLBGamesHandler(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query().Get("date")
	if d == "" {
		w.WriteHeader(400)
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

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gsJSON))
}

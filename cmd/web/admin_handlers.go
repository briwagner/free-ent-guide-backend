package main

import (
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/docker_importer"
	"log"
	"net/http"
)

// GamesImporter responds to POST to run the Docker importer
// for both MLB and NHL games on the date specified.
func GamesImporter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		enableCors(&w)
		// TODO: above is not enough for preflight options. Add below to default action?
		(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(200)
		return
	case "POST":
		enableCors(&w)

		d := r.URL.Query().Get("date")
		if d == "" {
			http.Error(w, "must pass a date", 400)
			return
		}

		DB := r.Context().Value(models.StorageContextKey).(models.Store)

		err := docker_importer.ImportNHL(DB, d)
		if err != nil {
			log.Print(err)
			http.Error(w, "importer failed", 500)
			return
		}

		err = docker_importer.ImportMLB(DB, d)
		if err != nil {
			log.Print(err)
			http.Error(w, "importer failed", 500)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// NHLImportHandler responds to the HTTP request and runs
// the Docker import for NHL games on the specified date.
func NHLImportHandler(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query().Get("date")
	if d == "" {
		http.Error(w, "must pass a date", 400)
		return
	}

	DB := r.Context().Value(models.StorageContextKey).(models.Store)

	err := docker_importer.ImportNHL(DB, d)
	if err != nil {
		log.Print(err)
		http.Error(w, "importer failed", 500)
		return
	}
	w.Write([]byte("ok"))
}

// MLBImportHandler responds to the HTTP request and runs
// the Docker import for MLB games on the specified date.
func MLBImportHandler(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query().Get("date")
	if d == "" {
		http.Error(w, "must pass a date", 400)
		return
	}

	DB := r.Context().Value(models.StorageContextKey).(models.Store)

	err := docker_importer.ImportMLB(DB, d)
	if err != nil {
		log.Print(err)
		http.Error(w, "importer failed", 500)
		return
	}
	w.Write([]byte("ok"))
}

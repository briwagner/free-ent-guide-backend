package main

import (
	"free-ent-guide-backend/pkg/docker_importer"
	"log"
	"net/http"
)

// NHLImportHandler responds to the HTTP request and runs
// the Docker import for NHL games on the specified date.
func NHLImportHandler(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query().Get("date")
	if d == "" {
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	err := docker_importer.ImportNHL(DB.DB, d)
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
		w.WriteHeader(400)
		w.Write([]byte("Must pass a date"))
		return
	}

	err := docker_importer.ImportMLB(DB.DB, d)
	if err != nil {
		log.Print(err)
		http.Error(w, "importer failed", 500)
		return
	}
	w.Write([]byte("ok"))
}

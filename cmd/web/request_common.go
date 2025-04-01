package main

import (
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// RequestCommon holds common data needed for most requests.
type RequestCommon struct {
	now       time.Time
	vars      map[string]string
	queryDate string
	queryZip  string
	queries   *modelstore.Queries
}

// prepareResponse prepares common data for most requests.
func prepareResponse(w http.ResponseWriter, r *http.Request) *RequestCommon {
	vars := mux.Vars(r)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	return &RequestCommon{
		now:       time.Now(),
		vars:      vars,
		queryDate: r.URL.Query().Get("date"),
		queryZip:  r.URL.Query().Get("zip"),
		queries:   r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries),
	}
}

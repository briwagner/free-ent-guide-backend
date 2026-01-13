package main

import (
	"context"
	"free-ent-guide-backend/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewRouter(app App) *mux.Router {
	mux := mux.NewRouter()
	mux.Handle("/v1/movies", otelhttp.NewHandler(
		http.HandlerFunc(app.GetMovies),
		"get-movies", // this becomes the title of the trace in Grafana
	)).Methods("GET")
	mux.Handle("/v1/discover", otelhttp.NewHandler(
		http.HandlerFunc(app.DiscoverMovies),
		"discover-movies",
	)).Methods("GET")
	mux.Handle("/v1/tv-movies", otelhttp.NewHandler(
		http.HandlerFunc(app.GetTvMovies),
		"tv-movies",
	)).Methods("GET")
	mux.Handle("/v1/tv-sports", otelhttp.NewHandler(
		http.HandlerFunc(app.GetTvSports),
		"tv-sports",
	)).Methods("GET")
	mux.Handle("/v1/tv-search", otelhttp.NewHandler(
		http.HandlerFunc(app.GetTvSearch),
		"tv-search",
	)).Methods("GET")
	mux.Handle("/v1/tv-show/{show_id}", otelhttp.NewHandler(
		http.HandlerFunc(app.GetTvShow),
		"tv-show-id",
	)).Methods("GET")
	mux.Handle("/v1/tv-show/episode/{id}", otelhttp.NewHandler(
		http.HandlerFunc(app.GetTvEpisode),
		"tv-show-episode",
	)).Methods("GET")

	mux.HandleFunc("/v1/users/create", app.UsersCreate)
	mux.HandleFunc("/v1/users/token", app.AuthHandler(http.HandlerFunc(app.UsersCreateToken)))
	mux.HandleFunc("/v1/users/revoke", app.UsersRevokeToken)

	// mux.HandleFunc("/v1/users/get-zip", app.AuthHandler(RoleHandler(http.HandlerFunc(UsersGetZip))))
	mux.HandleFunc("/v1/users/get-data", app.AuthHandler(http.HandlerFunc(app.UsersGetData)))
	mux.HandleFunc("/v1/users/add-zip", app.AuthHandler(http.HandlerFunc(app.UsersAddZip)))
	mux.HandleFunc("/v1/users/delete-zip", app.AuthHandler(http.HandlerFunc(app.UsersDeleteZip)))
	mux.HandleFunc("/v1/users/clear-zip", app.AuthHandler(RoleHandler(http.HandlerFunc(app.UsersClearZip)))).Methods("POST")
	mux.HandleFunc("/v1/users/add-show", app.AuthHandler(http.HandlerFunc(app.UsersAddShow)))
	mux.HandleFunc("/v1/users/delete-show", app.AuthHandler(http.HandlerFunc(app.UsersDeleteShow)))

	mux.Handle("/v1/sports/mlb/games", otelhttp.NewHandler(
		http.HandlerFunc(app.MLBGamesHandler), "mlb-games"))
	mux.Handle("/v1/sports/mlb/game/{game_id}", otelhttp.NewHandler(
		http.HandlerFunc(app.MLBGameHandler), "mlb-game-id"))
	mux.Handle("/v1/sports/mlb/team/{team_id}", otelhttp.NewHandler(
		http.HandlerFunc(app.MLBTeamHandler), "mlb-team"))

	mux.Handle("/v1/sports/nhl/games", otelhttp.NewHandler(
		http.HandlerFunc(app.NHLGamesHandler), "nhl-games"))
	mux.Handle("/v1/sports/nhl/game/{game_id}", otelhttp.NewHandler(
		http.HandlerFunc(app.NHLGameHandler), "nhl-game-id"))
	mux.Handle("/v1/sports/nhl/latest", otelhttp.NewHandler(
		http.HandlerFunc(app.NHLGamesLatest), "nhl-latest"))
	mux.Handle("/v1/sports/nhl/next", otelhttp.NewHandler(
		http.HandlerFunc(app.NHLGamesNext), "nhl-next"))

	// Metrics
	mux.Handle("/debug/vars", http.DefaultServeMux)

	// Middleware
	mux.Use(StorageHandler)

	return mux
}

// Middleware to add Storage ref to context.
func StorageHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), models.SqlcStorageContextKey, Queries)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

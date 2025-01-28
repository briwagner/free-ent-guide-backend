package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(app App) *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/v1/movies", GetMovies).Methods("GET")
	mux.HandleFunc("/v1/discover", DiscoverMovies).Methods("GET")
	mux.HandleFunc("/v1/tv-movies", GetTvMovies).Methods("GET")
	mux.HandleFunc("/v1/tv-sports", GetTvSports).Methods("GET")
	mux.HandleFunc("/v1/tv-search", GetTvSearch).Methods("GET")
	mux.HandleFunc("/v1/tv-show/{show_id}", GetTvShow).Methods("GET")
	mux.HandleFunc("/v1/tv-show/episode/{id}", GetTvEpisode).Methods("GET")

	mux.HandleFunc("/v1/users/create", UsersCreate)
	mux.HandleFunc("/v1/users/token", app.AuthHandler(http.HandlerFunc(app.UsersCreateToken)))
	mux.HandleFunc("/v1/users/revoke", app.UsersRevokeToken)

	// mux.HandleFunc("/v1/users/get-zip", app.AuthHandler(RoleHandler(http.HandlerFunc(UsersGetZip))))
	mux.HandleFunc("/v1/users/get-data", app.AuthHandler(http.HandlerFunc(UsersGetData)))
	mux.HandleFunc("/v1/users/add-zip", app.AuthHandler(http.HandlerFunc(UsersAddZip)))
	mux.HandleFunc("/v1/users/delete-zip", app.AuthHandler(http.HandlerFunc(UsersDeleteZip)))
	mux.HandleFunc("/v1/users/clear-zip", app.AuthHandler(RoleHandler(http.HandlerFunc(UsersClearZip)))).Methods("POST")
	mux.HandleFunc("/v1/users/add-show", app.AuthHandler(http.HandlerFunc(UsersAddShow)))
	mux.HandleFunc("/v1/users/delete-show", app.AuthHandler(http.HandlerFunc(UsersDeleteShow)))

	mux.HandleFunc("/v1/sports/mlb/games", MLBGamesHandler)
	mux.HandleFunc("/v1/sports/mlb/game/{game_id}", MLBGameHandler)

	mux.HandleFunc("/v1/sports/nhl/games", NHLGamesHandler)
	mux.HandleFunc("/v1/sports/nhl/game/{game_id}", NHLGameHandler)
	mux.HandleFunc("/v1/sports/nhl/latest", NHLGamesLatest)
	mux.HandleFunc("/v1/sports/nhl/next", NHLGamesNext)

	// Metrics
	mux.Handle("/debug/vars", http.DefaultServeMux)

	// Middleware
	mux.Use(StorageHandler)

	return mux
}

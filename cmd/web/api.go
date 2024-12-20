package main

import (
	"context"
	"expvar"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/cred"
	"log"
	"os"
	"os/signal"

	"net/http"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/shaj13/libcache/fifo"

	"github.com/shaj13/go-guardian/v2/auth"
)

var (
	c        cred.Cred
	counters *expvar.Map
	Queries  *modelstore.Queries
)

func main() {
	// Set-up application config.
	c.GetCreds("creds", ".")

	// Set-up database.
	Queries = modelstore.New(models.Setup(c))

	// Set-up authentication.
	var app App
	app.setupGoGuardian()

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

	mux.HandleFunc("/v1/sports/mlb/games", MLBGamesHandler)
	mux.HandleFunc("/v1/sports/mlb/game/{game_id}", MLBGameHandler)

	mux.HandleFunc("/v1/sports/nhl/games", NHLGamesHandler)
	mux.HandleFunc("/v1/sports/nhl/game/{game_id}", NHLGameHandler)
	mux.HandleFunc("/v1/sports/nhl/latest", NHLGamesLatest)
	mux.HandleFunc("/v1/sports/nhl/next", NHLGamesNext)

	// Metrics
	mux.Handle("/debug/vars", http.DefaultServeMux)
	counters = expvar.NewMap("counters")
	counters.Set("cache_hit", new(expvar.Int))
	counters.Set("cache_miss", new(expvar.Int))

	// Middleware
	mux.Use(StorageHandler)

	fmt.Printf("ENT API is live. Listening on port %v ...\n", c.GetPort())
	srv := &http.Server{
		Addr:         "0.0.0.0" + c.GetPort(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// do something before shutdown.

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if c.Env != "prod" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		return
	}

	(*w).Header().Set("Access-Control-Allow-Origin", c.CorsOrigin)
}

func cacheStatus(w *http.ResponseWriter, status bool) {
	if status {
		(*w).Header().Set("Cache", "HIT")
		counters.Add("cache_hit", 1)
	} else {
		(*w).Header().Set("Cache", "MISS")
		counters.Add("cache_miss", 1)
	}
}

// Validate user with basic auth.
func validateUser(ctx context.Context, r *http.Request, username, password string) (auth.Info, error) {
	u := &models.User{Email: username, Password: password}
	err := u.FindByEmail(Queries)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("failed to load user")
	}
	if u.CheckPasswordHash(Queries, password) {
		return auth.NewDefaultUser(u.Email, fmt.Sprintf("%d", u.ID), nil, nil), nil
	}

	return nil, fmt.Errorf("invalid credentials")
}

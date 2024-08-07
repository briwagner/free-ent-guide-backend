package main

import (
	"context"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/authenticator"
	"free-ent-guide-backend/pkg/cred"
	"log"

	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"

	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
)

var (
	c       cred.Cred
	Queries *modelstore.Queries
	author  authenticator.Authenticator
)

func main() {
	// Set-up application config.
	c.GetCreds("creds", ".")
	port := fmt.Sprintf(":%v", c.Port)

	// Set-up database.
	Sqlc := models.Setup(c)
	Queries = modelstore.New(Sqlc)

	// Set-up authentication.
	author = authenticator.Authenticator{
		Audience: jwt.SetAudience("any"),
		Issuer:   jwt.SetIssuer("ent_api"),
		Duration: c.TokenDuration,
	}
	author.AttachSecret(c.TokenSecret)
	setupGoGuardian()

	mux := mux.NewRouter()
	mux.HandleFunc("/v1/movies", GetMovies).Methods("GET")
	mux.HandleFunc("/v1/discover", DiscoverMovies).Methods("GET")
	mux.HandleFunc("/v1/tv-movies", GetTvMovies).Methods("GET")
	mux.HandleFunc("/v1/tv-sports", GetTvSports).Methods("GET")
	mux.HandleFunc("/v1/tv-search", GetTvSearch).Methods("GET")
	mux.HandleFunc("/v1/tv-show/{show_id}", GetTvShow).Methods("GET")
	mux.HandleFunc("/v1/tv-show/episode/{id}", GetTvEpisode).Methods("GET")

	mux.HandleFunc("/v1/users/create", UsersCreate)
	mux.HandleFunc("/v1/users/token", AuthHandler(http.HandlerFunc(UsersCreateToken)))
	mux.HandleFunc("/v1/users/revoke", UsersRevokeToken)

	// mux.HandleFunc("/v1/users/get-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersGetZip))))
	mux.HandleFunc("/v1/users/get-data", AuthHandler(http.HandlerFunc(UsersGetData)))
	mux.HandleFunc("/v1/users/add-zip", AuthHandler(http.HandlerFunc(UsersAddZip)))
	mux.HandleFunc("/v1/users/delete-zip", AuthHandler(http.HandlerFunc(UsersDeleteZip)))
	mux.HandleFunc("/v1/users/clear-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersClearZip)))).Methods("POST")
	mux.HandleFunc("/v1/users/add-show", AuthHandler(http.HandlerFunc(UsersAddShow)))

	mux.HandleFunc("/v1/sports/nhl/games", NHLGamesHandler)
	mux.HandleFunc("/v1/sports/mlb/games", MLBGamesHandler)
	mux.HandleFunc("/v1/sports/nhl/game/{game_id}", NHLGameHandler)
	mux.HandleFunc("/v1/sports/mlb/game/{game_id}", MLBGameHandler)
	mux.HandleFunc("/v1/sports/nhl/latest", NHLGamesLatest)

	// Docker importers: DEPRECATED
	// mux.HandleFunc("/v1/admin/add-games", AuthHandler(http.HandlerFunc(GamesImporter)))

	// Middleware
	mux.Use(StorageHandler)

	fmt.Printf("ENT API is live. Listening on port %v ...\n", port)
	http.ListenAndServe(port, mux)
}

// Middleware to add Storage ref to context.
func StorageHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), models.SqlcStorageContextKey, Queries)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
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
		fmt.Println("Cache HIT")
	} else {
		(*w).Header().Set("Cache", "MISS")
		fmt.Println("Cache MISS")
	}
}

var strategy union.Union
var jwtStrategy auth.Strategy
var cacheObj libcache.Cache

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

// Define strategies, set token expiration time.
func setupGoGuardian() {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Hour * 72)
	cacheObj.RegisterOnExpired(func(key, _ interface{}) {
		cacheObj.Peek(key)
	})
	basicStrategy := basic.NewCached(validateUser, cacheObj)

	jwtStrategy = jwt.New(cacheObj, author.Secret, author.Audience)
	strategy = union.New(jwtStrategy, basicStrategy)
}

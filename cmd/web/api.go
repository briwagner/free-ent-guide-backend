package main

import (
	"context"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/cred"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"

	"github.com/go-redis/redis"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/token"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
)

var c cred.Cred
var cacheClient *redis.Client

func main() {
	// Set-up application config.
	c.GetCreds("creds", ".")
	port := fmt.Sprintf(":%v", c.Port)

	// Set-up cache client for requests.
	cacheClient = redis.NewClient(&redis.Options{
		Addr:     c.RedisPort,
		Password: c.RedisPassword,
		DB:       c.RedisDB,
	})

	setupGoGuardian()

	mux := mux.NewRouter()
	mux.HandleFunc("/v1/movies", GetMovies).Methods("GET")
	mux.HandleFunc("/v1/discover", DiscoverMovies).Methods("GET")
	mux.HandleFunc("/v1/tv-movies", GetTvMovies).Methods("GET")
	mux.HandleFunc("/v1/tv-sports", GetTvSports).Methods("GET")
	mux.HandleFunc("/v1/tv-search", GetTvSearch).Methods("GET")

	mux.HandleFunc("/v1/users/create", UsersCreate)
	mux.HandleFunc("/v1/users/token", AuthHandler(http.HandlerFunc(UsersCreateToken)))
	mux.HandleFunc("/v1/users/revoke", UsersRevokeToken).Methods("POST")

	// mux.HandleFunc("/v1/users/get-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersGetZip))))
	mux.HandleFunc("/v1/users/get-zip", AuthHandler(http.HandlerFunc(UsersGetZip)))
	mux.HandleFunc("/v1/users/add-zip", AuthHandler(http.HandlerFunc(UsersAddZip)))
	mux.HandleFunc("/v1/users/delete-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersDeleteZip)))).Methods("POST")
	mux.HandleFunc("/v1/users/clear-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersClearZip)))).Methods("POST")

	fmt.Printf("ENT API is live. Listening on port %v ...\n", port)
	http.ListenAndServe(port, mux)
}

func getCache(key string) (string, error) {
	val, err := cacheClient.Get(key).Result()
	if err != nil {
		log.Printf("Key not found for: %v", key)
	}
	return val, err
}

func setCache(key string, val string, t time.Duration) error {
	err := cacheClient.Set(key, val, t).Err()
	if err != nil {
		fmt.Printf("Redis setCache: %s\n", err.Error())
	}
	return err
}

func enableCors(w *http.ResponseWriter) {
	if c.Env != "prod" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		return
	}
	(*w).Header().Set("Access-Control-Allow-Origin", "http://www.free-entertainment-guide.com")
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
var tokenStrategy auth.Strategy
var cacheObj libcache.Cache

// Validate user with basic auth.
func validateUser(ctx context.Context, r *http.Request, username, password string) (auth.Info, error) {
	u := &models.User{Name: username, Password: password}
	if u.CheckPasswordHash(cacheClient, password) {
		return auth.NewDefaultUser(u.Name, "1", nil, nil), nil
	}

	return nil, fmt.Errorf("Invalid credentials")
}

// Define strategies, set token expiration time.
func setupGoGuardian() {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Hour * 72)
	cacheObj.RegisterOnExpired(func(key, _ interface{}) {
		cacheObj.Peek(key)
	})
	basicStrategy := basic.NewCached(validateUser, cacheObj)
	tokenStrategy = token.New(token.NoOpAuthenticate, cacheObj)
	strategy = union.New(tokenStrategy, basicStrategy)
}

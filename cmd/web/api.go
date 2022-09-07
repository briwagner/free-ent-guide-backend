package main

import (
	"context"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/authenticator"
	"free-ent-guide-backend/pkg/cred"
	"free-ent-guide-backend/pkg/storage"

	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"

	"github.com/go-redis/redis"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
)

var c cred.Cred
var cacheClient *redis.Client

// var DB *gorm.DB
var DB storage.Store
var author authenticator.Authenticator

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

	// Set-up database.
	DB = storage.Setup(c)

	author = authenticator.Authenticator{}
	// Pass secret to setup JWT authenticator.
	setupGoGuardian(c.TokenSecret)

	mux := mux.NewRouter()
	mux.HandleFunc("/v1/movies", GetMovies).Methods("GET")
	mux.HandleFunc("/v1/discover", DiscoverMovies).Methods("GET")
	mux.HandleFunc("/v1/tv-movies", GetTvMovies).Methods("GET")
	mux.HandleFunc("/v1/tv-sports", GetTvSports).Methods("GET")
	mux.HandleFunc("/v1/tv-search", GetTvSearch).Methods("GET")

	mux.HandleFunc("/v1/users/create", UsersCreate)
	mux.HandleFunc("/v1/users/token", AuthHandler(http.HandlerFunc(UsersCreateToken)))
	mux.HandleFunc("/v1/users/revoke", UsersRevokeToken)

	// mux.HandleFunc("/v1/users/get-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersGetZip))))
	mux.HandleFunc("/v1/users/get-zip", AuthHandler(http.HandlerFunc(UsersGetZip)))
	mux.HandleFunc("/v1/users/add-zip", AuthHandler(http.HandlerFunc(UsersAddZip)))
	mux.HandleFunc("/v1/users/delete-zip", AuthHandler(http.HandlerFunc(UsersDeleteZip)))
	mux.HandleFunc("/v1/users/clear-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersClearZip)))).Methods("POST")

	// Middleware
	mux.Use(StorageHandler)

	fmt.Printf("ENT API is live. Listening on port %v ...\n", port)
	http.ListenAndServe(port, mux)
}

func StorageHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add Store on all requests.
		ctx := context.WithValue(r.Context(), storage.StorageContextKey, DB)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
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
var jwtStrategy auth.Strategy
var cacheObj libcache.Cache

// Validate user with basic auth.
func validateUser(ctx context.Context, r *http.Request, username, password string) (auth.Info, error) {
	u := &models.User{Name: username, Password: password}
	if u.CheckPasswordHash(cacheClient, password) {
		return auth.NewDefaultUser(u.Name, "1", nil, nil), nil
	}

	return nil, fmt.Errorf("invalid credentials")
}

// Define strategies, set token expiration time.
func setupGoGuardian(secret string) {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Hour * 72)
	cacheObj.RegisterOnExpired(func(key, _ interface{}) {
		cacheObj.Peek(key)
	})
	basicStrategy := basic.NewCached(validateUser, cacheObj)

	author.AttachSecret(secret)
	author.Audience = jwt.SetAudience("any")
	author.Issuer = jwt.SetIssuer(("ent_api"))
	author.Duration = c.TokenDuration
	jwtStrategy = jwt.New(cacheObj, author.Secret, author.Audience)
	strategy = union.New(jwtStrategy, basicStrategy)
}

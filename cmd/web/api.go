package main

import (
	"context"
	"errors"
	"fmt"
	"free-ent-guide-backend/models"
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
	"github.com/spf13/viper"
)

// Cred holds app credentials.
type Cred struct {
	Tms           string `mapstructure:"tms"`
	Moviedb       string `mapstructure:"moviedb"`
	Port          int    `mapstructure:"port"`
	Env           string `mapstructure:"env"`
	RedisPort     string `mapstructure:"redis_port"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
	Cache         bool   `mapstructure:"use_cache"`
	Timezone      string `mapstructure:"timezone"`
	TokenDuration int64  `mapstructure:"token_duration"`
	TokenSecret   string `mapstructure:"token_secret"`
}

var c Cred
var cacheClient *redis.Client

func main() {
	// Set-up application config.
	c.getCreds()
	port := fmt.Sprintf(":%v", c.Port)

	// Set-up cache client for requests.
	if c.Cache {
		cacheClient = redis.NewClient(&redis.Options{
			Addr:     c.RedisPort,
			Password: c.RedisPassword,
			DB:       c.RedisDB,
		})
	}

	setupGoGuardian()

	mux := mux.NewRouter()
	mux.HandleFunc("/v1/movies", GetMovies)
	mux.HandleFunc("/v1/discover", DiscoverMovies)
	mux.HandleFunc("/v1/tv-movies", GetTvMovies)
	mux.HandleFunc("/v1/tv-sports", GetTvSports)
	mux.HandleFunc("/v1/tv-search", GetTvSearch)

	mux.HandleFunc("/v1/users/create", UsersCreate)
	mux.HandleFunc("/v1/users/token", AuthHandler(http.HandlerFunc(UsersCreateToken))).Methods("GET")

	mux.HandleFunc("/v1/users/get-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersGetZip)))).Methods("GET")
	mux.HandleFunc("/v1/users/add-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersAddZip))))
	mux.HandleFunc("/v1/users/delete-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersDeleteZip))))
	mux.HandleFunc("/v1/users/clear-zip", AuthHandler(RoleHandler(http.HandlerFunc(UsersClearZip))))

	fmt.Printf("ENT API is live. Listening on port %v ...\n", port)
	http.ListenAndServe(port, mux)
}

func (c *Cred) getCreds() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// SetConfigFile() can error as it looks for absolute path.
	viper.SetConfigName("creds")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("File not loading: %v", err)
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("No credentials %v", err)
	}
}

func getCache(key string) (string, error) {
	if c.Cache != true {
		return "", errors.New("Caching is not enabled")
	}

	val, err := cacheClient.Get(key).Result()
	if err != nil {
		log.Printf("Key not found for: %v", key)
	}
	return val, err
}

func setCache(key string, val string, t time.Duration) error {
	if c.Cache != true {
		return errors.New("No storage found")
	}

	err := cacheClient.Set(key, val, t).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func enableCors(w *http.ResponseWriter) {
	if c.Env != "prod" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
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
	val, err := cacheClient.Get(fmt.Sprintf("user:%s", username)).Result()
	if err != nil {
		// TODO: what does this return?
		return nil, fmt.Errorf("Invalid credentials")
	}

	user := &models.User{Name: username, Password: password}
	if user.CheckPasswordHash(password, val) {
		return auth.NewDefaultUser(user.Name, "1", nil, nil), nil
	}

	return nil, fmt.Errorf("Invalid credentials")
}

func setupGoGuardian() {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Minute * 5)
	cacheObj.RegisterOnExpired(func(key, _ interface{}) {
		cacheObj.Peek(key)
	})
	basicStrategy := basic.NewCached(validateUser, cacheObj)
	tokenStrategy = token.New(token.NoOpAuthenticate, cacheObj)
	strategy = union.New(tokenStrategy, basicStrategy)
}

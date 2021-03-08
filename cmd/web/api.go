package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
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

	fmt.Printf("ENT API is live. Listening on port %v ...\n", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/movies", GetMovies)
	mux.HandleFunc("/v1/discover", DiscoverMovies)
	mux.HandleFunc("/v1/tv-movies", GetTvMovies)
	mux.HandleFunc("/v1/tv-sports", GetTvSports)
	mux.HandleFunc("/v1/tv-search", GetTvSearch)

	mux.HandleFunc("/v1/users/get-zip", UsersGetZip)
	mux.HandleFunc("/v1/users/add-zip", UsersAddZip)
	mux.HandleFunc("/v1/users/delete-zip", UsersDeleteZip)
	mux.HandleFunc("/v1/users/clear-zip", UsersClearZip)
	mux.HandleFunc("/v1/users/create", UsersCreate)

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

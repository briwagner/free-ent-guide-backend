package config

import (
	"fmt"
	"log"

	"free-ent-guide-backend/pkg/spaces"

	"github.com/spf13/viper"
)

// Config holds app config.
type Config struct {
	Tms           string        `mapstructure:"tms"`
	Moviedb       string        `mapstructure:"moviedb"`
	Port          int           `mapstructure:"port"`
	Env           string        `mapstructure:"env"`
	RedisPort     string        `mapstructure:"redis_port"`
	RedisPassword string        `mapstructure:"redis_password"`
	RedisDB       int           `mapstructure:"redis_db"`
	DB            string        `mapstructure:"db_string"`
	Timezone      string        `mapstructure:"timezone"`
	TokenDuration int64         `mapstructure:"token_duration"`
	TokenSecret   string        `mapstructure:"token_secret"`
	CorsOrigin    string        `mapstructure:"cors_origin"`
	SlackURL      string        `mapstructure:"slack_url"`
	Spaces        spaces.Config `mapstructure:"spaces"`
}

// GetCredsFromFile copies the configuration file into the config struct.
func GetCredsFromFile(fname string, fpath string) *Config {
	var c Config
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fpath)
	// SetConfigFile() can error as it looks for absolute path.
	viper.SetConfigName(fname)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("File not loading: %v", err)
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("No credentials %v", err)
	}
	return &c
}

func (c *Config) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

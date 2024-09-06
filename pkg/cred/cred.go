package cred

import (
	"fmt"
	"log"

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
	DB            string `mapstructure:"db_string"`
	Timezone      string `mapstructure:"timezone"`
	TokenDuration int64  `mapstructure:"token_duration"`
	TokenSecret   string `mapstructure:"token_secret"`
	CorsOrigin    string `mapstructure:"cors_origin"`
}

// GetCreds copies the configuration file into the cred struct.
func (c *Cred) GetCreds(fname string, fpath string) {
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
}

func (c *Cred) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

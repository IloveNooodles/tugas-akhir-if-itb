package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     int
	PostgresSSLMode  string
	ApiKey           string
	AdminAPIKey      string
	Port             int
	JWTKey           string
	FEURL            string
	RequestTimeout   int
	PollingTimeout   int
}

// Craete new config for the applications
func New() Config {
	v := viper.New()
	v.SetConfigFile(".env")
	v.AutomaticEnv()
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	v.AddConfigPath("../../")
	v.AddConfigPath("../../../")

	err := v.ReadInConfig()
	if err != nil {
		err := fmt.Errorf("config file not found: %s", err)
		panic(err)
	}

	cfg := Config{
		PostgresUser:     v.GetString("POSTGRES_USER"),
		PostgresDB:       v.GetString("POSTGRES_DB"),
		PostgresPassword: v.GetString("POSTGRES_PASSWORD"),
		PostgresHost:     v.GetString("POSTGRES_HOST"),
		PostgresPort:     v.GetInt("POSTGRES_PORT"),
		PostgresSSLMode:  v.GetString("POSTGRES_SSL_MODE"),
		ApiKey:           v.GetString("APP_API_KEY"),
		Port:             v.GetInt("APP_PORT"),
		JWTKey:           v.GetString("APP_JWT_KEY"),
		AdminAPIKey:      v.GetString("APP_ADMIN_API_KEY"),
		FEURL:            v.GetString("APP_FE_URL"),
		RequestTimeout:   v.GetInt("APP_REQUEST_TIMEOUT"),
		PollingTimeout:   v.GetInt("APP_POLLING_TIMEOUT"),
	}

	return cfg
}

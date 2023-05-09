package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GatewayRestPort       int    `mapstructure:"GATEWAY_REST_PORT"`
	GatewayRestAPIVersion string `mapstructure:"GATEWAY_REST_API_VERSION"`
	UserManagementURI     string `mapstructure:"USER_MANAGEMENT_SERVICE_URI"`
	SecretKey             string `mapstructure:"SECRET_KEY"`
}

func LoadConfig(path string) *Config {
	var cfg Config

	if path != "" {
		viper.AddConfigPath(path)
		viper.SetConfigType("env")
		viper.SetConfigName("app")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("ERROR read config", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("ERROR load config", err)
	}
	return &cfg
}

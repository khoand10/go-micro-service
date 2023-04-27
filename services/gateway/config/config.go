package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	RestPort       int    `mapstructure:"REST_PORT"`
	RestAPIVersion string `mapstructure:"API_VERSION"`
	//AllowOrigins   string `mapstructure:"CORS_ALLOW_ORIGINS_GATEWAY, required=true"`
	//RestPort       int    `mapstructure:"REST_PORT"`
	//RestAPIVersion string `mapstructure:"API_VERSION"`
	//JwtSecretKey               string `mapstructure:"JWT_SECRET_KEY"`
	//TokenExpirationHour        int    `mapstructure:"TOKEN_EXPIRATION_HOUR"`
	//RefreshTokenExpirationHour int    `mapstructure:"REFRESH_TOKEN_EXPIRATION_HOUR"`
}

func LoadConfig(path string) *Config {
	var cfg Config

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
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

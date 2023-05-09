package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	UserManagementPort         int    `mapstructure:"USER_MANAGEMENT_PORT"`
	UserManagementURI          string `mapstructure:"USER_MANAGEMENT_SERVICE_URI"`
	JwtSecretKey               string `mapstructure:"JWT_SECRET_KEY"`
	TokenExpirationHour        int    `mapstructure:"TOKEN_EXPIRATION_HOUR"`
	RefreshTokenExpirationHour int    `mapstructure:"REFRESH_TOKEN_EXPIRATION_HOUR"`
	MysqlDSN                   string `mapstructure:"MYSQL_DSN"`
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

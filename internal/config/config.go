package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBDatabase string
	SecretKey  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	config := &Config{
		DBUsername: viper.GetString("DB_USERNAME"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBDatabase: viper.GetString("DB_DATABASE"),
		SecretKey:  viper.GetString("SECRET_KEY"),
	}

	return config, nil
}

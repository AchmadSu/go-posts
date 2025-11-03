package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Config error:", err)
	}

	AppConfig = &Config{
		DBUser: viper.GetString("DB_USER"),
		DBPass: viper.GetString("DB_PASS"),
		DBHost: viper.GetString("DB_HOST"),
		DBName: viper.GetString("DB_NAME"),
	}
}

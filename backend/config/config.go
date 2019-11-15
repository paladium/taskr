package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init function configures the config variables
func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot load the server config")
	}
}

// GetConfig returns the current config
func GetConfig() *viper.Viper {
	return config
}

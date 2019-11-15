package config

import (
	"bytes"
	"log"

	"github.com/gobuffalo/packr"
	"github.com/spf13/viper"
)

var config *viper.Viper
var box packr.Box

// Init function configures the config variables
func Init(env string) {
	box = packr.NewBox("../static")
	config = viper.New()
	config.SetConfigType("yaml")
	configFile := box.Bytes("dev.yaml")
	err := config.ReadConfig(bytes.NewBuffer(configFile))
	if err != nil {
		log.Fatal("Cannot load the server config")
	}
}

// GetConfig returns the current config
func GetConfig() *viper.Viper {
	return config
}

// GetBox for static assets
func GetBox() packr.Box {
	return box
}

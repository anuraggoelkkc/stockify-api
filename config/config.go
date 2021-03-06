package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

const (
	GoogleApplicationCredentials = "GOOGLE_APPLICATION_CREDENTIALS"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) *viper.Viper {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	return config
}
func GetConfig() *viper.Viper {
	return config
}

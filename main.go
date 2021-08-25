package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"stockify-api/config"
	"stockify-api/src/constants"
	"stockify-api/src/logger"
	"stockify-api/src/server"
)

func main() {
	// Get environment flag from command-line
	environment := flag.String("e", "development", "environment")
	flag.Usage = func() {
		log.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	// Initialize app configs constants
	appConfig := config.Init(*environment)

	// set path of google cache file
	cachePath := appConfig.Get(config.GoogleApplicationCredentials).(string)
	jsonFile, _ := filepath.Abs(cachePath)
	os.Setenv(config.GoogleApplicationCredentials, jsonFile)

	//set gin mode (debug or release)
	if *environment == constants.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initalize logging
	logger.InitializeLogging(constants.ServerLogPath, *environment)

	server.Init()
}

package main

import (
	log "github.com/sirupsen/logrus"
	"mpesa-daraja-api-go/src/config"
	"os"
	"syscall"
)

// Set up a channel to listen to for interrupt signals
var runChannel = make(chan os.Signal, 1)

// init initializes system logging and sets up the log rotation and formatting
func init() {
	// Get the $LOG_FILE_PATH env variable and create the log path
	logFilePath, err := config.CreateLogPaths()

	if err != nil {
		// Kill the service as log file path creation has failed
		log.Fatalf("Service failed to start due to log file path creation error: %+v", err)
		runChannel <- syscall.SIGINT
	}

	// Initialize System logging and set-up log rotation
	config.SetupLogger(logFilePath)
}

// Initialize all System configs
func main() {
	configFilePath, err := config.ParseFlags()
	if err != nil {
		log.Fatal(err)
		runChannel <- syscall.SIGINT
	}

	configs, err := config.GetConfigurations(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Start the server
	configs.Run()
}

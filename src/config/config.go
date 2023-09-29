package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Config struct {
	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`
		// Port is the local machine TCP Port to bind the HTTP Server to
		Port    string `yaml:"port"`
		Timeout struct {
			// Server is the general server timeout to use
			// for graceful shutdowns
			Server int `yaml:"server"`

			// Write is the amount of time to wait until an HTTP server
			// write operation is cancelled
			Write int `yaml:"write"`

			// Read is the amount of time to wait until an HTTP server
			// read operation is cancelled
			Read int `yaml:"read"`

			// Read is the amount of time to wait
			// until an IDLE HTTP session is closed
			Idle int `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`

	Database struct {
		// Database connection name
		ConnectionName string `yaml:"connectionName"`
		// Database driver
		Driver string `yaml:"driver"`
		// The database Host
		Host string `yaml:"host"`
		// Database port
		Port int `yaml:"port"`
		// Database name
		Name string `yaml:"name"`
		// Database User
		User string `yaml:"user"`
		// Database password
		Password string `yaml:"password"`
		// Maximum idle connections
		MaxIdleConnections int `yaml:"maxIdleConnections"`
		// Maximum open connections
		MaxOpenConnections int `yaml:"maxOpenConnections"`
		// Maximum Idle Time
		MaxIdleTime int `yaml:"maxIdleTime"`
		// Maximum Life Time
		MaxLifeTime int `yaml:"maxLifeTime"`
		// Migrations Location
		MigrationsPath string `yaml:"migrationsPath"`
	} `yaml:"database"`
}

// GetConfigurations Gets config values from YAML file and adds them to a Config Struct
func GetConfigurations(configurationPath string) (*Config, error) {
	configuration := &Config{}

	file, err := os.Open(configurationPath)
	if err != nil {
		log.Errorf("Failed to open config file %s error: %+v", configurationPath, err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Unable to release file I/O due to err: %+v", err)
		}
	}(file)

	// init new YAML file to decode
	decode := yaml.NewDecoder(file)

	// start YAML decoding from file
	if err := decode.Decode(&configuration); err != nil {
		log.Errorf("Failed to decode YAML config file: %+v", err)
		return nil, err
	}
	return configuration, nil
}

// ValidateConfigPath function validates that the file path provided is a file
func ValidateConfigPath(filePath string) error {
	status, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if status.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", filePath)
	}
	return nil
}

// ParseFlags will create and parse the CLI flags and return the path to used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath = os.Getenv("CONFIG_FILE_PATH")

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		log.Errorf("Failed to validate config path %s and error %+v", configPath, err)
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

// NewRouter generates the router used in the HTTP Server
func NewRouter() *http.ServeMux {
	// Create a router and define routes and return that router
	router := http.NewServeMux()

	//Add routes here

	return router
}

// GenerateDatabaseDsn creates a database DSN URI
func GenerateDatabaseDsn(config Config) string {
	cfg := &mysql.Config{
		User:            config.Database.User,
		Passwd:          config.Database.Password,
		Net:             "tcp",
		Addr:            config.Database.Host + ":" + strconv.Itoa(config.Database.Port),
		DBName:          config.Database.Name,
		MultiStatements: true,
		ParseTime:       true,
	}
	log.Infof("Formatted Database DSN; %s", cfg.FormatDSN())
	return cfg.FormatDSN()
}

// Run will start the HTTP Server and initiate connection pool
func (configuration Config) Run() {
	// Set up a channel to listen to for interrupt signals
	var runChannel = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(configuration.Server.Timeout.Server),
	)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         configuration.Server.Host + ":" + configuration.Server.Port,
		Handler:      NewRouter(),
		ReadTimeout:  time.Duration(configuration.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(configuration.Server.Timeout.Write) * time.Second,
		IdleTimeout:  time.Duration(configuration.Server.Timeout.Idle) * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	time.Sleep(time.Second * 10)

	// Alert the user that the server is starting
	log.Infof("Server is starting on %s", server.Addr)

	_ = GenerateDatabaseDsn(configuration)

	// Call DB Connection pool start
	db, err := DatabaseConnectionPool(configuration)
	if err != nil {
		log.Fatalf("Unable to initialize database connection pool due to error: %+v", err)
		// Send kill command to close server as DB is not available
		runChannel <- syscall.SIGINT
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// Alert the user that the Database migrations are starting
	log.Info("Initializing Database Migrations")

	// Call DB Migrations
	err = RunDatabaseMigrations(db, configuration)
	if err != nil {
		log.Fatalf("Unable to initialize database migrations due to error: %+v", err)
		// Send kill command to close server as DB is not available
		runChannel <- syscall.SIGINT
	}

	// run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				// Normal interrupt operation, ignore
			} else {
				log.Fatalf("Server failed to start due to error %v", err)
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable, so we can let the user know why the server is shutting down
	interrupt := <-runChannel

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	log.Infof("Server is shutting down due to %+v", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}

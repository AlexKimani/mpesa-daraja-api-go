package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
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
		// Maximum batch size
		BatchSize int `yaml:"batchSize"`
		// Migrations Location
		MigrationsPath string `yaml:"migrationsPath"`
		// Generate Database Structs
		GenerateStructs bool `yaml:"generateStructs"`
	} `yaml:"database"`

	Mpesa struct {
		BaseUrl        string `yaml:"baseUrl"`
		ConsumerKey    string `yaml:"consumerKey"`
		ConsumerSecret string `yaml:"consumerSecret"`
		Urls           struct {
			Auth              string `yaml:"auth"`
			Express           string `yaml:"express"`
			ExpressQuery      string `yaml:"expressQuery"`
			C2bRegisterUrl    string `yaml:"c2bRegisterUrl"`
			B2c               string `yaml:"b2c"`
			TransactionStatus string `yaml:"transactionStatus"`
			AccountBalance    string `yaml:"accountBalance"`
			Reversal          string `yaml:"reversal"`
			TaxRemittance     string `yaml:"taxRemittance"`
			BusinessPayBill   string `yaml:"businessPayBill"`
			BusinessBuyGoods  string `yaml:"businessBuyGoods"`
		} `yaml:"urls"`
	} `yaml:"mpesa"`
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

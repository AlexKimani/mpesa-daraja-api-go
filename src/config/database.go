package config

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

// DatabaseConnectionPool initialize the Database connection pool and setup a pool of connections
func DatabaseConnectionPool(config Config, db *gorm.DB) (*sql.DB, error) {
	log.Info("Initializing Database Connection Pool")

	dbMaxIdleConnections := config.Database.MaxIdleConnections
	dbMaxOpenConnections := config.Database.MaxOpenConnections
	dbMaxIdleTime := config.Database.MaxIdleTime
	dbMaxLifeTime := config.Database.MaxLifeTime

	sqlDb, err := db.DB()
	sqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifeTime))
	sqlDb.SetMaxIdleConns(dbMaxIdleConnections)
	sqlDb.SetMaxOpenConns(dbMaxOpenConnections)
	sqlDb.SetConnMaxIdleTime(time.Duration(dbMaxIdleTime))

	if err != nil {
		log.Errorf("DatabaseConnectionPool Error: %+v", err)
		return nil, err
	}

	// This is for analyzing the stats after setting a connection
	log.Info("@OnboardingConnectionPool MYSQL MAX Open Connections: ",
		sqlDb.Stats().MaxOpenConnections)
	log.Info("@DatabaseConnectionPool MYSQL Open Connections: ",
		sqlDb.Stats().OpenConnections)
	log.Info("@DatabaseConnectionPool MYSQL InUse Connections: ",
		sqlDb.Stats().InUse)
	log.Info("@DatabaseConnectionPool MYSQL Idle Connections: ", sqlDb.Stats().Idle)
	return sqlDb, nil
}

// ConnectToDatabase connects to the Database and returns a single connection
func ConnectToDatabase(config Config) (*gorm.DB, error) {
	dbDsn := GenerateDatabaseDsn(config)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: config.Database.Driver,
		DSN:        dbDsn,
	}), &gorm.Config{
		CreateBatchSize: config.Database.BatchSize,
		PrepareStmt:     true,
		Logger:          ConfigureDatabaseLogger(),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
		return nil, err
	}
	return db, nil
}

// GenerateDatabaseDsn used to generate a database dsn URI
func GenerateDatabaseDsn(config Config) string {
	dbName := config.Database.Name
	dbUser := config.Database.User
	dbPassword := config.Database.Password
	dbPort := strconv.Itoa(config.Database.Port)
	// Create DB Connection string
	dbDsn := dbUser + ":" + dbPassword + "@tcp(" + config.Database.Host + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&multiStatements=true"
	return dbDsn
}

// ConfigureDatabaseLogger setup a custom logger for the database
func ConfigureDatabaseLogger() logger.Interface {
	var customLogger = &log.Logger{
		Out:       os.Stderr,
		Formatter: new(log.JSONFormatter),
		Hooks:     make(log.LevelHooks),
		Level:     log.DebugLevel,
	}
	newLogger := logger.New(
		customLogger, // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	return newLogger
}

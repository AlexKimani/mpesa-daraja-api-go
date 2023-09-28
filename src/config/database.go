package config

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// DatabaseConnectionPool initialize the Database connection pool and setup a pool of connections
func DatabaseConnectionPool(config Config) (bool, error) {
	log.Info("Initializing Database Connection Pool")

	dbName := config.Database.Name
	dbUser := config.Database.User
	dbPassword := config.Database.Password
	dbPort := strconv.Itoa(config.Database.Port)
	// Create DB Connection string
	dbDsn := dbUser + ":" + dbPassword + "@tcp(" + config.Database.Host + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: config.Database.Driver,
		DSN:        dbDsn,
	}), &gorm.Config{})

	dbMaxIdleConnections := config.Database.MaxIdleConnections
	dbMaxOpenConnections := config.Database.MaxOpenConnections
	dbMaxIdleTime := config.Database.MaxIdleTime
	dbMaxLifeTime := config.Database.MaxLifeTime

	sqlDb, err := db.DB()
	sqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifeTime))
	sqlDb.SetMaxIdleConns(dbMaxIdleConnections)
	sqlDb.SetMaxOpenConns(dbMaxOpenConnections)
	sqlDb.SetConnMaxIdleTime(time.Duration(dbMaxIdleTime))

	// This is for analyzing the stats after setting a connection
	log.Info("@OnboardingConnectionPool MYSQL MAX Open Connections: ",
		sqlDb.Stats().MaxOpenConnections)
	log.Info("@DatabaseConnectionPool MYSQL Open Connections: ",
		sqlDb.Stats().OpenConnections)
	log.Info("@DatabaseConnectionPool MYSQL InUse Connections: ",
		sqlDb.Stats().InUse)
	log.Info("@DatabaseConnectionPool MYSQL Idle Connections: ", sqlDb.Stats().Idle)
	if err != nil {
		log.Errorf("DatabaseConnectionPool Error: %+v", err)
		return false, err
	}
	return true, nil
}

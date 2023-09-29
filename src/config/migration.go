package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
)

// RunDatabaseMigrations runs database migrations
func RunDatabaseMigrations(db *sql.DB, config Config) error {
	log.Info("Connected to DB, About to initiate migrations")

	// Validate migration files
	migrations := &migrate.FileMigrationSource{
		Dir: config.Database.MigrationsPath,
	}

	// Execute Database Migrations and store the migrations information in the table gorp_migrations
	n, err := migrate.Exec(db, config.Database.Driver, migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Error on migration: %+v", err)
		return err
	}

	log.Infof("Applied %+v migrations", n)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Errorf("Error occurred while closing connection: %+v", err)
		}
	}(db)
	return nil
}

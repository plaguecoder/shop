package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"shop/logger"

	// pq driver for postgres
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file" // get db migration from path
)

const appDatabaseMigrationPath = "./db/migrations"
const connectionURL = "postgres://postgres:@localhost:5432/shop_test?sslmode=disable"
const migrationsPath = "file://./db/migrations"

func LoadDatabase() *sqlx.DB {
	db, err := sqlx.Open("postgres", connectionURL)
	if err != nil {
		logger.Logger.Fatalf("failed to load the database: %s", err)
	}

	if err = db.Ping(); err != nil {
		logger.Logger.Fatalf("ping to the database host failed: %s", err)
	}

	return db
}

func RunDatabaseMigrations() error {
	db, err := sql.Open("postgres", connectionURL)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)

	err = m.Up()
	if err != nil {
		return err
	}

	logger.Logger.Println("Migration successful")

	return nil
}

func RollbackLatestMigration() error {
	db, err := sql.Open("postgres", connectionURL)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)

	err = m.Steps(-1)
	if err != nil {
		return err
	}

	return nil
}

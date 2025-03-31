package sqlite

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Driver   string
	Name     string
	PORT     string
	Username string
	Password string
}

/*
Initialize a new Database connection.
In case of SQLite connection, only the driver and
the name are needed. The name corresponds to the path of the database file.
*/

func (db *Config) Inits() (*sql.DB, error) {

	// Establish a new database connection
	databaseConnection, err := sql.Open(db.Driver, db.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Apply migrations
	err = applyMigrations(databaseConnection)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Enable foreign key constraints
	_, err = databaseConnection.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return databaseConnection, nil
}

// Function to apply migrations
func applyMigrations(db *sql.DB) error {
	// Path to the migrations folder
	currentDir, _ := os.Getwd()
	migrationDir := currentDir + "/pkg/db/migrations/sqlite/"

	migrationPath := "file://" + migrationDir

	// Create a new instance of migrate
	m, err := migrate.New(migrationPath, "sqlite://forum.db")
	if err != nil {
		return err
	}

	defer m.Close()

	// Apply all available migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

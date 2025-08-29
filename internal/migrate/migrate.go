package migrate

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL string) error {
	newMigrate, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		return fmt.Errorf("Error creating new migrations: %v", err)
	}
	defer newMigrate.Close()

	if err := newMigrate.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Error running migrations: %v", err)
	}
	log.Println("Migrations done")
	return nil
}

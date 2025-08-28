package migrate

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	mgrt, err := migrate.New(
		"file://migrations",
		os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	if err := mgrt.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	log.Println("migrations applied")
}

package main

import (
	_ "database/sql"
	"firstTestTask/internal/config"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	newDB, err := config.NewDatabase()
	fmt.Println(newDB, err)
}

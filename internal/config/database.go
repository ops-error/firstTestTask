package config

import (
	"fmt"
	"log"
	"os"

	_ "database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database структура для работы с БД
type Database struct {
	*sqlx.DB
}

// NewDatabase создает новое подключение к БД
func NewDatabase() (*Database, error) {
	//проверяет наличие .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить файл .env")
	}
	//строка подключения
	dataSourсeName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	//Коннект к БД
	dataBase, err := sqlx.Connect("postgres", dataSourсeName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к бд: %v", err)
	}

	//Проверка подключения
	if err := dataBase.Ping(); err != nil {
		return nil, fmt.Errorf("Ошибка ping БД: %v", err)
	}
	fmt.Println("Успешное подключение к БД")

	//если подключение успешное, то возвращает готовую бд по структуре Database
	return &Database{dataBase}, nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}

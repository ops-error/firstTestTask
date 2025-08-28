package repository

import "firstTestTask/internal/config"

// BaseRepository структура базового репозиторя с общими методами
// Содержит указатель на
type BaseRepository struct {
	dataBase *config.Database
}

// создание нового репозитория
func NewBaseRepository(dataBase *config.Database) *BaseRepository {
	return &BaseRepository{dataBase}
}

// создаёт таблицу если её не существует
func (baseRepository *BaseRepository) CreateTable(query string) error {
	_, err := baseRepository.dataBase.Exec(query)
	return err
}

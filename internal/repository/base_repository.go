package repository

import "firstTestTask/internal/config"

// базовый репозиторий с общими методами
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

package repository

import (
	"firstTestTask/internal/config"
)

type DeliveryRepository struct {
	*BaseRepository
}

func NewDeliveryRepository(db *config.Database) *DeliveryRepository {
	return &DeliveryRepository{BaseRepository: NewBaseRepository(db)}
}

// создаёт таблицу доставки
func (repo *DeliveryRepository) CreateDeliveryTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS delivery (
	    name VARCHAR(255) NOT NULL,
	    phone VARCHAR(255) NOT NULL,
	    zip VARCHAR(255) PRIMARY KEY,
	    city VARCHAR(255) NOT NULL,
	    address VARCHAR(255) NOT NULL,
	    region VARCHAR(255) NOT NULL,
	    email VARCHAR(255) NOT NULL,
	)`

	return repo.CreateTable(query)
}

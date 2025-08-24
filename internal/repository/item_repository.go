package repository

import (
	"firstTestTask/internal/config"
	"firstTestTask/internal/models"
)

type ItemRepository struct {
	*BaseRepository
}

func NewItemRepository(db *config.Database) *ItemRepository {
	return &ItemRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (repo *ItemRepository) CreateItemTable() error {
	query := `
CREATE TABLE IF NOT EXISTS item (
    chrt_id VARCHAR(255) PRIMARY KEY,
    price NUMERIC,
    rid VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    sale VARCHAR(255) NOT NULL,
    size VARCHAR(255) NOT NULL,
    total_price VARCHAR(255) NOT NULL,
    nm_id VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
)`
	return repo.CreateTable(query)
}

func (repo *ItemRepository) GetItemByID(id int) (*models.Item, error) {
	var item models.Item
	query := `SELECT * FROM item WHERE chrt_id = ?`
}

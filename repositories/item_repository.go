package repositories

import "gin-fleamarket/models"

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

type itemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &itemMemoryRepository{items: items}
}

func (r *itemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}


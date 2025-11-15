package repositories

import (
	"errors"
	"gin-fleamarket/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
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

func (r *itemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}

	return nil, errors.New("item not found")
}
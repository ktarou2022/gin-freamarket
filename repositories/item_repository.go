package repositories

import (
	"errors"
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updaterItem models.Item) (*models.Item, error)
	Delate(itemId uint) error
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

func (r *itemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *itemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

func (r *itemMemoryRepository) Delate(itemId uint) error {
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

type itemRepository struct {
	db *gorm.DB
}

func (r *itemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error 
	}

	return &newItem, nil
}

// Delate implements IItemRepository.
func (i *itemRepository) Delate(itemId uint) error {
	panic("unimplemented")
}

// FindAll implements IItemRepository.
func (i *itemRepository) FindAll() (*[]models.Item, error) {
	panic("unimplemented")
}

// FindById implements IItemRepository.
func (i *itemRepository) FindById(itemId uint) (*models.Item, error) {
	panic("unimplemented")
}

// Update implements IItemRepository.
func (i *itemRepository) Update(updaterItem models.Item) (*models.Item, error) {
	panic("unimplemented")
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &itemRepository{db}
}

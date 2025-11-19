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

func (r *itemRepository) Delate(itemId uint) error {
	deleteItem, err := r.FindById(itemId)
	if err != nil {
		return err
	}

	result := r.db.Delete(&deleteItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *itemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return &items, nil
}

func (r *itemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, itemId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}

	return &item, nil
}

func (r *itemRepository) Update(updaterItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updaterItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updaterItem, nil
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &itemRepository{db}
}

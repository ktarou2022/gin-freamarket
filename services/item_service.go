package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(createItemInput dto.CreateTypeInput) (*models.Item, error)
}

type itemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &itemService{repository}
}

func (s *itemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *itemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}

func (s *itemService) Create(createItemInput dto.CreateTypeInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

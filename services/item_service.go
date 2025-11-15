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
	Update(itemId uint, updateItmeInput dto.UpdateItemInput) (*models.Item, error)
	Delate(itemId uint) error
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

func (s *itemService) Update(itemId uint, updateItmeInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := s.FindById(itemId)
	if err != nil {
		return nil, err
	}

	if updateItmeInput.Name != nil {
		targetItem.Name = *updateItmeInput.Name
	}

	if updateItmeInput.Price != nil {
		targetItem.Price = *updateItmeInput.Price
	}

	if updateItmeInput.Description != nil {
		targetItem.Description = *updateItmeInput.Description
	}

	if updateItmeInput.SoldOut != nil {
		targetItem.SoldOut = *updateItmeInput.SoldOut
	}
	
	return s.repository.Update(*targetItem)
}

func (s *itemService) Delate(itemId uint) error {
	return s.repository.Delate(itemId)
}
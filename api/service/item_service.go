package service

import (
	repository "primeauction/api/Repository"
	"primeauction/api/models"
)

type ItemService struct {
	itemRepo *repository.ItemRepository
}

func NewItemService(itemRepo *repository.ItemRepository) *ItemService {
	return &ItemService{itemRepo: itemRepo}
}
func (s *ItemService) CreateItem(item *models.Item) error {
	return s.itemRepo.CreateItem(item)
}
func (s *ItemService) GetItemById(id string) (*models.Item, error) {
	return s.itemRepo.GetItemById(id)
}
func (s *ItemService) UpdateItem(item *models.Item) error {
	return s.itemRepo.UpdateItem(item)
}
func (s *ItemService) DeleteItem(id string) error {
	return s.itemRepo.DeleteItem(id)
}
func (s *ItemService) GetAllItems() ([]*models.Item, error) {
	return s.itemRepo.GetAllItems()
}

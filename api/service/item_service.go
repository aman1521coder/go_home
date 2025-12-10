package service

import (
	"errors"
	repository "primeauction/api/Repository"
	"primeauction/api/models"
)

type ItemService struct {
	itemRepo *repository.ItemRepository
}

func NewItemService(itemRepo *repository.ItemRepository) *ItemService {
	return &ItemService{itemRepo: itemRepo}
}

// CreateItem validates and creates an item with user_id
func (s *ItemService) CreateItem(userID string, item *models.Item) error {
	// Validate user_id is provided
	if userID == "" {
		return errors.New("user_id is required")
	}

	// Validate item data
	if item.Name == "" {
		return errors.New("item name is required")
	}

	if item.Price < 0 {
		return errors.New("price cannot be negative")
	}

	if item.SellingPrice < item.Price {
		return errors.New("selling price must be greater than or equal to cost price")
	}

	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	// Set user_id from parameter (ensures user can only create items for themselves)
	item.UserId = userID

	return s.itemRepo.CreateItem(item)
}

// GetItemById retrieves an item by ID
func (s *ItemService) GetItemById(id string) (*models.Item, error) {
	if id == "" {
		return nil, errors.New("item id is required")
	}
	return s.itemRepo.GetItemById(id)
}

// UpdateItem updates an item (with authorization check)
func (s *ItemService) UpdateItem(userID string, item *models.Item) error {
	// Get existing item to check ownership
	existingItem, err := s.itemRepo.GetItemById(item.Id)
	if err != nil {
		return err
	}

	// Check if user owns the item
	if existingItem.UserId != userID {
		return errors.New("unauthorized: you can only update your own items")
	}

	// Validate updates
	if item.Name == "" {
		return errors.New("item name is required")
	}

	if item.Price < 0 {
		return errors.New("price cannot be negative")
	}

	if item.SellingPrice < item.Price {
		return errors.New("selling price must be greater than or equal to cost price")
	}

	// Ensure user_id cannot be changed
	item.UserId = userID

	return s.itemRepo.UpdateItem(item)
}

// DeleteItem deletes an item (with authorization check)
func (s *ItemService) DeleteItem(itemID, userID string) error {
	if itemID == "" {
		return errors.New("item id is required")
	}

	// Get item to check ownership
	item, err := s.itemRepo.GetItemById(itemID)
	if err != nil {
		return err
	}

	// Check if user owns the item
	if item.UserId != userID {
		return errors.New("unauthorized: you can only delete your own items")
	}

	return s.itemRepo.DeleteItem(itemID)
}

// GetAllItems retrieves all items
func (s *ItemService) GetAllItems() ([]*models.Item, error) {
	return s.itemRepo.GetAllItems()
}

// GetItemsByUserID retrieves all items for a specific user
func (s *ItemService) GetItemsByUserID(userID string) ([]*models.Item, error) {
	if userID == "" {
		return nil, errors.New("user_id is required")
	}
	return s.itemRepo.GetItemsByUserID(userID)
}

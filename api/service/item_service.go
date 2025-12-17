package service

import (
	"errors"
	"mime/multipart"
	repository "primeauction/api/Repository"
	"primeauction/api/models"
	"primeauction/api/utils"
)

type ItemService struct {
	itemRepo  *repository.ItemRepository
	imageRepo *repository.ItemImageRepository
}

func NewItemService(itemRepo *repository.ItemRepository) *ItemService {
	return &ItemService{
		itemRepo:  itemRepo,
		imageRepo: repository.NewItemImageRepository(itemRepo.GetDB()),
	}
}

// ValidateImages validates multiple image files
func (s *ItemService) ValidateImages(fileHeaders []*multipart.FileHeader) error {
	if len(fileHeaders) == 0 {
		return nil // Images are optional
	}

	if len(fileHeaders) > utils.MaxImages {
		return errors.New("maximum 10 images allowed per item")
	}

	for i, fileHeader := range fileHeaders {
		if err := utils.ValidateImageFile(fileHeader); err != nil {
			return errors.New("image " + string(rune(i+1)) + ": " + err.Error())
		}
	}

	return nil
}

// CreateItem validates and creates an item with user_id
func (s *ItemService) CreateItem(userID string, item *models.Item, imagePaths []string) error {
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

	// Set primary image if we have images
	if len(imagePaths) > 0 {
		item.Image = imagePaths[0]
	}

	// Create item first
	if err := s.itemRepo.CreateItem(item); err != nil {
		return err
	}

	// Save images if provided
	if len(imagePaths) > 0 {
		if err := s.imageRepo.CreateImages(item.Id, imagePaths); err != nil {
			// If image save fails, we should ideally rollback item creation
			// For now, we'll just return the error
			return errors.New("failed to save images: " + err.Error())
		}
	}

	return nil
}

// GetItemById retrieves an item by ID
func (s *ItemService) GetItemById(id string) (*models.Item, error) {
	if id == "" {
		return nil, errors.New("item id is required")
	}
	return s.itemRepo.GetItemById(id)
}

// UpdateItem updates an item (with authorization check)
func (s *ItemService) UpdateItem(userID string, item *models.Item, imagePaths []string) error {
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

	// Set primary image if we have new images
	if len(imagePaths) > 0 {
		item.Image = imagePaths[0]
		// Delete old images
		s.imageRepo.DeleteImagesByItemID(item.Id)
		// Save new images
		if err := s.imageRepo.CreateImages(item.Id, imagePaths); err != nil {
			return errors.New("failed to save images: " + err.Error())
		}
	}

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

	// Delete associated images
	s.imageRepo.DeleteImagesByItemID(itemID)

	// Delete image files
	for _, img := range item.Images {
		utils.DeleteImage(img.ImagePath)
	}
	if item.Image != "" {
		utils.DeleteImage(item.Image)
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

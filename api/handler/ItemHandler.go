package handler

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"primeauction/api/models"
	"primeauction/api/service"
	"primeauction/api/utils"
	"strconv"
)

type ItemHandler struct {
	ItemService *service.ItemService
}

func NewItemHandler(itemService *service.ItemService) *ItemHandler {
	return &ItemHandler{ItemService: itemService}
}
func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.ItemService.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Ensure we always return an array, even if nil
	if items == nil {
		items = []*models.Item{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form (max 50MB for multiple images)
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get userID from JWT token (set by auth middleware)
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Create item from form values
	item := models.Item{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Images:      []models.ItemImage{}, // Initialize empty array
	}

	// Parse price and selling_price
	if priceStr := r.FormValue("price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			item.Price = price
		}
	}
	if sellingPriceStr := r.FormValue("selling_price"); sellingPriceStr != "" {
		if sellingPrice, err := strconv.ParseFloat(sellingPriceStr, 64); err == nil {
			item.SellingPrice = sellingPrice
		}
	}
	if quantityStr := r.FormValue("quantity"); quantityStr != "" {
		if quantity, err := strconv.Atoi(quantityStr); err == nil {
			item.Quantity = quantity
		}
	}

	// Handle multiple image uploads
	var imagePaths []string
	form := r.MultipartForm
	files := form.File["images"] // Note: "images" (plural) in form

	if len(files) > 0 {
		// Validate images first
		fileHeaders := make([]*multipart.FileHeader, len(files))
		copy(fileHeaders, files)

		if err := h.ItemService.ValidateImages(fileHeaders); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Save images
		var err error
		imagePaths, err = utils.SaveMultipleImages(fileHeaders, userID)
		if err != nil {
			http.Error(w, "Failed to save images: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Create item with images
	if err := h.ItemService.CreateItem(userID, &item, imagePaths); err != nil {
		// Cleanup images on failure
		utils.DeleteMultipleImages(imagePaths)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Reload item to get images
	createdItem, err := h.ItemService.GetItemById(item.Id)
	if err != nil {
		http.Error(w, "Item created but failed to load: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdItem)
}
func (h *ItemHandler) GetItemById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if id == "" {
		id = r.URL.Query().Get("id") // Fallback for old style
	}
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	item, err := h.ItemService.GetItemById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		id = r.URL.Query().Get("id") // Fallback
	}
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Get userID from JWT token (set by auth middleware)
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get existing item to preserve data
	existingItem, err := h.ItemService.GetItemById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Create item from form values with fallback to existing data
	item := models.Item{
		Id:          id,
		Name:        existingItem.Name,
		Description: existingItem.Description,
		Image:       existingItem.Image, // Keep existing image by default
	}

	if name := r.FormValue("name"); name != "" {
		item.Name = name
	}
	if description := r.FormValue("description"); description != "" {
		item.Description = description
	}

	// Parse other fields with fallback to existing values
	item.Price = existingItem.Price
	if priceStr := r.FormValue("price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			item.Price = price
		}
	}

	item.SellingPrice = existingItem.SellingPrice
	if sellingPriceStr := r.FormValue("selling_price"); sellingPriceStr != "" {
		if sellingPrice, err := strconv.ParseFloat(sellingPriceStr, 64); err == nil {
			item.SellingPrice = sellingPrice
		}
	}

	item.Quantity = existingItem.Quantity
	if quantityStr := r.FormValue("quantity"); quantityStr != "" {
		if quantity, err := strconv.Atoi(quantityStr); err == nil {
			item.Quantity = quantity
		}
	}

	// Handle new image uploads
	var imagePaths []string
	form := r.MultipartForm
	files := form.File["images"]

	if len(files) > 0 {
		// Validate images
		fileHeaders := make([]*multipart.FileHeader, len(files))
		copy(fileHeaders, files)

		if err := h.ItemService.ValidateImages(fileHeaders); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Save new images first
		var err error
		imagePaths, err = utils.SaveMultipleImages(fileHeaders, userID)
		if err != nil {
			http.Error(w, "Failed to save images: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err := h.ItemService.UpdateItem(userID, &item, imagePaths); err != nil {
		// Cleanup new images on failure
		utils.DeleteMultipleImages(imagePaths)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// SUCCESS: Now delete old images ONLY if new ones were successfully saved
	if len(imagePaths) > 0 {
		for _, img := range existingItem.Images {
			utils.DeleteImage(img.ImagePath)
		}
		if existingItem.Image != "" {
			utils.DeleteImage(existingItem.Image)
		}
	}

	// Reload item to get updated images
	updatedItem, err := h.ItemService.GetItemById(id)
	if err != nil {
		http.Error(w, "Item updated but failed to load: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
}
func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		id = r.URL.Query().Get("id")
	}
	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}
	// Get userID from JWT token (set by auth middleware)
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if err := h.ItemService.DeleteItem(id, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item deleted successfully"})
}

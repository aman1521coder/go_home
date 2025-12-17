package models

import "time"

type Item struct {
	Id           string      `json:"id"`
	UserId       string      `json:"user_id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Price        float64     `json:"price"`
	SellingPrice float64     `json:"selling_price"`
	Image        string      `json:"image"`  // Primary/thumbnail image (backward compatibility)
	Images       []ItemImage `json:"images"` // All images for the item
	Quantity     int         `json:"quantity"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	IsSold       bool        `json:"is_sold"`
}

package models

import "time"

type ItemImage struct {
	Id          string    `json:"id"`
	ItemId      string    `json:"item_id"`
	ImagePath   string    `json:"image_path"`
	DisplayOrder int      `json:"display_order"`
	CreatedAt   time.Time `json:"created_at"`
}


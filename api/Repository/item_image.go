package repository

import (
	"database/sql"
	"primeauction/api/models"
)

type ItemImageRepository struct {
	db *sql.DB
}

func NewItemImageRepository(db *sql.DB) *ItemImageRepository {
	return &ItemImageRepository{db: db}
}

// CreateImages creates multiple image records for an item
func (r *ItemImageRepository) CreateImages(itemID string, imagePaths []string) error {
	query := `INSERT INTO item_images (item_id, image_path, display_order) 
		VALUES ($1, $2, $3)`

	for i, path := range imagePaths {
		_, err := r.db.Exec(query, itemID, path, i)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImagesByItemID retrieves all images for an item
func (r *ItemImageRepository) GetImagesByItemID(itemID string) ([]models.ItemImage, error) {
	query := `SELECT id, item_id, image_path, display_order, created_at 
		FROM item_images WHERE item_id = $1 ORDER BY display_order`

	rows, err := r.db.Query(query, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.ItemImage
	for rows.Next() {
		var img models.ItemImage
		err := rows.Scan(&img.Id, &img.ItemId, &img.ImagePath, &img.DisplayOrder, &img.CreatedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

// DeleteImagesByItemID deletes all images for an item
func (r *ItemImageRepository) DeleteImagesByItemID(itemID string) error {
	query := `DELETE FROM item_images WHERE item_id = $1`
	_, err := r.db.Exec(query, itemID)
	return err
}

// DeleteImageByID deletes a single image by ID
func (r *ItemImageRepository) DeleteImageByID(imageID string) error {
	query := `DELETE FROM item_images WHERE id = $1`
	_, err := r.db.Exec(query, imageID)
	return err
}

// GetImageByID retrieves a single image by ID
func (r *ItemImageRepository) GetImageByID(imageID string) (*models.ItemImage, error) {
	query := `SELECT id, item_id, image_path, display_order, created_at 
		FROM item_images WHERE id = $1`

	var img models.ItemImage
	err := r.db.QueryRow(query, imageID).Scan(
		&img.Id,
		&img.ItemId,
		&img.ImagePath,
		&img.DisplayOrder,
		&img.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &img, nil
}

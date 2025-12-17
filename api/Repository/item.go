package repository

import (
	"database/sql"
	"errors"
	"primeauction/api/models"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

// GetDB returns the database connection (for creating other repositories)
func (r *ItemRepository) GetDB() *sql.DB {
	return r.db
}
func (r *ItemRepository) CreateItem(item *models.Item) error {
	query := `INSERT INTO items (user_id, name, description, price, selling_price, image, quantity, is_sold)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		item.UserId,
		item.Name,
		item.Description,
		item.Price,
		item.SellingPrice,
		item.Image,
		item.Quantity,
		item.IsSold,
	).Scan(&item.Id, &item.CreatedAt, &item.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepository) GetItemById(id string) (*models.Item, error) {
	query := `SELECT id, user_id, name, description, price, selling_price, image, quantity, is_sold, created_at, updated_at 
	FROM items 
	WHERE id = $1`
	item := &models.Item{}

	err := r.db.QueryRow(query, id).Scan(&item.Id, &item.UserId, &item.Name, &item.Description, &item.Price, &item.SellingPrice, &item.Image, &item.Quantity, &item.IsSold, &item.CreatedAt, &item.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("item not found")
	}
	if err != nil {
		return nil, err
	}

	// Load images for this item
	imageRepo := NewItemImageRepository(r.db)
	images, err := imageRepo.GetImagesByItemID(item.Id)
	if err == nil {
		item.Images = images
		// Set primary image if not set and we have images
		if item.Image == "" && len(images) > 0 {
			item.Image = images[0].ImagePath
		}
	}

	return item, nil
}
func (r *ItemRepository) UpdateItem(item *models.Item) error {
	query := `UPDATE items 
		SET name=$1, description=$2, price=$3, selling_price=$4, image=$5, quantity=$6, is_sold=$7, updated_at=CURRENT_TIMESTAMP
		WHERE id=$8
		RETURNING updated_at`

	err := r.db.QueryRow(
		query,
		item.Name,
		item.Description,
		item.Price,
		item.SellingPrice,
		item.Image,
		item.Quantity,
		item.IsSold,
		item.Id,
	).Scan(&item.UpdatedAt)

	if err == sql.ErrNoRows {
		return errors.New("item not found")
	}
	if err != nil {
		return err
	}
	return nil
}
func (r *ItemRepository) DeleteItem(id string) error {
	query := "DELETE FROM items WHERE id=$1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to get rows affected")
	}
	if rowsAffected == 0 {
		return errors.New("the item does not exist")

	}
	return nil
}
func (r *ItemRepository) GetAllItems() ([]*models.Item, error) {
	query := `SELECT id, user_id, name, description, price, selling_price, image, quantity, is_sold, created_at, updated_at 
		FROM items 
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		item := &models.Item{} // Create new item for each row
		err := rows.Scan(
			&item.Id,
			&item.UserId,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.SellingPrice,
			&item.Image,
			&item.Quantity,
			&item.IsSold,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// GetItemsByUserID retrieves all items for a specific user
func (r *ItemRepository) GetItemsByUserID(userID string) ([]*models.Item, error) {
	query := `SELECT id, user_id, name, description, price, selling_price, image, quantity, is_sold, created_at, updated_at 
		FROM items 
		WHERE user_id = $1 
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	imageRepo := NewItemImageRepository(r.db)

	for rows.Next() {
		item := &models.Item{}
		err := rows.Scan(
			&item.Id,
			&item.UserId,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.SellingPrice,
			&item.Image,
			&item.Quantity,
			&item.IsSold,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Load images for each item
		images, err := imageRepo.GetImagesByItemID(item.Id)
		if err == nil {
			item.Images = images
			// Set primary image if not set and we have images
			if item.Image == "" && len(images) > 0 {
				item.Image = images[0].ImagePath
			}
		}

		items = append(items, item)
	}
	return items, nil
}

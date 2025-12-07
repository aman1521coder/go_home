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
func (r *ItemRepository) CreateItem(item *models.Item) error {
	query := `INSERT INTO items (user_id, name, description, price, selling_price, image, quantity, is_sold, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.Exec(query, item.UserId, item.Name, item.Description, item.Price, item.SellingPrice, item.Image, item.Quantity, item.IsSold, item.CreatedAt, item.UpdatedAt)
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
	return item, nil
}
func (r *ItemRepository) UpdateItem(item *models.Item) error {
	query:="UPDATE items SET name=$1, description=$2, price=$3, selling_price=$4, image=$5, quantity=$6, is_sold=$7, updated_at=$8 WHERE id=$9"
	_, err:=r.db.Exec(query, item.Name, item.Description, item.Price, item.SellingPrice, item.Image, item.Quantity, item.IsSold, item.UpdatedAt, item.Id)
	if err!=nil {
		return err
	}
	return nil
}
func (r *ItemRepository) DeleteItem(id string )error{
	query:="DELETE FROM items WHERE id=$1"
	result,err:=r.db.Exec(query,id)
	if err!=nil{
		return err
	}
	rowsAffected,err:=result.RowsAffected()
	if err!=nil{
		return  errors.New("failed to get rows affected")
	}
	if rowsAffected==0{
		return  errors.New("the item does not exist")
		
	}
	return nil
}
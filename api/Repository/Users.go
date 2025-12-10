package repository

import (
	"database/sql"
	"errors"
	"primeauction/api/models"
)
type UserRepository struct{
	db *sql.DB
}
func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{db: db}
}
func (r *UserRepository) CreateUser(user *models.User) error{
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	return err
}
func (r *UserRepository)GetUserByID(id string )(*models.User, error){
	query := `SELECT id, username, email, password FROM users WHERE id = $1`
	row:=r.db.QueryRow(query, id)
	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err !=nil{
		return  nil,errors.New("user not found")

	}
return &user,nil
}
func (r *UserRepository)GetUserByEmail(email string)(*models.User,error){
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	row:=r.db.QueryRow(query, email)
	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err !=nil{
		return  nil,errors.New("user not found")

	}
	return &user,nil
}
func (r *UserRepository)GetAllUsers()([]models.User,error){
	query := `SELECT id, username, email, password FROM users`
	rows,err:=r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	users:=[]models.User{}
	for rows.Next(){
		var user models.User
		err:=rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err!=nil{
			return nil,err
		}
		users=append(users,user)
	}
	return users,nil
}
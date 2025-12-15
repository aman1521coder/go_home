package service

import (
	"errors"
	repository "primeauction/api/Repository"
	"primeauction/api/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil

}
func (s *UserService) CreateUser(user *models.User) error {
	if user.Username == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashpassword)
	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}
	return nil
}
func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}
	user.Password = ""
	return user, nil
}
func (s *UserService) GetUserById(id string)(*models.User,error){
	user,err:=s.userRepo.GetUserByID(id)
	if err!=nil{
		return nil,err
	}
	return user,nil
}
func (s *UserService) UpdateUser(id string,user *models.User) error {
	if user.Username == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")

	}
	hashpassword,err :=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil{
		return  errors.New("failed to hash password")
	}
	user.Password=string(hashpassword)
	if err:=s.userRepo.UpdateUser(id,user);err!=nil{
		return err
	}
	return nil

}
func (s *UserService)DeleteUser(id string) error{
	if err:=s.userRepo.DeleteUser(id);err!=nil{
		return err
	}
	return  nil
}
package service

import repository "primeauction/api/Repository"

type UserService struct {
	userRepo *repository.UserRepository
}
func NewUserService(userRepo *repository.UserRepository) *UserService{
	return &UserService{userRepo: userRepo}
}

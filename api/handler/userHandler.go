package handler
import (
	"encoding/json"
	"net/http"
	"primeauction/api/models"
	"primeauction/api/service"
)
type UserHandler struct{
	UserService *service.UserService
}
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}
func (h *UserHandler) GetUsers(w http.ResponseWriter,r *http.Request){
	users,err:=h.UserService.GetAllUsers()
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
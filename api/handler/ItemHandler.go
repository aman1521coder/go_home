package handler

import (
	"encoding/json"
	"net/http"
	"primeauction/api/models"
	"primeauction/api/service"
)
type ItemHandler struct{
	ItemService *service.ItemService
}
func NewItemHandler(itemService *service.ItemService) *ItemHandler {
	return &ItemHandler{ItemService: itemService}
}
func (h *ItemHandler) GetAllItems(w http.ResponseWriter,r *http.Request){
	items,err:=h.ItemService.GetAllItems()
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	  return
		}	
	json.NewEncoder(w).Encode(items)
}
func (h *ItemHandler)  createItem(w http.ResponseWriter,r *http.Request){
	var item models.Item
	if err:=json.NewDecoder(r.Body).Decode(&item);err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	userID:="tem-user-id"
	if err:=h.ItemService.CreateItem(userID,&item);err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
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
func (h *ItemHandler)  CreateItem(w http.ResponseWriter,r *http.Request){
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
func (h *ItemHandler) GetItemById(w http.ResponseWriter,r *http.Request){

	id:=r.URL.Query().Get("id")
	if id==""{
		http.Error(w,"ID is required",http.StatusBadRequest)
		return
	}
	item,err:=h.ItemService.GetItemById(id)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
func (h *ItemHandler)UpdateItem(w http.ResponseWriter,r *http.Request){
	id:=r.URL.Query().Get("id")
	if id==""{
		http.Error(w,"ID is required",http.StatusBadRequest)
		return
	}
	var item models.Item
	if err:=json.NewDecoder(r.Body).Decode(&item);err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	if err:=h.ItemService.UpdateItem(id,&item);err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
func (h *ItemHandler)DeleteItem(w http.ResponseWriter,r *http.Request){
	id:=r.URL.Query().Get("id")
	if id==""{
   http.Error(w,"Id is required",http.StatusBadRequest)
   return
	}
	// we get the user id from  the middleware for now lets ude temp user id
	userID:="temp-user-id"
	if err:=h.ItemService.DeleteItem(id,userID);err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"Item deleted successfully"})
}

package routes

import (
	"net/http"
	"primeauction/api/handler"
	"primeauction/api/middleware"
)

type Route struct {
	Path    string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func SetupRoutes(itemHandler *handler.ItemHandler, userHandler *handler.UserHandler) []Route {
	return []Route{
		// Public routes (no authentication required)
		{Path: "/api/auth/register", Method: "POST", Handler: userHandler.Register},
		{Path: "/api/auth/login", Method: "POST", Handler: userHandler.Login},
		{Path: "/api/items", Method: "GET", Handler: itemHandler.GetAllItems},      // Public: view all items
		{Path: "/api/items/{id}", Method: "GET", Handler: itemHandler.GetItemById}, // Public: view single item

		// Protected routes (require authentication)
		{Path: "/api/items", Method: "POST", Handler: middleware.AuthMiddleware(itemHandler.CreateItem)},
		{Path: "/api/items/{id}", Method: "PUT", Handler: middleware.AuthMiddleware(itemHandler.UpdateItem)},
		{Path: "/api/items/{id}", Method: "DELETE", Handler: middleware.AuthMiddleware(itemHandler.DeleteItem)},

		// User routes (protected)
		{Path: "/api/users", Method: "GET", Handler: middleware.AuthMiddleware(userHandler.GetUsers)},
		{Path: "/api/users/{id}", Method: "GET", Handler: middleware.AuthMiddleware(userHandler.GetUserById)},
		{Path: "/api/users/{id}", Method: "PUT", Handler: middleware.AuthMiddleware(userHandler.UpdateUser)},
		{Path: "/api/users/{id}", Method: "DELETE", Handler: middleware.AuthMiddleware(userHandler.DeleteUser)},
	}
}
func RegisterRoutes(routes *[]Route) {
	for _, route := range *routes {
		r := route
		http.HandleFunc(r.Path, func(w http.ResponseWriter, req *http.Request) {
			if req.Method == r.Method {
				r.Handler(w, req)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
	}
}

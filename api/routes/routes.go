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
		// Admin-only: Create items
		{Path: "/api/items", Method: "POST", Handler: middleware.AuthMiddleware(middleware.AdminMiddleware(itemHandler.CreateItem))},
		// Authenticated users: Update and delete items
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
	// Group routes by path to handle method routing
	pathHandlers := make(map[string]map[string]http.HandlerFunc)

	for _, route := range *routes {
		if pathHandlers[route.Path] == nil {
			pathHandlers[route.Path] = make(map[string]http.HandlerFunc)
		}
		pathHandlers[route.Path][route.Method] = route.Handler
	}

	// Register each unique path with a method router
	for path, methods := range pathHandlers {
		path := path       // Capture for closure
		methods := methods // Capture for closure

		http.HandleFunc(path, middleware.CORSMiddleware(func(w http.ResponseWriter, r *http.Request) {
			handler, exists := methods[r.Method]
			if !exists {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			handler(w, r)
		}))
	}
}

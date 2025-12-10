package routes
import(
	"net/http"
"primeauction/api/handler"

)
type Route struct {
	Path string
	Method string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func setupRoutes(itemHandler *handler.ItemHandler,userHandler *handler.UserHandler) [] Route{
	return [] Route{
		{Path: "/items", Method: "GET", Handler: itemHandler.GetAllItems},
		{Path: "/items", Method: "POST", Handler: itemHandler.CreateItem},
		{Path: "/items/{id}", Method: "GET", Handler: itemHandler.GetItemById},
		{Path: "/items/{id}", Method: "PUT", Handler: itemHandler.UpdateItem},
		{Path: "/items/{id}", Method: "DELETE", Handler: itemHandler.DeleteItem},
		{Path: "/users", Method: "GET", Handler: userHandler.GetUsers},
		{Path: "/users", Method: "POST", Handler: userHandler.C},
		{Path: "/users/{id}", Method: "GET", Handler: userHandler.GetUserById},
		{Path: "/users/{id}", Method: "PUT", Handler: userHandler.UpdateUser},
		{Path: "/users/{id}", Method: "DELETE", Handler: userHandler.DeleteUser},
	}
}
func registerRotes(routes *[]Route){
	for _, route := range *routes {
		r:=route
		http.HandleFunc(r.Path, func(w http.ResponseWriter, req *http.Request) {
			if req.Method == r.Method {
				r.Handler(w, req)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
	}
}


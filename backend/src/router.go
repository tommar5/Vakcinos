package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/tommar5/Saitinai/app/controllers"
	r "github.com/tommar5/Saitinai/app/repositories"
)

var Ucontroller = &c.UserController{UserRepo: r.UserRepo{}}
var Vcontroller = &c.VaccineController{VaccineRepo: r.VaccineRepo{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Authentication",
		"POST",
		"/get-token",
		Ucontroller.GetToken,
	},
	Route{
		"Index",
		"GET",
		"/user",
		Ucontroller.Index,
	},
	Route{
		"AddUser",
		"POST",
		"/user/add",
		c.AuthenticationMiddleware(Ucontroller.AddUser),
	},
	Route{
		"UpdateUser",
		"PUT",
		"/user/update",
		c.AuthenticationMiddleware(Ucontroller.UpdateUser),
	},
	// Get User by {id}
	Route{
		"GetUser",
		"GET",
		"/user/{id}",
		Ucontroller.GetUser,
	},
	// Delete User by {id}
	Route{
		"DeleteUser",
		"DELETE",
		"/user/{id}/delete",
		c.AuthenticationMiddleware(Ucontroller.DeleteUser),
	},
	// Search user with string
	Route{
		"SearchUser",
		"GET",
		"/user/search/{query}",
		Ucontroller.SearchUser,
	}}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

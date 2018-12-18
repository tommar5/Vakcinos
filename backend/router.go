package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Ucontroller = &UserController{UserRepo: UserRepo{}}
var Vcontroller = &VaccineController{VaccineRepo: VaccineRepo{}}

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

package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Login",
		"POST",
		"/authenticate",
		CreateTokenEndpoint,
	},
	Route{
		"Protected",
		"GET",
		"/protected",
		ProtectedEndpoint},
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
		Ucontroller.AddUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/user/update",
		Ucontroller.UpdateUser,
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
		Ucontroller.DeleteUser,
	},
	// Search user with string
	Route{
		"SearchUser",
		"GET",
		"/user/search/{query}",
		Ucontroller.SearchUser,
	},
	Route{
		"Index",
		"GET",
		"/vaccine",
		Vcontroller.Index,
	},
	Route{
		"AddVaccin",
		"POST",
		"/vaccine/add",
		Vcontroller.AddVaccine,
	},
	Route{
		"UpdateVaccin",
		"PUT",
		"/vaccine/update",
		Vcontroller.UpdateVaccine,
	},
	// Get User by {id}
	Route{
		"GetVAccin",
		"GET",
		"/vaccine/{id}",
		Vcontroller.GetVaccine,
	},
	// Delete User by {id}
	Route{
		"DeleteVaccin",
		"DELETE",
		"/vaccine/{id}/delete",
		Vcontroller.DeleteVaccine,
	},
	// Search user with string
	Route{
		"SearchVaccin",
		"GET",
		"/vaccine/search/{query}",
		Vcontroller.SearchVaccine,
	}}

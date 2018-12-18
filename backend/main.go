package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	mgo "gopkg.in/mgo.v2"
)

var vaccines *mgo.Collection
var users *mgo.Collection

func main() {
	session, err := mgo.Dial("172.17.0.1:27018")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("localhost err")
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Get posts collection
	vaccines = session.DB("vaccine").C("vaccines")
	users = session.DB("vaccine").C("users")

	router := NewRouter() // create routes

	if err := http.ListenAndServe(":8080", cors.AllowAll().Handler(router)); err != nil {
		log.Fatal(err)
	}
}

package app

import (
	"fmt"
	"log"
	"strings"

	db "github.com/tommar5/Saitinai/app/config"
	usr "github.com/tommar5/Saitinai/app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserRepo ...
type UserRepo struct{}

var userId = 1

// GetUsers returns the list of Users
func (r UserRepo) GetUsers() usr.Users {
	session, err := mgo.Dial(db.SERVER)

	if err != nil {

		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(db.DBNAME).C(db.COLLECTION)
	results := usr.Users{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// GetUserByID returns a unique User
func (r UserRepo) GetUserByID(id int) usr.User {
	session, err := mgo.Dial(db.SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(db.DBNAME).C(db.COLLECTION)
	var result usr.User

	fmt.Println("ID in GetUserById", id)

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// GetUsersByString takes a search string as input and returns users
func (r UserRepo) GetUsersByString(query string) usr.Users {
	session, err := mgo.Dial(db.SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(db.DBNAME).C(db.COLLECTION)
	result := usr.Users{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"secondname": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// AddUser adds a User in the DB
func (r UserRepo) AddUser(user usr.User) bool {
	session, err := mgo.Dial(db.SERVER)
	defer session.Close()

	userId++
	user.ID = userId
	session.DB(db.DBNAME).C(db.COLLECTION).Insert(user)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New User ID- ", user.ID)

	return true
}

// UpdateUser updates a User in the DB
func (r UserRepo) UpdateUser(user usr.User) bool {
	session, err := mgo.Dial(db.SERVER)
	defer session.Close()

	err = session.DB(db.DBNAME).C(db.COLLECTION).UpdateId(user.ID, user)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated User ID - ", user.ID)

	return true
}

// DeleteUser deletes an User
func (r UserRepo) DeleteUser(id int) string {
	session, err := mgo.Dial(db.SERVER)
	defer session.Close()

	// Remove user
	if err = session.DB(db.DBNAME).C(db.COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted User ID - ", id)
	// Write status
	return "OK"
}

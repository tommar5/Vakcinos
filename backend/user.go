package main

import (
	"time"

	"github.com/fxtlabs/date"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Firstname string        `json:"firstname" bson:"firstname"`
	Lastname  string        `json:"lastname" bson:"lastname"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password"`
	Birthday  date.Date     `json:"birthday" bson:"birthday"`
	Phone     string        `json:"phone" bson:"phone"`
	City      string        `json:"city" bson:"city"`
	Address   string        `json:"address" bson:"address"`
	Role      int           `json:"role" bson:"role"`
	CreatedAt time.Time     `json:"createdAt" bson:"created_at"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Users []User

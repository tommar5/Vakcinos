package main

import (
	"github.com/fxtlabs/date"
	"gopkg.in/mgo.v2/bson"
)

//Vaccine ...
type Vaccine struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	LOT        string        `json:"lot" bson:"lot"`
	ExpiryDate date.Date     `json:"expirydate" bson:"expirydate"`
	UseFromAge int           `json:"usefromage" bson:"usefromage"`
	Cost 	   string	 `json:"cost" bson:"cost"`
}

//Vaccines ...
type Vaccines []Vaccine

/*
gamintojas,
serijos numeris (LOT),
galiojimo laikas,
wiek od ktorego mozna skiepyti,
to ze kobieta neberemenna,
to ze 2 tyg nie pil antybiotykow,
to ze nie zrobil dwoch zywych wakcyn w jeden dzien,
to ze nie zobil wiecej niz 4 niezywe wakcyny w jeden dzien,
to ze od innych zrobionych wakcyn przeszlo choc 4 tyg
*/

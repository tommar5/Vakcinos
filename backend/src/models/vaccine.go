package app

import "github.com/fxtlabs/date"

//Vaccine ...
type Vaccine struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Producer   string    `json:"producer"`
	LOT        string    `json:"lot"`
	ExpiryDate date.Date `json:"expirydate"`
	UseFromAge int8      `json:"usefromage"`
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

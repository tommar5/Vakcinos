package app

import (
	"fmt"
	"log"
	"strings"

	db "github.com/tommar5/Saitinai/app/config"
	vac "github.com/tommar5/Saitinai/app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type VaccineRepo struct{}

var vaccineId = 10

// GetVaccines returns the list of Vaccines
func (r VaccineRepo) GetVaccines() vac.Vaccines {
	session, err := mgo.Dial(db.SERVER)

	if err != nil {

		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(db.DBNAME).C(db.COLLECTION)
	results := vac.Vaccines{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// GetVaccineById returns a unique Vaccine
func (r VaccineRepo) GetVaccineById(id int) vac.Vaccine {
	session, err := mgo.Dial(db.SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(db.DBNAME).C(db.COLLECTION)
	var result vac.Vaccine

	fmt.Println("ID in GetVaccineById", id)

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// GetVaccinesByString takes a search string as input and returns vaccines
func (r VaccineRepo) GetVaccinesByString(query string) vac.Vaccines {
	session, err := mgo.Dial(db.SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(db.DBNAME).C(db.COLLECTION)
	result := vac.Vaccines{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"title": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// AddVaccine adds a Vaccine in the DB
func (r VaccineRepo) AddVaccine(vaccine vac.Vaccine) bool {
	session, err := mgo.Dial(db.SERVER)
	defer session.Close()

	vaccineId += 1
	vaccine.ID = vaccineId
	session.DB(db.DBNAME).C(db.COLLECTION).Insert(vaccine)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New Vaccine ID- ", vaccine.ID)

	return true
}

// UpdateVaccine updates a Vaccine in the DB
func (r VaccineRepo) UpdateVaccine(vaccine vac.Vaccine) bool {
	session, err := mgo.Dial(db.SERVER)
	defer session.Close()

	err = session.DB(db.DBNAME).C(db.COLLECTION).UpdateId(vaccine.ID, vaccine)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated Vaccine ID - ", vaccine.ID)

	return true
}

// DeleteVaccine deletes an Vaccine
func (r VaccineRepo) DeleteVaccine(id int) string {
	session, err := mgo.Dial(db.SERVER)
	defer session.Close()

	// Remove vaccine
	if err = session.DB(db.DBNAME).C(db.COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted Vaccine ID - ", id)
	// Write status
	return "OK"
}

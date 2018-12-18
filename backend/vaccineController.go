package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//VaccineController ...
type VaccineController struct {
	VaccineRepo VaccineRepo
}

// Index GET /
func (c *VaccineController) Index(w http.ResponseWriter, r *http.Request) {
	result := []Vaccine{}
	if err := vaccines.Find(nil).Sort("name").All(&result); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		responseJSON(w, result)
	}
}

// AddVaccine POST /
func (c *VaccineController) AddVaccine(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	ua := r.Header.Get("Content-Type")
	if !strings.Contains(ua, "application/json") {
		responseCode(w, http.StatusUnsupportedMediaType)
		return
	}
	vaccine := &Vaccine{}
	err = json.Unmarshal(data, vaccine)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := Vaccine{}
	err = vaccines.Find(bson.M{"name": vaccine.Name}).One(&result)
	log.Printf("1")
	if err == nil {
		responseError(w, "Vaccine already exists", http.StatusConflict)
		log.Printf("2")
		return
	}
	log.Printf("4")

	vaccines.Insert(vaccine)
	if err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseCode(w, http.StatusCreated)
}

// SearchVaccine GET /
func (c *VaccineController) SearchVaccine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	query := vars["query"] // param query
	log.Println("Search Query - " + query)

	vaccines := c.VaccineRepo.GetVaccinesByString(query)
	data, _ := json.Marshal(vaccines)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	responseJSON(w, data)
}

// UpdateVaccine PUT /
func (c *VaccineController) UpdateVaccine(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	vaccine := &Vaccine{}
	err = json.Unmarshal(data, vaccine)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := vaccines.UpdateId(bson.ObjectIdHex(params["id"]), vaccine); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON(w, vaccine)
}

// GetVaccine GET - Gets a single vaccine by ID /
func (c *VaccineController) GetVaccine(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := Vaccine{}
	err := vaccines.Find(bson.M{"_id": bson.ObjectIdHex(params["id"])}).One(&result)
	if err != nil {
		responseError(w, "Invalid Vaccine ID", http.StatusBadRequest)
		return
	}
	responseJSON(w, result)
}

// DeleteVaccine DELETE /
func (c *VaccineController) DeleteVaccine(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	valid := bson.IsObjectIdHex(params["id"])
	if valid != true {
		responseCode(w, http.StatusNotFound)
		return
	}

	if err := vaccines.RemoveId(bson.ObjectIdHex(params["id"])); err != nil {
		responseError(w, err.Error(), http.StatusNotFound)
		return
	}
	responseCode(w, http.StatusNoContent)
}

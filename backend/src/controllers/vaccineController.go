package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	v "github.com/tommar5/Saitinai/app/models"
	vR "github.com/tommar5/Saitinai/app/repositories"
)

//VaccineController ...
type VaccineController struct {
	VaccineRepo vR.VaccineRepo
}

// Index GET /
func (c *VaccineController) Index(w http.ResponseWriter, r *http.Request) {
	vaccines := c.VaccineRepo.GetVaccines() // list of all vaccines
	data, _ := json.Marshal(vaccines)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddVaccine POST /
func (c *VaccineController) AddVaccine(w http.ResponseWriter, r *http.Request) {
	var vaccine v.Vaccine
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	log.Println(body)

	if err != nil {
		log.Fatalln("Error AddVaccine", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddVaccine", err)
	}

	if err := json.Unmarshal(body, &vaccine); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddVaccine unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(vaccine)
	success := c.VaccineRepo.AddVaccine(vaccine) // adds the product to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
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
	return
}

// UpdateVaccine PUT /
func (c *VaccineController) UpdateVaccine(w http.ResponseWriter, r *http.Request) {
	var vaccine v.Vaccine
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error UpdateVaccine", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateVaccine", err)
	}

	if err := json.Unmarshal(body, &vaccine); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateVaccine unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(vaccine.ID)
	success := c.VaccineRepo.UpdateVaccine(vaccine) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// GetVaccine GET - Gets a single vaccine by ID /
func (c *VaccineController) GetVaccine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	id := vars["id"] // param id
	log.Println(id)

	vaccineid, err := strconv.Atoi(id)

	if err != nil {
		log.Fatalln("Error GetVaccine", err)
	}

	vaccine := c.VaccineRepo.GetVaccineById(vaccineid)
	data, _ := json.Marshal(vaccine)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// DeleteVaccine DELETE /
func (c *VaccineController) DeleteVaccine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	id := vars["id"] // param id
	log.Println(id)

	vaccineid, err := strconv.Atoi(id)

	if err != nil {
		log.Fatalln("Error GetVaccine", err)
	}

	if err := c.VaccineRepo.DeleteVaccine(vaccineid); err != "" { // delete a vaccine by id
		log.Println(err)
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

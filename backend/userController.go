package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//Controller ...
type UserController struct {
	UserRepo UserRepo
}

/* Middleware handler to handle all requests for authentication */
func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	ua := req.Header.Get("Content-Type")
	log.Print(ua)
	if !strings.Contains(ua, "application/json") {
		responseCode(w, http.StatusUnsupportedMediaType)
		return
	}

	user := &User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := User{}
	err = users.Find(bson.M{"email": user.Email}).One(&result)
	if err != nil {
		responseError(w, "User not found", http.StatusNotFound)
		return
	}
	rez := CheckPasswordHash(user.Password, result.Password)
	if !rez {
		// responseError(w, "Wrong password!", http.StatusUnauthorized)
		responseError(w, "Wrong password!", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"password": user.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

// Get Authentication token GET /
func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
	// params := req.URL.Query()
	// log.Print(req.Header.Get("token"))
	token, _ := jwt.Parse(req.Header.Get("authorization"), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
	}
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Index GET /
func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	result := []User{}
	if err := users.Find(nil).Sort("-created_at").All(&result); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		responseJSON(w, result)
	}
}

// AddUser POST /
func (c *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
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
	user := &User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := User{}
	err = users.Find(bson.M{"email": user.Email}).One(&result)
	log.Printf("1")
	if err == nil {
		responseError(w, "User already exists", http.StatusConflict)
		log.Printf("2")
		return
	}
	log.Printf("4")
	user.CreatedAt = time.Now().UTC()
	password := user.Password

	user.Password, err = HashPassword(password)
	if err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users.Insert(user)
	if err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseCode(w, http.StatusCreated)
}

// SearchProduct GET /
func (c *UserController) SearchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	query := vars["query"] // param query
	log.Println("Search Query - " + query)

	products := c.UserRepo.GetUsersByString(query)
	data, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// UpdateUser PUT /
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := users.UpdateId(bson.ObjectIdHex(params["id"]), user); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON(w, user)
}

// GetUser GET - Gets a single user by ID /
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := User{}
	err := users.Find(bson.M{"_id": bson.ObjectIdHex(params["id"])}).One(&result)
	if err != nil {
		responseError(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	responseJSON(w, result)
}

// DeleteUser DELETE /
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	valid := bson.IsObjectIdHex(params["id"])
	if valid != true {
		responseCode(w, http.StatusNotFound)
		return
	}

	if err := users.RemoveId(bson.ObjectIdHex(params["id"])); err != nil {
		responseError(w, err.Error(), http.StatusNotFound)
		return
	}
	responseCode(w, http.StatusNoContent)
}

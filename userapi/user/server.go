/**
* API for accessing and manipulating user information
 */

package main

/* Imports */
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"
	// "net"
	// "os"
	// "strings"
)

/*var mongodb_server = os.Getenv("MONGO_SERVER")
var mongodb_database = os.Getenv("MONGO_DATABASE")
var mongodb_collection = os.Getenv("MONGO_COLLECTION")
var mongo_admin_database = os.Getenv("MONGO_ADMIN_DATABASE")
var mongo_username = os.Getenv("MONGO_USERNAME")
var mongo_password = os.Getenv("MONGO_PASS")*/

var mongodb_server = "localhost"
var mongodb_database = "userdb"
var mongodb_collection = "users"

func newUserServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	n.UseHandler(router)
	return n
}

/* Initializing resource URI */
func initRoutes(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/users", getAllUsers).Methods("GET") // Get all users OR a single user with given email ID
	router.HandleFunc("/users/{id}", getUserById).Methods("GET") // Get user with given ID
	router.HandleFunc("/users/signup", createUser).Methods("POST")	// Create new user
	router.HandleFunc("/users/signin", userSignIn).Methods("POST") // User login
	//router.HandleFunc("/users", deleteUser).Methods("DELETE") // Delete user with given email ID
	router.HandleFunc("/users/{id}", deleteUserById).Methods("DELETE") // Delete user with given user ID
	// router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/ping", checkPing(formatter)).Methods("GET") // Ping-pong test API
}

/* Setup response headers */
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

/* Handler for /ping */
func checkPing(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "User API is alive!"
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

/* Generates hash of password */
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

/* Compare and verify password hash */
func verifyPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func logger(message string) {
	fmt.Println(message)
}

/* 
 * Handler for /user. Fetch all users 
 * OR 
 * returns matching user record if "email" URL param is provided
 */
func getAllUsers(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	w.Header().Set("Content-Type", "application/json")

	/* Open DB session */
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		message := struct{ Message string }{"Error while connecting to database"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var result []User // Result set to store records

	/* Open DB connection*/
	c := session.DB(mongodb_database).C(mongodb_collection)

	query := bson.M{} // Empty query to fetch all records

	keys, _ := req.URL.Query()["email"] // Get 'email' URL param, if present any
	if len(keys) >= 1 {
		query = bson.M{"email": string(keys[0])}
	}

	err = c.Find(query).All(&result) // Fetch all users

	if err != nil {
		message := struct{ Message string }{"Error while fetching users"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode(result)
}

/**
 * Fetch record of given user
 * Handler for /user/{id}
 */
func getUserById(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	w.Header().Set("Content-Type", "application/json")

	/* Open DB session */
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		message := struct{ Message string }{"Error while connecting to database"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	params := mux.Vars(req)

	var id string = params["id"]

	query := bson.M{"id": id}

	var result bson.M

	if id == "" {
		message := struct{ Message string }{"User ID not provided"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	} else {
		/* Obtain DB connection */
		c := session.DB(mongodb_database).C(mongodb_collection)

		/* Execute query */
		err = c.Find(query).One(&result)

		/* Handle errors and return apt response*/
		if err != nil && err != mgo.ErrNotFound {
			text := "Error while fetching user information" + id
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(text)
			return
		}

		if err != nil && err == mgo.ErrNotFound {
			text := "No user found for ID: " + id
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(text)
			return
		}
	}

	json.NewEncoder(w).Encode(result)
}

/**
 * Handler for /users/signup URI
 */
func createUser(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	/* Open DB session */
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		message := "Error while connecting to database"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	/* Obtain DB connection */
	c := session.DB(mongodb_database).C(mongodb_collection)

	/* Validate that email ID is not empty */
	if len(user.Email) <= 0 {
		message := "Email not provided"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* Check for duplicate email address */
	query := bson.M{"email": user.Email}
	var result bson.M
	err = c.Find(query).One(&result)
	if err != nil && err != mgo.ErrNotFound {
		message := "Error while fetching data"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	} else if result != nil {
		message := "User with this email ID already exists"
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* Generate password hash */
	pHash, err := hashPassword(user.Password)
	if err != nil {
		message := "Error creating password hash"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* Assign new user ID */
	uniqueId := uuid.NewV4()
	user.Id = uniqueId.String()

	/* Set password hash */
	user.Password = pHash

	/* Commit new user info */
	err = c.Insert(user)
	if err != nil {
		message := "Error while creating new user"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* Return newly created user */
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func deleteUserById(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	w.Header().Set("Content-Type", "application/json")

	/* Open DB session */
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		message := struct{ Message string }{"Some error occured while connecting to database!!"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	/* Obtain DB connection */
	c := session.DB(mongodb_database).C(mongodb_collection)

	params := mux.Vars(req)

	if len(params) == 0 {
		message := "Email ID not provided"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	email:= params["email"]

	query := bson.M{"email": params["email"]}

	/* First, fetch the record with given email ID */
	var user User
	err = c.Find(query).One(&user)
	if err != nil && err != mgo.ErrNotFound {
		message := "Error while fetching the record from database"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	} else if err == mgo.ErrNotFound {
		message := "User not found"
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* Now, remove the record from database */
	err = c.Remove(query)
	if err != nil && err != mgo.ErrNotFound {
		message := "Error while deleting the record from database"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* Return the deleted record */
	json.NewEncoder(w).Encode(user)
}

/*
func updateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)
	params := mux.Vars(req)
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		message := struct{ Message string }{"Some error occured while connecting to database!!"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}
	err = session.DB(mongo_admin_database).Login(mongo_username, mongo_password)
	if err != nil {
		message := struct{ Message string }{"Some error occured while login to database!!"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{"id": params["id"]}
	updator := bson.M{
		"$set": bson.M{
			"firstname": person.Firstname,
			"lastname":  person.Lastname,
			"address":   person.Address,
			"password":  person.Password}}
	err = c.Update(query, updator)
	if err != nil && err != mgo.ErrNotFound {
		message := struct{ Message string }{"Some error occured while querying to database!!"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	} else if err == mgo.ErrNotFound {
		message := struct{ Message string }{"User not found!!"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(message)
		return
	}
	json.NewEncoder(w).Encode(struct{ Message string }{"user with id:" + params["id"] + " was Updated"})
}*/

/* Handler for /users/signin resource */
func userSignIn(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	/* Open DB session */
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		message := struct{ Message string }{"Error while connecting to database"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	/* Obtain DB connection */
	c := session.DB(mongodb_database).C(mongodb_collection)

	query := bson.M{"email": user.Email}

	var result User

	/* Check if user already exists */
	err = c.Find(query).One(&result)

	if err != nil && err != mgo.ErrNotFound {
		message := "Error while fetching user information"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* User not found in database */
	if err == mgo.ErrNotFound {
		message := "User not found"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* Verify password */
	tempPwd := user.Password
	passwordHash := result.Password
	if !verifyPasswordHash(tempPwd, passwordHash) {
		message := "Incorrect password"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* User verified */
	userData := bson.M{
		"email":     result.Email,
		"firstName": result.Firstname,
		"lastName":  result.Lastname,
		"address":   result.Address,
		"id":        result.Id}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userData)
}

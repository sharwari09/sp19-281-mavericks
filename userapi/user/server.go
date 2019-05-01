/**
* API for accessing and manipulating user information
 */

package main

/* Imports */
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"
)

var mongodb_server = os.Getenv("MONGO_SERVER")
var mongodb_database = os.Getenv("MONGO_DATABASE")
var mongodb_collection = os.Getenv("MONGO_COLLECTION")
//var allowed_origin = os.Getenv("ALLOWED_ORIGIN")
var dashboard_url = os.Getenv("DASHBOARD_URL")

func newUserServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router))
	return n
}

/* Initializing resource URI */
func initRoutes(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/users", getAllUsers).Methods("GET")            // Get all users OR a single user with given email ID
	router.HandleFunc("/users/{id}", getUserById).Methods("GET")       // Get user with given ID
	router.HandleFunc("/users/signup", createUser).Methods("POST")     // Create new user
	router.HandleFunc("/users/signin", userSignIn).Methods("POST")     // User login
	router.HandleFunc("/users", deleteUserByEmail).Methods("DELETE")   // Delete user with given email ID
	router.HandleFunc("/users/{id}", deleteUserById).Methods("DELETE") // Delete user with given user ID
	router.HandleFunc("/ping", checkPing(formatter)).Methods("GET")    // Ping-pong test API
	router.HandleFunc("/users/signin", optionsHandler(formatter)).Methods("OPTIONS")
}

/* Setup response headers */
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

/* Handler for /ping GET */
func checkPing(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "User API is alive!"
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

//API Options Handler
func optionsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		fmt.Println("options handler PREFLIGHT Request")
		return
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

/* Send new user ID to dashboard API */
func makeRequest(uid string) {
	url := dashboard_url + "?bucket=eventbrite&username=" + uid

	resp, err := http.Get(url)
	if err != nil {
		logger("Error while posting to dashboard")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger("Failed to post to dashboard")
		logger(string(body))
	}
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

	if (*req).Method == "OPTIONS" {
		fmt.Println("PREFLIGHT Request")
		return
	}

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

	keys, _ := req.URL.Query()["email"] // Get 'email' URL param, if present any
	if len(keys) >= 1 {
		query := bson.M{"email": string(keys[0])}
		var r User
		err = c.Find(query).One(&r) // Fetch only one users
		if err != nil && err != mgo.ErrNotFound {
			message := "Error while fetching users"
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(message)
			return
		}

		if err != nil && err == mgo.ErrNotFound {
			message := "No user found"
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(message)
			return
		}

		json.NewEncoder(w).Encode(r)
		return
	}

	/* Return all records */
	query := bson.M{}                // Empty query to fetch all records
	err = c.Find(query).All(&result) // Fetch all users

	if err != nil && err != mgo.ErrNotFound {
		message := "Error while fetching users"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	if err != nil && err == mgo.ErrNotFound {
		message := "No user found"
		w.WriteHeader(http.StatusNotFound)
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

	if (*req).Method == "OPTIONS" {
		fmt.Println("PREFLIGHT Request")
		return
	}

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

	if (*req).Method == "OPTIONS" {
		fmt.Println("PREFLIGHT Request")
		return
	}

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
	uniqueId, err := uuid.NewV4()
	if err != nil {
		logger("Error while creating unique ID for new user")
		return
	}
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

	/* Send new user ID to dashboard API */
	makeRequest(uniqueId.String())

	/* Return newly created user */
	w.WriteHeader(http.StatusCreated)
	userData := bson.M{
		"email": user.Email,
		"id":    user.Id}
	json.NewEncoder(w).Encode(userData)
}

/*
 * Deletes user with given email ID. Email ID is provided as URL param
 * Handler of /user DELETE
 */
func deleteUserByEmail(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	w.Header().Set("Content-Type", "application/json")

	if (*req).Method == "OPTIONS" {
		fmt.Println("PREFLIGHT Request")
		return
	}

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

	keys, _ := req.URL.Query()["email"] // Get 'email' URL param, if present any
	if len(keys) < 1 {
		message := "Email ID not provided"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	/* Get email ID */
	query := bson.M{"email": string(keys[0])}

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

/**
 * Delete user with given user ID
 * Handler of /user/{id} DELETE
 */
func deleteUserById(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	w.Header().Set("Content-Type", "application/json")

	if (*req).Method == "OPTIONS" {
		fmt.Println("PREFLIGHT Request")
		return
	}

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

	params := mux.Vars(req)
	if len(params) == 0 {
		message := "User ID not provided"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	query := bson.M{"id": params["id"]}

	/* First, fetch the record with given user ID */
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

/* Handler for /users/signin resource */
func userSignIn(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	w.Header().Set("Content-Type", "application/json")

	if (*req).Method == "OPTIONS" {
		fmt.Println("PREFLIGHT Request")
		return
	}

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
		w.WriteHeader(http.StatusNotFound)
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
		"id":        result.Id,
		"email":     result.Email,
		"firstname": result.Firstname}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userData)
}

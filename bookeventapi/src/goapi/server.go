package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	// "github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	// "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "127.0.0.1:27017"
var mongodb_database = "cmpe281"
var mongodb_collection = "bookings"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/book", bookHandler(formatter)).Methods("POST")
	mx.HandleFunc("/book", optionsHandler(formatter)).Methods("OPTIONS")
	mx.HandleFunc("/booking", bookingsHandler(formatter)).Methods("GET")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Go API version 1.0 alive!"})
	}
}

func optionsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		fmt.Println("options handler PREFLIGHT Request")
			return
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// API Process Orders 
func bookHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var b Bookings
		setupResponse(&w, req)
		fmt.Println("request body : ", req.Body)
		_ = json.NewDecoder(req.Body).Decode(&b)
		fmt.Println(req.Body)
		fmt.Println("booking details:", b.EventName)
		// Open MongoDB Session
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

        book_entry := Bookings{
			EventName: b.EventName,
			EventID: b.EventID,
			UserID: b.UserID,
			Price: b.Price}
		fmt.Println("book_entry", book_entry)
        fmt.Println( "EventID: ", book_entry.EventID , "Price: ", book_entry.Price)
        err = c.Insert(book_entry)
        if err != nil {
        		fmt.Println("Error while booking event: ", err)
                log.Fatal(err)
                formatter.JSON(w, http.StatusInternalServerError, 
				struct{ Response error }{err})
        }

		// Return booking Status
		formatter.JSON(w, http.StatusOK, struct{ Response string }{"Event successfully booked"})
	}
}
func bookingsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"get bookings API alive!"})
		setupResponse(&w, req)
		fmt.Println("inside getbookingsHandler")
		
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

        var results []Bookings
        // params := mux.Vars(req)
        // var userID string = params["userID"]
        // fmt.Printf( "User ID: %s", userID)
        // err = c.Find(bson.M{"userID": userID}).All(&results)
        err = c.Find(nil).All(&results)
        if err != nil {
        		fmt.Println("Error while getting booked events: ", err)
                log.Fatal(err)
                formatter.JSON(w, http.StatusInternalServerError, 
				struct{ Response error }{err})
        }
        fmt.Println(results)
        response := EventResponse{
			Count: len(results),
			BookedEvents: results}
		if len(results) > 0 {
			formatter.JSON(w, http.StatusOK, response)
		}else{
			formatter.JSON(w, http.StatusNoContent,
				struct{ Response string }{"No booked Events found"})
		}
	}
}


package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"log"
	"github.com/gorilla/handlers"
	"net/http"
	// "github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"bytes"
)

// MongoDB Config
//var mongodb_server = "52.52.92.197:27017"
//var mongodb_database = "cmpe281"
//var mongodb_collection = "bookings"
var mongodb_server = os.Getenv("MONGO_SERVER") + ":27017"
var mongodb_database = os.Getenv("DATABASE")
var mongodb_collection = os.Getenv("COLLECTION")
var dashboard_url = os.Getenv("DASHBOARD_URL")

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(mx))
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/book", bookHandler(formatter)).Methods("POST")
	mx.HandleFunc("/book", optionsHandler(formatter)).Methods("OPTIONS")
	mx.HandleFunc("/booking/{userID}", getuserBookings(formatter)).Methods("GET")
	mx.HandleFunc("/bookings/{eventID}", geteventBookings(formatter)).Methods("GET")
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

/* Send new booking data to dashboard */
func makeRequest(e *Bookings) {
	url := dashboard_url

	// Construct the message to be sent in request body
	message := map[string]interface{}{
		"bucket":    "eventbrite",
		"user_uuid": e.UserID,
		"orgId":     e.OrgId,
		"eventId":   e.EventID,
		"eventName": e.EventName,
		"location":  e.Location,
		"timeOfBooking": "25-04-2019",
		"date":      e.Date,
	}

	// Marshal message into JSON format
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println("Sent new booked event details to dashboard")
		log.Fatal(err)
		return
	}
}
// API to book an event
func bookHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var b Bookings
		setupResponse(&w, req)
		fmt.Println("request body : ", req.Body)
		_ = json.NewDecoder(req.Body).Decode(&b)
		fmt.Println("booking details:", b.EventName)
		// Open MongoDB Session
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var match Bookings

		err = c.Find(bson.M{"$and": []bson.M{bson.M{"eventID": b.EventID}, bson.M{"userID": b.UserID}}}).One(&match)

		if err == nil {
			fmt.Printf("Booking already exists!")
			formatter.JSON(w, http.StatusConflict, struct{ Response string }{"Booking already exists"})

		} else {

			bookID, _ := uuid.NewV4()

			book_entry := Bookings{
				EventName: b.EventName,
				EventID:   b.EventID,
				UserID:    b.UserID,
				Price:     b.Price,
				OrgId:     b.OrgId,
				Date:      b.Date,
				Location:  b.Location,
				BookID:    bookID.String()}
			fmt.Println("book_entry", book_entry)
			fmt.Println("EventID: ", book_entry.EventID, "Price: ", book_entry.Price)
			err = c.Insert(book_entry)
			if err != nil {
				fmt.Println("Error while booking event: ", err)
				log.Fatal(err)
				formatter.JSON(w, http.StatusInternalServerError,
					struct{ Response error }{err})
			}
			makeRequest(&book_entry)
			// Return booking Status
			formatter.JSON(w, http.StatusOK, struct{ Response string }{"Event successfully booked"})
		}

	}
}

// API to get booked event by userID
func getuserBookings(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		fmt.Println("inside getuserBookings")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var results []Bookings
		params := mux.Vars(req)
		var userID string = params["userID"]
		fmt.Printf("User ID: %s", userID)
		err = c.Find(bson.M{"userID": userID}).All(&results)
		if err != nil {
			fmt.Println("Error while getting booked events: ", err)
			log.Fatal(err)
			formatter.JSON(w, http.StatusInternalServerError,
				struct{ Response error }{err})
		}
		fmt.Println(results)
		response := EventResponse{
			Count:        len(results),
			BookedEvents: results}

		if len(results) > 0 {
			formatter.JSON(w, http.StatusOK, response)
		} else {
			formatter.JSON(w, http.StatusNoContent,
				struct{ Response string }{"No booked Events found"})
		}
	}
}

// API to get booked event by eventID

func geteventBookings(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		fmt.Println("inside geteventBookings")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var results []Bookings
		params := mux.Vars(req)
		var eventID string = params["eventID"]
		fmt.Printf("Event ID: %s", eventID)
		err = c.Find(bson.M{"eventID": eventID}).All(&results)
		if err != nil {
			fmt.Println("Error while getting booked events: ", err)
			log.Fatal(err)
			formatter.JSON(w, http.StatusInternalServerError,
				struct{ Response error }{err})
		}
		fmt.Println(results)
		response := EventResponse{
			Count:        len(results),
			BookedEvents: results}

		if len(results) > 0 {
			formatter.JSON(w, http.StatusOK, response)
		} else {
			formatter.JSON(w, http.StatusNoContent,
				struct{ Response string }{"No booked Events found"})
		}
	}
}


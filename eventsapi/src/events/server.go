package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/codegangsta/negroni"
	_ "github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	_ "github.com/satori/go.uuid"
	 "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var mongodb_server = "localhost:27017"
var mongodb_database = "eventbrite"
var mongodb_collection = "events"

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD",  "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	n.UseHandler(handlers.CORS(allowedHeaders,allowedMethods , allowedOrigins)(mx))
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/events", postEventHandler(formatter)).Methods("POST")
	mx.HandleFunc("/events", getAllEventsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/events/{eventName}", getEventHandler(formatter)).Methods("GET")
	mx.HandleFunc("/events/{eventName}", deleteEventhandler(formatter)).Methods("DELETE")
}

func postEventHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var e Event
		// Open MongoDB Session
		_ = json.NewDecoder(req.Body).Decode(&e)
		fmt.Println("Event: ", e)
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
		}
		defer session.Close()
        session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var match Event

		err = c.Find(bson.M{"date": time.Unix(e.Date, 0)}).One(&match)
		fmt.Println("Match: ", match)
		if err == nil{
			fmt.Printf("Event %s is already scheduled at the same time provided!", match.EventName)
		} else {
		event_entry := ScheduledEvent{
			EventId: e.EventId,
			EventName: e.EventName,
			Organizer: e.Organizer,	
			Date: time.Unix(e.Date, 0),
			Location: e.Location}

		err = c.Insert(event_entry)				
						
		if err != nil {
			fmt.Println("Error while adding Events: ", err)
			formatter.JSON(w, http.StatusInternalServerError, 
				struct{ Response error }{err})
		} else {
			formatter.JSON(w, http.StatusOK, 
				struct{ Response string }{"Event successfully added"})}
		
		}
	}
}

// API Get All Events Handler
func getEventHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Events API (Get) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var results []ScheduledEvent
		params := mux.Vars(req)
		fmt.Println("Params: ", params)
		var eventName string = params["eventname"]
		fmt.Printf( "Event Name: %s", eventName )
		err = c.Find(bson.M{"eventname": eventName}).All(&results)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)
		response := EventResponse{
			Count: len(results),
			Events: results}
		
		if len(results) > 0 {
			formatter.JSON(w, http.StatusOK, response)
		}else{
			formatter.JSON(w, http.StatusNoContent, 
				struct{ Response string }{"No Events found for the given ID"})
		}
	}
}

// API Get All Events Handler
func getAllEventsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Events API (Get) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var results []ScheduledEvent
		err = c.Find(nil).All(&results)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)
		if len(results) > 0 {
			formatter.JSON(w, http.StatusOK, results)
		}else{
			formatter.JSON(w, http.StatusNoContent, 
				struct{ Response string }{"No Events found for the given ID"})
		}
	}
}

func deleteEventhandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var eventName string = params["eventName"]
		fmt.Println("Event To Delete is: ", eventName)
		err = c.Remove(bson.M{"eventName": eventName})
		if err!=nil{
			fmt.Println("Event not found")
			formatter.JSON(w, http.StatusNotFound, "Event Not Found")
			return
		}
		formatter.JSON(w, http.StatusOK, "Event: " +
			eventName + " Deleted")
	}
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
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}
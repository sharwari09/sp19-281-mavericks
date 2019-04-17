package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/satori/go.uuid"
	_ "github.com/streadway/amqp"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

/*TODO: Support environment variables for Mongo Config*/

var mongodbServer = "localhost:27017"
var mongodbDatabase = "eventbrite"
var mongodbCollection = "events"

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

/*TODO: Connect to MongoDb only when admin user is provided*/

func postEventHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var e Event
		// Open MongoDB Session
		_ = json.NewDecoder(req.Body).Decode(&e)
		fmt.Println("Event: ", e)
		session, err := mgo.Dial(mongodbServer)
        if err != nil {
                panic(err)
		}
		defer session.Close()
        session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(mongodbCollection)
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
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Events API (Get) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(mongodbCollection)

		var results []ScheduledEvent
		params := mux.Vars(req)
		var eventName string = params["eventName"]
		fmt.Printf( "Event Name: %s", eventName )
		err = c.Find(bson.M{"eventName": eventName}).All(&results)
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
		session, err := mgo.Dial(mongodbServer)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Events API (Get) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(mongodbCollection)
		var results []ScheduledEvent
		err = c.Find(nil).All(&results)

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
				struct{ Response string }{"No Events found"})
		}
	}
}

func deleteEventhandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodbServer)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodbDatabase).C(mongodbCollection)
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
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	_ "github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	_ "github.com/satori/go.uuid"
	 "gopkg.in/mgo.v2"
    _"gopkg.in/mgo.v2/bson"
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
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/events", postEventHandler(formatter)).Methods("POST")
	//mx.HandleFunc("/events", getAllEventsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/events", getEventHandler(formatter)).Methods("GET").Queries("eventname")
}

func postEventHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var e Event
		// Open MongoDB Session
		_ = json.NewDecoder(req.Body).Decode(&e)
		fmt.Println("Event is: ", e.EventName)
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
		}
		defer session.Close()
        session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		event_entry := Event{
			EventName: e.EventName,
			Organiser: e.Organiser,	
			Schedule: e.Schedule}

		err = c.Insert(event_entry)				
						
		if err != nil {
			fmt.Println("Error while adding Events: ", err)
			formatter.JSON(w, http.StatusInternalServerError, 
				struct{ Response error }{err})
		} else {
			formatter.JSON(w, http.StatusOK, 
				struct{ Response string }{"Event successfully added"})
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

		var results []Event
		params := mux.Vars(req)
		fmt.Println("Params: ", params)
		var path string = params["path"]
		fmt.Println("Path is: ", path)
		//if len(eventName) != 0 {
		//	fmt.Printf( "Event Name: %s", eventName )
		//	err = c.Find(bson.M{"eventname": eventName}).All(&results)
		//} else {
		err = c.Find(nil).All(&results)
		//}
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
		var results []Event
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
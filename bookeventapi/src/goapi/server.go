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
// type bookings struct {
// 	Id             	bson.ObjectId `json:"id" bson:"_id"`	
// 	eventName	 	string		  `json:"eventName" bson:"eventName"`
// 	eventID			string        `json:"eventID" bson:"eventID"`
// 	userID			string	      `json:"userID" bson:"userID"`
// }
// RabbitMQ Config
//var rabbitmq_server = "rabbitmq"
//var rabbitmq_port = "5672"
//var rabbitmq_queue = "gumball"
//var rabbitmq_user = "guest"
//var rabbitmq_pass = "guest"

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
	// mx.HandleFunc("/gumball", gumballUpdateHandler(formatter)).Methods("PUT")
	// mx.HandleFunc("/order", gumballNewOrderHandler(formatter)).Methods("POST")
	// mx.HandleFunc("/order/{id}", gumballOrderStatusHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/order", gumballOrderStatusHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/orders", gumballProcessOrdersHandler(formatter)).Methods("POST")
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
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// // API Update Gumball Inventory
// func gumballUpdateHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
//     	var m gumballMachine
//     	_ = json.NewDecoder(req.Body).Decode(&m)		
//     	fmt.Println("Update Gumball Inventory To: ", m.CountGumballs)
// 		session, err := mgo.Dial(mongodb_server)
//         if err != nil {
//                 panic(err)
//         }
//         defer session.Close()
//         session.SetMode(mgo.Monotonic, true)
//         c := session.DB(mongodb_database).C(mongodb_collection)
//         query := bson.M{"SerialNumber" : "1234998871109"}
//         change := bson.M{"$set": bson.M{ "CountGumballs" : m.CountGumballs}}
//         err = c.Update(query, change)
//         if err != nil {
//                 log.Fatal(err)
//         }
//        	var result bson.M
//         err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
//         if err != nil {
//                 log.Fatal(err)
//         }        
//         fmt.Println("Gumball Machine:", result )
// 		formatter.JSON(w, http.StatusOK, result)
// 	}
// }

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


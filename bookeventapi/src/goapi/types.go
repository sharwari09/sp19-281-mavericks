package main
import (
    _ "gopkg.in/mgo.v2/bson"
)
// type gumballMachine struct {
// 	Id             	int 	
// 	CountGumballs   int    	
// 	ModelNumber 	string	    
// 	SerialNumber 	string	
// }

type Bookings struct {
	Id             	string `json:"id" bson:"_id,omitempty"`	
	EventName	 	string		  `json:"eventName" bson:"eventName"`
	EventID			string        `json:"eventID" bson:"eventID"`
	UserID			string		  `json:"userID" bson:"userID"`
	Price			int        	  `json:"price" bson:"price"`
}

//var orders map[string] order

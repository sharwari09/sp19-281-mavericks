package main

import "gopkg.in/mgo.v2/bson"

type Event struct{
	Id      	bson.ObjectId `json:"id" bson:"_id,omitempty"`
	EventName string
	Organiser string	
	Schedule string		
}

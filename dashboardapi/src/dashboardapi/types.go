package main

import "gopkg.in/mgo.v2/bson"

type Dashboard struct{
	Id      	bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
	UserId		string		  	`json:"_id" bson:"_id,omitempty"`
	Events		[]Event			`json:"event" bson:"event,omitempty"`
}

type Event struct {
    Id      	bson.ObjectId 	 `json:"id" bson:"_id,omitempty"`
	NumberOfBookings	int		 `json:"numberOfBookings" bson:"numberOfBookings,omitempty"`
	NumberOfViews		int		 `json:"numberOfViews" bson:"numberOfViews,omitempty"`
	Location			string   `json:"location" bson:"location"`
	StillOpen			bool     `json:"stillOpen" bson:"stillOpen"`
}


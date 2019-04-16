package main

import "gopkg.in/mgo.v2/bson"

type Dashboard struct{
	Id      	bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
	UserId		string		  	`json:"_id" bson:"_id,omitempty"`
	events		[]Event			`json:"event" bson:"event,omitempty"`
}

type Event struct {
    Id      	bson.ObjectId 	 `json:"id" bson:"_id,omitempty"`
	numberOfBookings	int		 `json:"numberOfBookings" bson:"numberOfBookings,omitempty"`
	numberOfViews		int		 `json:"numberOfViews" bson:"numberOfViews,omitempty"`
	location			string   `json:"location" bson:"location"`
	stillOpen			bool     `json:"stillOpen" bson:"stillOpen"`
}


package main

import (
	_ "gopkg.in/mgo.v2/bson"
)

type Bookings struct {
	Id        string `json:"id"`
	EventName string `json:"eventName" bson:"eventName"`
	EventID   string `json:"eventID" bson:"eventID"`
	UserID    string `json:"userID" bson:"userID"`
	BookID    string `json:"bookID" bson:"bookID"`
	Price     int    `json:"price" bson:"price"`
        OrgId     string `json:"orgId" bson:"orgId"`
	Date      string `json:"date" bson:"date"`
	Location  string `json:"location" bson:"location"`
}

type EventResponse struct {
	Count        int        `json:"count"`
	BookedEvents []Bookings `json:"bookedevents"`
}

//var orders map[string] order

package main

import "time"

type EventPayload struct{
	Id        string   `json:"id"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	EventId   string    `json:"eventId" bson:"eventId"`
	EventName string	`json:"eventName" bson:"eventName"`
	Location string		`json:"location" bson:"location"`
	Date int64		`json:"date" bson:"date"`
}

type ScheduledEvent struct{
	Id        string   `json:"id"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	EventId   string    `json:"eventId" bson:"eventId"`
	EventName string	`json:"eventName" bson:"eventName"`
	Location string		`json:"location" bson:"location"`
	Date time.Time		`json:"date" bson:"date"`
}

type EventResponse struct{
	Count	int		`json:"count"`
	Events  []ScheduledEvent		`json:"events"`
}

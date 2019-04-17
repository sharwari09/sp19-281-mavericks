package main

import "time"

type Event struct{
	EventId   string    `json:"eventId" bson:"eventId"`
	EventName string	`json:"eventName" bson:"eventName"`
	Organizer string	`json:"organizer" bson:"organizer"`
	Location string		`json:"location" bson:"location"`
	Date int64		`json:"date" bson:"date"`
}

type ScheduledEvent struct{
	EventId   string    `json:"eventId" bson:"eventId"`
	EventName string	`json:"eventName" bson:"eventName"`
	Organizer string	`json:"organizer" bson:"organizer"`
	Location string		`json:"location" bson:"location"`
	Date time.Time		`json:"date" bson:"date"`
}

type EventResponse struct{
	Count	int		`json:"count"`
	Events  []ScheduledEvent		`json:"events"`
}

package main

import "time"

type EventPayload struct {
	OrgId      string  `json:"orgId" bson:"orgId"`
	EventId    string  `json:"eventId" bson:"eventId"`
	EventName  string  `json:"eventName" bson:"eventName"`
	Location   string  `json:"location" bson:"location"`
	Date       string  `json:"date" bson:"date"`
	BucketName string  `json:"bucketname" bson:"bucketname"`
	Price      float32 `json:"price"`
}

type ScheduledEvent struct {
	OrgId      string    `json:"orgId"`
	EventId    string    `json:"eventId" bson:"eventId"`
	EventName  string    `json:"eventName" bson:"eventName"`
	Location   string    `json:"location" bson:"location"`
	Date       time.Time `json:"date" bson:"date"`
	BucketName string    `json:"bucketname" bson:"bucketname"`
	Price      float32   `json:"price"`
}

type EventResponse struct {
	Count  int              `json:"count"`
	Events []ScheduledEvent `json:"events"`
}

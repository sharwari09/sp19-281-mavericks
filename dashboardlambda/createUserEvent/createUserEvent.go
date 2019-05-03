package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

// AWS lambda request
type MyRequest struct {
	Bucket    string `json:"bucket"`    // Bucket in Riak
	Key       string `json:"user_uuid"` // user_id of user who is posting event
	EventID   string `json:"eventId"`   // id of the event
	EventName string `json:"eventName"` // name of the event
	Location  string `json:"location"`  // location
	Date      string `json:"date"`      // date
	Price     int    `json:"price"`     // price
}

// AWS lambda response
type MyResponse struct {
	Status bool `json:"status"`
}

// RIAK Network Load Balancer Reesponse

type Dashboard struct {
	PostedEvents []PostedEvent `json:"postedEvents"`
	BookedEvents []BookedEvent `json:"bookedEvents"`
}

type PostedEvent struct {
	OrgID            string `json:"orgId"`   // id of the organizer
	EventID          string `json:"eventId"` // id of the event
	EventName        string `json:"eventName"`
	Location         string `json:"location"`
	Date             string `json:"date"`
	NumberOfViews    int    `json:"numberOfviews"`
	NumberOfBookings int    `json:"numberOfBookings"`
	Price            int    `json:"price"` // price
}

type BookedEvent struct {
	OrgID         string `json:"orgId"`
	EventName     string `json:"eventName"`
	EventID       string `json:"eventId"`
	Date          string `json:"date"`
	TimeOfBooking string `json:"timeOfBooking"`
	Location      string `json:"location"`
	Price         int    `json:"price"` // price
}

func createUserEvent(request MyRequest) (MyResponse, error) {

	var nlb = os.Getenv("riak_cluster_nlb")
	var port = os.Getenv("nlb_port")
	var bucket = request.Bucket
	var key = request.Key
	var location = request.Location
	var eventID = request.EventID
	var date = request.Date
	var eventName = request.EventName
	var price = request.Price

	// Getting the details of the user from RIAK
	fmt.Printf("Request Object\n")
	fmt.Printf("%v\n", request)
	var url = fmt.Sprintf("http://%s:%s/buckets/%s/keys/%s", nlb, port, bucket, key)
	responseUserDetails, err := http.Get(url)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return MyResponse{Status: false}, err
	}
	defer responseUserDetails.Body.Close()
	responseUserDetailsBody, _ := ioutil.ReadAll(responseUserDetails.Body)
	var dashboard Dashboard
	errResponseUserDetailsBody := json.Unmarshal(responseUserDetailsBody, &dashboard)
	if errResponseUserDetailsBody != nil {
		log.Fatalf("ERROR: %s", errResponseUserDetailsBody)
		return MyResponse{Status: false}, errResponseUserDetailsBody
	}

	createEvent := PostedEvent{
		OrgID:            key,
		EventID:          eventID,
		NumberOfViews:    0,
		NumberOfBookings: 0,
		Location:         location,
		Date:             date,
		EventName:        eventName,
		Price:            price,
	}

	dashboard.PostedEvents = append(dashboard.PostedEvents, createEvent)
	dashboardString, _ := json.Marshal(dashboard)

	payload := bytes.NewBuffer(dashboardString)
	req, _ := http.NewRequest("PUT", url+"?returnbody=true", payload)
	req.Header.Add("Content-Type", "application/json")
	responsePostingEvent, errPostingEvent := http.DefaultClient.Do(req)

	if errPostingEvent != nil {
		log.Fatalf("ERROR: %s", errPostingEvent)
		fmt.Println("ERROR " + errPostingEvent.Error())
		return MyResponse{Status: false}, err
	}

	defer responsePostingEvent.Body.Close()
	body, _ := ioutil.ReadAll(responsePostingEvent.Body)

	fmt.Printf("Results: %s\n", string(body))
	return MyResponse{Status: true}, nil

}

func main() {
	lambda.Start(createUserEvent)
}

/*

API Gateway URL:
# request
{
	"bucket"    : "eventbrite",
	"user_uuid" : "",
	"eventId"   : "",
	"eventName" : "",
	"location"  : "",
	"date"      : "",
	"price"		: 233,
}


# response
{
	"status" : bool
}
*/

/*
nlb_port 80
riak_cluster_nlb riak-cluster-network-lb-d90a8ac266b9ee92.elb.us-west-2.amazonaws.com
*/

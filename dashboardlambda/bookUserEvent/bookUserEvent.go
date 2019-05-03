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
	Bucket        string `json:"bucket"`
	Key           string `json:"user_uuid"` // The one who is booking the event
	OrgID         string `json:"orgId"`     // Organizer of the event
	EventID       string `json:"eventId"`   // The event to be booked
	EventName     string `json:"eventName"`
	Date          string `json:"date"`
	TimeOfBooking string `json:"timeOfBooking"`
	Location      string `json:"location"`
	Price         int    `json:"price"` // price
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

type IncrementBooking struct {
	Bucket  string `json:"bucket"`
	Key     string `json:"user_uuid"`
	EventID string `json:"eventId"`
}

func bookUserEvent(request MyRequest) (MyResponse, error) {

	var nlb = os.Getenv("riak_cluster_nlb")
	var port = os.Getenv("nlb_port")
	var bucket = request.Bucket
	var key = request.Key
	var orgID = request.OrgID
	var eventID = request.EventID
	var eventName = request.EventName
	var date = request.Date
	var timeOfBooking = request.TimeOfBooking
	var location = request.Location
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

	bookEvent := BookedEvent{
		OrgID:         orgID,
		EventID:       eventID,
		EventName:     eventName,
		Date:          date,
		TimeOfBooking: timeOfBooking,
		Location:      location,
		Price:         price,
	}

	// Adding to list of booked Events
	dashboard.BookedEvents = append(dashboard.BookedEvents, bookEvent)
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

	go incrementBookingOfEvent(orgID, eventID, bucket)
	fmt.Printf("Results: %s\n", string(body))
	return MyResponse{Status: true}, nil

}

func incrementBookingOfEvent(orgID string, eventID string, bucket string) {
	fmt.Println("****Increment booking event****")
	bookingIncrementURL := os.Getenv("booking_increment_url")
	fmt.Println(bookingIncrementURL)
	data := IncrementBooking{
		Bucket:  bucket,
		Key:     orgID,
		EventID: eventID,
	}
	//data := map[string]string{"bucket": bucket, "user_uuid": orgID, "eventId": eventID}
	payload, _ := json.Marshal(data)
	var jsonStr = []byte(payload)
	req, _ := http.NewRequest("POST", bookingIncrementURL, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)

	if response.StatusCode != http.StatusOK {
		fmt.Println("The error is")
		fmt.Println(err.Error())
	}

	defer response.Body.Close()
	fmt.Println("****Increment booking event end****")
}

func main() {
	lambda.Start(bookUserEvent)
}

/*

API Gateway URL:
# request
{
	"bucket": "eventbrite",
  "user_uuid": "4a6bd3f7-9fa9-44e7-a79c-2258594fe0c",
  "orgId": "d4a03e22-2055-4167-8209-aaad98f29bcc",
  "eventId" : "",
  "eventName": "Arihant's Birthday",
  "date": "27-06-2019",
  "timeOfBooking": "25-04-2019",
  "location": "Brooklyn,NY"
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

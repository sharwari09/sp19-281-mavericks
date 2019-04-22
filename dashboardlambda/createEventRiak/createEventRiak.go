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

// AWS API Gateway request
type MyRequest struct {
	Bucket    string `json:"bucket"`
	Key       string `json:"user_uuid"`
	Location  string `json:"location"`
	Schedule  string `json:"schedule"`
	Organizer string `json:"organizer"`
	EventID   string `json:"event_id"`
}

// AWS API GATEWAY response
type MyResponse struct {
	Status bool `json:"status"`
}

// RIAK Network Load Balancer Reesponse

type Dashboard struct {
	PostedEvents []PostedEvent `json:"posted_events"`
	BookedEvents []BookedEvent `json:"booked_events"`
}

type PostedEvent struct {
	EventID          string `json:"event_id"`
	NumberOfViews    int    `json:"number_of_views"`
	NumberOfBookings int    `json:"number_of_bookings"`
	Location         string `json:"location"`
	Schedule         string `json:"schedule"`
	Organizer        string `json:"organizer"`
}

type BookedEvent struct {
	EventID       string `json:"event_id"`
	TimeOfBooking uint64 `json:"time_of_booking"`
}

func putUser(request MyRequest) (MyResponse, error) {

	var nlb = os.Getenv("riak_cluster_nlb")
	var port = os.Getenv("nlb_port")
	var bucket = request.Bucket
	var key = request.Key
	var location = request.Location
	var schedule = request.Schedule
	var organizer = request.Organizer
	var eventID = request.EventID

	// Getting the details of the user from RIAK
	var url = fmt.Sprintf("http://%s:%s/buckets/%s/keys/%s?returnbody=true", nlb, port, bucket, key)
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

	//payload := strings.NewReader("{\"posted_events\" : [],\"booked_events\" : []}")
	createEvent := PostedEvent{
		EventID:          eventID,
		NumberOfViews:    0,
		NumberOfBookings: 0,
		Location:         location,
		Schedule:         schedule,
		Organizer:        organizer,
	}

	dashboard.PostedEvents = append(dashboard.PostedEvents, createEvent)
	dashboardString, _ := json.Marshal(dashboard)
	//b := bytes.NewBuffer(dashboardString)
	//payload := strings.NewReader(dashboardString)
	payload := bytes.NewBuffer(dashboardString)
	req, _ := http.NewRequest("PUT", url, payload)
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
	lambda.Start(putUser)
}

/*

API Gateway URL:
# request
{
	Bucket    string `json:"bucket"`
	Key       string `json:"user_uuid"`
	Location  string `json:"location"`
	Schedule  string `json:"schedule"`
	Organizer string `json:"organizer"`
	EventID   string `json:"event_id"`
}


# response
{
	"status" : bool
}
*/

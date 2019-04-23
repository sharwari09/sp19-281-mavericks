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
	Bucket    string `json:"bucket"`
	Key       string `json:"user_uuid"`
	OrgID     string `json:"orgId"`
	EventName string `json:"eventName"`
	Location  string `json:"location"`
	Date      string `json:"date"`
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
	OrgID            string `json:"orgId"`
	EventName        string `json:"eventName"`
	Location         string `json:"location"`
	Date             string `json:"date"`
	NumberOfViews    int    `json:"numberOfviews"`
	NumberOfBookings int    `json:"numberOfBookings"`
}

type BookedEvent struct {
	OrgID         string `json:"orgId"`
	EventName     string `json:"eventName"`
	Date          string `json:"date"`
	TimeOfBooking uint64 `json:"timeOfBooking"`
}

func createUserEvent(request MyRequest) (MyResponse, error) {

	var nlb = os.Getenv("riak_cluster_nlb")
	var port = os.Getenv("nlb_port")
	var bucket = request.Bucket
	var key = request.Key
	var location = request.Location
	var orgID = request.OrgID
	var date = request.Date

	// Getting the details of the user from RIAK
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
		OrgID:            orgID,
		NumberOfViews:    0,
		NumberOfBookings: 0,
		Location:         location,
		Date:             date,
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
	Bucket    string `json:"bucket"`
	Key       string `json:"user_uuid"`
	OrgID     string `json:"orgId"`
	EventName string `json:"eventName"`
	Location  string `json:"location"`
	Date      string `json:"date"`
}


# response
{
	"status" : bool
}
*/


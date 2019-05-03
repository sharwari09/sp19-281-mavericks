package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyRequest struct {
	Bucket string `json:"bucket"`
	Key    string `json:"user_uuid"`
}

type MyResponse struct {
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

func getUserEvents(request MyRequest) (MyResponse, error) {

	var nlb = os.Getenv("riak_cluster_nlb")
	var port = os.Getenv("nlb_port")
	var bucket = request.Bucket
	var key = request.Key

	fmt.Printf("Request Object\n")
	fmt.Printf("%v\n", request)
	var url = fmt.Sprintf("http://%s:%s/buckets/%s/keys/%s", nlb, port, bucket, key)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return MyResponse{}, err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var data MyResponse
	errInUnmarshall := json.Unmarshal(body, &data)

	if errInUnmarshall != nil {
		fmt.Println("whoops:", err)
		return MyResponse{}, err
	}

	fmt.Printf("Results: %v\n", data)
	return data, nil

}

func main() {
	lambda.Start(getUserEvents)

}

/*

API Gateway URL:
# request
{
    "bucket" : "eventbrite"
	"user_uuid" : ""
}


# response
{
        "postedEvents":[
        {
            "orgId":"a418b7f2-1aec-4e70-a0c7-984fc12ff587",
            "eventId":"806ef8b7-8261-459e-903e-0abed74e1a6e",
            "eventName":"Summer bash",
            "location":"San Jose,CA",
            "date":"05/28/2019",
            "numberOfviews":0,
            "numberOfBookings":0
        }
    ],
    "bookedEvents":[]
}
*/

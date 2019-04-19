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
	key    string `json:"userid"`
}

type PostedEvent struct {
	EventID          string `json:"event_id"`
	NumberOfViews    int    `json:"number_of_views"`
	NumberOfBookings int    `json:"number_of_bookings"`
	TimeOfPosting    uint64 `json:"time_of_posting"`
}

type BookedEvent struct {
	EventID       string `json:"event_id"`
	TimeOfBooking uint64 `json:"time_of_booking"`
}

type MyResponse struct {
	PostedEvents []PostedEvent `json:"posted_events"`
	BookedEvents []BookedEvent `json:"booked_events"`
}

func dashboard(request MyRequest) (MyResponse, error) {
	/* http://{{riak-cluster-nlb}}:{{riak-cluster-nlb-port}}/buckets/{{bucket-name}}}/keys/{{key}} */
	var nlb = os.Getenv("riak-cluster-nlb")
	var port = os.Getenv("nlb-port")
	var bucket = request.Bucket
	var key = request.key

	fmt.Println("nlb %s, port, %s, bucket %s, key %s", nlb, port, bucket, key)
	var url = fmt.Sprintf("http://%s:%s/buckets/%s/keys/%s", nlb, port, bucket, key)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return nil, err
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		var data MyResponse
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("whoops:", err)
			return nil, err
		} else {
			fmt.Printf("Results: %v\n", data)
			return data, nil
		}

	}

}

func main() {
	lambda.Start(dashboard)
}

/*

API Gateway URL:
# request
{
    "bucket": "eventbrite",
    "key": "asp"
}

{"events":[{"1235":12},{"12456":12}]}
# response
{
    "posted_events" : [
    	{
			"event_id" : "12345",
			"number_of_views" : 20,
			"number_of_bookings" : 30,
			"time_of_posting" : 123894943444
		},
        {
    		"event_id" : "12385",
			"number_of_views" : 20,
			"number_of_bookings" : 30,
			"time_of_posting" : 123894943444
		}
	],

	"booked_events" : [
		{
			"event_id" : "12344",
			"time_of_booking" : 123894923334
		},
        {
    		"event_id" : "12346",
			"time_of_booking" : 123894923334
		},
        {
    		"event_id" : "18344",
			"time_of_booking" : 123894923334
		}
	]
}
*/

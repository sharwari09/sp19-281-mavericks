package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyRequest struct {
	Bucket string `json:"bucket"`
	Key    string `json:"username"`
}

type MyResponse struct {
	Status bool `json:"status"`
}

func putUser(request MyRequest) (MyResponse, error) {

	var nlb = os.Getenv("riak_cluster_nlb")
	var port = os.Getenv("nlb_port")
	var bucket = request.Bucket
	var key = request.Key

	var url = fmt.Sprintf("http://%s:%s/buckets/%s/keys/%s?returnbody=true", nlb, port, bucket, key)
	fmt.Println("" + url)
	payload := strings.NewReader("{\"posted_events\" : [],\"booked_events\" : []}")
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
		fmt.Println("ERROR " + err.Error())
		return MyResponse{Status: false}, err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

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
	"bucket" : "eventbrite"
	"username" : "arihant"
}


# response
{
	"status" : bool
}
*/

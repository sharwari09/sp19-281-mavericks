package main
 
import (
        "fmt"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyRequest struct {
        Name string `json:"What is your name?"`
        Age int     `json:"How old are you?"`
}
 
type MyResponse struct {
        Message string `json:"Answer:"`
}
 
func sample(request MyRequest) (MyResponse, error) {
        return MyResponse{Message: fmt.Sprintf("%s is %d years old!", request.Name, request.Age)}, nil
}
 
func main() {
        lambda.Start(sample)
}    

/*
# request
{
    "What is your name?": "Anoop",
    "How old are you?": 28
} 

# response
{
    "Answer": "Jim is 33 years old!"
} 
*/

// var eventURL= 'http://40.113.234.119:3000/'
// var userURL= 'http://40.122.67.229:5000/'
// var bookURL= 'http://10.250.184.217:4000/'
// var dashboardURL="https://lho0n8hxa3.execute-api.us-west-2.amazonaws.com/getUserEventDetails"
// var bucket = "eventbrite"

var eventURL= 'http://54.149.88.143:8000/eventapi/'
var userURL= "http://54.149.88.143:8000/userapi/"
var bookURL= 'http://54.149.88.143:8000/bookapi/'
var dashboardURL="https://lho0n8hxa3.execute-api.us-west-2.amazonaws.com/getUserEventDetails"
var bucket = "eventbrite"
var incrementEventBookingURL="https://7v6pqirtai.execute-api.us-west-2.amazonaws.com/increment-event-booking-prod"
var incrementEventViewURL="https://nprke4h3j8.execute-api.us-west-2.amazonaws.com/incrementUserEventView"

export{
    bookURL,
    userURL,
    eventURL,
    dashboardURL,
    bucket,
    incrementEventBookingURL,
    incrementEventViewURL
}
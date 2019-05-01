# Creating a USER key in RIAK

METHOD: GET

URL: ```https://lho0n8hxa3.execute-api.us-west-2.amazonaws.com/putUserDetails?bucket=eventbrite&username={username}```

Request
```
{
  "bucket": "eventbrite",
  "username": "83d81f94-f0c2-4bbe-a4bc-d5411b85b477"
}
```

Response
```
{
    "status": true
}
```

# Creating a Event for User in Riak

METHOD: POST

URL: ```https://k3gku1lix8.execute-api.us-west-2.amazonaws.com/createUserEvent```

Request
```
{
  "bucket": "eventbrite",
  "user_uuid": "a418b7f2-1aec-4e70-a0c7-984fc12ff587",
  "eventId": "806ef8b7-8261-459e-903e-0abed74e1a6e",
  "eventName": "Summer bash",
  "location": "San Jose,CA",
  "date": "05/28/2019"
}
```

Response
```
{
    "status": true
}
```


# List all events in riak

METHOD: POST

URL: ```https://lho0n8hxa3.execute-api.us-west-2.amazonaws.com/getUserEventDetails```

Request
```
{
    "bucket": "eventbrite",
    "user_uuid": "74840ad8-72ca-43e6-9654-1e6c470dcf7b"
}
```

Response
```
{
    "postedEvents": [
        {
            "orgId": "",
            "eventId": "",
            "eventName": "Holi Hai",
            "location": "Cupertino",
            "date": "05/27/2019",
            "numberOfviews": 0,
            "numberOfBookings": 0
        },
        {
            "orgId": "a418b7f2-1aec-4e70-a0c7-984fc12ff587",
            "eventId": "654ab43c-48be-4954-8018-4f3bd62f9b90",
            "eventName": "Holi Hai",
            "location": "Cupertino",
            "date": "05/27/2019",
            "numberOfviews": 0,
            "numberOfBookings": 0
        },
        {
            "orgId": "a418b7f2-1aec-4e70-a0c7-984fc12ff587",
            "eventId": "806ef8b7-8261-459e-903e-0abed74e1a6e",
            "eventName": "Summer bash",
            "location": "San Jose,CA",
            "date": "05/28/2019",
            "numberOfviews": 0,
            "numberOfBookings": 0
        }
    ],
    "bookedEvents": []
}
```


# Book User Event

METHOD: POST

URL: ```https://7g1vnr3vy6.execute-api.us-west-2.amazonaws.com/bookUserEvent```

Request
```
{
  "bucket": "eventbrite",
  "user_uuid": "4a6bd3f7-9fa9-44e7-a79c-2258594fe0c", 
  "orgId": "d4a03e22-2055-4167-8209-aaad98f29bcc",
  "eventId" : "d4a03e22-2055-4167-8209-aaad98f29bcc",
  "eventName": "Arihant's Birthday",
  "date": "27-06-2019",
  "timeOfBooking": "25-04-2019",
  "location": "Brooklyn,NY"
}
```

Response
```
{
    "postedEvents": [
        {
            "orgId": "",
            "eventId": "",
            "eventName": "Holi Hai",
            "location": "Cupertino",
            "date": "05/27/2019",
            "numberOfviews": 0,
            "numberOfBookings": 0
        },
        {
            "orgId": "a418b7f2-1aec-4e70-a0c7-984fc12ff587",
            "eventId": "654ab43c-48be-4954-8018-4f3bd62f9b90",
            "eventName": "Holi Hai",
            "location": "Cupertino",
            "date": "05/27/2019",
            "numberOfviews": 0,
            "numberOfBookings": 0
        },
        {
            "orgId": "a418b7f2-1aec-4e70-a0c7-984fc12ff587",
            "eventId": "806ef8b7-8261-459e-903e-0abed74e1a6e",
            "eventName": "Summer bash",
            "location": "San Jose,CA",
            "date": "05/28/2019",
            "numberOfviews": 0,
            "numberOfBookings": 0
        }
    ],
    "bookedEvents": []
}
```

# Get all events from User RIAK key

METHOD: POST

URL: ```https://lho0n8hxa3.execute-api.us-west-2.amazonaws.com/getUserEventDetails```

Request
```
{
    "bucket": "eventbrite",
    "user_uuid": "a418b7f2-1aec-4e70-a0c7-984fc12ff587"
}
```

Response
```
{
    "postedEvents": [
        {
            "orgId": "",
            "eventId": "",
            "eventName": "Holi Hai",
            "location": "Cupertino",
            "date": "05/27/2019",
            "numberOfviews": 0,
            "numberOfBookings": 0
        },
    ],
    "bookedEvents": []
}
```

# Increment User event Booking
Method: POST

URL: ```https://7v6pqirtai.execute-api.us-west-2.amazonaws.com/increment-event-booking-prod```

Request
```
{
  "bucket": "eventbrite",
  "user_uuid": "a418b7f2-1aec-4e70-a0c7-984fc12ff587",
  "eventId": "806ef8b7-8261-459e-903e-0abed74ee"
}
```

Response
```
{
    "status": true
}
```

# Increment User event Views
Method: POST

URL: ```https://nprke4h3j8.execute-api.us-west-2.amazonaws.com/incrementUserEventView```

Request
```
{
  "bucket": "eventbrite",
  "user_uuid": "a418b7f2-1aec-4e70-a0c7-984fc12ff587",
  "eventId": "806ef8b7-8261-459e-903e-0abed74ee"
}
```

Response
```
{
    "status": true
}
```


DB Structure
Key: 
`user_uuid`

```
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
}

type BookedEvent struct {
	OrgID         string `json:"orgId"`
	EventName     string `json:"eventName"`
	EventID       string `json:"eventId"`
	Date          string `json:"date"`
	TimeOfBooking string `json:"timeOfBooking"`
	Location      string `json:"location"`
}
```






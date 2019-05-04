## eventsapi API Documentation 

eventsapi collection contains POST, GET, and DELETE requests for /events api service

**GET Ping to events service**
```shell
http://{{server}}:3000/ping
```
-----------------------------------------------
POST Create Event
```shell
http://{{server}}:3000/events
```
HEADERS
Content-Typeapplication/json
```json
BODY
{
  "eventName": "Garba 2019",
  "orgId": "8D46ff8A44-3EAA-4019-BF7C-ffffd",
  "bucketname": "eventbrite",
  "location": "San Jose",
  "date": "02-02-2019",
  "price": 200
}
```
Request

| Parameter     | Type      | Description |
| ------------- |-----------| ------------|
| eventName     | string    | Event ID
| orgId         | string    | Organizer ID|
| bucketname    | string    | Bucketname  |
| location      | string    | Location    |
| date          | string    | Date of the event|
| price         | int       | Price of event ticket|

Response Parameters for Success (Status code: 200)

-----------------------------------------------
GET Get all events
```shell
http://{{server}}:3000/events
```
GET request gets all the create events

Response Parameters for Success (Status code: 200)

-----------------------------------------------

DEL Delete Event
```shell
http://{{server}}:3000/events/{{eventId}}
```
Deletes an event using its ID

Response Parameters for Success (Status code: 200)

-----------------------------------------------

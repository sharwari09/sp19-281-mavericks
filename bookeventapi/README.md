# BookEvent API


### Golang BookEvent API runs on Azure Kubernetes Engine

Kong URL for BookEvent API:
```http://54.149.88.143:8000/bookapi/```

**Attributes**

|Attribute Name    | Type    | Description|
|---------------|-------|------------|
|paymentId |String|Payment| Id         |
|userId |    String  |    User  |     Id |
|orderId |    String | Unique | Order Number |
|totalAmount |Double |Total | Amount Paid |
|status    | Boolean |    Payment Status (True = Paid, False = Cancelled) |
|paymentDate |DateTime    |Paid date|

## 1. Ping the API endpoint
   
**Request**
   
```
GET /ping
Content-Type: application/json
```
    
**Parameters**

None


**Response**

```
{
"Test": "Go API version 1.0 alive!"
}
```
    
## 2. Book an event

**Request**

```
POST /book
Content-Type: application/json
```

**Parameters**

|Parameter    |Type |    Description|
|-----|-----|------|
|eventName    |String|    Event Name|
|eventID|    String|    Event ID|
|userID|    String|    User ID|
|price|    Int|    Total amount of event tickets|
|orgId|    String|    Event Organizer ID|
|date|    String|    Date of Event|
|location|    String|    Location of the event|

**Response**

Parameters for Success (Status code: 200)
```
  {
   "Response": "Event successfully booked"
  }
```
## 3. Get booked events List by eventID

**Request**

```
GET /bookings/:eventID
Content-Type: application/json
```

**Parameters**

None

**Response**

Parameters for Success (Status code: 200)
```
{
"count": 1,
"bookedevents": [
    {
        "id": "",
        "eventName": "RSA conference",
        "eventID": "c9b4f552-c6e6-a957-2c385c08e",
        "userID": "a638ef3e-7c5d-4f9a-08-1",
        "bookID": "decbf878-c0e7-450c-bc11-be50ada439c3",
        "price": 80,
        "orgId": "a638ef3e-c5-4f9a-8108",
        "date": "2019-05-13",
        "location": "San Francisco"
    }
]
}
```

## 4. Get booked events List by userID

**Request**


 ```
 GET /booking/:userID
 Content-Type: application/json
 ```


**Parameters**

|Parameter    |Type |    Description|
|-----|-----|------|
|userID    |String|    User Id|

**Response**

Parameters for Success (Status code: 200)

```
{
"count": 1,
"bookedevents": [
    {
        "id": "",
        "eventName": "RSA conference",
        "eventID": "c9b4f552-c6e6-a957-2c385c08e",
        "userID": "a638ef3e-7c5d-4f9a-08-1",
        "bookID": "decbf878-c0e7-450c-bc11-be50ada439c3",
        "price": 80,
        "orgId": "a638ef3e-c5-4f9a-8108",
        "date": "2019-05-13",
        "location": "San Francisco"
    }
]
}
```

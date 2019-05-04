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

1. Ping the API endpoint
   **Request**
   
    ```GET /ping
    Content-Type: application/json```
    
   **Parameters**

   None


   **Response**
   
    ```{
    "Test": "Go API version 1.0 alive!"
    }```
    


2. Book an event

   **Request**
   
    ```POST /book
    Content-Type: application/json```
    
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

  {
   "Response": "Event successfully booked"
  }

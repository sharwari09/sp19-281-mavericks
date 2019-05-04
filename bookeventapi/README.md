# BookEvent API

This microservice handles booking of an event on our app. <br/>
The user can choose number of the seats and register for the event. <br/>
The service will contain API implementation of booking the events and storing them into the MongoDB cluster.<br/>

![Booking](images/bookmyevent-frontend/10.png)


### Golang BookEvent API runs on Azure Kubernetes Engine

Kong URL for BookEvent API:
```http://54.149.88.143:8000/bookapi/```

## Running GOLANG API locally

- Open a terminal
- Set your GOPATH to the project directoy
```shell
export GOPATH="Your Project directory"
```

Note you might need to setup your environment before running the API
- Get all the packages

```shell
make go-get
```

- Build your app
```shell
make go-build
```

- Run the app from terminal
```shell
make go-run
```

- Check whether the app is running
```shell
[negroni] listening on :4000
```

## Running the GO API in EC2 using docker
1. Install Docker
2. Start Docker
sudo systemctl start docker
sudo systemctl is-active docker
3. Login to your docker hub account
sudo docker login
4. Create Docker file
sudo vim Dockerfile

```shell
FROM golang:latest 
EXPOSE 4000
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
RUN cd /app ; go install events
CMD ["/app/bin/events"]
```

5. Build the docker image locally
```shell
sudo docker build -t bookevent .
sudo docker images
```

6. Push docker image to dockerhub
```shell
docker push bookevent:latest
```

7. Create Public EC2 Instance
Configuration:
```shell
AMI: CentOS 7 (x86_64) - with Updates HVM
Instance Type: t2.micro
VPC: cmpe281
Network: Public subnet (us-west-1c)
Auto Public IP: Yes
SG Open Ports: 22, 80, 8080, 3000, 8000
Key Pair: cmpe281-us-west-1
```
8. ssh to your ec2 instance, user name is centos
9. Create docker-compose yml file (with the environment variables set up)
10. Deploy go API for order sevice
```shell
docker-compose up
```
11. Clean Up docker environment when finished
```shell
docker stop events
docker rm events
docker rmi {imageid}
```


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

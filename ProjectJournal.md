# Team Mavericks
Team project for CMPE-281, Cloud Computing <br>
Submitted To: [Prof Paul Nguyen](https://github.com/paulnguyen)

Team members:
1. [Arihant Sai](https://github.com/Arihant1467)
2. [Pratik Bhandarkar](https://github.com/pratikb25)
3. [Sayali Patil](https://github.com/SayaliPatil)
4. [Sharwari Phadnis](https://github.com/sharwari09)
5. [Thol Chidambaram](https://github.com/thol)

## Key Contributions:
### Pratik Bhandarkar: (Users Microservice, Kong API Gateway)
- Design and deployment of a GO microservices to handle new user signup and user login
- Deployment of MongoDB Shard Clusters for this microservice
- Setup of Kong API gateway to expose all microservices over public network
- Worked in discussing architecture for the BookMyEvent application.
- Worked on finalizing backend databases for the app.

## Project Architecture Diagram :

![mavericks_latest](https://user-images.githubusercontent.com/4371600/57172019-6f088d00-6dcf-11e9-9d91-78b74062195c.jpg)

## Description :


#### 1. Frontend - User

**Technology Stack:** ReactJs, CSS 

The frontend User Client has been used by a user to log in to the application and the corresponding request will be transfered to corresponding user GO API via Kong API Gateway.

#### 2. Frontend - Event

**Technology Stack:** ReactJs, CSS 

The frontend Event client has been used by a logged in user to register a new event with all the details i.e eventname,date, location, price per ticket etc. in to the application and the corresponding request will be transferred to the corresponding event GO API via Kong Gateway.

#### 3. Frontend - Eventbooking

**Technology Stack:** ReactJs, CSS. 

The frontend EventBooking client has been used by a logged in user to book with all the details i.e eventname, date, location, price per ticket etc. in to the application and the corresponding request will be transferred to the corresponding event GO API via Kong Gateway. 

#### 4. Frontend - Dashboard 

**Technology Stack:** ReactJs, CSS

This page describes the association of the user with the app. It lists out the `Posted Events` and `Booked Events` of the user. 
**Posted Events:** These are the events which are posted by the user. This also provides analytics of the events
**Booked Events:** These are the events which are booked by the user

#### 5. Kong API Gateway

The Kong API Gateway is used to route the frontend request to the External Load Balancer for respective  GO APIS deployed on Azure Kubernetes Service (AKS).

#### 6. Go APIs

i] User API service has below features :

 Add a new user <br/>
 Delete a user <br/>
 Update user details <br/>
 Get user by userID <br/>
 
ii] Event API service has below features :

 Register a new event <br/>
 Delete an event <br/>
 Get all events <br/>
 Get an event by eventId <br/>
 
 iii] EventBooking API service has below features :
 
 Book a registered event <br/>
 Get bookings by userID <br/>
 Get bookings by userID <br/>
 
 iv] Dashboard API service has below features :
 
 View Posted Events <br/>
 View Booked Events <br/>
 View Analytics Regarding Posted Events <br/>
 
 
#### 7. Mongo DB Sharded cluster

The MongoDb sharded cluster consists of a replica set of 2 config server AWS EC2 instances, 2 shard server instances with 1 node in each shard server and 1 mongos instance as a query router to which respective GO API will send request.
For the details regarding sharded cluster, refer [BookEventAPI  sharded cluster](https://github.com/nguyensjsu/sp19-281-mavericks/blob/master/bookeventapi/MongoDB_sharding_for_bookeventAPI.md)

#### 8. Riak Cluster

The riak cluster consists of 5 nodes.

# AKF Scale Cube  :

## X-axis Scaling: 

 x-axis scaling or Horizontal duplication refers to running multiple identical copies of the application behind a load  
 balancer. In x-axis scaling, each server runs an identical copy of the service. It has been impleneted with multiple clones  
 of our APIs (i.e pods) behind an External Load Balancer in Azure Kubernetes Serivce (AKS).<br/>
     

## Y-axis Scaling:

 Y axis scaling refers to functional decomposition of a monolith service i.e. creating microservices. <br/>
 It has been implemented by separating all the services independently i.e userAPI, eventAPI and bookeventAPI with pods 
 deployed on Azure Kubernetes Service(AKS) <br/>

## Z-axis Scaling:

 Z axis scaling refers to splitting similar data into different servers.<br/>
 It has been implemented by using MongoDB sharded cluster with 2 config servers, 2 sharded replica sets and 1 mongos   
 server. MongoDb has been used to store user details, events details and booking details. <br/>





# Microservices Distribution
1. User signup - (Owner: Pratik Bhandarkar)<br/>
![User Microservice](images/users-service.png)
   - This microservice allows a user to sign up with our app.
   - During sign up a user has to provide his/her email id (which can later be used to login) and a password.
   - The user details and credentials are stored in the MongoDB cluster. A user, on signed up, can book register for an
     event and create his/her own events to be hosted.
2. Create Event - (Owner: Sharwari Phadnis)<br/>
![Create Event Microservice](images/events-service.jpg)
   - This microservice handles creation of events on our app. 
   - The user/organiser has to add details such as name, schedule, venue related to a particular 
     event that they want to host. 
   - The service will contain API implementation of creating events and storing them into the MongoDB cluster.
3. Browse Event - Thol
4. Book event - (Owner: Sayali)<br/>
![Book Event Microservice](images/bookevents-service.png)
   - This microservice handles booking of an event on our app.
   - The user can book the ticket and process it with payment
   - The service will contain API implementation of booking the events and storing them into the MongoDB cluster.
5. Dashboard - Arihant <br/>
![Dashboard Microservice](images/dashboard-service.png)
   - This microservice handles the dashboard of the user
   - Here the user will be able to view events posted by him and booked by him and analytics
   - The service will contain API implementation of dashboarding the events and retrieving from RIAK cluster

# [Network Partition](https://docs.hazelcast.org/docs/2.0/manual/html/ch13s08.html)
A partition is an event of node or network failure

Partition Tolerance : A partition tolrant system refers to a system which is functional despite of network or node failures

**CAP Theorem**
- C : Consistency

A guarantee that every node in a distributed cluster returns the same, most recent, successful write. Consistency refers to every client having the same view of the data

- A : Availibility

Every non-failing node returns a response for all read and write requests in a reasonable amount of time. The key word here is every. To be available, every node on (either side of a network partition) must be able to respond in a reasonable amount of time


- P : Partitioning

he system continues to function and upholds its consistency guarantees in spite of network partitions. Network partitions are a fact of life. Distributed systems guaranteeing partition tolerance can gracefully recover from partitions once the partition heals

It is a widely used concept in distributed system.


# Creativity in the use and application of topics and tools discussed in class

We have deployed our front end client on Heroku.
We have used kong as our microservice API gateway.
All our microservices are deployedas pods using Azure Kubernetes Service. All the microservices pods are running behind external load balancer on AKS achieving x-axis and y-axis as part of AKF scale cube.


# Website 
[http://cmpe281-mavericks-eventbrite.herokuapp.com/](http://cmpe281-mavericks-eventbrite.herokuapp.com/)

- Logging into the system
![Login](images/bookmyevent-frontend/1.png)

- Browsing thrugh the list
![Browse List](images/bookmyevent-frontend/2.png)

- Creating the event
![Creating the event](images/bookmyevent-frontend/3.png)

- Successful response
![Success in creation](images/bookmyevent-frontend/4.png)

- List gets updated
![New event added](images/bookmyevent-frontend/5.png)

- Update of Posted Events in dashboard
![Posted events update](images/bookmyevent-frontend/6.png)

- Signing out 
![Signing out](images/bookmyevent-frontend/7.png)

- Logging in as another user 
![Another user](images/bookmyevent-frontend/8.png)

- Browsing through the list and checking for the event which we posted earlier 
![checking for posted event](images/bookmyevent-frontend/9.png)

- Booking for the event we browsed 
![Booking](images/bookmyevent-frontend/10.png)

- Successful Booking of event
![Successful booking](images/bookmyevent-frontend/12.png)

- Checking the dashboard for the updates 
![Dashboard booked updates](images/bookmyevent-frontend/13.png)




# Meeting #1
## Week-1 (8 April 2019 to 14 April 2019)
## Discussion Points:

### Ideas: Few project ideas that have been thought and put on the table
    - A door dash kind of system for dry cleaners 
    - An online event-planning site from which you can create an event page, register attendees, and sell tickets online
    - Rental car application to book a car on rent after checking the car details
    - A eCommerce platform like Shopify for online stores and retail point-of-sale systems for selling and marketing
    - A online ticket system like Fandango for reserving seats and booking tickets for Movies, Plays and Shows around the user
    - Home automation along with community monitoring platform (smart communities)
    - Platform for drone management - control and operate drones for agriculture, monitoring and delivery.
    - Finalized Project Idea :
    - An online event-planning site from which you can create an event page, register attendees, and sell tickets online

### Action Items
    - A kanban board has been developed for the project to list various tasks of each member.
    - All the members have been assigned different microservices of the project

### Challenges
    - The API and database schema of the project
    - Listing out services which needs to be leveraged from different cloud providers for the project
    - Developing more scalable application and avoiding the single point of failure
    - Dashboard to monitor the end-to-end status of the system
    - Setup continuous integration and delivery

# Meeting #2 : April 10, 2019
## Discussion Points:
 - Assign microservices tasks to all
 - No need for MySQL, we can use MongoDB for all db needs
 - Load balancer
 - Network or HAProxy?
 - Create sharded MongoDB in a VPC share it with team
 - VPC peering for High Availibility
 - Application components in AKF Scaling
     - X-axis - Data replication
     - Y-axis - Split of service
     - Z-axis - Horizontal data partitioning
 - Redis for browsing the events to improve response time
 - REST API documentation / design per microservice
 - Create Postman REST API collection and share
 - Implementing kong or AWS API gateway
 - discussion regarding wow factor - Amazon EKS service
 - MongoDB: Cluster for storing Users, events, payments information

### Action Items :

 - decided scalable architecture of the application
 - decided how partition will affect nosql system
 
### Challenges :

- Finalizing replication of microservices
- Implementing kubernetes service


# Meeting #3 : April 13, 2019

## Discussion Points:

- Decided global redis for all microservices
- Riak db for dashboard
- List of microservices :

   Users : 	(Pratik) <br/>
      Login
      Signup
      update
   Events :    (Sharwari & Sayali)<br/>
       create event 
       list event
       book event
   Dashboard :  (Arihant & Thol)<br/>
      booked events
      Event and number of views
      No. of events posted
      Number of bookings

- event details :

   Event name 
   Fees
   Location
   Time
   Organizer

### Action Items:

  - Updated the KanBan style task board
  - Completed API development and bug fixes
  - Developed frontend application and integrated it with the backend

### Challenges :

  - CORS issue while connecting GO APIs from frontend
  - Communication between microservices


## References
- [CAP theorem](https://dzone.com/articles/understanding-the-cap-theorem)
- [Network Partition](https://aphyr.com/posts/281-jepsen-on-the-perils-of-network-partitions)
- [AWS Lambda](https://aws.amazon.com/lambda/)
- [Serverless computation](https://martinfowler.com/articles/serverless.html)
- [Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes)
- [Azure Kubernetes Service](https://docs.microsoft.com/en-us/azure/)
- [Prof Paul Nguyen](https://github.com/paulnguyen)


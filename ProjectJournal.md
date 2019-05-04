# Team Mavericks
Team project for CMPE-281, Cloud Computing <br>
Submitted To: [Prof Paul Nguyen](https://github.com/paulnguyen)

Team members:
1. [Arihant Sai](https://github.com/Arihant1467)
2. [Pratik Bhandarkar](https://github.com/pratikb25)
3. [Sayali Patil](https://github.com/SayaliPatil)
4. [Sharwari Phadnis](https://github.com/sharwari09)
5. [Thol Chidambaram](https://github.com/thol)


## Description :


#### 1. Frontend - User

Technology Stack: ReactJs, CSS. The frontend User Client has been used by a user to log in to the application and the corresponding request will be transfered to corresponding user GO API via Kong API Gateway.

#### 2. Frontend - Event

Technology Stack: ReactJs, CSS. The frontend Event client has been used by a logged in user to register a new event with all the details i.e eventname, date, location, price per ticket etc. in to the application and the corresponding request will be transferred to the corresponding event GO API via Kong Gateway.

#### 3. Frontend - Eventbooking

Technology Stack: ReactJs, CSS. The frontend EventBooking client has been used by a logged in user to book with all the details i.e eventname, date, location, price per ticket etc. in to the application and the corresponding request will be transferred to the corresponding event GO API via Kong Gateway. 

#### 4. Frontend - Dashboard 
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

#### 8. Riak Cluster

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




The riak cluster consists of 3 nodes.
# Microservices Distribution
1. User signup - (Owner: Pratik Bhandarkar)<br/>
![User Microservice](images/users-service.png)
   - This microservice allows a user to sign up with our app.
   - During sign up a user has to provide his/her email id (which can later be used to login) and a password.
   - The user details and credentials are stored in the MongoDB cluster. A user, on signed up, can book register for an
     event and create his/her own events to be hosted.
2. Create Event - (Owner: Sharwari Phadnis)<br>
![Create Event Microservice](images/events-service.jpg)
   - This microservice handles creation of events on our app. 
   - The user/organiser has to add details such as name, schedule, venue related to a particular 
     event that they want to host. 
   - The service will contain API implementation of creating events and storing them into the MongoDB cluster.
3. Browse Event - Thol
4. Book event - (Owner: Sayali)<br>
![Book Event Microservice](images/bookevents-service.png)
   - This microservice handles booking of an event on our app.
   - The user can book the ticket and process it with payment
   - The service will contain API implementation of booking the events and storing them into the MongoDB cluster.
5. Dashboard - Arihant
![Dashboard Microservice](images/dashboard-service.png)
   - This microservice handles the dashboard of the user
   - Here the user will be able to view events posted by him and booked by him and analytics
   - The service will contain API implementation of dashboarding the events and retrieving from RIAK cluster

# Minutes of Meeting
## Week-1 (8 April 2019 to 14 April 2019)
### Ideas: Few project ideas that have been thought and put on the table
    - A door dash kind of system for dry cleaners 
    - An online event-planning site from which you can create an event page, register attendees, and sell tickets online
    - Rental car application to book a car on rent after checking the car details
    - A eCommerce platform like Shopify for online stores and retail point-of-sale systems for selling and marketing
    - A online ticket system like Fandango for reserving seats and booking tickets for Movies, Plays and Shows around the user
    - Home automation along with community monitoring platform (smart communities)
    - Platform for drone management - control and operate drones for agriculture, monitoring and delivery.

### Finalized Project Idea
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

### April 10, 2019
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
 - Redis key-value store: To improve lookup performance while listing events
 - decided scalable architecture of the application
 - decided how partition will affect nosql system

### April 13, 2019

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


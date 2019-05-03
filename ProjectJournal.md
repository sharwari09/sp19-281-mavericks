# Team Mavericks
Team project for CMPE-281, Cloud Computing <br>
Submitted To: [Prof Paul Nguyen](https://github.com/paulnguyen)

Team members:
1. [Arihant Sai](https://github.com/Arihant1467)
2. [Pratik Bhandarkar](https://github.com/pratikb25)
3. [Sayali Patil](https://github.com/SayaliPatil)
4. [Sharwari Phadnis](https://github.com/sharwari09)
5. [Thol Chidambaram](https://github.com/thol)


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
   - This microservice allows a user to sign up with our app.
   - During sign up a user has to provide his/her email id (which can later be used to login) and a password.
   - The user details and credentials are stored in the MongoDB cluster. A user, on signed up, can book register for an
     event and create his/her own events to be hosted.
2. Create Event - (Owner: Sharwari Phadnis)<br>
   - This microservice handles creation of events on our app. 
   - The user/organiser has to add details such as name, schedule, venue related to a particular 
     event that they want to host. 
   - The service will contain API implementation of creating events and storing them into the MongoDB cluster.
3. Browse Event - Thol
4. Book event - (Owner: Sayali)<br>
   - This microservice handles booking of an event on our app.
   - The user can book the ticket and process it with payment
   - The service will contain API implementation of booking the events and storing them into the MongoDB cluster.
5. Dashboard - Arihant

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
 - VPC pairing for replication
 - Application components in AKF Scaling
     - X-axis - replication
     - Y-axis - microservices
     - Z-axis - mongodb sharding
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


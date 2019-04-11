# Team Mavericks
Team project for CMPE-281, Cloud Computing <br>
Submitted To: [Prof Paul Nguyen](https://github.com/paulnguyen)

Team members:
1. [Arihant Sai](https://github.com/Arihant1467)
2. [Pratik Bhandarkar](https://github.com/pratikb25)
3. [Sayali Patil](https://github.com/SayaliPatil)
4. [Sharwari Phadnis](https://github.com/sharwari09)
5. [Thol Chidambaram](https://github.com/thol)

# Microservices Distribution:

1. User signup - Pratik
2. Create Event - Sharwari<br>
   - This microservice handles creation of events on our app. 
   - The user/organiser has to add details such as name, schedule, venue related to a particular 
     event that they want to host. 
   - The service will contain API implementation of creating events and storing them into the MongoDB cluster.
3. Browse Event - Thol
4. Booking event - Sayali
5. Dashboard - Arihant


# Minutes of Meeting :

## Week-1 (8 April, 2019)
- Ideas: Few project ideas that have been thought and put on the table
    - A door dash kind of system for dry cleaners 
    - An online event-planning site from which you can create an event page, register attendees, and sell tickets online
    - Rental car application to book a car on rent after checking the car details
    - A eCommerce platform like Shopify for online stores and retail point-of-sale systems for selling and marketing
    - A online ticket system like Fandango for reserving seats and booking tickets for Movies, Plays and Shows around the user
    - Home automation along with community monitoring platform (smart communities)
    - Platform for drone management - control and operate drones for agriculture, monitoring and delivery.

- Finalized Project
    - An online event-planning site from which you can create an event page, register attendees, and sell tickets online

- Action Items:
    - A kanban board has been developed for the project to list various tasks of each member.
    - All the members have been assigned different microservices of the project

- Challenges
    - The API and database schema of the project
    - Listing out services which needs to be leveraged from different cloud providers for the project
    - Developing more scalable application and avoiding the single point of failure
    - Dashboard to monitor the end-to-end status of the system
    - Setup continuous integration and delivery

## Apr 10, 2019
 - Assign microservices tasks to all
 - No need for MySQL, we can use 
 - Assign micro-services tasks for each member
 - Assign micro-services tasks for each member
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
 - REST API documentation / design
 - Create Postman REST API collection and share


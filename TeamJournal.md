# TEAM JOURNAL #
<br>

## Counter Burger Application ##
## Team Members ##
- Aditi Kumari
- Aditya Doshatti
- Anand Kumar
- Ramya Iyer
- Sanjna Dhamejani

<br>

## Application structure ##
We have followed service-oriented architecture as a result of which the application has been decomposed into five microservices namely,
* Payments
* menu
* Order
* Location
* User

## Weekly development ##

### Week 1 (March 25 - March 31) ###
#### MoM (notes) ####
* Researched the various possible applications that we can develop
* Decided on the Application Domain
* Discussed the depth of scope that we want to cover
#### Result ####
* Decided to develop Counter burger Application
* Decided to use Microservices development technique

### Week 2 (April 1 - April 7) ###
#### MoM (notes) ####
* Decided the services to decompose 
* Discussed about the requirements and resources needed.
* Constructed the ncessary folder structure 
* Created Individual Journals to log personal progress
* Discussed about the relevant Database options
#### Result ####
* Everyone was assigned a microservice to work on.
* Go was decided to be the language for writing APIs.
* Decided on suitable databases for each microservice

### Week 3 (April 8 - April 14) ###
#### MoM (notes) ####
* Creation and configuration of databases
* started writing APIs
* Discussed the front-end technologies to be used.
* Flow of the website was decided

#### Result ####
* Narrowed down on React for developing UI

**Aditi**
* Created a MongoDB cluster for storing order details
* Setup backend to connect to the mongodb cluster and created a Ping Handler REST API for order service health check
* Created intital setup for the react front end and started building the microservice UI

**Aditya**
* Created a Sharded MongoDB cluster to store payments information 
* Setup backend to connect to the mongodb cluster and created a Ping Handler REST API for order service health chec
* Created intital setup for the react front end and started building the microservice UI

**Anand**
* Created a RIAK KV database cluster 
* Setup go backend to connect to RIAK cluster and tested it using PING handler API. 
* Created intital setup for the react front end and started building the microservice UI

**Ramya Iyer**
* Setup AWS RDS MySQL Database for storing user data
* Setup backend to connect to mysql and tested connection with the RDS was tested using API
* Created intital setup for the react front end and started building the microservice UI

**Sanjna**
* Setup MongoDB cluster with sharding
* Setup backend to connect to the mongodb cluster and created a Ping Handler REST API for order service health check
* Created intital setup for the react front end and started building the microservice UI

### Week 4 (April 15 - April 21) ###
#### MoM (notes) ####
* Discussed the necessary APIs for running the application

#### Result ####
* The ensuing week saw us working on the major functional APIs, the specifics of which are as follows

**Aditi**
* Created API for getting, posting and deleting the order details for cart 
* Completed UI development for the microservice

**Aditya**
* Created API for getting, posting and deleting the payment details  
* Completed UI development for the microservice

**Anand**
* Created API for getting, posting and deleting the Location details  
* Completed UI development for the microservice

**Ramya**
* Created API for getting, posting and deleting the User details  
* Completed UI development for the microservice

**Sanjna**
* Created API for getting, posting and deleting the Menu details  
* Completed UI development for the microservice

### Week 5 (April 22 - April 28) ###
#### MoM (notes) ####
* Discuss the appropriate deployment methods for individual services
* weighed pros and cons of various potential means

#### Result ####

**Aditi**
* Created docker hosts on two virtual machines with a network load balancer on top

**Aditya**
* Deployed the application on Google Kubernetes Engine

**Anand**
* Deployed the application on Amazon Elastic Container Service with an Application load balancer with Auto Scaling group

**Ramya**
* Deployed using dockerized containers on EC2 and classic load balancer 

**Sanjna**
* Deployed the application on Google Kubernetes Engine

### Week 6 (April 29 - May 3) ###

#### MoM (notes) ####
* integration of microservices using KONG API Gateway
* Front-end deployment startegy discussion
* End to end testing
#### Result ####
* Spin-up two EC2 instances with KONG containers with a elb to distribute load and avoid single point of failure for API Gateway
* Deployed the front-end using Heroku 
* End to end testing(manual) conducted successfully.


## Architecture Diagram ##

## AKF Scale Cube ##

## Demonstrating our application's ability to handle a network partition ##

1. A user is logged in
2. The user wants to search for location based on zip code
3. Even after a partition is created in the location database, the location details are always available.

For this purpose, an AP system - Riak KV is chosen which is always available even under partiton

Figure :- ELB showing all riak nodes in the cluster as working
![Screenshot from 2019-05-03 15-57-56](https://user-images.githubusercontent.com/43103509/57172496-abd78280-6dd5-11e9-9ead-9c605f84e551.png)

Figure :- Testing the location api from front end before partition
![Screenshot from 2019-05-03 15-59-18](https://user-images.githubusercontent.com/43103509/57172523-ee00c400-6dd5-11e9-96d4-185fb5502d06.png)

Figure :- ELB showing riak nodes after one node stops working
![Screenshot from 2019-05-03 15-59-51](https://user-images.githubusercontent.com/43103509/57172549-551e7880-6dd6-11e9-9dde-447842aa4733.png)

Figure :- Testing the location api from front end after partition
![Screenshot from 2019-05-03 16-00-52](https://user-images.githubusercontent.com/43103509/57172553-73847400-6dd6-11e9-8226-39a8e027d001.png)

## Testing ##








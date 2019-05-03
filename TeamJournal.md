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

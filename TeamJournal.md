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

![Architecture Diagram](./Images/ArchitectureDiagram.png "Architecture Diagram")

### Menu Component

1. APIs are running on a Kubernetes Cluster Running in GCP with 3 nodes in the Cluster.
2. Database used in MongoDb, Sharded over the shard key as objectId.
3. All the Database replica set in Private subnet in AWS and Mongos query router is in public subnet in AWS which is querried from GCP.

### Orders Component

1. APIs are running on Docker hosts on EC2 Instances in AWS which are under a Network Load Balancer.
2. APIs are talking with MongoDB replica Set in private subnet under same VPC.
3. A replica set of 5 EC2 instances form the Database Cluster

### Payments Component

1. APIs are on Private Kubernestes Cluster Running in GCP with 3 pods in Deployment running on 3 nodes in private subnet.
2. Only the Kubernetes Service is exposed with public IP.
3. Sharded MongoDB Database Replica set as DB running in the Private Subnet in the Google Cloud Platform.
4. The sharding is done on PaymentId as shar key which is UUID generated by Go API

### Users Component

1. Users APIs are running on 2 docker hosts running in AWS under a Application Load Balancer.
2. The Database used is SQL, using Relational Database Service provided by RDS
3. A 2 node RDS cluster which is replica set.

### Locations Component

1. APIs are running in Amazon Elastic Container Service having 2 tasks running
2. The ECS connects with Load balancer over the Riak KV cluster
3. Database used in Riak KV, which has best resiliency. A cluster of 3 nodes of Riak instances running in AWS are used.

## AKF Scale Cube ##

#### X Axis: Replication
 We have implemented replication of data in all the five microservices. 

   | Micoservice | Type	of Replication  |
   |----|----|
   | User | RDS Replica Set |
   | Location | Riak Cluster |
   | Menu | MongoDB Replica Set |
   | Order | MongoDB Replica Set |
   | Payment | MongoDB Replica Set |
   
  > The series of steps to show the replication of data in MongoDB Replica Set:
   
   1. Posting a order from Postman:
   
   ![Screenshot (434)](https://user-images.githubusercontent.com/28626925/57173229-3c4faa80-6e4a-11e9-8363-2b29adf9302c.png)
   
   
   
   2. Checking the record from the primary node of MongoDB Replica Set.
   
   ![aws_mongo1](https://user-images.githubusercontent.com/28626925/57173184-62c11600-6e49-11e9-8653-2576a44f8b01.png)
   
   
   3. Checking the record from secondary nodes of the replica set.
   
![aws img](https://user-images.githubusercontent.com/28626925/57173178-4b822880-6e49-11e9-95b1-224981c52b6b.png)

   
   ![Screenshot from 2019-05-04 08-34-53](https://user-images.githubusercontent.com/28626925/57173165-24c3f200-6e49-11e9-987a-469b0e82e6e1.png)
   
   


#### Y Axis: Microservices

 We have split our counterburger application into five microservices.
  ```
  User
  Location
  Menu
  Order
  Payment
  ```
 We have developed our application in such a way that even if one microservice goes down, the users will be able to use   
 the rest features of the applications without any difficulty.

#### Z Axis: Sharding

 We have implemented these series of steps to test sharding of data in Menu and Order microservices.
 
 These are series of steps taken to shard the data:
 1. Inserting the data in primary node and fetching the data from secondary nodes. This proves the consistency of data is  maintained. 
 2. Isolating a secondary node from its cluster and inserting the data in the cluster. This will help to prove if we read from the isolated secondary node, we will get the stale data.
 3. Removing the isolation mechanism used and connect the secordary node back to its cluter. If we read now, we will get the updated data. This will prove the replication of data.   
 
 Picture to show multiple EC2 instances in AWS for MongoDB Sharding (In Private Subnet - all instances, only mongos in public subnet)
 
 <img width="1440" alt="Screen Shot 2019-05-03 at 6 49 53 PM" src="https://user-images.githubusercontent.com/43122063/57173435-cb29dc00-6de4-11e9-9266-247378fe6db2.png">
 
 Picture to show shards from config server nodes.

<img width="1440" alt="Screen Shot 2019-05-03 at 7 21 21 PM" src="https://user-images.githubusercontent.com/43122063/57173396-23141300-6de4-11e9-964f-d920cf3d4b0a.png">

<img width="1440" alt="Screen Shot 2019-05-03 at 7 22 21 PM" src="https://user-images.githubusercontent.com/43122063/57173399-3626e300-6de4-11e9-904f-e948545201e5.png">

Picture to show Menu collection data from Mongos Instance 

<img width="1440" alt="Screen Shot 2019-05-03 at 7 29 24 PM" src="https://user-images.githubusercontent.com/43122063/57173406-59519280-6de4-11e9-9ac1-afc997fe50b2.png">

Picture to show shard distribution across multiple shard servers. Shard key for Menu database is Object Id.

<img width="1440" alt="Screen Shot 2019-05-03 at 7 29 38 PM" src="https://user-images.githubusercontent.com/43122063/57173452-04fae280-6de5-11e9-8e8d-5aa58fd9b704.png">

Picture to show data split in Shard server 1.1 (rs0 - primary)

<img width="1440" alt="Screen Shot 2019-05-03 at 7 36 34 PM" src="https://user-images.githubusercontent.com/43122063/57173467-2bb91900-6de5-11e9-8b78-341ffe5563cc.png">

Picture to show data split in Shard server 1.1 (rs0 - secondary)

<img width="1440" alt="Screen Shot 2019-05-03 at 7 37 26 PM" src="https://user-images.githubusercontent.com/43122063/57173487-70dd4b00-6de5-11e9-997e-052706c57d05.png">

Picture to show data split in Shard server 1.1 (rs1 - primary)

<img width="1440" alt="Screen Shot 2019-05-03 at 7 37 42 PM" src="https://user-images.githubusercontent.com/43122063/57173490-82265780-6de5-11e9-8950-1b40bab8d72a.png">

Picture to show data split in Shard server 1.1 (rs1 - secondary)

<img width="1440" alt="Screen Shot 2019-05-03 at 7 38 15 PM" src="https://user-images.githubusercontent.com/43122063/57173500-a2eead00-6de5-11e9-983c-1b1333d585e9.png">

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

The microservices are not interdependent. Thus, even if one service is down, the others work fine.

Figure :- Making Payments down
![Screen Shot 2019-05-03 at 8 30 02 PM](https://user-images.githubusercontent.com/43103509/57173327-cbc17300-6de2-11e9-8c66-00bb699b9bea.png)

Figure :- Cart working fine

![Screen Shot 2019-05-03 at 8 36 28 PM](https://user-images.githubusercontent.com/43103509/57173351-32df2780-6de3-11e9-95a3-c079db4fe343.png)


## Testing ##

#### [Complete Testing with Screenshots](./DemoScreenShots.md)

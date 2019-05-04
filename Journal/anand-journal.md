Project kickoff
Today we had our first team project meet in which we discussed about the potential domains on which we can work and the scope of our project.

### discussion highlights ###
-> project domain: transaction management <br/>
-> project architecture: Service Oriented Architecture(SOA) <br/>
-> Potential choice of database: MongoDB <br/>
-> Development model: Agile <br/><br>

Date: 04/09/2019

-> Possible microservices
-> Distribution of work and guaging suitability of members based on their background
-> Discussion on possible wow factor implementation
-> going through previous year's work to have a better understanding of the project


-> Created a RIAK cluster of three database instances
-> Placed a load balancer on top of the database machines
-> inserted data in the cluster

-> made an image of the application
-> tagged and pushed the image on dockerhub

#ECS Creation#
-> created a container definition using the pushed docker image
-> created a task definition specifying the details about task
-> created service by defining various details like number of tasks and load balancing details
-> Added an auto scaling group with load balancer

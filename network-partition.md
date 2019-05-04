## Network Partitioning on the Dashboard Service of Book My Event App

### SSH into Jumpbox
`ssh -i "cmpe281-nyu-aws-oregon.pem.pem" ec2-user@ec2-54-186-66-83.us-west-2.compute.amazonaws.com`

### SSH into our RIAK clusters through jumpbox
```
ssh -i "cmpe281-nyu-aws-oregon.pem" ec2-user@10.0.1.169
ssh -i "cmpe281-nyu-aws-oregon.pem" ec2-user@10.0.1.104
ssh -i "cmpe281-nyu-aws-oregon.pem" ec2-user@10.0.1.7
ssh -i "cmpe281-nyu-aws-oregon.pem" ec2-user@10.0.2.215
ssh -i "cmpe281-nyu-aws-oregon.pem" ec2-user@10.0.2.172
```

Status of RIAK Cluster
![RIAK clusters](images/riak-network-partition/1.png)

Pinging into RIAK cluster for checking health
```
curl -X GET http://10.0.1.169:8098/ping
curl -X GET http://10.0.1.104:8098/ping
curl -X GET http://10.0.1.7:8098/ping
curl -X GET http://10.0.2.215:8098/ping
curl -X GET http://10.0.2.172:8098/ping
```
![Health check](images/riak-network-partition/2.png)

Dashboard service before Partition
![Dashboard service before Partition](images/riak-network-partition/3.png)

Stopping one of the instance of the cluster
![Stopping on instance](images/riak-network-partition/4.png)

Pinging into RIAK cluster for checking health
```
curl -X GET http://10.0.1.169:8098/ping
curl -X GET http://10.0.1.104:8098/ping
curl -X GET http://10.0.1.7:8098/ping
curl -X GET http://10.0.2.215:8098/ping
curl -X GET http://10.0.2.172:8098/ping
```
One of the instance `http://10.0.2.172:8098/ping` did not respond.

![Status](images/riak-network-partition/5.png)

Dashboard status after Stopping one of the instance
![Dashboard Status](images/riak-network-partition/6.png)




## MongoDB sharding for BookEventAPI :

A MongoDB sharded cluster consists of the following components:

1] shard:  Each shard contains a subset of the sharded data. Each shard can be deployed as a replica set.

2] mongos: The mongos acts as a query router, providing an interface between client applications and the sharded cluster.

3] config servers: Config servers store metadata and configuration settings for the cluster. As of MongoDB 3.4, config servers 			   must be deployed as a replica set (CSRS).


Reference : https://docs.mongodb.com/manual/sharding/


## Below are 5 instances for MongoDB sharding:

	config-server1 : 10.0.1.192
	config-server2 : 10.0.1.224
	shard-1	       : 10.0.1.185
	shard-2        : 10.0.1.73
	mongos	       : 52.52.92.197


<img width="1440" alt="sharding" src="https://user-images.githubusercontent.com/4371600/57172710-69fc0b80-6dd8-11e9-80fe-5b9cf14b041c.png">


## Start Mongos	: 

sudo mongos --config /etc/mongod.conf --fork --logpath /var/log/mongodb/mongod.log

mongo -port 27017

## Add Shards :

<img width="1153" alt="addshard_rs0" src="https://user-images.githubusercontent.com/4371600/57172486-93676800-6dd5-11e9-80fe-434cd5d9d7ab.png">

<img width="1026" alt="addshard_rs1" src="https://user-images.githubusercontent.com/4371600/57172502-b2fe9080-6dd5-11e9-89b9-6f4e7b1e5349.png">


## List Shards :

<img width="852" alt="listshard" src="https://user-images.githubusercontent.com/4371600/57172479-7df23e00-6dd5-11e9-966e-9a42eadceceb.png">

## Use Database

## Add Shard Key
1. Enable Sharding on database
2. Add shard key

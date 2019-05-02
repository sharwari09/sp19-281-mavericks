#!/bin/sh

### Run this file on CentOS 7 EC2 instance in "Public" gateway and expects env variables UPSTREAM_URL and REQUEST_PATH ###

if [ -z "$UPSTREAM_URL" ]; then 
	echo "Pls export UPSTREAM URL as 'export UPSTREAM_URL=http://google.com:5000'"
	exit 1
fi

if [ -z "$REQUEST_PATH" ]; then 
	echo "Pls export REQUEST_PATH as 'export REQUEST_PATH=\"goapi\"" 
	exit 1
fi

if [ -z "$USERNAME" ]; then 
	echo "Pls export USERNAME as 'export USERNAME=\"cmpe281\"" 
	exit 1
fi

# Install Docker-CE
echo "Step 1: Installing docker CE"
sudo yum install -y yum-utils device-mapper-persistent-data  lvm2
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install -y docker-ce docker-ce-cli containerd.io


# Enable docker
echo "Step 2: Enable docker"
sudo systemctl enable docker
sudo systemctl start docker


# Create new network
echo "Step 3: Create new network"
docker network create --drive bridge gumball-net


# Setup Cassandra for Kong
echo "Step 4: Set up Cassandra"
docker run -d --name kong-database --network=gumball-net cassandra:2.2


# Setup kong-database
echo "Step 5: Setup kong-database"
docker run --rm --network=gumball-net -e "KONG_DATABASE=cassandra" -e "KONG_PG_HOST=kong-database" -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" kong:0.9.9 kong migrations up


# Start kong:
echo "Step 6: Start Kong"
docker run -d --name kong --network gumball-net -e "KONG_DATABASE=cassandra" -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    -e "KONG_PG_HOST=kong-database" \
    -p 8000:8000 \
    -p 8443:8443 \
    -p 8001:8001 \
    -p 7946:7946 \
    -p 7946:7946/udp \
    kong:0.9.9


# Create request path
echo "Step 7: Create request path"
curl -i -X POST --url http://localhost:8001/apis/ -d 'name=gumball-api2' -d "upstream_url=$UPSTREAM_URL" -d "request_path=/$REQUEST_PATH/" -d 'strip_request_path=true'


# Create consumer (user)
echo "Step 8: Create consumer (user)"
curl -i -X POST --url http://localhost:8001/consumers/ --data "username=$USERNAME"


# Create key-auth
echo "Step 9: Create auth-key"
curl -i -X POST --url http://localhost:8001/consumers/$USERNAME/key-auth --data ''


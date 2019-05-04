## Events Microservice

- This microservice handles creation of events on our app.
- The user/organiser has to add details such as name, schedule, venue related to a particular event that they want to host.
- The service will contain API implementation of creating events and storing them into the MongoDB cluster.



![Events Microservice](https://github.com/nguyensjsu/sp19-281-mavericks/blob/master/images/events-service.jpg)



## GOLANG REST API for Events service

### Running GOLANG API locally
- Open a terminal

- Set your GOPATH to the project directoy
```shell
export GOPATH="Your Project directory"
```

Note you might need to setup your environment before running the API
- Get all the packages
```shell
make go-get
```

- Build your app
```shell
make go-build
```

- Run the app from terminal
```shell
make go-run
```

- Check whether the app is running
```shell
[negroni] listening on :3000
```

5. Test APIs using postman
Try pinging the API
```shell
curl -X GET \
  http://localhost:3000/ping \
  -H 'Postman-Token: 081a833d-eb64-4e46-b7c9-8321499d03c2' \
  -H 'cache-control: no-cache'

{
    "Test": "API version 1.0 alive!"
}
```

## Running the GO API in EC2 using docker

1. Install Docker
2. Start Docker
sudo systemctl start docker
sudo systemctl is-active docker
3. Login to your docker hub account
sudo docker login
4. Create Docker file
sudo vim Dockerfile

```shell
FROM golang:latest 
EXPOSE 3000
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
RUN cd /app ; go install events
CMD ["/app/bin/events"]
```

5. Build the docker image locally
```shell
sudo docker build -t events .
sudo docker images
```

6. Push docker image to dockerhub
```shell
docker push events:latest
```

7. Create Public EC2 Instance
Configuration:
```shell
AMI: CentOS 7 (x86_64) - with Updates HVM
Instance Type: t2.micro
VPC: cmpe281
Network: Public subnet (us-west-1c)
Auto Public IP: Yes
SG Open Ports: 22, 80, 8080, 3000, 8000
Key Pair: cmpe281-us-west-1
```
8. ssh to your ec2 instance, user name is centos
9. Create docker-compose yml file (with the environment variables set up)
10. Deploy go API for order sevice
```shell
docker-compose up
```
11. Clean Up docker environment when finished
```shell
docker stop events
docker rm events
docker rmi {imageid}
```


## Running the GO API on Microsoft Azure Kubernetes Service

Please refer the following for deploying Microservices on AKS:

https://github.com/nguyensjsu/sp19-281-mavericks/tree/master/azure_kubernetes_service 

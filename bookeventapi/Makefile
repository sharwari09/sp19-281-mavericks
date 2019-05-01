all: clean

clean: 
	find . -name 'bookevent' -type f -exec rm -f {} \; 
	go clean
go-get:
	rm -rf src/github.com
	go get -v github.com/codegangsta/negroni
	go get -v github.com/streadway/amqp
	go get -v github.com/gorilla/mux
	go get -v github.com/unrolled/render
	go get -v github.com/satori/go.uuid
	go get -v gopkg.in/mgo.v2
	go get -v gopkg.in/mgo.v2/bson
	go get -v github.com/gorilla/handlers

go-format:
	go fmt bookevent

go-install:
	go install bookevent

go-build:
	go build bookevent

go-run: go-build
	./bookevent

docker-run:
	docker run --name bookevent -e MONGO_SERVER=$(mongodb) -e DATABASE=$(db) -e COLLECTION=$(collection) -p 3000:3000 -td bookevent	

docker-stop:
	docker stop $(name)

docker-rm:
	docker rm $(name)

docker-clean:
	docker stop $(name)
	docker rm $(name)

docker-tag:
	docker tag bookevent sayalipatil/bookevent:$(v)

docker-push:
	docker push sayalipatil/bookevent:$(v)

docker-run-bookevent:
	docker build -t bookevent .
	docker tag events sayalipatil/bookevent:$(v)
	docker push sayalipatil/bookevent:$(v)
	docker run --name bookevent -e MONGO_SERVER=$(mongodb) -e DATABASE=$(db) -e COLLECTION=$(collection) -p 3000:3000 -td sayalipatil/bookevent:$(v)


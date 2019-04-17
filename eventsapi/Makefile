all: clean

clean: 
	find . -name 'gumball' -type f -exec rm -f {} \; 
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

go-format:
	go fmt events

go-install:
	go install events

go-build:
	go build events

go-run: go-build
	./events

docker-run:
	docker run --network $(network) --name events -p 3000:3000 -td events	

docker-stop:
	docker stop $(name)

docker-rm:
	docker rm $(name)

docker-clean:
	docker stop $(name)
	docker rm $(name)



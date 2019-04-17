all: clean

clean: 
	find . -name 'user' -type f -exec rm -f {} \; 
	go clean

go-get:
	rm -rf src/github.com
	go get -v golang.org/x/crypto/bcrypt
	go get -v github.com/codegangsta/negroni
	go get -v github.com/gorilla/mux
	go get -v github.com/unrolled/render
	go get -v github.com/satori/go.uuid
	go get -v gopkg.in/mgo.v2
	go get -v gopkg.in/mgo.v2/bson

run:
	go run src/user/$(app).go

main:
	go run src/user/main.go

format:
	go fmt user

install:
	go install user

build:
	go build user

#!/bin/sh

export MONGO_SERVER="localhost" 
export MONGO_DATABASE="userdb" 
export MONGO_COLLECTION="users" 
go run .

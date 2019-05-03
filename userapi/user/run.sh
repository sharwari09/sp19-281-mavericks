#!/bin/sh

export MONGO_SERVER="34.212.50.122" 
#export MONGO_SERVER="localhost" 
export MONGO_DATABASE="userdb" 
export MONGO_COLLECTION="users" 
export DASHBOARD_URL="https://k3gku1lix8.execute-api.us-west-2.amazonaws.com/createUserEvent"
go run .

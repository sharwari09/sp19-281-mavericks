/*
	Users API
	Port mapping and initiate the user services
*/

package main

import (
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := newUserServer()
	server.Run(":" + port)
}

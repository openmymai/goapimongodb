package main

import (
	"fmt"
	"gonetflixapimongo/router"
	"log"
	"net/http"
)

const portNumber = ":8001"

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	log.Fatal(http.ListenAndServe(portNumber, r))
}

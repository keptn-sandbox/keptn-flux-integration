package main

import (
	"log"
	"net/http"
	"os"

	"github.com/keptn-sandbox/keptn-flux-integration/handlers"
)

func handleRequests() {
	log.Print("Starting the service...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	handleRequests()
}

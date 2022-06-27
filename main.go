package main

import (
	"net/http"
	"os"

	"github.com/keptn-sandbox/keptn-flux-integration/handlers"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func handleRequests() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log.Print("Starting the service...")

	port := os.Getenv("KEPTN_FLUX_INTEGRATION_PORT")
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

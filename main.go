package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/keptn-sandbox/keptn-flux-integration/models"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var fluxPayload models.FluxPayload
	json.Unmarshal(body, &fluxPayload)

	fmt.Println(string(body))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fluxPayload)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func main() {
	handleRequests()
}

package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/keptn-sandbox/keptn-flux-integration/models"
)

func home(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("Invoked the Home Endpoint...")
	log.Print(string(body))

	var fluxPayload models.FluxPayload
	if err := json.Unmarshal(body, &fluxPayload); err != nil {
		log.Print(err)
	}

	// payload := provider.GetCloudEvent(fluxPayload.InvolvedObject.Name)
	// if err := notifier.PostMessage("", payload); err != nil {
	// 	log.Print(err)
	// }

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fluxPayload)
}

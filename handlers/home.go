package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/keptn-sandbox/keptn-flux-integration/models"
	"github.com/keptn-sandbox/keptn-flux-integration/pkg/provider"
)

func home(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("*************")
	log.Print(string(body))

	var fluxPayload models.FluxPayload
	if err := json.Unmarshal(body, &fluxPayload); err != nil {
		log.Print(err)
	}

	payload := provider.GetCloudEvent(fluxPayload.InvolvedObject.Name)
	log.Print(payload)
	// if err := notifier.PostMessage("", payload); err != nil {
	// 	log.Print(err)
	// }

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fluxPayload)
}

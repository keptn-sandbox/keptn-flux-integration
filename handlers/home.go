package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/keptn-sandbox/keptn-flux-integration/pkg/notifier"
	"github.com/keptn-sandbox/keptn-flux-integration/pkg/provider"
)

func home(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalf("Could not read request body: %s", err)
	}

	log.Printf("Flux Payload Received: %s", string(body))

	var fluxPayload provider.FluxPayload
	if err := json.Unmarshal(body, &fluxPayload); err != nil {
		log.Errorf("Error unmarshalling flux payload: %s", err)
	}

	event := provider.GetCloudEvent(fluxPayload.InvolvedObject.Name, fluxPayload.InvolvedObject.Namespace)

	if err := notifier.PostMessage(event); err != nil {
		log.Errorf("Error sending event, payload %s error: %s", event, err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fluxPayload)
}

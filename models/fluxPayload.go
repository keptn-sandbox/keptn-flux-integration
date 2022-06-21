package models

import "time"

type FluxPayload struct {
	InvolvedObject struct {
		Kind            string `json:"kind"`
		Namespace       string `json:"namespace"`
		Name            string `json:"name"`
		UID             string `json:"uid"`
		APIVersion      string `json:"apiVersion"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"involvedObject"`
	Severity  string    `json:"severity"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Reason    string    `json:"reason"`
	Metadata  struct {
		Revision string `json:"revision"`
		Summary  string `json:"summary"`
	} `json:"metadata"`
	ReportingController string `json:"reportingController"`
	ReportingInstance   string `json:"reportingInstance"`
}

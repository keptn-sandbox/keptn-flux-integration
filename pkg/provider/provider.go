package provider

import (
	"bytes"
	"context"
	"os"
	"time"

	texttemplate "text/template"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

type KeptnEvent struct {
	Address string                      `json:"address"`
	Headers map[interface{}]interface{} `json:"headers"`
	Body    string                      `json:"body"`
}

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

func GetCloudEvent(name string, namespace string) KeptnEvent {
	// Flux Provider CRD Details
	group := os.Getenv("FLUX_PROVIDER_GROUP")
	version := os.Getenv("FLUX_PROVIDER_VERSION")
	resource := os.Getenv("FLUX_PROVIDER_RESOURCE")

	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	dynamic := dynamic.NewForConfigOrDie(config)

	resourceId := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}

	log.Printf("Using Keptn group: %s", group)
	log.Printf("Using Keptn version: %s", version)
	log.Printf("Using Keptn resource: %s", resource)
	log.Printf("Using Keptn name: %s", name)
	log.Printf("Using Keptn namespace: %s", namespace)

	provider, err := dynamic.Resource(resourceId).Namespace(namespace).
		Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		log.Errorf("Error getting Provider: %s", err)
	}

	clientset := kubernetes.NewForConfigOrDie(config)
	secret, err := clientset.CoreV1().Secrets(namespace).
		Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		log.Errorf("Error getting Secret list: %s", err)
	}

	var keptnEvent KeptnEvent
	keptnURL := os.Getenv("KEPTN_URL")
	keptnEvent.Address = os.Getenv("KEPTN_URL")

	log.Printf("Using Keptn URL: %s", keptnURL)

	var template string
	for k, v := range secret.Data {
		switch k {
		case "body":
			template = string(v)
		case "headers":
			headers := make(map[interface{}]interface{})
			err := yaml.Unmarshal(v, &headers)

			if err != nil {
				log.Errorf("Error parsing headers: %s", err)
			}

			keptnEvent.Headers = headers
		}
	}

	tmpl, err := texttemplate.New("labels").Parse(template)
	if err != nil {
		log.Errorf("Error parsing template: %s", err)
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, provider.GetLabels()); err != nil {
		log.Errorf("Error replacing template: %s", err)
	}

	keptnEvent.Body = b.String()

	return keptnEvent
}

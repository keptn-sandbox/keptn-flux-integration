package provider

import (
	"bytes"
	"context"
	"os"

	texttemplate "text/template"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
)

func GetProviderPayload(name string) string {
	namespace := os.Getenv("FLUX_NAMESPACE")
	group := os.Getenv("PROVIDER_GROUP")
	version := os.Getenv("PROVIDER_VERSION")
	resource := os.Getenv("PROVIDER_RESOURCE")

	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	dynamic := dynamic.NewForConfigOrDie(config)

	resourceId := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}

	provider, err := dynamic.Resource(resourceId).Namespace(namespace).
		Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		log.Errorf("Error getting Provider: %s", err)
	}

	template := GetTemplateFromSecret(config, ctx, name, namespace)
	var b bytes.Buffer

	tmpl, err := texttemplate.New("labels").Parse(template)
	if err != nil {
		log.Errorf("Error parsing template: %s", err)
	}

	if err := tmpl.Execute(&b, provider.GetLabels()); err != nil {
		log.Errorf("Error replacing template: %s", err)
	}

	return b.String()
}

func GetTemplateFromSecret(config *rest.Config, ctx context.Context,
	name string, namespace string) string {

	clientset := kubernetes.NewForConfigOrDie(config)
	secret, err := clientset.CoreV1().Secrets(namespace).
		Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		log.Errorf("Error getting Secret list: %s", err)
	}

	var template string
	for k, v := range secret.Data {
		switch k {
		case "body":
			template = string(v)
		}
	}

	return template
}

# Installation

### 1. Add your Keptn Cloud Event details in Secret
The secret is linked to the flux provider it contains the url to the keptn-flux-integration application, the keptn headers required for the cloud event, and the event itself. We are using Go Templating so the key value pairs will be passed and applied through labels in the flux provider. 

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: podtato-head-podtato-kustomize
  namespace: flux-system
stringData:
  address: http://keptn-flux-integration.flux-system.svc.cluster.local:80
  headers: |
     Content-Type: application/json
     x-token: <keptn-api-token>
  body: |
    {
      "data": {
        "configurationChange": {},
        "labels": {
          "servicename": "{{.keptnService}}"
        },
        "project": "{{.keptnProject}}",
        "service": "{{.keptnService}}",
        "stage": "{{.keptnStage}}"
      },
      "source": "flux",
      "specversion": "1.0",
      "type": "{{.keptnType}}",
      "shkeptnspecversion": "0.2.3"
    }
```
### 2. Add Helm Repo

```bash
helm repo add keptn-flux-integration https://keptn-sandbox.github.io/keptn-flux-integration
```

### 3. Install Helm Chart
```bash
helm install keptn-flux-integration --set keptn.url=https://keptn.ortelius.io/api/v1/event -n flux-system
```

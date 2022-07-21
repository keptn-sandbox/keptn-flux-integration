# Installation

### 1. Bootstrap flux (skip if you have already)

In order to bootstrap flux you will need to download the cli and have a GitHub Token.

```bash
export GITHUB_TOKEN=<your-gh-token-here>
```

```bash
flux bootstrap github \
--owner=bradmccoydev \
--repository=flux-demo \
--path=clusters/gitops \
--personal
```

### 1. Add your Keptn Cloud Event details in Secret
The secret is linked to the flux provider it contains the url to the keptn-flux-integration application, the keptn headers required for the cloud event, and the event itself. We are using Go Templating so the key value pairs will be passed and applied through labels in the flux provider.

Replace the Keptn API token below, you can get it here:

```bash
kubectl get secret -n keptn keptn-api-token -o jsonpath="{.data.keptn-api-token}" | base64 --decode
```

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

Update if required
```bash
helm repo update keptn-flux-integration
```

### 3. Install Helm Chart
```bash
helm install keptn-flux-integration keptn-flux-integration/keptn-flux-integration --set keptn.url=https://keptn.ortelius.io/api/v1/event -n flux-system
```

### 4. Create Flux Alert and Provider
```yaml
apiVersion: notification.toolkit.fluxcd.io/v1beta1
kind: Alert
metadata:
  name: podtato-head-podtato-kustomize
  namespace: podtato-kustomize
spec:
  summary: default
  providerRef:
    name: podtato-head-podtato-kustomize
  eventSeverity: info
  eventSources:
    - kind: HelmRelease
      name: "*"
    - kind: Kustomization
      name: "podtato-head-podtato-kustomize"
      namespace: podtato-kustomize
      matchLabels:
        app.kubernetes.io/name: podtato-head
---
apiVersion: notification.toolkit.fluxcd.io/v1beta1
kind: Provider
metadata:
  name: podtato-head-podtato-kustomize
  namespace: podtato-kustomize
  labels:
    keptnProject: gsoc
    keptnService: podtato-head
    keptnStage: qa
    keptnType: sh.keptn.event.qa.delivery.triggered
spec:
  type: generic
  address: http://keptn-flux-integration.flux-system.svc.cluster.local:80
  secretRef:
    name: podtato-head-podtato-kustomize
```

### 5. Keptn - Create Project gsoc 

```yaml
apiVersion: "spec.keptn.sh/0.2.2"
kind: "Shipyard"
metadata:
  name: "shipyard-delivery"
spec:
  stages:
    - name: "qa"
      sequences:
        - name: "delivery"
          tasks:
            - name: "je-deployment"
            - name: "je-test"
    - name: "production"
      sequences:
        - name: "delivery"
          triggeredOn:
            - event: "qa.delivery.finished"
          tasks:
            - name: "je-deployment"
```

### 6. Create Service podtato-head
Create Service in Keptn UI which is the same as the provider labels in line 101 above
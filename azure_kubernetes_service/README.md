
## Microservices on Azure Kubernetes Services

- Create Microsoft Azure account.

- Install Azure CLI on macOS:
```shell
brew update && brew install azure-cli
```

- Login into your Azure account
```shell
azure login
```
- Create Resource Group
```shell
az group create --name eventbrite --location eastus
```
This will take some time to create the resource group

- Create a service principal
```shell
az ad sp create-for-rbac --skip-assignment

{
  "appId": "e7596ae3-6864-4cb8-94fc-20164b1588a9",
  "displayName": "azure-cli-2018-06-29-19-14-37",
  "name": "http://azure-cli-2018-06-29-19-14-37",
  "password": "52c95f25-bd1e-4314-bd31-d8112b293521",
  "tenant": "72f988bf-86f1-41af-91ab-2d7cd011db48"
}
```

- Create a Kubernetes cluster
```shell
az aks create \
    --resource-group eventbrite \
    --name eventsAKS \
    --node-count 1 \
    --service-principal <appId> \
    --client-secret <password> \
    --generate-ssh-keys
```
Provide values of appId and password obtained from creating service principle

- Connect to cluster using kubectl:
```shell
az aks get-credentials --resource-group eventbrite --name eventsAKS
```

- To verify the connection to your cluster
```shell
kubectl get nodes
```

- To verify the connection to your cluster
```shell
kubectl get nodes
```

- Deploy Application or Microservice on AKS Cluster by creating a deployment and a service:
```shell
cat events-deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: events
spec:
  selector:
    matchLabels:
      app: events
  replicas: 2
  template:
    metadata:
      labels:
        app: events
    spec:
      containers:
      - name: events
        image: sphadnis09/events:v10
        ports:
        - containerPort: 3000
        env:
        - name: SERVER
          value: "3.210.40.136"
        - name: DATABASE
          value: "eventbrite"
        - name: COLLECTION
          value: "events"
        - name: DASHBOARD_URL
          value: "https://k3gku1lix8.execute-api.us-west-2.amazonaws.com/createUserEvent"
---
apiVersion: v1
kind: Service
metadata:
  name: events
spec:
  type: LoadBalancer
  ports:
  - port: 3000
  selector:
    app: events

```

- Deploy the above yaml as follows:
```shell
kubectl apply -f events-deployment.yaml
```

- Open a new tab and execute the following to run the dashboard:
```shell
kubectl create clusterrolebinding kubernetes-dashboard --clusterrole=cluster-adn --serviceaccount=kube-system:kubernetes-dashboard

az aks browse --resource-group eventbrite --name eventsAKS
```
Once the dashboard starts running successfully, you will see the following:

![AKS Dashboard](https://github.com/nguyensjsu/sp19-281-mavericks/blob/master/images/aks-dashboard.png)



## References

- https://docs.microsoft.com/en-us/azure/aks/tutorial-kubernetes-deploy-application

- https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest

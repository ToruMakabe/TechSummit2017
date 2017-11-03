# Sample code for MS Tech Summit 2017 Japan

## Azure Container Instances Demo
aci/
* Azure Resource Manager template for containers deployment to ACI

```
$ az group deployment create -g yourResourceGroup --template-file ./azuredeploy.json
$ az container logs -g yourResourceGroup -n ts17ContainerGroup --container-name sid
ecar
$ az container delete -g yourResourceGroup -n ts17ContainerGroup
```

go/
* Golang sample apps 
  * server: Simple HTTP server
    * Expose port 8080
    * Display own hostname, all IPs on network interfaces, environemt variables and Azure instance metadatas
  * client: Sidecar container
    * Monitor HTTP server via localhost and log it to stdout
    * Write own hostname and all IPs on network interfaces to stdout when container start
* Dockerfile
  * Multi-stage build

## Azure Distributed Data Science Toolkit
aztk/
* Sample cluster configuration of AZTK
  * Modify parameters of sample.secret.yaml
  * Rename filename from sample.secret.yaml to secret.yaml

```
$ aztk spark cluster create --id yourClusterId --size-low-pri 4
$ aztk spark cluster get --id yourClusterId
$ aztk spark cluster ssh --id yourClusterId
$ aztk spark cluster delete --id yourClusterId
```

## Azure Container Service - AKS
aks/
* Sample K8s service manifest
  * simple vote application
    * Vote web app container with Azure LB service
    * Redis container

```
$ kubectl create -f ./azure-vote.yml
$ kubectl delete deploy azure-vote-front
$ kubectl delete deploy azure-vote-back
```

## ACS-Engine
acs-engine/default/
* Sample K8s cluster definition for ACS-Engine

```
$ acs-engine generate ./kubernetes.json
```
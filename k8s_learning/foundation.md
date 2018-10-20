# basic concept about k8s

## k8s components  

### master  

Used to manage the nodes.  
  
* kube-apiserver:  
  to reflect to the request, including creating/updating/deleting resources and so on
* etcd:  
  to backup the status for rebuilding?
* kube-schedular:  
  to watch a newly created pod and distribe it to a node.  
* kube-controller-manager:  

> * node controller: watch and respond when a node down.  
> * replicaion controller:  maintain the correct number of pods  
> * endpoint controller  
> * service account & Token controlker

* cloud-controller-manager:
  not acomplished.

### node component  

* kubelet:  
  make sure containers are running in a pod.  
* kube-proxy:  
* container runtime:  
  software that is responsible for runing containers including docker, rkt and so on.  

### Addons  

need further learning.

* DNS  
* Web UI  
* Container Resource Monitoring  
* Cluster-level Logging  

## k8s API  

Used to manage specific resource via kube-apiserver.?  

Each resource in each version should have its api, including crd.  

We can use kubectl to call api.  

Api always in path: /apis/$GROUP_NAME/$VERSION in a project, and used as **apiVersion: $GROUP_NAME/$VERSION** in yaml.  

## k8s Objects  

Contains typemeta/metadata/spec/status  

* typemeta  
  Version and kind of resource.
* metadata  
  Name, UID and labels.
* spec  
  The desired state.
* status  
  Current state. Filled automatically.  

### namespace  

Used when there is many groups or projects.  
Twe same resources with the same name cannot exist in one namespace, but can exist across multiple namesaces.  
There are also some resources are not in a namespace. (An example is needed ??)  

### label and selector  

Label: some key/value pairs used to identify one or some resources.  
Label key name is suggested to use <project_name>.<group_name>/<label_key>, like apps.tensorstack.io/version  

Selector: Used to search resources via labels. It can be used in kubectl like "kubectl get pods -l '$condition'". It is also used in a yaml (For example, we can set labels in pods and set selector in a service, then the service is connected to some specific pods).  

### annotation  

Annotation is like label, it is also some key/value pairs, but it is not used to identify resources.  

Annotation stores some informations about the resource, these informations may not be used by k8s but we can get it with k8s api or via "kubectl describe $resource"  

### field selector  

Used to filter resources with the fields in resources like metadata.name, spec.replicas and status.phase.  

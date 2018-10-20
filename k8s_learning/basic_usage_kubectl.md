# Basic usage of kubectl  

## rough usage  

```go
// Create object using image  
kubectl run example --image example_image
kubectl create deployment example --image example_image
// In this way, it will create an object alone
// For example, if we try to create a server, we need also a service resource to expose the port, so we cannot use this way  
```

```go
// Operate using configuration file
kubectl create -f example.yaml
// If we want to create two resources using one config file, we should replace name with generateName.
kubectl delete -f example_1.yaml -f example_2.yaml  
// While deleting a resource, there should be name in metadata instead of generateName.
// update
kubectl replace -f example.yaml
```

## Using imperative commands

### create

* run : create a deployment to run containers in pods.
* expose : create a service  
* autoscale : create an autoscaler to scale a controller ??
* create \<objecttype> [\<subtype>] \<isntancename>: create a resource.  

### update  

* scale : scale a controller or adjust replica count.  
* annotate : add or remove an annotation.  
* label : add or remove a label.  
* set : set a aspect of an object.  
* edit : edit the raw configuration.  
* patch : modify specific fields of a live object.  

### delete  

kubectl delete \<type>/\<name>  

### view an object  

* get : print basic information  
* describe : print detailed information  
* logs : print the stdout and stderr  

### multiple  

multiple creating and setting

```go
kubectl create service clusterip my-svc --clusterip="None" -o yaml --dry-run  \
| kubectl set selector --local -f - 'environment=qa' -o yaml  \
| kubectl create -f -
// create a service into a yaml, set a selector to it and create it in cluseter
```

```go
kubectl create service clusterip my-svc --clusterip="None" -o yaml --dry-run > /tmp/srv.yaml
kubectl create --edit -f /tmp/srv.yaml
```

## Imperative management using configuration files  

```go
kubectl create -f <filename|url>
kubectl replace -f <filename|url>
kubectl delete -f <filename|url>
kubectl get -f <filename|url>
```

```go
kubectl create -f <url> --edit
// get object from url and edit it before creating
```

Transfer from imperative command management into imperative object configuration.  

```go
// transfer imperative commands to imperative object configuration
1. kubectl get <kind>/<name> -o yaml --export > <kind>_<name>.yaml
2. remove the status field
3. kubectl replace -f <kind>_<name>.yaml
```

## Declarative management using configuration files  

Creating/updating object all uses **"kubectl apply -f \<directory/>"**  
Deleting uses **"kubectl apply -f \<directory/> --prune -l your=label"**  

> it is not suggest to use apply to delete object  

Viewing an object also uses **"kubectl get -f <filename|url> -o yaml"**  

### update with apply  

When we use apply to update, we actually use patch command.  

It will compare the config and the last-applied-configuration in annotation, and update object vie these two config.  

There may be something wrong while we use default field and update it.  

> If two default field are relational, and if we update one of them, and the other will fail to get find it.  

### change management method  

Transfer from imperative command management into imperative object configuration.  

Transfer from imperative object configuration into declarative object configuration:  

```go
// Set last-applied-configuration using --save-config
kubectl replace --save-config -f <kind>_<name>.yaml

kubectl apply ....
```  

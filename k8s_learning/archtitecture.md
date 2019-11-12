# Kubernetes Architecture  

## Nodes  

Managed by note controller.  

```shell  
# get node information
k get nodes [node_name] -o yaml
```

### Address  

ExternalIP: in spec, available from outside the cluster.  
InternalIP: in status.addresses, is routable only within the cluster.  
Hostname:   in status.addresses, stand for internalIP??  

### Condition  

Tell the condition of a node:  
&#8195;&#8195;whether the memory, disk space, PID is sufficient;  
&#8195;&#8195;whether kubelet is ready  

It is in status.conditions.  

### Capacity  

Tell the resources available on the node, including CPU, memory and the maximum number of pods.  


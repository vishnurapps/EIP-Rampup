# Vishnu implementation

make a folder inside the src folder

```shell
mkdir demo
```

now change to that directory and initialize go module

```shell
cd demo
go mod init vishnu
```

create our go application

## Dockerfile creation

- Make sure to use the golang image instead of linux images as we need to do many configurations in the later case.
- set the working directory `WORKDIR /app`
- Copy everything from the current directory to the Working Directory inside the container `COPY . .`
- Build the Go app `RUN go build -o main .` Dependencies will be downloaded at this stage
- Expose port 9091 to the outside world `EXPOSE 9091`
- Command to run the executable `CMD ["./main"]`

## Interact with container from shell

### Sample GET

```shell
curl http://localhost:9091/all
```
### Sample POST

```shell
curl --header "Content-Type: application/json" --request POST --data '{"Title":"Football","Desc":"Champions League","Content":"Liverpool won Champions League"}' http://localhost:9091/all
```

## Kubernetes deployment

### Create cluster and download kubectl

```
$ minikube start --wait=false
* minikube v1.8.1 on Ubuntu 18.04
* Using the none driver based on user configuration
* Running on localhost (CPUs=2, Memory=2460MB, Disk=145651MB) ...
* OS release is Ubuntu 18.04.4 LTS
* Preparing Kubernetes v1.17.3 on Docker 19.03.6 ...
  - kubelet.resolv-conf=/run/systemd/resolve/resolv.conf
* Launching Kubernetes ...
* Enabling addons: default-storageclass, storage-provisioner
* Configuring local host environment ...
* Done! kubectl is now configured to use "minikube"
```

### List nodes

```
$ kubectl get nodes
```

Output
```
NAME       STATUS   ROLES    AGE   VERSION
minikube   Ready    master   25s   v1.17.3
```

### Create Dockerfile

Make a directory called demo in /root/go and go to that folder
```
mkdir /root/go
cd /root/go
mkdir demo
cd demo/
```
Do go mode init to generate the go.mod file
```
go mod init vishnu
```

Create the go file and do go build to generat the go.sum file 
```
go build
```

Create a docker file with the above contents and build it to get the crud:1.0 image

```
docker build -t crud:1.0 .
```

Issue docker images command to see the created image

```
$ docker images | grep crud
crud                                            1.0                 f31712143232        42 minutes ago      855MB
```

### Launch a deplyment

```shell
$ kubectl run demo --image crud:1.0 --replicas=1
```
output

```
kubectl run --generator=deployment/apps.v1 is DEPRECATED and will be removed in a future version. Use kubectl run --generator=run-pod/v1or kubectl create instead.
deployment.apps/demo created
```

### Status of deployment

```shell
$ kubectl get deployments
```

output

```
NAME   READY   UP-TO-DATE   AVAILABLE   AGE
demo   1/1     1            1           11s
```

### Detailed description of running pod

```shell
$ kubectl describe deployment demo
```

output

```
Name:                   demo
Namespace:              default
CreationTimestamp:      Wed, 30 Sep 2020 05:22:34 +0000
Labels:                 run=demo
Annotations:            deployment.kubernetes.io/revision: 1
Selector:               run=demo
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  run=demo
  Containers:
   demo:
    Image:        crud:1.0
    Port:         <none>
    Host Port:    <none>
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   demo-66dd77d4d8 (1/1 replicas created)
Events:
  Type    Reason             Age    From                   Message
  ----    ------             ----   ----                   -------
  Normal  ScalingReplicaSet  4m57s  deployment-controller  Scaled up replica set demo-66dd77d4d8 to 1
```

### Expose a port of the deployment

```shell
$ kubectl expose deployment demo --external-ip="172.17.0.18" --port=8000 --target-port=80
```

output
```
service/demo exposed
```

### Testing the exposed port

```shell
$ curl http://172.17.0.18:8000
```
In my case the port 80 was not exposed so I got connection refused

```
curl: (7) Failed to connect to 172.17.0.18 port 8000: Connection refused
```

### Creating an exposed pod

 ```
 $ kubectl run crudexposed --image=crud:1.0 --replicas=1 --port=9091 --hostport=9091
 ```
 
 Output
 ```
kubectl run --generator=deployment/apps.v1 is DEPRECATED and will be removed in a future version. Use kubectl run --generator=run-pod/v1or kubectl create instead.
deployment.apps/crudexposed created
```

### Testing the exposed pod

```
$ curl http://172.17.0.18:9091/all
```

Output

```
[{"Title":"Cricket","Desc":"Worldcup","Content":"India won worldcup"}]
```

### List pods

```
$ kubectl get pods
```

Output
```
NAME                           READY   STATUS    RESTARTS   AGE
crudexposed-59496fbb75-k4jd8   1/1     Running   0          8m10s
demo-66dd77d4d8-hkp7w          1/1     Running   0          23m
httpexposed-68cb8c8d4-4fvd5    1/1     Running   0          11m
```

### Upscale the pod count

```
$ kubectl scale --replicas=3 deployment crudexposed
```

Output
```
deployment.apps/crudexposed scaled
```
### List pods

```
$ kubectl get pods
```
Output
```
NAME                           READY   STATUS    RESTARTS   AGE
crudexposed-59496fbb75-gpqb6   0/1     Pending   0          75s
crudexposed-59496fbb75-k4jd8   1/1     Running   0          14m
crudexposed-59496fbb75-ztxx8   0/1     Pending   0          75s
demo-66dd77d4d8-hkp7w          1/1     Running   0          29m
httpexposed-68cb8c8d4-4fvd5    1/1     Running   0          17m
```

### Interact with pod

POST new article

```
curl --header "Content-Type: application/json" --request POST --data '{"Title":"F1 Race","Desc":"F1 League","Content":"Bottas won F1 League"}' http://172.17.0.18:9091/all
```
Fetch all articles

```
$ curl http://172.17.0.18:9091/all[{"Title":"Cricket","Desc":"Worldcup","Content":"India won worldcup"},{"Title":"Football","Desc":"Champions League","Content":"Liverpoolwon Champions League"},{"Title":"F1 Race","Desc":"F1 League","Content":"Bottas won F1 League"}]
```

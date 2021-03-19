
######Install Minikube on ubuntu
 * https://phoenixnap.com/kb/install-minikube-on-ubuntu



##kubectl
Kubectl uses the Kubernetes API to interact with the cluster

#### kubectl Basic Command
|Command | Use |
|--------|-----|
|`minikube version`|minikube version command|
|`minikube start`||
|`minikube dashboard`||
|`kubectl version`||
|`kubectl cluster-info`|To view the cluster details|
|`kubectl create deployment` |We need to provide the deployment name and app image location (include the full repository url for images hosted outside Docker hub) . </br>`kubectl create deployment kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1`|
|------|-------|
|`kubectl get nodes`|shows all nodes that can be used to host our application|
|`kubectl get deployment`|to view the deployment|
|`kubectl get pods`|view the pods|
|`kubectl get pods -o wide`||
|`kubectl get`| list resources|
|`kubectl get nodes`|To view the nodes in the cluster|
|`kubectl get deployments`|To list your deployments|
|`kubectl get rs`| see the replica set created by deployment|
|`kubectl get services`|View the Service you just created|
|`kubectl get pod,svc -n kube-system`|View the Pod and Service you just created|
|`kubectl get nodes --help   `||
|`kubectl get pods $POD_NAME -o json `&#124;` jq `||
|--------|--------|
|`kubectl config view`|view the kubectl configuration|
|`kubectl delete service <service name>`|To delete a service|
|`kubectl delete deployment <deployment name`||
|`kubectl delete rs <replica name>`|delete a replica set|
|`minikube delete`|delete the minikube virtual machine|
|`minikube addons list`|List the currently supported addons|
|`minikube addons disable <addon name>`|To disable an addon|
|`minikube addons enable <addon name>`|To enable an addon|
|`minikube stop`|stop the minikube virtual machine|
|`kubectl logs`|print the logs from a container in a pod|
|`kubectl exec `|execute a command on a container in a pod|
|` kubectl exec <pod name> env `| list the environment variable inside the pod |
|` kubectl exec -ti <pod name> bash `|  start a bash session inside the Pod’s container |
|`kubectl exec <podName> -it -- /bin/sh`|exec in a pod|
|______|_______|
|`kubectl describe `| show detailed information about a resource|
|`kubectl describe deployment`||
|` kubectl describe node <node name>  `|  |
|`kubectl -n kube-system get all`|get the dashboard info in terminal|
|`curl $(minikube ip):$NODE_PORT`| to see the previously exposed IP and PORT|
|______|__________|
|`kubectl scale`||
|`kubectl scale deployments/<deployment-name> --replicas=4`||
|`kubectl scale deployments/<deployment name> --replicas=2  `||
|--------|--------|
|` kubectl rollout status deployments/<deploymenty name>` ||
|` kubectl rollout undo deployments/<deployment name> `| back to previous |
|`   `||
|`  kubectl delete all --all `||
|`kubectl delete pv --all`||
|`kubectl get all`||



```
$ minikube version
```





######Hello Minicube
 * https://kubernetes.io/docs/tutorials/hello-minikube/#objectives



#### Kind Command   [Resources](https://kind.sigs.k8s.io/docs/user/quick-start/)

|Command | Use |
|--------|-------|
|` kind create cluster --name <cluster name>  `| to create a cluster . Don't user --name flag if you want to get the default value "kind"|
|` kind  get clusters`| list all the cluster name |
|` kind get nodes  `| list all the nodes |
|` kind get pods  `| list all the pods |
|`kubectl cluster-info --context kind-<clustername>`|interact with a specific cluster, you only need to </br>specify the cluster name as a context in kubectl|
|`kind delete cluster`|If the flag --name is not specified, kind will use the default cluster </br> context name kind and delete that cluster|
|`kind load docker-image <your image name> --name <cluster name>` </br> Ex : `kind load docker-image pranganmajumder/go-basic-restapi:1.0.4 --name kind`|Docker images can be loaded into your </br> cluster nodes with: kind load docker-image my-custom-image. </br>[How to see docker images that are loaded](https://stackoverflow.com/questions/60487792/kind-cluster-how-to-see-docker-images-that-are-loaded)|



## Orchestration tools offer , inspired from [Pulak Kanti Bhowmik](https://github.com/pkbhowmick/study-material)
# Kubernetes study materials


## Orchestration tools offer
- High Availability or no downtime
- Scalability or high performance
- Disaster recovery - backup and restore

## Kubernetes Basic Interactive tutorial

Resource: https://kubernetes.io/docs/tutorials/kubernetes-basics/

### Step #01 : Create a kubernetes cluster

- Install minikube from [here](https://minikube.sigs.k8s.io/docs/start/)
- ```$ minikube version``` [to check minikube is installed properly]
- ```$ minikube start``` [to run kubernetes cluster locally. Minikube started a VM and now a kubernetes cluster is running in that VM]
- Install kubectl from [here](https://kubernetes.io/docs/tasks/tools/install-kubectl/). kubectl is kubernetes command line interface which uses the kubernetes API to interact with the cluster.
- ```$ kubectl version``` [to check kubectl is installed properly]
- ```$ kubectl cluster-info``` [to view the cluster details]
- ```$ kubectl get nodes``` [to view the nodes in the cluster]

### Step #02 : Deploy an app

- ```$ kubectl create deployment <deployment_name> --image=<image_location>``` Note: Include the full repository url for images outside Docker hub.
- ```$ kubectl create deployment go-api-server --image=pkbhowmick/go-rest-api:latest``` [to deploy go-rest-api image in kubernetes]
- ```$ kubectl get deployments``` [to see the list of all deployments]
- ```$ kubectl proxy``` Note: Pods running inside kubernetes running on a private, isolated network and by default they are visible within the same kubernetes cluster. kubectl command can create a proxy that will provide API endpoint to communicate with our application.

### Step #03 : Explore the app

- ```$ kubectl get pods``` [to see the existing pods]
- ```$ kubectl describe pods``` [to view what containers are inside that Pod and what images are used to build those containers]
- ```$ kubectl logs $POD_NAME``` [to view the logs of the pod]
- ```$ kubectl exec $POD_NAME env``` [to view the list of environment variables of the pod]
- ```$ kubectl exec -ti $POD_NAME sh``` [to start a shell session inside the pod]

### Step #04 : Expose the app publicly

- ```$ kubectl get services``` [to see the list of services] Note: Although each Pod has a unique IP address, those IPs are not exposed outside the cluster without a Service. Services allow your applications to receive traffic.
- ```$ kubectl expose deployment/go-api-server --type="NodePort" --port=8080``` Note: We have a Service called kubernetes that is created by default when minikube starts the cluster. To create a new service and expose it to external traffic we’ll use the expose command with NodePort as parameter
- ```$ kubectl describe services/go-rest-api``` [To find out what port was opened externally (by the NodePort option)]
- ```$ export NODE_PORT=$(kubectl get services/go-api-server -o go-template='{{(index .spec.ports 0).nodePort}}')``` [Create an environment variable called NODE_PORT that has the value of the Node port assigned]
- ```$ echo NODE_PORT```
- ```$ curl --user admin:demo $(minikube ip):$NODE_PORT/api/users``` Note: Now APIs are accessible publicly by nodeport and minikube ip

### Step #05 : Scale up the app

- ```$ kubectl get rs``` [To see the ReplicaSet created by the Deployment]
- ```$ kubectl scale deployments/go-api-server --replicas=4``` [To scale up we'll use the kubectl scale command, followed by the deployment type, name and desired number of instances]

## Kubernetes Architecture

### Master node

- API server
- Scheduler
- Controller manager
- ETCD

### Worker node

- kubelet
- kube-proxy
- pod
- container runtime


## Kubernetes Object Management

- Imperative commands
- Imperative object configuration
- Declarative object configuration



## Minikube

## Kubectl

## Master

## Node

## Pod

- The smallest unit to interact & configure.
## Service

## Add Label to service

## Running multiple instance (Scaling)






#Replication Controller
##Concept
 * Ensures that a specified number of pods are running at any time
 * Advantages
##Review Demo
 * Manifest file
   >Configuration file yaml or json file. Contains 4 major parts apiVersion, kind, metadata, spec
 * Deploy application with Replication Controller
    * Once created a rc file run the following command </br>
      ```kubectl create -f nginx-rc.yaml```
      
    * `kubectl get pods` will show the all pods with newly created pods or run ` kubectl get po -l app=nginx-app` . here "nginx-app" is label
    * print the complete detail of nginx-rc replication controller `kubectl describe rc nginx-rc`
    
 * Test : Stop a running node and see how the replica set behave or create more node to see the result.
 * Display and validate
 * Scale up and Scale down the application : for scaling up `kubectl scale rc nginx-rc --replicas=5` , scale down `kubectl scale rc nginx-rc --replicas=2`
 * Clean up : `kubectl delete -f nginx-rc.yaml` will delete all the replication controller of nginx-rc.yaml manifest file







#Replica Set [According to Resources](https://www.youtube.com/watch?v=1WM-LsH6tKc&list=PLMPZQTftRCS8Pp4wiiUruly5ODScvAwcQ&index=20)
##Concept
* ReplicaSet : Ensures that a specified number of pods are running at any time   
* Labels & Selectors
* Equality-based and Set-based selectors
##Review Demo
* Manifest file
  >Configuration file yaml or json file. Contains 4 major parts apiVersion, kind, metadata, spec
* Deploy application with Replication Controller
   * Once created a rc file run the following command </br>
     ```kubectl create -f nginx-rs.yaml```

   * `kubectl get pods` will show the all pods with newly created pods or run ` kubectl get po -l tier=frontend`  which will display only "tier=frontend" pods
   * print the complete detail of nginx-rc replication controller `kubectl describe rc nginx-rs`

* Test : Stop a running node and see how the replica set behave or create more node to see the result.
* Display  : to get pods wide output `kubectl get po -o wide`  
* Scale up and Scale down the application : </br>
  for scaling up `kubectl scale rc nginx-rs --replicas=5` </br>
  here nginx-rs is Replica set name which is mentioned in metadata </br>
  scale down `kubectl scale rc nginx-rs --replicas=2`
* Clean up : `kubectl delete -f nginx-rs.yaml` will delete all the replication controller of nginx-rc.yaml manifest file



#Deployment 
###Feature
* Multiple Replicas
* Upgrade
* Rollback
* Scale up or down
* Pause and Resume

###Deployment Type
* Recreate
* Rolling Update ( Ramped / Incremental)
* Canary
* Blue / Green
### Review Demo
* Manifest file
* Deploy app : 
  `kubeclt create -f nginx-deploy.yaml`
* Display and validate : </br>
  `kubectl get deploy -l app=nginx-app` </br>
  `kubectl get rs -l app = nginx-app` </br>
  `kubectl get po -l app=nginx-app` </br>
  `kubectl describe deploy nginx-deploy` here nginx-deploy is manifest file name.
* Test-Upgrade : </br>
  `kubectl set image deploy <deploymentName> <containerName>=new image` </br>
  Ex: `kubectl set image deploy nginx-deployment nginx-container=nginx:1.9.1` </br>
  roll out status `kubectl rollout status deployment/<deploymentName>` 
  Ex : `kubectl rollout status deployment/nginx-deployment` </br>
  Get all deployment available on master node `kubectl get deploy`
* Test-Rollback :  </br>
  to check status `kubectl rollout deployment/nginx-deployment` </br>
  check rollout history `kubectl rollout history deployment/nginx-deployment` </br>
  back to previous version `kubectl rollout undo deployment/nginx-deployement`
  
* Test-Scale up & Down pod instances : </br>
    scale up `kubectl scale deployment <deploymentName in metadata> --replicas=5` </br>
    Ex : `kubectl scale deployment nginx-deployment --replicas=5` </br>
    to check `kubectl get po -l app=nginx-app` it will show the running replica labeled app=nginx-app </br>
    
    scale down `kubectl scale deployment <deploymentName in metadata> --replicas=3` </br>

* Clean : </br>
  delete all the instances of manifest file `kubectl delete -f nginx-deploy.yaml` </br>
  check again `kubectl get po -l app=nginx-app`
  


#Service:
###Concept:
* A service is a grouping of pods that are running on the cluster
* Type :
   * ClusterIp: reachable only within the cluster. So  Scope is confined only within the cluster. 
   * NodePort : Expose app outside the internet world .
   * LoadBalancer: To access all the nodes use LoadBalancer .
    
###Resources:
* [Kubernetes Services Explained](https://www.youtube.com/watch?v=T4Z7visMM4E)
* [Cluster Ip Services](https://www.youtube.com/watch?v=dVDElh_Kd48&list=PLMPZQTftRCS8Pp4wiiUruly5ODScvAwcQ&index=25)
* [NodePort Services](https://www.youtube.com/watch?v=eth7osiCryc&list=PLMPZQTftRCS8Pp4wiiUruly5ODScvAwcQ&index=23)
* [Load Balancing Services](https://www.youtube.com/watch?v=xCsz9IOt-fs&list=PLMPZQTftRCS8Pp4wiiUruly5ODScvAwcQ&index=24)



    
###NodePort:
#### Concept:
    * Why we need : 
    * nodePort:
    * Port Types: 
         * NodePort [30000 to 32767 ] : to expose on the internet . It's a port on the node where pod is running.
         * ServicePort / Port : this is a port on the service itself. 
         * TargetPort: Typically ServicePort & TargetPort are same [80]. This is a port on actual pod where the app is running.

    * Scenarious: 
         * You can only have once service per port
         * You can only use ports 30000 - 32767
         * if your Node/Vm IP addrs change, you need to deal with that

 
#### Review Demo: 
* Manifest file: write manifest file for nodeport service </br> 
* Create objects:</br>
    * at first deployment `kubectl create -f nginx-deploy.yaml`
    * then expose it create nodeport service `kubectl create -f nginx-svc-np.yaml`
        * after editing manifest file you need to run the following command:
          `kubectl apply -f nginx-deploy.yaml` && `kubectl apply -f nginx-svc-np.yaml`
    
* Display :
    * get services `kubectl get service -l app=nginx-app` or `kubectl get svc -l app=nginx-app` or `kubectl describe services/my-service` or </br>
    `kubectl describe svc my-service` . Here my-service stands for service name 
  
* Validate: 
* Test use cases :
* Clean up: </br>
Pod will be automatically deleted if you delete the respective service. </br>
  delete service `kubectl delete svc my-service` </br>
  again check if there is any running pods `kubectl get pods`



#LoadBalancer : [Resources](https://www.youtube.com/watch?v=xCsz9IOt-fs&list=PLMPZQTftRCS8Pp4wiiUruly5ODScvAwcQ&index=24)

  
#DaemonSet :  [Resources](https://www.youtube.com/watch?v=yYeUic8B6fM)
###Concept:
* Controller that comes close to DaemonSet
* DaemonSet Overview
* In-depth

###Review Demo:
* Manifest file
* Deploy application with DaemonSet : </br>
    * deploy the daemonSet  `kubectl create -f fluentd-daemonset.yaml` & check which node it is running on `kubectl get po -o wide`
    * get daemonSet `kc get ds` 
* Display and validate DaemonSet : </br>
    * run `kubectl get node` to get all the nodes available. 
    * get detail `kc describe ds fluentd-ds` . here 'fluent-ds' is the name of the object mentioned in metadata. 
    
* Clean Up : </br>
    * please make sure delete the deployment not just pod `kubectl delete ds fluentd-ds` . if you delete the 
    pod before daemonSet, the pod will be recreated.



#Job [Resources](https://www.youtube.com/watch?v=j1EnBbxSz64)
###Type:
    * Jobs (Run to completion) : Bash processing
        * Job is a Controller in k8s, which supervises Pods for carrying out certain tasks.
        * Each Job creates one/more Pods
        * Ensures they are successfully terminated
        * when job complet, no more pods are created. Keeping them around to view the logs for checking issue.
        * Job object also remain after it is completed. 
        * when one delete job , all the pod associated with it get deleted.
        * Job controller restarts/ rescheduled if a pod/node fails during execution
        * Use Cases:
            * one time initialization of resources such as database
            * Multiple workers to process message in queue

    * CronJob (Scheduled) : 

### Review Demo:
* Manifest file: write manifest file for Job </br>
* Create objects:</br>
    * at first deployment `kubectl create -f countdown-jobs.yaml`
    
* Display : </br>
    * `kubectl get jobs` </br>
    * `kubectl get po -o wide` </br>
    * to see pod logs `kubectl logs <podName>`
    * `kubectl describe jobs <jobName>`  . Here job name is countdown
    * `pods=$(kubectl get pods --selector=job-name=pi --output=jsonpath='{.items[*].metadata.name}')
       echo $pods`

* Validate:
* Test use cases :
* Clean up: </br>
  Pod will be automatically deleted if you delete the respective Job. </br>
  delete Job `kubectl delete jobs <jobName>` . Here job name is countdown </br> 
  again check if there is any running pods created from job `kubectl get pods`



#CronJobs: [Resources](https://www.youtube.com/watch?v=OZAhYSDkhsI)



### Review Demo:
* Manifest file: write manifest file for Job </br>
* Create objects:</br>
    * at first deployment `kubectl create -f fileName.yaml`

* Display : </br>
    * `kubectl get cronjobs` </br>
    * `kubectl get po -o wide` </br>
    * to see pod logs `kubectl logs <podName>`
    * `kubectl describe cronjobs <jobName>`  . Here job name is countdown
   

* Validate:
* Test use cases :
* Clean up: </br>
  Pod will be automatically deleted if you delete the respective Job. </br>
  delete Job `kubectl delete cronjobs <cronjobName>` . Here job name is countdown </br>
  again check if there is any running pods created from job `kubectl get pods`

###Cron schedule syntex:
```
# ┌───────────── minute (0 - 59)
# │ ┌───────────── hour (0 - 23)
# │ │ ┌───────────── day of the month (1 - 31)
# │ │ │ ┌───────────── month (1 - 12)
# │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
# │ │ │ │ │                                   7 is also Sunday on some systems)
# │ │ │ │ │
# │ │ │ │ │
# * * * * *
```


#Storage:
## Volume:
### Type:
#### emptyDir: [Resources](https://www.youtube.com/watch?v=6K6gSwJcuV4&list=PLMPZQTftRCS8Pp4wiiUruly5ODScvAwcQ&index=27)

#### hostPath: 
  * Resources
    * [Srinath Challa](https://www.youtube.com/watch?v=gVuYmOlnu3A&list=PLMPZQTftRCS8Pp4wiiUruly5ODScvAwcQ&index=28)
  
##### Concept:
* mounts a file of directory from the host node's filesystem into your Pod
* Remains ever after the pod is terminated
* Similar to docker volume
* Data inside the hostPath volume remain even after the pod is terminated

##### Review Demo:
* Manifest file: hostpath.yaml is your manifest file
* Create object : `kc create -f hostpath.yaml` it will create a pod named redis-hostpath
* Display: Same as before
* Validate: go to node `docker exec nodeName` & create a file . The same file will be created inside the pod's test-mnt directory


## Static Volume Provisioning:
#### Step:
* Persistent Volume (PV) : 
    


## Dynamic Provisioning :
#### Lifecycle:
* Admin setups up storage provisioner
* Admin creates a StorageClass
* User creates a PVC
* K8s automatically creates a PV
* Pod uses the PVC as a volume

## Warning:
* Delete pv before deleting PVC , In case of deleting PVC before PV , follow the given [link](https://stackoverflow.com/questions/51358856/kubernetes-cant-delete-persistentvolumeclaim-pvc)
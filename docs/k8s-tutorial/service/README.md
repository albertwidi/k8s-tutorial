# Service

**From the docs:** A Kubernetes Service is an abstraction which defines a logical set of Pods and a policy by which to access them - sometimes called a micro-service. The set of Pods targeted by a Service is (usually) determined by a Label Selector (see below for why you might want a Service without a selector).

To check more about service, please go [here](https://kubernetes.io/docs/concepts/services-networking/service/)

## TLDR

Service is an abstraction to communicate with pods

## Service Configuration

```yaml
apiVersion: v1
kind: Service
metadata:
  name: tutorial-backend
spec:
  selector:
    app: tutorial-backend
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9000
```

## Cases 

### Creating Service

For creating service, you can use this command: `kubectl apply -f service.yaml`

The command above should give this result:

```shell
$ kubectl apply -f service.yaml 
service/tutorial-backend created
```

Check if the service is created by using command: `kubectl get svc -n tutorial`

```shell
$ kubectl get svc -n tutorial
NAME               TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
tutorial-backend   ClusterIP   10.105.123.198   <none>        80/TCP    52s
```

### Check Service Configuration and Endpoints

To check service configuration you can use this command: `kubectl get svc tutorial-backend -n tutorial -o yaml`

Result of above command:

```yaml
$ kubectl get svc tutorial-backend -n tutorial -o yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"name":"tutorial-backend","namespace":"tutorial"},"spec":{"ports":[{"port":80,"protocol":"TCP","targetPort":9000}],"selector":{"app":"tutorial-backend"}}}
  creationTimestamp: 2019-05-08T09:58:03Z
  name: tutorial-backend
  namespace: tutorial
  resourceVersion: "167074"
  selfLink: /api/v1/namespaces/tutorial/services/tutorial-backend
  uid: c527f6ca-7177-11e9-b634-8092b38543a0
spec:
  clusterIP: 10.105.123.198
  ports:
  - port: 80
    protocol: TCP
    targetPort: 9000
  selector:
    app: tutorial-backend
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
```

To check if the service has a valid selector and endpoints: `kubectl describe svc tutorial-backend -n tutorial`

```yaml
$ kubectl describe svc tutorial-backend -n tutorial
Name:              tutorial-backend
Namespace:         tutorial
Labels:            <none>
Annotations:       kubectl.kubernetes.io/last-applied-configuration={"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"name":"tutorial-backend","namespace":"tutorial"},"spec":{"ports":[{"port":80,"protoco...
Selector:          app=tutorial-backend
Type:              ClusterIP
IP:                10.105.123.198
Port:              <unset>  80/TCP
TargetPort:        9000/TCP
Endpoints:         172.17.0.7:9000
Session Affinity:  None
Events:            <none>
```

As we can see from above output, the service has 1 endpoint. The service is valid to deliver traffic to our backend

### Port Forward Traffic From Local To Service

To port-forward service, run this command: `kubectl port-forward -n tutorial svc/tutorial-backend 9000:80`

Above command means, we will port forward a service named `tutorial-backend` in `tutorial` namespace and forward `port 80` of the service to `localhost 9000`

The command above will show this output: 

```shell
$ kubectl port-forward -n tutorial svc/tutorial-backend 9000:80
Forwarding from 127.0.0.1:9000 -> 9000
Forwarding from [::1]:9000 -> 9000
```

Then try to hit the service using: `curl localhost:9000/v1/content`

The command aboce should give this output:

```shell
$ curl localhost:9000/v1/content
{"dynamic":"","static":"this is static content"}
```

### Invalid Selector

It is possible to have an invalid selector, which will give 0 endpoints in result.

For example, apply this command: `kubectl apply -f service_invalid.yaml`

Result:

```shell
$ kubectl apply -f service_invalid.yaml 
service/tutorial-backend configured
```

Then check/describe your service by using this command: `kubectl describe service -n tutorial tutorial-backendv`

Result:

```yaml
$ kubectl describe service -n tutorial tutorial-backend 
Name:              tutorial-backend
Namespace:         tutorial
Labels:            <none>
Annotations:       kubectl.kubernetes.io/last-applied-configuration={"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"name":"tutorial-backend","namespace":"tutorial"},"spec":{"ports":[{"port":80,"protoco...
Selector:          app=tutorial-backend-v2
Type:              ClusterIP
IP:                10.105.123.198
Port:              <unset>  80/TCP
TargetPort:        9000/TCP
Endpoints:         <none>
Session Affinity:  None
Events:            <none>
```

In the description of `service` we can see `Endpoints: <none>` which means it doesn't have any endpoints or backend

### Load Balancer Type Service

Service with load-balancer type means it will open a public IP and your service is exposed to the `internet` or `internal vpc`.

To create a service with `LoadBalancer` type, run this command: `kubcetl apply -f service_lb.yaml`

Command Result:

```shell
$ kubectl apply -f service_lb.yaml
service/tutorial-backend configured
```

Try to get list of service with this command: `kubectl get svc -n tutorial`

The command should give you this result:

```shell
$ kubectl get svc -n tutorial
NAME               TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
tutorial-backend   LoadBalancer   10.105.123.198   <pending>     80:32382/TCP   147m
```

The `<pending>` state of `CLUSTER-IP` is expected if you use `minikube`, but if you are using cloud provider, you should see the IP number.

To enable the external IP on `minikube` you can run `minikube tunnel`. Then try to get your service list again using `kubectl get svc -n tutorial`.

You should see the `CLUSTER-IP` now, for example:

```shell
$ kubectl get svc -n tutorial
NAME               TYPE           CLUSTER-IP       EXTERNAL-IP      PORT(S)        AGE
tutorial-backend   LoadBalancer   10.105.123.198   10.105.123.198   80:32382/TCP   150m
```

Try to curl the IP with this command: `curl 10.105.123.198/v1/content`. It should give you the correct result:

```shell
$ curl 10.105.123.198/v1/content
{"dynamic":"","static":"this is static content"}
```

### Service With Multiple Pods

What if we have multiple pods? How it will behave?

Let's create a new pod first by running this command: `kubectl apply -f pod_new.yaml`

Result of the command:

```shell
$ kubectl apply -f pod_new.yaml
pod/tutorial-backend-new created
```

Now check the pod by running this command: `kubectl get po -n tutorial`

```shell
$ kubectl get po -n tutorial
NAME                   READY     STATUS    RESTARTS   AGE
tutorial-backend       1/1       Running   0          3h59m
tutorial-backend-new   1/1       Running   0          54s
```

We can see that both old and new pod are running.

Let's describe the service once again by running: `kubectl describe svc -n tutorial tutorial-backend`

```shell
$ kubectl describe svc -n tutorial tutorial-backend
Name:                     tutorial-backend
Namespace:                tutorial
Labels:                   <none>
Annotations:              kubectl.kubernetes.io/last-applied-configuration={"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"name":"tutorial-backend","namespace":"tutorial"},"spec":{"ports":[{"port":80,"protoco...
Selector:                 app=tutorial-backend
Type:                     LoadBalancer
IP:                       10.105.123.198
LoadBalancer Ingress:     10.105.123.198
Port:                     <unset>  80/TCP
TargetPort:               9000/TCP
NodePort:                 <unset>  32382/TCP
Endpoints:                172.17.0.7:9000,172.17.0.8:9000
Session Affinity:         None
External Traffic Policy:  Cluster
Events:
  Type    Reason  Age   From                Message
  ----    ------  ----  ----                -------
  Normal  Type    3h    service-controller  ClusterIP -> LoadBalancer
```

You should notice that the endpoints is now hsa 2 IP and type is `LoadBalancer` due our last changes to the service

Try to hit the service again using: `curl 10.105.123.198/v1/content`. You will receive two different response, because we are deploying two pods:

```shell
$ curl 10.105.123.198/v1/content
{"dynamic":"haloha","static":"this is static content"}
```

and 

```shell
$ curl 10.105.123.198/v1/content
{"dynamic":"","static":"this is static content"}
```
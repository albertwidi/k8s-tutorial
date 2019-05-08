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

### Invalid Selector

### Looad Balancer Type Service
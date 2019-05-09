# Replicaset

**From the docs:** A ReplicaSet’s purpose is to maintain a stable set of replica Pods running at any given time. As such, it is often used to guarantee the availability of a specified number of identical Pods.

To check more about replicaset, please go [here](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/)

## TLDR

Replicaset help you to control the number of pod replica

## Replicaset Config

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: tutorial-backend 
  namespace: tutorial
  labels:
    app: tutorial-backend
spec:
  # modify replicas according to your case
  replicas: 3
  selector:
    matchLabels:
      app: tutorial-backend 
```

## Cases

### Pre-requisite

Before begin the tutorial, let's create our pod first with this command: `kubectl apply -f pod.yaml -f pod_new.yaml`

```shell
$ kubectl apply -f pod.yaml -f pod_new.yaml
pod/tutorial-backend created
pod/tutorial-backend-new created
```

You should see two pods if you run this command: `kubectl get po -n tutorial`

```shell
$ kubectl get po -n tutorial
NAME                   READY     STATUS    RESTARTS   AGE
tutorial-backend       1/1       Running   0          30s
tutorial-backend-new   1/1       Running   0          30s
```

### Creating Replicaset

To create a replicaset, run this command: `kubectl apply - replicaset.yaml`

```shell
$ kubectl apply -f replicaset.yaml 
replicaset.apps/tutorial-backend created
```

Then try to get the pod list using: `kubectl get po -n tutorial`

You should see 3 pods running from the list, for example:

```shell
$ k get po -n tutorial
NAME                     READY     STATUS    RESTARTS   AGE
tutorial-backend         1/1       Running   0          95m
tutorial-backend-flp4n   1/1       Running   0          33s
tutorial-backend-new     1/1       Running   0          95m
```

### Deleting Pod

Now let's try to delete a pod and see what happen with the replicaset. 

But before that, let's check our `replicaset` configuration by using: `kubectl get replicaset -n tutorial tutorial-backend -o yaml`

```yaml
$ k get replicaset -n tutorial tutorial-backend -o yaml
apiVersion: extensions/v1beta1
kind: ReplicaSet
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"apps/v1","kind":"ReplicaSet","metadata":{"annotations":{},"labels":{"app":"tutorial-backend"},"name":"tutorial-backend","namespace":"tutorial"},"spec":{"replicas":3,"selector":{"matchLabels":{"app":"tutorial-backend"}},"template":{"metadata":{"labels":{"app":"tutorial-backend"}},"spec":{"containers":[{"image":"albertwidi/k8s-tutorial-backend:latest","imagePullPolicy":"Never","livenessProbe":{"httpGet":{"path":"/healtcheck?probe=liveness","port":9000},"initialDelaySeconds":15,"timeoutSeconds":1},"name":"tutorial-backend","ports":[{"containerPort":9000,"name":"http","protocol":"TCP"}]}]}}}}
  creationTimestamp: 2019-05-08T18:12:17Z
  generation: 1
  labels:
    app: tutorial-backend
  name: tutorial-backend
  namespace: tutorial
  resourceVersion: "203619"
  selfLink: /apis/extensions/v1beta1/namespaces/tutorial/replicasets/tutorial-backend
  uid: d0817912-71bc-11e9-b634-8092b38543a0
spec:
  replicas: 3
  selector:
    matchLabels:
      app: tutorial-backend
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: tutorial-backend
    spec:
      containers:
      - image: albertwidi/k8s-tutorial-backend:latest
        imagePullPolicy: Never
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /healtcheck?probe=liveness
            port: 9000
            scheme: HTTP
          initialDelaySeconds: 15
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        name: tutorial-backend
        ports:
        - containerPort: 9000
          name: http
          protocol: TCP
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
status:
  availableReplicas: 3
  fullyLabeledReplicas: 3
  observedGeneration: 1
  readyReplicas: 3
  replicas: 3
```

As we can see from the ocnfiguration, we are expecting 3 replica to run at one time and using `selector` with `matchLabels` of `app: tutorial-backend`. Which mean it will make sure all `pods` that have label `app: tutorial-backend` to have at least 3 replica.

Use this command to delete a pod: `kubectl delete po -n tutorial tutorial-backend`

```shell
$ kubectl delete po -n tutorial tutorial-backend
pod "tutorial-backend" deleted
```

Then check the pod list again: `kubectl get po -n tutorial`

```
$ kubectl get po -n tutorial
NAME                     READY     STATUS    RESTARTS   AGE
tutorial-backend-flp4n   1/1       Running   0          5m52s
tutorial-backend-new     1/1       Running   0          100m
tutorial-backend-xxjpd   1/1       Running   0          31s
```

You should see a brand new pod created, because `replicaset` making sure it has at least 3 replica.
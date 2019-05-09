# Pod

**From the docs:** Pod are the smallest deployable unit of computing that can be created and managed by Kubernetes

To check more about pod, please go [here](https://kubernetes.io/docs/concepts/workloads/pods/pod/)

## TLDR

Pod is where your container is being put and running inside kubernetes. In a single pod, can consist one or more container.

## Pod Configuration

Example of pod configuration

```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: liveness-http
  namespace: tutorial
spec:
  containers:
  - args:
    - /server
    image: k8s.gcr.io/liveness
    livenessProbe:
      httpGet:
        # when "host" is not defined, "PodIP" will be used
        # host: my-host
        # when "scheme" is not defined, "HTTP" scheme will be used. Only "HTTP" and "HTTPS" are allowed
        # scheme: HTTPS
        path: /healthz
        port: 8080
        httpHeaders:
        - name: X-Custom-Header
          value: Awesome
      initialDelaySeconds: 15
      timeoutSeconds: 1
    name: liveness
```

## Pod Lifecycle

In this section we will discussing about the pod lifecycle, liveness and readiness.

### Liveness

Liveness probe needed because application might not start because of some problem

### Readiness

Readiness probe needed because application might up but not ready yet to receive trraffic.

## Cases

In this section we will discuss several cases

### Applying Pod To Kubernetes

To apply `pod` into our kubernetes cluster, you can use `kubectl` command.

Lets try to apply `pod.yaml` that exists within the pod tutorial directory

Type `kubectl apply -f pod.yaml` and press enter. The command should give you this result:

```shell
$ k apply -f pod.yaml         
pod/tutorial-backend created
```

To check the `pod`, we can use this command: `kubectl get po -n tutorial`

```shell
$ k get po -n tutorial
NAME               READY     STATUS    RESTARTS   AGE
tutorial-backend   1/1       Running   0          5s
```

To check the `pod` configuration, we can use this command: `kubectl get po tutorial-backend -n tutorial -o yaml`

```yaml
$ kubectl get po tutorial-backend -n tutorial -o yaml
apiVersion: v1
kind: Pod
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Pod","metadata":{"annotations":{},"labels":{"app":"tutorial-backend"},"name":"tutorial-backend","namespace":"tutorial"},"spec":{"containers":[{"image":"albertwidi/k8s-tutorial-backend:latest","imagePullPolicy":"Never","livenessProbe":{"httpGet":{"path":"/healtcheck?probe=liveness","port":9000},"initialDelaySeconds":15,"timeoutSeconds":1},"name":"tutorial-backend","ports":[{"containerPort":9000,"name":"http","protocol":"TCP"}]}]}}
  creationTimestamp: 2019-05-08T08:43:02Z
  labels:
    app: tutorial-backend
  name: tutorial-backend
  namespace: tutorial
  resourceVersion: "161618"
  selfLink: /api/v1/namespaces/tutorial/pods/tutorial-backend
  uid: 4a74ed2e-716d-11e9-b634-8092b38543a0
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
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-vrsxm
      readOnly: true
  dnsPolicy: ClusterFirst
  enableServiceLinks: true
  nodeName: minikube
  priority: 0
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-vrsxm
    secret:
      defaultMode: 420
      secretName: default-token-vrsxm
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: 2019-05-08T08:43:02Z
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: 2019-05-08T08:43:04Z
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: 2019-05-08T08:43:04Z
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: 2019-05-08T08:43:02Z
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://fca3d341be1f4b081c655ad647fb4ed94930e21ee440e27fbcce8b7b92f3063b
    image: albertwidi/k8s-tutorial-backend:latest
    imageID: docker://sha256:fceff374104de97adbbbd78fabef0cc803db02e97859da222b7d94fb02c7b40b
    lastState: {}
    name: tutorial-backend
    ready: true
    restartCount: 0
    state:
      running:
        startedAt: 2019-05-08T08:43:03Z
  hostIP: 192.168.122.79
  phase: Running
  podIP: 172.17.0.7
  qosClass: BestEffort
  startTime: 2019-05-08T08:43:02Z
```

### Port Forward Traffic From Local to Pod

To port-forward traffic from your local computer to pod, please use this command: `$ k port-forward -n tutorial tutorial-backend 9000:9000`

The command above should give this output:

```shell
$ k port-forward -n tutorial tutorial-backend 9000:9000
Forwarding from 127.0.0.1:9000 -> 9000
Forwarding from [::1]:9000 -> 9000
```

Try giving a call to the service by hitting `$ curl localhost:9000/v1/content`

Result of curl should be:

```shell
$ curl localhost:9000/v1/content
{"dynamic":"","static":"this is static content"}
```

### Pod CrashLoopBackOff

### Destroying Pod, Why Pod Isn't Coming Back?

Let's try to delete a pod, use this command to delete the pod: `kubectl delete -f pod.yaml`

```shell
$ k delete -f pod.yaml 
pod "tutorial-backend" deleted
```

And then try to get all the pod in `tutorial` namespace: `$ k get po -n tutorial`

```shell
$ k get po -n tutorial
No resources found.
```

As you see in above result, there are no resources for pods found in the namespace because pods is not managed by anybody. Pod is just a unit that need to be managed by another resource.
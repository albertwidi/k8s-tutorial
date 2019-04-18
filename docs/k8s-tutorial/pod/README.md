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

### Liveness

### Readiness

## Cases

In this section we will discuss several cases

### Applying Pod To Kubernetes

To apply `pod` into our kubernetes cluster, you can use `kubectl` command.

Lets try to apply `pod.yaml` that exists within the pod tutorial directory

Type `kubectl apply -f pod.yaml` and press enter. The command should give you this result:

```yaml
```

To check the `pod`, we can use this command: `kubectl get po -n tutorial`

```yaml
```

To check the `pod` configuration, we can use this command: `kubectl get po liveness-http -n tutorial`

```yaml
```

### Giving Traffic To Pod

### Giving Traffic To Pod, While Pod Not Ready

### Pod CrashLoopBackOff

### Destroying Pod, Why Pod Isn't Coming Back?
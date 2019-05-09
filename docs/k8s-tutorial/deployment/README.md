# Deployment

**From the doc:** A Deployment controller provides declarative updates for Pods and ReplicaSets.

You describe a desired state in a Deployment object, and the Deployment controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.

To check  more about deployment, please go [here](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

## TLDR

Deployment will control your `Pod` and `ReplicaSet`

## Deployment Configuration

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tutorial-backend-deploy
  labels:
    app: tutorial-backend-deploy
spec:
  replicas: 3
  selector:
    matchLabels:
      app: tutorial-backend
  template:
    metadata:
      labels:
        app: tutorial-backend
    spec:
      containers:
      - name: tutorial-backend
        image: albertwidi/k8s-tutorial-backend:latest
        imagePullPolicy: Never
        ports:
        - containerPort: 9000
          name: http
          protocol: TCP
        livenessProbe:
          httpGet:
            # when "host" is not defined, "PodIP" will be used
            # host: my-host
            # when "scheme" is not defined, "HTTP" scheme will be used. Only "HTTP" and "HTTPS" are allowed
            # scheme: HTTPS
            path: /healtcheck?probe=liveness
            port: 9000 
          initialDelaySeconds: 15
          timeoutSeconds: 1
```

## Cases


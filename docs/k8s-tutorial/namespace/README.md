# Namespace

**From the docs:** Kubernetes supports multiple virtual clusters backed by the same physical cluster. These virtual clusters are called namespaces.

To learn more about namespace, pleasae go [here](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)

## Namespace Configuration

This is an example of namespace configuration

```yaml
apiVersion: v1
kind: Namespace
metadata:
  labels:
    name: tutorial
  name: tutorial
```

## Cases

### Get List Of All Namespace

To get list of all namespace, type this command: `kubectl get namespaces`. The command should give you this kind of output:

```yaml
```

### Creating Namespace

To create a new namespace, lets use `namespace.yaml` file that exists within the namespace tutorial directory and use this command:

`kubectl apply -f namespace.yaml`

The command should give you this output:

```
```

Lets check the namespace list once again using `kubectl get namespaces`, and you will see a new `tutorial` namespace:

```yaml
```
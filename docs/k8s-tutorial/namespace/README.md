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

```sh
$ k get namespaces
NAME              STATUS    AGE
default           Active    76d
kube-node-lease   Active    2d14h
kube-public       Active    76d
kube-system       Active    76d
```

### Creating Namespace

To create a new namespace, lets use `namespace.yaml` file that exists within the namespace tutorial directory and use this command:

`kubectl apply -f namespace.yaml`

The command should give you this output:

```
$ k apply -f docs/k8s-tutorial/namespace/namespace.yaml 
namespace/tutorial created
```

Lets check the namespace list once again using `kubectl get namespaces`, and you will see a new `tutorial` namespace:

```sh
$ k get namespaces
NAME              STATUS    AGE
default           Active    76d
kube-node-lease   Active    2d14h
kube-public       Active    76d
kube-system       Active    76d
tutorial          Active    25s
```

### Deleting Namespace

To delete the namespace we can simply running this command:

`kubectl delete -f namespace.yaml`

Or if you want to delete a spesific namespace, use this command:

`kubectl delet namespace tutorial`

The command above will give this output:

```sh
$ k delete -f docs/k8s-tutorial/namespace/namespace.yaml
namespace "tutorial" deleted
```
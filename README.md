# Kubernetes Tutorial

This kubernetes tutorial is using `minikube` and will cover this following agenda:

1. Install helm into `minikube`

2. Install jenkins via `stable/helm` charts

3. Build web application using `go`

4. Build an application using Jenkins

5. Deploy an application to kubernetes cluster using helm

## Helm

For helm to be fully functional, you need to `install helm` on your local and then `install tiller` to your k8s-cluster

### Installing Helm

Go to this [page](https://docs.helm.sh/using_helm/#installing-helm) or follow the steps bellow:

**Binary Release**

1. Go to [here](https://github.com/helm/helm/releases) and download desired version
2. Unpack `tar -zxvf helm-v2.0.0-linux-amd64.tgz`
3. Move it `mv linux-amd64/helm /usr/local/bin/helm`

**Linux**

Snap `sudo snap install helm --classic`

**MacOS**

Use homebrew `brew install kubernetes-helm`

### Installing Tiller

Tiller is server for helm that running inside your kubernetes-cluster

Go to [this](https://docs.helm.sh/using_helm/#installing-tiller) page or follow the step bellow:

`helm init`

then check your `kube-system` namespace, it should show you a `tiller` pod

```bash
tiller-deploy-69ffbf64bc-p2mbl         1/1       Running   1          10s
```

## Jenkins

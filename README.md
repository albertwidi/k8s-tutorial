# Kubernetes Tutorial

## Disclaimer

This tutorial is for **beginner** in kubernetes, not for managing kubernetes cluster and its maintenance. This all about introduction to container, CI/CD and deploying application to kubernetes.

## Agenda

This kubernetes tutorial is using `minikube` and will cover this following agenda:

1. Install helm into `minikube`

2. Install jenkins via `stable/helm` charts

3. Build web application using `go`

4. Build an application using Jenkins

5. Deploy an application to kubernetes cluster using helm

## Pre-Requisite

To be able to run this tutorial, you need to install several things:

### System Requirement

This tutorial is build on-top `ubuntu 18.04` and is using around 4-12Gb depends on what is running on minikube.

The safe point is around 8Gb

### Install Docker

**Ubuntu**

Please look into [this](https://docs.docker.com/install/linux/docker-ce/ubuntu/) page and follow the instruction to install docker CE linux for ubuntu.

Or

Simply use this script, use it at your own risk:

```bash
$ curl -fsSL https://get.docker.com -o get-docker.sh
$ sudo sh get-docker.sh
```

After docker is installed, run `  sudo usermod -aG docker your-user` to run `docker` command as non-root user.

**MacOS**

Please [visit](https://hub.docker.com/editions/community/docker-ce-desktop-mac) this link to install docker CE Deskop-Mac for macOS.

### Install Kubectl

**Debian & Ubuntu**

Run this script:

```bash
sudo apt-get update && sudo apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee -a /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubectl
```

**Binary**

1. Run `curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl`. To download specific version, please replace the `$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)` with specific version number, for example v.1.13.0.
2. Make the kubectl binary executeable `chmod +x ./kubectl`
3. Move the binary to your path `sudo mv ./kubectl /usr/local/bin/kubectl`

## Minikube

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

## Backend And Frontend Application



## GoCD

GoCD is a continous integration and delivery tool from Thoughtworks. 

To learn more about GoCD please go into their website [here](https://www.gocd.org/getting-started/part-1/)

### Building Application With GoCD

## DroneIO

## Jenkins
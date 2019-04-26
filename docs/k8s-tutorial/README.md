# Kubernetes Tutorial

## Disclaimer

This tutorial is for **beginner** in kubernetes, not for managing kubernetes cluster and its maintenance. This all about introduction to container, basic of kubernetes, CI/CD and deploying application to kubernetes.

## Agenda

### Installing Pre-Requisite

1. Install `docker`. Please look into the tutorial [here](https://github.com/albertwidi/k8s-tutorial/tree/master/docs/k8s-tutorial#install-docker)
2. Install `minikube`. Please look into the tutorial [here](https://github.com/albertwidi/k8s-tutorial/tree/master/docs/k8s-tutorial#installing-minikube)

### Kubernetes Basic

1. Namespace
2. Pod
3. Deployment
4. Service
5. StatefulSet
6. ReplicaSet
7. DaemonSet
8. Job
9. Endpoint
10. Ingress

### CI/CD With Kubernetes

This kubernetes tutorial is using `minikube` and will cover this following agenda:

1. Install helm into `minikube`

2. Install jenkins via `stable/helm` charts

3. Build web application using `go`

4. Build an application using GoCD

5. Deploy an application to kubernetes cluster using helm

## Pre-Requisite

To be able to run this tutorial, you need to install several things:

### System Requirement

This tutorial is build on-top `ubuntu 18.04` and is using around 4-12Gb depends on what is running on minikube.

The safe point is around 8Gb

### Install Go

This tutorial is using `go` to create programs, please make sure you have `go` inside your computer.

To install go, please follow [this](https://golang.org/doc/install) link to go to golang.org tutorial for installing `go`

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

Minikube is a tool that makes it easy to run kubernetes locally.

### Installing Minikube

To install minikube, please follow [this](https://kubernetes.io/docs/tasks/tools/install-minikube/) link. Or follow the steps below:

**Linux**

Install KVM2 driver, the KVM2 driver is intended to replace the KVM driver. KVM driver is actually deprecated:

```sh
# Install libvirt and qemu-kvm on your system, e.g.
# Debian/Ubuntu (for older Debian/Ubuntu versions, you may have to use libvirt-bin instead of libvirt-clients and libvirt-daemon-system)
sudo apt install libvirt-clients libvirt-daemon-system qemu-kvm
# Fedora/CentOS/RHEL
sudo yum install libvirt-daemon-kvm qemu-kvm

# Add yourself to the libvirt group so you don't need to sudo
# NOTE: For older Debian/Ubuntu versions change the group to `libvirtd`
sudo usermod -a -G libvirt $(whoami)

# Update your current session for the group change to take effect
# NOTE: For older Debian/Ubuntu versions change the group to `libvirtd`
newgrp libvirt
```

Then install the driver itself:

```sh
curl -LO https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-kvm2 \
  && sudo install docker-machine-driver-kvm2 /usr/local/bin/
```

**MacOS**

For MacOS user, you can use Kubernetes in Docker for Mac. Or if you still want to install minikube, you can read the steps below:

This text below is a copy of [this](https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#vmware-unified-driver) link

Using VMware unified driver

The VMware unified driver will eventually replace the existing vmwarefusion driver. The new unified driver supports both VMware Fusion (on macOS) and VMware Workstation (on Linux and Windows)

To install the vmware unified driver, head over at [here](https://github.com/machine-drivers/docker-machine-driver-vmware/releases) and download the release for your operating system.

The driver must be:

1. Stored in $PATH
2. Named docker-machine-driver-vmware
3. Executable (chmod +x on UNIX based platforms)
4. If you're running on macOS with Fusion, this is an easy way install the driver:

```sh
export LATEST_VERSION=$(curl -L -s -H 'Accept: application/json' https://github.com/machine-drivers/docker-machine-driver-vmware/releases/latest | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/') \
&& curl -L -o docker-machine-driver-vmware https://github.com/machine-drivers/docker-machine-driver-vmware/releases/download/$LATEST_VERSION/docker-machine-driver-vmware_darwin_amd64 \
&& chmod +x docker-machine-driver-vmware \
&& mv docker-machine-driver-vmware /usr/local/bin/
```

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

## GoCD

GoCD is a continous integration and delivery tool from Thoughtworks. 

To learn more about GoCD please go into their website [here](https://www.gocd.org/getting-started/part-1/)

## DroneIO

## Jenkins

## The Tutorial

In this tutorial, we will use `go` as our main language to build programs, and 
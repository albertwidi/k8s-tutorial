# Docker Tutorial

## Origin of container

### Cgroups

### Namespace

## Install Docker

To install docker

## Docker Image

### Building Docker Image

`docker build -f ./backend/Dockerfile -t albertwidi/k8s-tutorial-backend . --rm`

and 

`cd backend && docker build -f Dockerfile -t albertwidi/k8s-tutorial-backend . --rm`

Is different, because docker will always assume the image is being built from the relative path and use `/var/lib/docker/tmp` instead of current directory.

### Push Docker Image to Image Registry
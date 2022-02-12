#!/usr/bin/env sh

# This is meant to be executed on the machine holding the kubernetes cluster!

kubectl apply -f db-service.yaml,microservice-tcp-service.yaml,db-deployment.yaml,microservice_net-networkpolicy.yaml,microservice-deployment.yaml,microservice-claim0-persistentvolumeclaim.yaml

#!/bin/bash
kubectl delete -n natss -f ~/go/src/github.com/knative/eventing/contrib/natss/config/broker/natss.yaml
kubectl delete  ns natss
kubectl delete -f https://github.com/knative/eventing/releases/download/v0.5.0/natss.yaml
kubectl delete -f https://github.com/knative/eventing/releases/download/v0.5.0/eventing.yaml
kubectl delete -f https://github.com/knative/serving/releases/download/v0.5.0/serving.yaml
kubectl delete -f https://github.com/knative/serving/releases/download/v0.5.0/istio.yaml
kubectl delete -f https://github.com/knative/serving/releases/download/v0.5.0/istio-crds.yaml



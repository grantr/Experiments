# #!/bin/bash
# kubectl apply -f https://github.com/knative/serving/releases/download/v0.4.0/istio-crds.yaml
# kubectl apply -f https://github.com/knative/serving/releases/download/v0.4.0/istio.yaml
# kubectl apply -f https://github.com/knative/serving/releases/download/v0.4.0/serving.yaml
# kubectl apply -f https://github.com/knative/eventing/releases/download/v0.4.0/eventing.yaml
# kubectl apply -f https://github.com/knative/eventing/releases/download/v0.4.0/natss.yaml
# kubectl create namespace natss
# kubectl label namespace natss istio-injection=enabled
# kubectl apply -n natss -f ~/go/src/github.com/knative/eventing/contrib/natss/config/broker/natss.yaml

kubectl apply -f https://github.com/knative/serving/releases/download/v0.5.1/istio-crds.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/v0.5.1/istio.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/v0.5.1/serving.yaml
kubectl apply -f https://github.com/knative/eventing/releases/download/v0.5.0/eventing.yaml
kubectl apply -f https://github.com/knative/eventing/releases/download/v0.5.0/natss.yaml
kubectl create namespace natss
kubectl label namespace natss istio-injection=enabled
kubectl apply -n natss -f ~/go/src/github.com/knative/eventing/contrib/natss/config/broker/natss.yaml
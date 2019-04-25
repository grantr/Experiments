gcloud container clusters create $1\
  --zone=us-west1-b \
  --cluster-version=latest \
  --num-nodes=3 \
  --machine-type=n1-standard-4 \
  --enable-autoscaling --min-nodes=1 --max-nodes=3 \
  --enable-autorepair \
  --scopes=service-control,service-management,compute-rw,storage-ro,cloud-platform,logging-write,monitoring-write,pubsub,datastore

kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole=cluster-admin \
  --user=$(gcloud config get-value core/account)

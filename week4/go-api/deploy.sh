#!/bin/bash
DIRPATH=/Users/mariia.rubina13/Projects/cloud/week4
#delete old deployments and services
kubectl delete -f "$DIRPATH/deployment/api.yaml" 
# kubectl delete -f "$DIRPATH/deployment/service.yaml" 
#apply new
kubectl apply -f "$DIRPATH/deployment/api.yaml"
# kubectl apply -f "$DIRPATH/deployment/service.yaml"
# find external ip
IP=$(kubectl get service api-svc -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
echo "$IP"
while [ "$IP" = "pending" ]
do
echo "$IP"
sleep 30
IP=$(kubectl get service api-svc -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
done
echo "$IP"
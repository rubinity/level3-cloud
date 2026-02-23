#!/bin/bash
DIRPATH=/Users/mariia.rubina13/Projects/cloud/week5/vue/redis
#delete old deployments and services
kubectl delete -f "$DIRPATH/deployment/redis-ui.yaml" 
# kubectl delete -f "$DIRPATH/deployment/service.yaml" 
#apply new
kubectl apply -f "$DIRPATH/deployment/redis-ui.yaml"
# kubectl apply -f "$DIRPATH/deployment/service-ui.yaml"
# find external ip
IP=$(kubectl get service redis-svc -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
echo "$IP"
while [ "$IP" = "pending" ]
do
echo "$IP"
sleep 30
IP=$(kubectl get service redis-svc -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
done
echo "$IP"
terraform import openstack_networking_network_v2.public 991250c0-a2f4-4d9e-858f-867d176092ba

mariia-cluster


export STACKIT_SERVICE_ACCOUNT_KEY=/Users/mariia.rubina13/Projects/cloud/week3/credentials/credentials.json
stackit auth activate-service-account
export KUBERNETES_SERVICE_HOST=https://api.mar-cluster.38183bc8e9.s.ske.eu01.onstackit.cloud
ske id 261006cb-88f1-4a5a-b4cb-1341dad5f39b 
access via cubctl
install 
kubectl on mac
curl -LO "https://dl.k8s.io/release/v1.35.0/bin/darwin/amd64/kubectl"

curl http://188.34.100.217:8001/version

188.34.100.217:8001


kubectl run net-debug --rm -it --image=curlimages/curl:8.5.0 --restart=Never -- sh

curl -I https://registry.ske.eu01.stackit.cloud
kubectl exec -it <pod-name> -- sh
curl http://kubernetes-bootcamp-658f6cbd58-85dvd.default.svc.cluster.local:8080


delet oml
kubectl delete namespace olm
kubectl delete namespace operators


kubectl get crds -o name | grep -E 'operators.coreos.com|olm|catalogsources' | xargs -r kubectl delete
kubectl get clusterrole -o name | grep -E 'olm|operator-lifecycle-manager' | xargs -r kubectl delete
kubectl get clusterrolebinding -o name | grep -E 'olm|operator-lifecycle-manager' | xargs -r kubectl delete

kubectl get sa -n olm -o name | xargs -r kubectl delete
kubectl get sa -n operators -o name | xargs -r kubectl delete

kubectl get all --all-namespaces | grep -E 'olm|catalog'
kubectl get crds | grep -E 'operators.coreos.com|olm|catalogsources'
kubectl get clusterrole | grep -E 'olm|operator-lifecycle-manager'
kubectl get clusterrolebinding | grep -E 'olm|operator-lifecycle-manager'


check
kubectl get pods -n olm
kubectl logs redis-operator-5b45f955b7-zc9nk -n olm
wget --ca-certificate=$CA --header="Authorization: Bearer $TOKEN" https://kubernetes.default.svc/version
wget --no-check-certificate --header="Authorization: Bearer $TOKEN" -qO- https://kubernetes.default.svc/version



http://[pod_name].default.svc.cluster.local:[podport]

https://ske.api.stackit.cloud/v2/projects/261006cb-88f1-4a5a-b4cb-1341dad5f39b/regions/eu01/clusters


https://api.mar-cluster.38183bc8e9.s.ske.eu01.onstackit.cloud
api.mar-cluster.38183bc8e9.internal.ske.eu01.stackit.cloud
https://api.mar-cluster.38183bc8e9.internal.ske.eu01.stackit.cloud
kubectl -n olm patch deployment catalog-operator \
  --type='json' -p='[{
    "op": "add",
    "path": "/spec/template/spec/containers/0/env/-",
    "value": {
      "name": "KUBERNETES_SERVICE_HOST",
      "value": "api.mar-cluster.38183bc8e9.s.ske.eu01.onstackit.cloud"
    }
  }]'
kubectl run debug-pod --rm -i --tty --image=bitnami/kubectl --restart=Never
kubectl -n olm run debug-pod --rm -i --tty --image=bitnami/kubectl --restart=Never \
  --env="KUBERNETES_SERVICE_HOST=api.mar-cluster.38183bc8e9.s.ske.eu01.onstackit.cloud" \
  --overrides='
{
  "spec": {
    "securityContext": {
      "runAsNonRoot": true,
      "runAsUser": 1000,
      "runAsGroup": 3000,
      "fsGroup": 2000
    },
    "containers": [{
      "name": "debug-pod",
      "image": "bitnami/kubectl",
      "securityContext": {
        "allowPrivilegeEscalation": false,
        "capabilities": {"drop":["ALL"]},
        "runAsNonRoot": true,
        "seccompProfile": {"type":"RuntimeDefault"}
      },
      "stdin": true,
      "tty": true
    }]
  }
}'
curl -k https://$KUBERNETES_SERVICE_HOST

Error: error configuring catalog operator: Get "https://api.mar-cluster.38183bc8e9.internal.ske.eu01.stackit.cloud:443/apis/operators.coreos.com/v1/operatorgroups": dial tcp: lookup api.mar-cluster.38183bc8e9.internal.ske.eu01.stackit.cloud: i/o timeout
Error: error configuring catalog operator: Get "https://api.mar-cluster.38183bc8e9.internal.ske.eu01.stackit.cloud:443/apis/operators.coreos.com/v1/operatorgroups": dial tcp: lookup api.mar-cluster.38183bc8e9.internal.ske.eu01.stackit.cloud: i/o timeout

redis validation
kubectl describe --namespace ot-operators pods
kubectl get pods -n ot-operators

redis-cli health check
INFO replication
PING
ROLE
SET healthcheck ok
GET healthcheck
INFO clients
INFO SERVER 

check delete behavior
kubectl get pods -n ot-operators -w
kubectl delete pod -n ot-operators redis-replication-2

stackit ske cluster hibernate mar-cluster

Activate service account authentication in the STACKIT CLI using a service account key which includes the private key
  $ stackit auth activate-service-account --service-account-key-path path/to/service_account_key.json

  Activate service account authentication in the STACKIT CLI using the service account key and explicitly providing the private key in a PEM encoded file, which will take precedence over the one in the service account key
  $ stackit auth activate-service-account --service-account-key-path path/to/service_account_key.json --private-key-path path/to/private.key

  Activate service account authentication in the STACKIT CLI using the service account token
  $ stackit auth activate-service-account --service-account-token my-service-account-token

  Only print the corresponding access token by using the service account token. This access token can be stored as environment variable (STACKIT_ACCESS_TOKEN) in order to be used for all subsequent commands.
  $ stackit auth activate-service-account --service-account-token my-service-account-token --only-print-access-token



  kubectl run net-debug \
  --rm -it \
  --restart=Never \
  --image=curlimages/curl:8.5.0 \
  -- sh

  nslookup bootcamp-service.default.svc.cluster.local

  http://188.34.74.142
kubectl exec -it kubernetes-bootcamp-658f6cbd58-vsp4k -- netstat -tlnp
kubectl exec -it kubernetes-bootcamp-658f6cbd58-vsp4k -- ss -tln
89e2ca32-744b-465a-9888-c0e43c75c54a mar
 057e3708-444b-4e14-a903-26fe80f51d3b

  IPv4: allow all incoming udp traffic with port range 30000-32767 │ udp      │ ingress   │ IPv4       │ 30000-32767 │ 0.0.0.0/0    

  curl http://188.34.74.71
  redis-cli -h 188.34.73.3 -p 6379
  brew install redis
  redis-cli --version

  redis-ext    LoadBalancer   100.82.77.115    188.34.73.3   6379:31548/TCP   12m
  simple-web   LoadBalancer   100.82.156.129   188.34.74.71   80:30676/TCP   60m
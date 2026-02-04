terraform import openstack_networking_network_v2.public 991250c0-a2f4-4d9e-858f-867d176092ba
https://docs.stackit.cloud/products/runtime/kubernetes-engine/basics/basics/
https://registry.terraform.io/providers/stackitcloud/stackit/latest/docs/resources/ske_cluster
https://docs.api.eu01.stackit.cloud/documentation/ske
https://github.com/stackitcloud/stackit-cli/blob/main/docs/stackit.md
https://github.com/stackitcloud/stackit-cli/blob/main/docs/stackit_ske.md

https://medium.com/developingnodes/mastering-kubernetes-operators-your-definitive-guide-to-starting-strong-70ff43579eb9
https://www.youtube.com/watch?v=mTC3UZ8bHJc
https://github.com/OT-CONTAINER-KIT/redis-operator
https://ot-redis-operator.netlify.app/docs/installation/validation/

https://redis.io/docs/latest/commands/info/
mariia-cluster

stackit auth activate-service-account --key-file /Users/mariia.rubina13/Projects/cloud/week3/credentials/credentials.json
stackit auth activate-service-account --private-key-path /Users/mariia.rubina13/.ssh/sa-key-dadde60a-e135-4d29-a4dd-c996b87c51cf-private.pem
export STACKIT_SERVICE_ACCOUNT_KEY=/Users/mariia.rubina13/Projects/cloud/week3/credentials/credentials.json
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
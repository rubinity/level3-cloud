# helm_release for redis operator corresponds to  the following CLI command:
# helm install redis-replication ot-helm/redis-replication \
#   --set redisreplication.clusterSize=3 --namespace ot-operators
resource "helm_release" "redis-operator" {
  name       = "redis-operator"
  namespace   = "ot-operators"
  repository = "https://ot-container-kit.github.io/helm-charts/"
  chart      = "redis-operator"
# for upgrade insert the correct version into the next line, uncomment it and apply using terrafor
#   version    = ""

  set = [
    {
    name  = "featureGates.GenerateConfigInInitContainer"
    value = true
    }
  ]
}
# redis cluster
# $ helm upgrade redis ot-helm/redis -f custom-values.yaml \
#     --install --namespace ot-operators
# $ helm upgrade redis-cluster ot-helm/redis-cluster -f custom-values.yaml \
#   --set redisCluster.clusterSize=3 --install --namespace ot-operators
# helm_release for redis replication corresponds to  the following CLI command:
# helm install redis-operator ot-helm/redis-operator --namespace ot-operators --set featureGates.GenerateConfigInInitContainer=true
# resource "helm_release" "redis-replication" {
#   name       = "redis-replication"
#   namespace   = "ot-operators"
#   repository = "https://ot-container-kit.github.io/helm-charts/"
#   chart      = "redis-replication"
#   values = [
#     file("${path.module}/manifest/resources/values.yaml")
#   ]
# }

# redis standalone

# helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
# helm repo update
# kubectl create ns ingress-nginx
# helm install ingress-nginx ingress-nginx/ingress-nginx --namespace ingress-nginx


resource "helm_release" "ingress-nginx" {
  name       = "ingress-nginx"
  namespace   = "ingress-nginx"
  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"
# for upgrade insert the correct version into the next line, uncomment it and apply using terrafor
#   version    = ""
}

# helm repo add jetstack https://charts.jetstack.io
# helm install cert-manager --namespace cert-manager jetstack/cert-manager --set webhook.timeoutSeconds=15 --set crds.enabled
resource "helm_release" "redis-operator" {
  name       = "redis-operator"
  namespace   = "ot-operators"
  repository = "https://ot-container-kit.github.io/helm-charts/"
  chart      = "redis-operator"
# for upgrade insert the correct version into the next line, uncomment it and apply using terrafor
#   version    = ""

  set = [
    {
    name  = "featureGates.GenerateConfigInInitContainer"
    value = true
    }
  ]
}
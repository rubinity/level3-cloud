resource "kubernetes_namespace_v1" "test-namespace" {
  metadata {
    name = "test-namespace"
  }
}

resource "kubernetes_namespace_v1" "ot-operators" {
  metadata {
    name = "ot-operators"
  }
}

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

resource "helm_release" "redis-replication" {
  name       = "redis-replication"
  namespace   = "ot-operators"
  repository = "https://ot-container-kit.github.io/helm-charts/"
  chart      = "redis-replication"

  set = [
    {
    name  = "redisreplication.clusterSize"
    value = "3"
    }
  ]
}

# helm install redis-replication ot-helm/redis-replication \
#   --set redisreplication.clusterSize=3 --namespace ot-operators

# helm repo add ot-helm https://ot-container-kit.github.io/helm-charts/
# $ helm install redis-operator ot-helm/redis-operator --namespace ot-operators --set featureGates.GenerateConfigInInitContainer=true
# ...

# data "kubernetes_endpoints_v1" "endpoints" {
#     metadata {
#       name      = "my-service"
#       namespace = "default"
#     }
# }

# output "k8s_api_endpoint" {
#   value = kubernetes.host
# }

# kubernetes_endpoints_v1
# output "ske_api_endpoint" {
#   description = "The Kubernetes API endpoint for the STACKIT cluster"
#   value       = stackit_ske_cluster.mariia-cluster.api_address
# }

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

# helm_release for redis replication corresponds to  the following CLI command:
# helm install redis-operator ot-helm/redis-operator --namespace ot-operators --set featureGates.GenerateConfigInInitContainer=true
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

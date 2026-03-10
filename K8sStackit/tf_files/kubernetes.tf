resource "kubernetes_namespace_v1" "ot-operators" {
  metadata {
    name = "ot-operators"
  }
}

# resource "kubernetes_api_service_v1" "example" {
#   metadata {
#     name = "terraform-example"
#   }
#   spec {
#     version = kubernetes_api_service_v1
#     version_priority = 1
#     group = redis
#     group_priority_minimum = 1
#     selector  = {
#       app = "${kubernetes_pod.example.metadata.0.labels.app}"
#     }
#     session_affinity = "ClientIP"
#     port {
#       port        = 8080
#       target_port = 80
#     }

#     type = "LoadBalancer"
#   }
# }

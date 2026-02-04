resource "stackit_ske_cluster" "mariia-cluster" {
  project_id             = "261006cb-88f1-4a5a-b4cb-1341dad5f39b"
  name                   = "mar-cluster"
  kubernetes_version_min = "1.34.3"
  node_pools = [
    {
      name               = "np-mariia"
      machine_type       = "g1a.2d"
      os_name            = "ubuntu"
      os_version_min     = "2204.20250728.0"
      minimum            = "2"
      maximum            = "2"
      availability_zones = ["eu01-3"]
    }
  ]

  maintenance = {
    enable_kubernetes_version_updates    = true
    enable_machine_image_version_updates = true
    start                                = "01:00:00Z"
    end                                  = "02:00:00Z"
  }
}

resource "stackit_ske_kubeconfig" "kconfig" {
  project_id   = "261006cb-88f1-4a5a-b4cb-1341dad5f39b"
  cluster_name = stackit_ske_cluster.mariia-cluster.name

  refresh        = true
  expiration     = 7200 # 2 hours

}

output "api_endpoint_from_kubeconfig" {
  value = yamldecode(
    stackit_ske_kubeconfig.kconfig.kube_config
  )["clusters"][0]["cluster"]["server"]
  sensitive = true
}

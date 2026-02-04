terraform {
  required_providers {
    stackit = {
      source = "stackitcloud/stackit"
      version = "0.52.0"
    }
  }
}

provider "stackit" {
  # Configuration options
credentials_path = "/Users/mariia.rubina13/Projects/cloud/week3/credentials/credentials.json"
}

provider "kubernetes" {
  config_path    = "~/.kube/config"
  # host                   = yamldecode(stackit_ske_kubeconfig.example.kube_config).clusters.0.cluster.server
  # client_certificate     = base64decode(yamldecode(stackit_ske_kubeconfig.kconfig).users.0.user.client-certificate-data)
  # client_key             = base64decode(yamldecode(stackit_ske_kubeconfig.kconfig).users.0.user.client-key-data)
  # cluster_ca_certificate = base64decode(yamldecode(stackit_ske_kubeconfig.kconfig).clusters.0.cluster.certificate-authority-data)
}

provider "helm" {
  kubernetes = {
      config_path    = "~/.kube/config"
    }
  }


terraform {
  backend "s3" {
    bucket = "mariia"
    key    = "mariia/terraform/state/terraform.tfstate"
    endpoints = {
      s3 = "https://object.storage.eu01.onstackit.cloud"
    }
    region                      = "eu01"
    skip_credentials_validation = true
    skip_region_validation      = true
    skip_s3_checksum            = true
    skip_requesting_account_id  = true
    force_path_style            = true
    secret_key                  = "JADMs+l4kW/1B/jOdv2jE/+u+GsfHpvYsnLmccbr"
    access_key                  = "K7CCF20V2V9GS9CIBTX4"
  }
}
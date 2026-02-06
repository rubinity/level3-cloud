terraform {
  required_version = ">= 1.14.4"
}


terraform {
  required_providers {
    stackit = {
      source = "stackitcloud/stackit"
      version = "0.52.0"
    }
  }
}

provider "stackit" {
credentials_path = "/Users/mariia.rubina13/Projects/cloud/week3/credentials/credentials.json"
}

provider "kubernetes" {
  config_path    = "~/.kube/config"
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
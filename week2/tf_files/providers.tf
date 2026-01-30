terraform {
  required_providers {
    openstack = {
      source = "terraform-provider-openstack/openstack"
      version = "~>3.4.0"
    }
  }
}

provider "openstack" {
  user_name   = "admin"
  tenant_name = "cluster"
  password    = "TenTrees"
  auth_url    = "http://10.0.0.119/identity"
  region      = "RegionOne"
}

# local network
resource "openstack_networking_network_v2" "cluster_network_1" {
  name           = "cluster_network_1"
  admin_state_up = "true"
  shared = false
}

# external network
resource "openstack_networking_network_v2" "public" {
  name           = "public"
}

# subnet associated with a local network
resource "openstack_networking_subnet_v2" "cluster_subnet_1" {
  name       = "cluster_subnet_1"
  network_id = openstack_networking_network_v2.cluster_network_1.id
  cidr       = "10.0.0.64/26"
  ip_version = 4
  dns_nameservers = ["8.8.8.8", "1.1.1.1"]
  allocation_pool {
    start = "10.0.0.66"
    end = "10.0.0.120"
  }
}

# subnet associated with a public network
resource "openstack_networking_subnet_v2" "public_subnet" {
  name       = "public_subnet"
  network_id = openstack_networking_network_v2.public.id
  # cidr       = "172.24.4.0/24"
  # ip_version = 4
  # dns_nameservers = ["8.8.8.8", "1.1.1.1"]
}


# create router attached to external network
resource "openstack_networking_router_v2" "cluster_router" {
  name           = "cluster_router"
  admin_state_up = "true"
  # external_network_id = "991250c0-a2f4-4d9e-858f-867d176092ba"
  external_network_id = openstack_networking_network_v2.public.id
}

# attach private subnet to router
resource "openstack_networking_router_interface_v2" "int_2" {
  router_id = openstack_networking_router_v2.cluster_router.id
  subnet_id = openstack_networking_subnet_v2.cluster_subnet_1.id
}


resource "openstack_networking_port_v2" "port_2" {
  name           = "port_2"
  network_id     = openstack_networking_network_v2.cluster_network_1.id
  # security_group_ids = [openstack_networking_secgroup_v2.cluster_secgroup.id]
  port_security_enabled = true
  fixed_ip {
    subnet_id  = openstack_networking_subnet_v2.cluster_subnet_1.id
    ip_address = "10.0.0.121"
  }
  # admin_state_up = "true"
}

#create floating ip
resource "openstack_networking_floatingip_v2" "floatip_1" {
  pool = "public"
}

#attaching floating ip to a port
resource "openstack_networking_floatingip_associate_v2" "fip_1" {
  floating_ip = openstack_networking_floatingip_v2.floatip_1.address
  port_id    = openstack_networking_port_v2.port_2.id
}

data "openstack_networking_secgroup_v2" "secgroup" {
  name = "default"
  tenant_id = data.openstack_identity_project_v3.cluster_project.id
}
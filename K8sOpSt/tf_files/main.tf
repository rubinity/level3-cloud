

#create flavor
resource "openstack_compute_flavor_v2" "ubuntu-max-flavor" {
  name  = "u.max"
  ram   = "4096"
  vcpus = "2"
  disk  = "25"
  flavor_id = "43"
  is_public = true
}

resource "openstack_compute_flavor_v2" "ubuntu-worker-flavor" {
  name  = "u.worker"
  ram   = "4096"
  vcpus = "1"
  disk  = "15"
  flavor_id = "44"
  is_public = true
}


# ubuntu VM
resource "openstack_compute_instance_v2" "ClusterVM" {
  name = "ClusterVM"
  image_id = openstack_images_image_v2.ubuntu-22-04-server-cloudimg-amd64.id
  flavor_id = "43"
  key_pair = "clusterkeys"
  security_groups = ["default"]
  tags = ["Cluster"]
  network {
    port = openstack_networking_port_v2.port_2.id
  }
  user_data = file("${path.module}/init.sh")
}


# add image to glance
resource "openstack_images_image_v2" "ubuntu-22-04-server-cloudimg-amd64" {
  name             = "ubuntu-22-04-server-cloudimg-amd64"
  image_source_url = "https://cloud-images.ubuntu.com/releases/jammy/release-20260119/ubuntu-22.04-server-cloudimg-amd64.img"
  container_format = "bare"
  disk_format      = "qcow2"
  visibility = "public"
}

data "openstack_identity_project_v3" "cluster_project" {
  name = "cluster"
}


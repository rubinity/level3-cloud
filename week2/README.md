# 1. Install terraform
# 1.1. Preparation to terraform installation
```
#creating a new project
source /opt/stack/devstack/openrc admin admin
openstack project create cluster
openstack role add --project cluster --user admin admin
source /opt/stack/devstack/openrc admin cluster
#installation
mkdir terraform
cd terraform
curl -o terraform.zip https://releases.hashicorp.com/terraform/1.14.3/terraform_1.14.3_linux_amd64.zip
unzip terraform.zip
sudo mv terraform /usr/local/bin

```
# 1.2. Validation
```
terraform -version
terraform validate
```
# 1.3. Init terraform
```
//added provider to main.tf
terraform init
//edited provider in main.tf
terraform init -upgrade

```
# 1.4. Creating tf files and provisioning k8s cluster

//Script for easy tf files transfer
```
#copy files from local to remote
scp week2/tf_files/*.tf stack@188.34.101.189:~/terraform
```

//importing existing networks
terraform import openstack_networking_network_v2.public 991250c0-a2f4-4d9e-858f-867d176092ba
terraform import openstack_networking_subnet_v2.public_subnet 8c8242c1-b4d3-4de7-af97-a9e11fed3887
#creating a keypair
openstack keypair create --type ssh clusterkeys > cluster

# 2. k8s install
[k8s install](/scripts/init.sh)

some commands for checking health
kubectl get nodes
kubectl get componentstatuses
kubectl get pods -n kube-system
kubectl get --raw='/readyz?verbose'
kubectl cluster-info

#test with a simple pod
kubectl run smoke-test --image=busybox --restart=Never -- sleep 10
kubectl get pod smoke-test


# get api running
#copy kubectl config to accessible folder
mkdir -p $HOME/.kube
sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
#untain the node because it's not only control plane node
kubectl taint nodes clustervm node-role.kubernetes.io/control-plane-
#apply CNI
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
#mount bpffs - required for Calico eBPF dataplane (kernel-level prerequisite)
echo "bpffs /sys/fs/bpf bpf defaults 0 0" | sudo tee -a /etc/fstab
#restart
sudo systemctl restart kubelet
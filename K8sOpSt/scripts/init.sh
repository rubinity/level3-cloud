#!/bin/bash
# script with external downloads
# add hostname
# HOSTNAME="clustervm"
IP=$(hostname -I | awk '{print $1}')
# create post init script
echo 'mkdir -p $HOME/.kube' > /home/ubuntu/post-init.sh
echo 'sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config' >> /home/ubuntu/post-init.sh
echo 'sudo chown $(id -u):$(id -g) $HOME/.kube/config' >> /home/ubuntu/post-init.sh
echo 'kubectl apply -f calico.yaml' >> /home/ubuntu/post-init.sh
echo 'sudo systemctl restart kubelet' >> /home/ubuntu/post-init.sh
sudo chmod +x /home/ubuntu/post-init.sh
# echo "$IP $HOSTNAME" | sudo tee -a /etc/hosts
# 1. container runtime (containerd)
# You need to install a container runtime into each node in the cluster so that Pods can run there. 
# sysctl params required by setup, params persist across reboots
# mount bpffs - required for Calico eBPF dataplane (kernel-level prerequisite)
echo "bpffs /sys/fs/bpf bpf defaults 0 0" | sudo tee -a /etc/fstab
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.ipv4.ip_forward = 1
EOF
# Apply sysctl params without reboot
sudo sysctl --system
# download containerd
curl -LO https://github.com/containerd/containerd/releases/download/v2.2.1/containerd-2.2.1-linux-amd64.tar.gz
sudo tar Cxzvf /usr/local containerd-2.2.1-linux-amd64.tar.gz
echo "installing containerd:"
sudo mkdir -p /usr/local/lib/systemd/system
sudo curl -L -o /usr/local/lib/systemd/system/containerd.service https://raw.githubusercontent.com/containerd/containerd/main/containerd.service 
sudo systemctl daemon-reload
sudo systemctl enable --now containerd
#installing runc
echo "installing runc:"
sudo curl -LO https://github.com/opencontainers/runc/releases/download/v1.4.0/runc.amd64
sudo install -m 755 runc.amd64 /usr/local/sbin/runc
#Installing CNI plugins
echo "installing CNI plugins:"
sudo curl -LO https://github.com/containernetworking/plugins/releases/download/v1.9.0/cni-plugins-linux-amd64-v1.9.0.tgz
sudo mkdir -p /opt/cni/bin
sudo tar Cxzvf /opt/cni/bin cni-plugins-linux-amd64-v1.9.0.tgz
curl -Lo /home/ubuntu/calico.yaml https://docs.projectcalico.org/manifests/calico.yaml
# configuration file
echo "editing config:"
sudo mkdir -p /etc/containerd
sudo containerd config default | sudo tee /etc/containerd/config.toml
# change value in config for groups
sudo sed -i -e "s/\(SystemdCgroup *= *\).*/\1true/" /etc/containerd/config.toml
sudo systemctl restart containerd
echo "containerd is installed"
# Now install kubeadm, kubelet, kubectl
# Update the apt package index and install packages needed to use the Kubernetes apt repository:
echo "installing kubelet kubeadm kubctl:"
sudo apt-get update
# apt-transport-https may be a dummy package; if so, you can skip that package
# Download the public signing key for the Kubernetes package repositories. 
# If the directory `/etc/apt/keyrings` does not exist, it should be created before the curl command, read the note below.
# sudo mkdir -p -m 755 /etc/apt/keyrings
sudo curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.35/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
# Add the appropriate Kubernetes apt repository.
# This overwrites any existing configuration in /etc/apt/sources.list.d/kubernetes.list
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.35/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
sudo systemctl enable --now kubelet
# before creating a cluster with kubeadm
echo "editing cubeadm config:"
sudo kubeadm config print init-defaults | sudo tee /etc/kubernetes/kubeadm-config.yaml
sudo sed -i -e "s/\(taints *: *\).*/\1[]/" /etc/kubernetes/kubeadm-config.yaml
sudo sed -i -e "s/\(advertiseAddress *: *\).*/\1$IP/" /etc/kubernetes/kubeadm-config.yaml
printf -- "---\nkind: KubeletConfiguration\napiVersion: kubelet.config.k8s.io/v1beta1\ncgroupDriver: systemd\n"| sudo tee -a /etc/kubernetes/kubeadm-config.yaml
sudo kubeadm init --config /etc/kubernetes/kubeadm-config.yaml
#after this script has finished running login as ubintu and run ./post-init.sh (without sudo!!!!)




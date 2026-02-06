
# Developing PaaS product

- [Developing PaaS product](#developing-paas-product)
- [1. Preparation](#1-preparation)
- [1.1. Service accounts](#11-service-accounts)
- [1.2. Configuring terraform](#12-configuring-terraform)
- [1.3. Set stackit provider](#13-set-stackit-provider)
- [2. Provisioning a cluster](#2-provisioning-a-cluster)
- [2.1 Creating a cluster](#21-creating-a-cluster)
- [2.2 Kubectl installation and Cluster validation](#22-kubectl-installation-and-cluster-validation)
- [2.3 Test deployment using k8s](#23-test-deployment-using-k8s)
- [3. Provisioning of an operator](#3-provisioning-of-an-operator)
- [3.1 Helm installation using terraform](#31-helm-installation-using-terraform)
- [3.2 Redis operator installation](#32-redis-operator-installation)
- [3.3 redis connectivity](#33-redis-connectivity)


# 1. Preparation
# 1.1. Service accounts
Create a service account on the Stackit portal in AIM AND MANAGEMENT -> Service accounts  
Via AIM AND MANAGEMENT -> Access give owner access to this account  
Create a key for the service account and download a json service account key file 
create a credential file with path to the json key  


# 1.2. Configuring terraform
Installed terraform  
```
terraform -version
terraform validate
```
in terraform added terraform provider

# 1.3. Set stackit provider
in terraform add stackit provider  
Use the path to credential file  
`terraform init`

# 2. Provisioning a cluster
# 2.1 Creating a cluster
add cluster resourse to main.tf  
```
terraform plan
terraform apply
```
wait until the cluster is ready
 
[tf_files](tf_files) 


# 2.2 Kubectl installation and Cluster validation

```
#kubectl installation on the remote machine
curl -Lo ~/bin/kubctl "https://dl.k8s.io/release/v1.35.0/bin/darwin/amd64/kubectl"
#some commands for checking health
kubectl get nodes
kubectl get pods -n kube-system
kubectl get --raw='/readyz?verbose'
kubectl cluster-info
```

# 2.3 Test deployment using k8s

```
#test with a simple pod
kubectl run smoke-test --image=busybox --restart=Never -- sleep 10
kubectl get pod smoke-test
kubectl delete pod smoke-test
```


# 3. Provisioning of an operator
# 3.1 Helm installation using terraform

add helm and kubernetes to terraform providers  
[providers](tf_files/providers.tf)  
`terraform init -upgrade`


# 3.2 Redis operator installation
create helm_release in terraform corresponding to redis operator  
[terraform](tf_files/helm.tf)   
apply 

# 3.3 redis connectivity
create load balancer service yaml  
[redis](tf_files/manifest/resources/redis.yaml)  

```
kubectl apply -f manifest/resources/redis.yaml
#install redis-cli to desktop
brew install redis
redis-cli --version
#look up for service ip
kubectl get service -n ot-operators
redis-cli -h [ip] -p 6379 
redis-cli -h 188.34.74.168 -p 6379
#inside client
info
```


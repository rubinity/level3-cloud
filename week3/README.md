
# Developing PaaS product

- [Developing PaaS product](#developing-paas-product)
- [1. Preparation](#1-preparation)
- [1.1. Service accounts](#11-service-accounts)
- [1.2. Configuring terraform](#12-configuring-terraform)
- [1.3. to do](#13-to-do)
- [1.4. to do](#14-to-do)
- [2. Provisioning a cluster](#2-provisioning-a-cluster)
- [2.1 Creating a cluster](#21-creating-a-cluster)
- [2.2 Kubectl instalation and Cluster validation](#22-kubectl-instalation-and-cluster-validation)
- [2.3 Test deployment using k8s](#23-test-deployment-using-k8s)
- [3. Provisioning of an operator](#3-provisioning-of-an-operator)
- [3.1 Helm installation using terraform](#31-helm-installation-using-terraform)
- [3.2 Redis operator installation](#32-redis-operator-installation)
- [4.](#4)


# 1. Preparation
# 1.1. Service accounts
Create a service account on the Stackit portal in AIM AND MANAGEMENT -> Service accounts
Via AIM AND MANAGEMENT -> Access give owner access to this account


```
stackit config set --project-id xxxx-xxxx-xxxxx
stackit config unset --project-id 
stackit config set --project-id 261006cb-88f1-4a5a-b4cb-1341dad5f39b
stackit ske enable -p [Project ID] 

```

# 1.2. Configuring terraform

```
terraform -version
terraform validate
```

# 1.3. to do


# 1.4. to do
added provider to main.tf
`terraform init`
edited provider in main.tf
`terraform init -upgrade`

# 2. Provisioning a cluster
# 2.1 Creating a cluster

```
terraform plan
terraform apply
```
wait until the cluster is ready
 
[tf_files](tf_files) 


# 2.2 Kubectl instalation and Cluster validation

```
#
kubectl cluster-info
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

curl http://kubernetes-bootcamp.default.svc.cluster.local:8080
curl http://localhost:8001/api/v1/namespaces/default/pods/test-service:8080/proxy/

```
qqqqq
```


# 3. Provisioning of an operator
# 3.1 Helm installation using terraform

add helm to terraform providers
[providers](tf_files/providers.tf) 

init upgrade


# 3.2 Redis operator installation
create helm_release in terraform corresponding to redis operator 
[terraform](tf_files/helm.tf)
apply 

# 4.
```
tttt

```

kubectl delete service -n ot-operators redis-ext
redis-cli -h 188.34.95.51 -p 6379 
kubectl get svc redis-ext -n ot-operators -o yaml
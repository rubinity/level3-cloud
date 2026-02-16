package main

// import rr "github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"

import (
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/api"
	// "os/exec"

	 
	// "k8s.io/api"

)



func main() {
	
	cli, clientset := api.InitClients()
	if err := api.Router(cli, clientset).Run("0.0.0.0:80"); err != nil{
		panic(err.Error())
	}


	
}


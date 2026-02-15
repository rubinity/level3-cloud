package main

// import rr "github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"

import (
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/api"
	// "os/exec"
	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	 
	// "k8s.io/api"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	scheme := runtime.NewScheme()
	v1beta2.AddToScheme(scheme)
	cli, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		panic(err.Error())
	}
	router := gin.Default()
	router.POST("/create", api.CreateReplHandler(cli))
	router.DELETE("/delete", api.DeleteReplHandler(cli))
	router.GET("/list/:ns", api.ListReplHandler(cli))
	router.GET("/connection/:ns/:name", api.ConnectionHandler(cli, clientset))
	router.Run("0.0.0.0:80")
}


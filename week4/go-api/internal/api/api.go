package api

import (

	// "log"
	// "os/exec"

	"github.com/gin-gonic/gin"

	"k8s.io/client-go/rest"

	// "k8s.io/api"
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"

	// "fmt"

	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Router(cli client.Client, clientset kubernetes.Interface) *gin.Engine {

	router := gin.Default()
	router.POST("/create", CreateReplHandler(cli))
	router.DELETE("/delete", DeleteReplHandler(cli))
	router.GET("/list/:ns", ListReplHandler(cli))
	router.GET("/connection/:ns/:name", ConnectionHandler(cli, clientset))
	// router.Run("0.0.0.0:80")
	return router
}

func InitClients() (cli client.Client, clientset kubernetes.Interface) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	scheme := runtime.NewScheme()
	v1beta2.AddToScheme(scheme)
	cli, err = client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		panic(err.Error())
	}
	return cli, clientset
}

func createNamespace(clientset kubernetes.Interface, ctx context.Context, namespace string) {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}

	_, _ = clientset.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	// !!todo if ns exist
	// if err != nil {
	// 	panic(err.Error())
	// }
}

func GetRepl(cli client.Client, ctx context.Context, namespace string, name string) (*v1beta2.RedisReplication, error) {
	var repl v1beta2.RedisReplication
	var key client.ObjectKey

	key.Name = name
	key.Namespace = namespace
	err := cli.Get(ctx, key, &repl)
	return &repl, err
}

func getlist(cli client.Client, ctx context.Context, namespace string, name string) *v1beta2.RedisReplicationList {
	var redislist v1beta2.RedisReplicationList
	var err error
	if namespace == "" && name == "" {
		err = cli.List(ctx, &redislist)
	} else {
		opts := setOpts(namespace, name)
		err = cli.List(ctx, &redislist, opts)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return &redislist
}

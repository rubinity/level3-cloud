package api

import (

	"log"
	"os"
	// "os/exec"

	"github.com/gin-gonic/gin"

	"k8s.io/client-go/rest"

	// "k8s.io/api"
	"context"
	"fmt"
"log/slog"
	// "os"
	"k8s.io/apimachinery/pkg/runtime"

	// "fmt"

	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"github.com/gin-contrib/cors"
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/auth"
	// "github.com/golangcompany/JWT-Authentication/routes"

)

func Router(cli client.Client, clientset kubernetes.Interface) *gin.Engine {
	
	if os.Getenv("ACCESS_SECRET") == "" {
		log.Fatalf("%s not set", "ACCESS_SECRET")
	}

	endpoint, _ := getEndpoint(clientset)
	println("endp", endpoint)
	rds := auth.NewRedis(endpoint)
	router := gin.Default()
	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
        	return true
    	},
		// AllowOrigins:     []string{os.Getenv("FRONTEND_ORIGIN")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	router.POST("/api/create",  auth.AuthMiddleware(rds), CreateReplHandler(cli, logger))
	router.DELETE("/api/delete", auth.AuthMiddleware(rds), DeleteReplHandler(cli, logger))
	router.GET("/api/list/:ns", auth.AuthMiddleware(rds), ListReplHandler(cli, logger))
	router.POST("/api/auth", AuthHandler(rds, clientset))
	router.POST("/api/logout", LogoutHandler(rds))
	router.GET("/api/connection/:ns/:name", auth.AuthMiddleware(rds), ConnectionHandler(cli, clientset, logger))
	return router
}

func TestRouter(cli client.Client, clientset kubernetes.Interface) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))	
	router.POST("/api/create", CreateReplHandler(cli, logger))
	router.DELETE("/api/delete", DeleteReplHandler(cli, logger))
	router.GET("/api/list/:ns", ListReplHandler(cli, logger))
	router.GET("/api/connection/:ns/:name", ConnectionHandler(cli, clientset, logger))
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

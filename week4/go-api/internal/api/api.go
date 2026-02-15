package api

import (

	// "log"
	// "os/exec"

	"context"
	"fmt"
	// "fmt"

	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)



func createNamespace(clientset *kubernetes.Clientset, ctx context.Context, namespace string){
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


func getRepl(cli client.Client, ctx context.Context, namespace string, name string) (*v1beta2.RedisReplication, error){
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
	if err != nil{
		fmt.Println(err.Error())
	}
	return &redislist
}






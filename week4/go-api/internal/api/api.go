package api

import (

		// "log"
	// "os/exec"
	"context"
	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	
	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)



func createNamespace(clientset *kubernetes.Clientset, namespace string){
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}

	_, _ = clientset.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	// !!todo if ns exist
		// if err != nil {
		// 	panic(err.Error())
		// }
}







func getRepl(cli client.Client, namespace string, name string) *v1beta2.RedisReplication{
	var repl v1beta2.RedisReplication
	var key client.ObjectKey

	key.Name = name
	key.Namespace = namespace
	cli.Get(context.TODO(), key, &repl)
	return &repl
}


func getlist(cli client.Client, namespace string, name string) *v1beta2.RedisReplicationList {
	var redislist v1beta2.RedisReplicationList
	var err error
	if namespace == "" && name == "" {
		err = cli.List(context.TODO(), &redislist)
	} else {
		opts := setOpts(namespace, name)
		err = cli.List(context.TODO(), &redislist, opts)
	}
	if err != nil {
		panic(err.Error())
	}
	return &redislist
}







func deleteRepl(cli client.Client, repl *v1beta2.RedisReplication) error{
	return cli.Delete(context.TODO(), repl);
}



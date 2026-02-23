func setResource(rr *v1beta2.RedisReplication, name string, namespace string, replnum int32){
	rr.Name = name
	rr.Namespace = namespace
	size := replnum
rau := int64(1000)
fsg := int64(1000)
var psc  corev1.PodSecurityContext

psc.RunAsUser = &rau
psc.FSGroup = &fsg
rr.Spec.Size = &size
rr.APIVersion = "redis.redis.opstreelabs.in/v1beta2"
rr.Kind = "redisReplication"
rr.Spec.KubernetesConfig.Image = "quay.io/opstree/redis:latest"
rr.Spec.KubernetesConfig.ImagePullPolicy = "IfNotPresent"
rr.Spec.PodSecurityContext = &psc
rr.Spec.Storage = &vb.Storage{}
rr.Spec.Storage.VolumeClaimTemplate.Spec.AccessModes = []corev1.PersistentVolumeAccessMode{
    corev1.ReadWriteOnce,
}
// rr.Spec.Storage.VolumeClaimTemplate.Spec.AccessModes = append(
//     rr.Spec.Storage.VolumeClaimTemplate.Spec.AccessModes,
//     corev1.ReadWriteOnce,
// )
rr.Spec.Storage.VolumeClaimTemplate.Spec.Resources.Requests = make(corev1.ResourceList)
rr.Spec.Storage.VolumeClaimTemplate.Spec.Resources.Requests[corev1.ResourceStorage] = resource.MustParse("1Gi")
var exporter  vb.RedisExporter
exporter.Image =  "quay.io/opstree/redis-exporter:latest"
exporter.Enabled = false
rr.Spec.RedisExporter = &exporter

}


// func createReplicationAPI(clientset *kubernetes.Clientset, dynclient *dynamic.DynamicClient) gin.HandlerFunc {

//     return func (c *gin.Context){
// 		log.Println("createReplicationAPI HIT")
//         var req ReplicationRequest
//        if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(400, gin.H{"error": err.Error()})
// 		createNamespace(clientset, req.Namespace)
// 		repl := createReplicationStructure(req.Namespace, int64(req.Number))
// 		gvr := schema.GroupVersionResource{
// 			Group:    "redis.redis.opstreelabs.in",
// 			Version:  "v1beta2",
// 			Resource: "redisreplications",   // plural form
// 		}

// 		repl, err = dynclient.Resource(gvr).Namespace(req.Namespace).Create(context.TODO(), repl, metav1.CreateOptions{})
// 		if err != nil {
// 			panic(err)
// 		}
// 		c.String(http.StatusOK, "replication %s created \n", repl.GetName())
// 		// fmt.Println("Created repl", repl.GetName())
//         // return
//     	}
// 	}
// }
// {"number": 3, "namespace": "myspace"}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	router.GET("/list", api.ListAPI(clientset))

	func ListAPI(clientset *kubernetes.Clientset) gin.HandlerFunc {

	return func(c *gin.Context) {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		c.String(http.StatusOK, "There are %d pods in the cluster\n", len(pods.Items))
		_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "api-app", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			c.String(http.StatusOK, "Pod api-app not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
			c.String(http.StatusOK, "Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found api-app pod in default namespace\n")
			c.String(http.StatusOK, "Found api-app pod in default namespace\n")
		}

	}
}

func HelloAPI(c *gin.Context) {
	c.String(http.StatusOK, "hello world!!\n")
}

// func ListReplHandler(cli client.Client) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		var redislist v1beta2.RedisReplicationList
// 		namespace := c.Param("ns")
// 		err := cli.List(context.TODO(), &redislist)
// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		c.String(http.StatusOK, "number of items %d \n", len(redislist.Items))
// 		fmt.Printf("number of items %d \n", len(redislist.Items))
// 		for _, item := range redislist.Items {
// 			// image := item.Spec.KubernetesConfig.Image
// 			var image string
// 			if item.Spec.RedisExporter != nil {
// 				image = item.Spec.RedisExporter.Image
// 			} else {
// 				image = "no"
// 			}

// 			c.String(http.StatusOK, "%s , %s, %s  \n", item.Name, image, item.Spec.KubernetesConfig.ImagePullPolicy)
// 		}
// 	}
// }
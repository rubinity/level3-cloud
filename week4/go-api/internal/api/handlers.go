package api

import (
	// "log"
	"net/http"

	// "os/exec"
	// "encoding/json"
	"fmt"

	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ListReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data []RRInfo
		namespace := c.Param("ns")
		// fmt.Println(namespace)
		list := getlist(cli, c.Request.Context(), namespace, "")
		for _, item := range list.Items {
			itemdata := RRInfo{
				Kind:   item.Kind,
				Name:   item.Name,
				Status: item.Status,
			}
			data = append(data, itemdata)
		}
		if data == nil {
			data = make([]RRInfo, 0)
		}
		c.JSON(http.StatusOK, data)
	}
}

func ConnectionHandler(cli client.Client, clientset *kubernetes.Clientset ) gin.HandlerFunc {

	return func(c *gin.Context) {
		namespace := c.Param("ns")
		name := c.Param("name")
		repl, _ := getRepl(cli, c.Request.Context(), namespace, name)
		service_name := name + "-additional"
		service, _ := clientset.CoreV1().Services(namespace).Get(c.Request.Context(), service_name, metav1.GetOptions{})
		c.JSON(http.StatusOK, gin.H{
			"connection": repl.Status.ConnectionInfo,
			"public_ip": service.Status.LoadBalancer.Ingress[0].IP,
		})
	}
}

func CreateReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		var repl v1beta2.RedisReplication
		var req ReplicationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		fmt.Println(req.Name)
		fmt.Println(req.Namespace)
		fmt.Println(req.Size)
		size := int32(req.Size)
		if size < 1 || size > 3 {
			size = 3
		}
		setResource(&repl, req.Name, req.Namespace, size)
		// data, _ := json.MarshalIndent(repl.Spec, "", "  ")
		// fmt.Println(string(data))
		err := cli.Create(c.Request.Context(), &repl)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"name":   repl.Name,
				"status": fmt.Sprintf("%s not created. Error: %s", repl.Name, err.Error()),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name":   repl.Name,
				"status": "created",
			})
		}

	}
}

func jsonError(c *gin.Context, name string, err string) {
	c.JSON(http.StatusOK, gin.H{
		"name":   name,
		"status": fmt.Sprintf("%s not deleted. Error: %s", name, err),
	})
}

func DeleteReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req DeleteRequest
		var err error
		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		repl, err := getRepl(cli, c.Request.Context(), req.Namespace, req.Name)
		if err == nil {
			err = cli.Delete(c.Request.Context(), repl)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"name":   repl.Name,
					"status": "deleted",
				})
				return
			}
		}
		jsonError(c, req.Name, err.Error())
	}
}

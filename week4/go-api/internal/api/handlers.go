package api

import (
	// "log"
	"net/http"

	// "os/exec"
	// "encoding/json"
	"fmt"

	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

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

func ConnectionHandler(cli client.Client, clientset kubernetes.Interface) gin.HandlerFunc {

	return func(c *gin.Context) {
		namespace := c.Param("ns")
		name := c.Param("name")
		repl, err := GetRepl(cli, c.Request.Context(), namespace, name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": fmt.Sprintf("%s not found Error: %s", name, err.Error()),
			})
			return
		}
		service_name := name + "-additional"
		service, err := clientset.CoreV1().Services(namespace).Get(c.Request.Context(), service_name, metav1.GetOptions{})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": fmt.Sprintf("service not found. Error: %s", err.Error()),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"connection": repl.Status.ConnectionInfo,
			"public_ip":  service.Status.LoadBalancer.Ingress[0].IP,
		})
	}
}

func CreateReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		var repl v1beta2.RedisReplication
		var req ReplicationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": fmt.Sprintf("Not created. Error: %s", err.Error()),
			})
			return
		}
		size := int32(req.Size)
		if size < 1 || size > 3 {
			size = 3
		}
		setResource(&repl, req.Name, req.Namespace, size)
		err := cli.Create(c.Request.Context(), &repl)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": fmt.Sprintf("%s not created. Error: %s", req.Name, err.Error()),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": fmt.Sprintf("%s created.", req.Name),
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
		repl, err := GetRepl(cli, c.Request.Context(), req.Namespace, req.Name)
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

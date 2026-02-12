package api

import (
	"log"
	"net/http"

	// "os/exec"
	"context"
	// "encoding/json"
	"fmt"

	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"github.com/gin-gonic/gin"
	// "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)


func ListReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		namespace := c.Param("ns")
		fmt.Println(namespace)
		list := getlist(cli, namespace, "")
		size := len(list.Items)
		c.String(http.StatusOK, "number of items %d \n", size)
		for _, item := range list.Items {
			c.String(http.StatusOK, "%s \n", item.Name)
		}
	}
}


func CreateReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		var rr v1beta2.RedisReplication
		var req ReplicationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		fmt.Println(req.Name)
		fmt.Println(req.Namespace)//check if exists!!!
		fmt.Println(req.Size)
		size := int32(req.Size)
		if size < 1 || size > 3{
			size = 3
		}
		setSpecs(&rr, req.Name, req.Namespace, size)
		// data, _ := json.MarshalIndent(rr.Spec, "", "  ")
		// fmt.Println(string(data))
		err := cli.Create(context.TODO(), &rr)
		if err != nil {
			panic(err.Error())
		}
		c.String(http.StatusOK, "created")
	}
}

func DeleteReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req DeleteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		repl := getRepl(cli, req.Namespace, req.Name)
		// !!todo error handling if item doesn't exist
		c.String(http.StatusOK, " item to be deleted %s  \n", repl.Name)
		if err := deleteRepl(cli, repl); err != nil {
			log.Printf("not deleted %s \n", err.Error())
			c.String(http.StatusSeeOther, "failed to delete %s  \n", repl.Name)
			return
		}
		c.String(http.StatusSeeOther, "deleted %s  \n", repl.Name)
	}
}

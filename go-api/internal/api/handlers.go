package api

import (
	// "log"
	"net/http"
	"strconv"

	// "os/exec"
	// "encoding/json"
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/auth"
	"fmt"

	"context"

	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// "github.com/joho/godotenv"
)

// @Summary Get list
// @Description Returns a list of replications in a namespace
// @Tags redis list
// @Produce json
// @Param body body ReplicationRequest true "Replication info"
// @Success 200 {object} RRInfo "Replication list"
// @Router /list{ns} [get]
func ListReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data []RRInfo
		namespace := c.Param("ns")
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

var demoUser = Namespace{ Namespace: "test2", Password: "pass"}
func AuthHandler(rds *auth.Redis, clientset kubernetes.Interface) gin.HandlerFunc {

	return func(c *gin.Context) {
		var in struct {
			Namespace string `json:"namespace"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid json"})
			return
		}
		if in.Namespace != demoUser.Namespace || in.Password != demoUser.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		toks, err := auth.IssueTokens(demoUser.Namespace)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not issue tokens"})
			return
		}
		if err := auth.Persist(c, rds, toks); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not persist tokens"})
			return
		}
		auth.SetAuthCookies(c, toks)
		c.JSON(http.StatusOK, gin.H{"ok": true})
	}
}

func LogoutHandler(rds *auth.Redis) gin.HandlerFunc {
	return func(c *gin.Context) {
		acc, _ := c.Cookie("access_token")
		ctx := context.Background()

		if acc != "" {
			if claims, err := auth.ParseAccess(acc); err == nil {
				_ = rds.DelJTI(ctx, claims.ID)
			}
		}
		auth.ClearAuthCookies(c)
		c.JSON(http.StatusOK, gin.H{"ok": true})	
	}
}


// @Summary Get connection info
// @Description Returns connection information by namespace and name
// @Tags connection
// @Produce json
// @Param body body ReplicationRequest true "Replication info"
// @Success 200 {object} map[string]interface{}"Creation result"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Not found"
// @Router /connection/{ns}/{name} [get]
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

func getEndpoint(clientset kubernetes.Interface) (endpoint string, err error){
	service, err := clientset.CoreV1().Endpoints("redis-auth").Get(context.TODO(), "store-master", metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	ip := service.Subsets[0].Addresses[0].IP
	port := int(service.Subsets[0].Ports[0].Port)
	endpoint = ip + ":" + strconv.Itoa(port)
	return endpoint, nil
}

// @Summary Create replication
// @Description Creates a redis replication cluster by namespace, name and size. The size can't be less than 1 or greater than 3 and is set to 3 if the value is wrong or undefined
// @Tags redis, replication
// @Accept json
// @Produce json
// @Param body body ReplicationRequest true "Creation info"
// @Success 200 {object} map[string]interface{}"Creation result"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Not found"
// @Router /create [post]
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
				"status": fmt.Sprintf("Replication %s created. Cluster size: %d", req.Name, size),
			})
		}
	}
}

func jsonError(c *gin.Context, code int, name string, err string) {
	c.JSON(code, gin.H{
		"status": fmt.Sprintf("%s not deleted. Error: %s", name, err),
	})
}

// @Summary Delete replication
// @Description Deletes a redis replication cluster by name and namespace
// @Tags redis, replication
// @Accept json
// @Produce json
// @Param body body DeleteRequest true "Replication info"
// @Success 200 {object} map[string]interface{} "Deletion result"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Not found"
// @Failure 409 {object} map[string]interface{} "Conflict"
// @Router /delete [post]
func DeleteReplHandler(cli client.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req DeleteRequest
		var err error
		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		repl, err := GetRepl(cli, c.Request.Context(), req.Namespace, req.Name)
		if err != nil {
			jsonError(c, http.StatusNotFound, req.Name, err.Error())
			return
		}
		err = cli.Delete(c.Request.Context(), repl)
		if err != nil {
			jsonError(c, http.StatusConflict, req.Name, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": fmt.Sprintf("%s deleted.", req.Name),
		})
	}
}

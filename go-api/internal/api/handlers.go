package api

import (
	// "log"
	"log/slog"
	"net/http"
	"strconv"

	// "os/exec"
	// "encoding/json"
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/auth"
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/logging"
	"fmt"

	"context"

	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	redislibsent "github.com/OT-CONTAINER-KIT/redis-operator/api/redissentinel/v1beta2"
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
func ListReplHandler(cli client.Client, logger *slog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data []RRInfo
		namespace := c.Param("ns")
		ev := logging.AuditEvent{
			Namespace: namespace,
			Action:    "List",
			Result:    "Fail",
			IP:        c.ClientIP(),
		}
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
		ev.Result = "Success"
		logging.AuditLog(c, ev, logger)
		c.JSON(http.StatusOK, data)
	}
}

var demoUser = Namespace{Namespace: "test2", Password: "level3cloud"}

func AuthHandler(rds *auth.Redis, clientset kubernetes.Interface, logger *slog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		var in struct {
			Namespace string `json:"namespace"`
			Password  string `json:"password"`
		}
		ev := logging.AuditEvent{
			Namespace: in.Namespace,
			Action:    "Authorization",
			Result:    "Fail",
			IP:        c.ClientIP(),
		}
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusBadRequest)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
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
func ConnectionHandler(cli client.Client, clientset kubernetes.Interface, logger *slog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		namespace := c.Param("ns")
		name := c.Param("name")
		ev := logging.AuditEvent{
			Namespace: namespace,
			Action:    "GetConnection",
			Result:    "Fail",
			IP:        c.ClientIP(),
		}
		repl, err := GetRepl(cli, c.Request.Context(), namespace, name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": fmt.Sprintf("%s not found Error: %s", name, err.Error()),
			})
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusBadRequest)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
			return
		}
		service_name := name + "-additional"
		service, err := clientset.CoreV1().Services(namespace).Get(c.Request.Context(), service_name, metav1.GetOptions{})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": fmt.Sprintf("service not found. Error: %s", err.Error()),
			})
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusBadRequest)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"connection": repl.Status.ConnectionInfo,
			"public_ip":  service.Status.LoadBalancer.Ingress[0].IP,
		})
		ev.Result = "Success"
		logging.AuditLog(c, ev, logger)
	}
}

func getEndpoint(clientset kubernetes.Interface) (endpoint string, err error) {
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
func CreateReplHandler(cli client.Client, logger *slog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		var repl v1beta2.RedisReplication
		var sent redislibsent.RedisSentinel
		var req ReplicationRequest
		ev := logging.AuditEvent{
			Namespace: "",
			Action:    "Create",
			Result:    "Fail",
			IP:        c.ClientIP(),
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": fmt.Sprintf("Not created. Error: %s", err.Error()),
			})
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusBadRequest)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
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
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusNotFound)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
		} else {
			setSentinel(&sent, req.Name, req.Namespace)
			err = cli.Create(c.Request.Context(), &sent)
			if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": fmt.Sprintf("Sentinel for %s not created. Error: %s", req.Name, err.Error()),
			})
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusNotFound)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
			} 
			c.JSON(http.StatusOK, gin.H{
				"status": fmt.Sprintf("Replication %s created. Cluster size: %d", req.Name, size),
			})
			ev.Result = "Success"
			logging.AuditLog(c, ev, logger)
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
func DeleteReplHandler(cli client.Client, logger *slog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req DeleteRequest
		var err error
		ev := logging.AuditEvent{
			Namespace: req.Namespace,
			Action:    "Delete",
			Result:    "Fail",
			IP:        c.ClientIP(),
		}
		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusBadRequest)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
			return
		}
		repl, err := GetRepl(cli, c.Request.Context(), req.Namespace, req.Name)
		if err != nil {
			jsonError(c, http.StatusNotFound, req.Name, err.Error())
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusNotFound)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
			return
		}
		err = cli.Delete(c.Request.Context(), repl)
		if err != nil {
			jsonError(c, http.StatusConflict, req.Name, err.Error())
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusConflict)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
			return
		}
		sname := req.Name + "-sentinel"
		sent, err := GetSent(cli, c.Request.Context(), req.Namespace, sname)
		if err != nil {
			jsonError(c, http.StatusNotFound, sname, err.Error())
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusNotFound)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
			return
		}
		err = cli.Delete(c.Request.Context(), sent)
		if err != nil {
			jsonError(c, http.StatusConflict, sname, err.Error())
			ev.ErrorType = logging.ErrorTypeFromStatus(http.StatusConflict)
			ev.ErrorMessage = err.Error()
			logging.AuditLog(c, ev, logger)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": fmt.Sprintf("%s deleted.", req.Name),
		})
		ev.Result = "Success"
		logging.AuditLog(c, ev, logger)
	}
}

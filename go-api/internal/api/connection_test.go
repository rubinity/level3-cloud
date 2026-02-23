package api_test

import (
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/api"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/runtime"
	k8fake "k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

// type Response struct {
// 	Message string `json:"message"`
// }

type Connection struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type ConnectionResponse struct {
	Connection Connection `json:"connection"`
	Public_ip string `json:"public_ip"`
}

type ErrResponse struct {
	Status string `json:"status"`
}

func SetService() * corev1.Service{
// Fake Service for testing
	fakeService := &corev1.Service{
    ObjectMeta: metav1.ObjectMeta{
        Name:      "realname-additional",   // match the name in your request
        Namespace: "realns",     // match the namespace in your request
    },
    Spec: corev1.ServiceSpec{
        Type: corev1.ServiceTypeLoadBalancer,
    },
    Status: corev1.ServiceStatus{
        LoadBalancer: corev1.LoadBalancerStatus{
            Ingress: []corev1.LoadBalancerIngress{
                {
                    IP: "1.2.3.4",  // fake public IP
                },
            },
        },
    },
	}
	return fakeService
}

func SetRepl() * v1beta2.RedisReplication{
	fakeRepl := &v1beta2.RedisReplication{
    ObjectMeta: metav1.ObjectMeta{
        Name:      "realname",   // match the name in your request
        Namespace: "realns",     // match the namespace in your request
    },
	}
	return fakeRepl
}

// func SetUndelRepl() * v1beta2.RedisReplication{
// 	fakeRepl := &v1beta2.RedisReplication{
//     ObjectMeta: metav1.ObjectMeta{
//         Name:      "realname",   // match the name in your request
//         Namespace: "realns",     // match the namespace in your request
//     },
// 	}
// 	fakeRepl.Dele
// 	return fakeRepl
// }

func InitClients(fakeService * corev1.Service, fakeRepl * v1beta2.RedisReplication)(cli client.Client, clientset kubernetes.Interface){
	scheme := runtime.NewScheme()
	v1beta2.AddToScheme(scheme)
	if fakeService != nil{
		clientset = k8fake.NewSimpleClientset(fakeService)
	} else {
		clientset = k8fake.NewSimpleClientset()
	}

	if fakeRepl != nil{
		cli = fake.NewClientBuilder().WithScheme(scheme).WithObjects(fakeRepl).Build()
	} else {
		cli = fake.NewClientBuilder().WithScheme(scheme).Build()
	}
	return cli, clientset
}

// {"connection":{"host":"repl4-master.test2.svc.cluster.local","port":6379},"public_ip":"188.34.110.196"}

// happy path
func TestConnection(t *testing.T) {
	fakeService := SetService()
	fakeRepl := SetRepl()
	cli, clientset := InitClients(fakeService, fakeRepl)

	router := api.TestRouter(cli, clientset)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/api/connection/realns/realname", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
	var response ConnectionResponse
	body := recorder.Body.String()
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}
	if response.Public_ip != "1.2.3.4" {
		t.Fatalf("bad ipo: %s", response.Public_ip)
	}
}

// not found because of the wrong namespace
func TestConnectionFakeNS(t *testing.T) {
	fakeService := SetService()
	fakeRepl := SetRepl()
	cli, clientset := InitClients(fakeService, fakeRepl)

	router := api.TestRouter(cli, clientset)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/api/connection/falsens/realname", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
	var response ErrResponse
	body := recorder.Body.String()
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}
	if !strings.Contains(response.Status, "realname not found") {
		t.Fatalf("bad status: %s", response.Status)
	}
}

// not found because doesn't exist
func TestConnectionFakename(t *testing.T) {
	fakeService := SetService()
	fakeRepl := SetRepl()
	cli, clientset := InitClients(fakeService, fakeRepl)

	router := api.TestRouter(cli, clientset)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/api/connection/realns/fakename", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
	var response ErrResponse
	body := recorder.Body.String()
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}
	if !strings.Contains(response.Status, "fakename not found") {
		t.Fatalf("bad status: %s", response.Status)
	}
}

//no associated service
func TestConnectionNoService(t *testing.T) {
	// fakeService := SetService()
	fakeRepl := SetRepl()
	cli, clientset := InitClients(nil, fakeRepl)

	router := api.TestRouter(cli, clientset)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/api/connection/realns/realname", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
	var response ErrResponse
	body := recorder.Body.String()
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}
	if !strings.Contains(response.Status, "service not found") {
		t.Fatalf("bad status: %s", response.Status)
	}
}
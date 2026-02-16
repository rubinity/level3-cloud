package api_test

import (
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/api"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	// "strings"
)

// {"connection":{"host":"repl4-master.test2.svc.cluster.local","port":6379},"public_ip":"188.34.110.196"}

func TestCreate(t *testing.T) {
	req := api.ReplicationRequest{
		Size:      3,
		Namespace: "realnamespace",
		Name:      "realname",
	}
	cli, clientset := InitClients(nil, nil)

	router := api.Router(cli, clientset)
	recorder := httptest.NewRecorder()

	reqbytes, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}

	request, err := http.NewRequest("POST", "/create", bytes.NewReader(reqbytes))
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
	var response ErrResponse
	body := recorder.Body.String()
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}
	if response.Status != "realname created." {
		t.Fatalf("bad status: %s", response.Status)
	}
}

type WrongRequest struct {
	Size      int    `json:"size"`
	Lamespace string `json:"lamespace"`
	Name      string `json:"name"`
}

func TestCreateRequest(t *testing.T) {
	req := WrongRequest{
		Size:      3,
		Lamespace: "realnamespace",
		Name:      "realname",
	}
	cli, clientset := InitClients(nil, nil)

	router := api.Router(cli, clientset)
	recorder := httptest.NewRecorder()

	reqbytes, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}

	request, err := http.NewRequest("POST", "/create", bytes.NewReader(reqbytes))
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
	repl, err := api.GetRepl(cli, t.Context(), req.Lamespace, req.Name)
	if err == nil {
		t.Fatalf("found in wrong namespace: %s", repl.Namespace)
	}
}

func TestCreateDuplicate(t *testing.T) {
	req := api.ReplicationRequest{
		Size:      3,
		Namespace: "realns",
		Name:      "realname",
	}
	fakeRepl := SetRepl()
	cli, clientset := InitClients(nil, fakeRepl)

	router := api.Router(cli, clientset)
	recorder := httptest.NewRecorder()

	reqbytes, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}

	request, err := http.NewRequest("POST", "/create", bytes.NewReader(reqbytes))
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("duplicate creation: %d", recorder.Code)
	}
	// var response ErrResponse
	// body := recorder.Body.String()
	// if err := json.Unmarshal([]byte(body), &response); err != nil {
	// 	t.Fatalf("parsing json response: %v", err)
	// }
	// if response.Status != "realname created." {
	// 	t.Fatalf("bad status: %s", response.Status)
	// }
}

// func TestCreateFakeNS(t *testing.T) {
// 	fakeService := SetService()
// 	fakeRepl := SetRepl()
// 	cli, clientset := InitClients(fakeService, fakeRepl)

// 	router := api.Router(cli, clientset)
// 	recorder := httptest.NewRecorder()
// 	request, err := http.NewRequest("GET", "/connection/falsens/realname", nil)
// 	if err != nil {
// 		t.Fatalf("building request: %v", err)
// 	}
// 	router.ServeHTTP(recorder, request)
// 	if recorder.Code != http.StatusNotFound {
// 		t.Fatalf("bad status code: %d", recorder.Code)
// 	}
// 	var response ErrResponse
// 	body := recorder.Body.String()
// 	if err := json.Unmarshal([]byte(body), &response); err != nil {
// 		t.Fatalf("parsing json response: %v", err)
// 	}
// 	if !strings.Contains(response.Status, "realname not found") {
// 		t.Fatalf("bad status: %s", response.Status)
// 	}
// }

// func TestConnectionFakename(t *testing.T) {
// 	fakeService := SetService()
// 	fakeRepl := SetRepl()
// 	cli, clientset := InitClients(fakeService, fakeRepl)

// 	router := api.Router(cli, clientset)
// 	recorder := httptest.NewRecorder()
// 	request, err := http.NewRequest("GET", "/connection/realns/fakename", nil)
// 	if err != nil {
// 		t.Fatalf("building request: %v", err)
// 	}
// 	router.ServeHTTP(recorder, request)
// 	if recorder.Code != http.StatusNotFound {
// 		t.Fatalf("bad status code: %d", recorder.Code)
// 	}
// 	var response ErrResponse
// 	body := recorder.Body.String()
// 	if err := json.Unmarshal([]byte(body), &response); err != nil {
// 		t.Fatalf("parsing json response: %v", err)
// 	}
// 	if !strings.Contains(response.Status, "fakename not found") {
// 		t.Fatalf("bad status: %s", response.Status)
// 	}
// }

// func TestConnectionNoService(t *testing.T) {
// 	// fakeService := SetService()
// 	fakeRepl := SetRepl()
// 	cli, clientset := InitClients(nil, fakeRepl)

// 	router := api.Router(cli, clientset)
// 	recorder := httptest.NewRecorder()
// 	request, err := http.NewRequest("GET", "/connection/realns/realname", nil)
// 	if err != nil {
// 		t.Fatalf("building request: %v", err)
// 	}
// 	router.ServeHTTP(recorder, request)
// 	if recorder.Code != http.StatusNotFound {
// 		t.Fatalf("bad status code: %d", recorder.Code)
// 	}
// 	var response ErrResponse
// 	body := recorder.Body.String()
// 	if err := json.Unmarshal([]byte(body), &response); err != nil {
// 		t.Fatalf("parsing json response: %v", err)
// 	}
// 	if !strings.Contains(response.Status, "service not found") {
// 		t.Fatalf("bad status: %s", response.Status)
// 	}
// }

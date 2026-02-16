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
// happy path: creating a replication
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
	if response.Status != "Replication realname created. Cluster size: 3" {
		t.Fatalf("bad status: %s", response.Status)
	}
}

type WrongRequest struct {
	Size      int    `json:"size"`
	Namespace string `json:"lamespace"`
	Name      string `json:"name"`
}

// testing a request with a wrong json field name
func TestCreateRequest(t *testing.T) {
	req := WrongRequest{
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
	repl, err := api.GetRepl(cli, t.Context(), req.Namespace, req.Name)
	if err == nil {
		t.Fatalf("found in wrong namespace: %s", repl.Namespace)
	}
	if repl.Namespace != "" {
		t.Fatalf("Namespace %s not empty", repl.Namespace)
	}
}

// happy path: testing size adjustment
func TestWrongSize(t *testing.T) {
	req := api.ReplicationRequest{
		Size:      25,
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
	repl, err := api.GetRepl(cli, t.Context(), req.Namespace, req.Name)
	if err != nil {
		t.Fatalf("not found: %s", repl.Name)
	}
	if *repl.Spec.Size != 3 {
		t.Fatalf("wrong size: %s", repl.Name)
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
}

type BadRequest struct {
	Size      string   `json:"size"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}


func TestBadRequest(t *testing.T) {
	req := BadRequest{
		Size:      "3",
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
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
}

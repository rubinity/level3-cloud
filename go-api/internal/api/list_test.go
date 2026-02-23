
package api_test

import (
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/api"
	// "bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	// "strings"


)

func TestList(t *testing.T) {

	fakeRepl := SetRepl()
	cli, clientset := InitClients(nil, fakeRepl)

	router := api.TestRouter(cli, clientset)
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/api/list/realns", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("bad status code: %d", recorder.Code)
	}

	var data []api.RRInfo
	body := recorder.Body.String()
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}
	
	if len(data) == 0{
		t.Fatalf("empty list")
	}

	if data[0].Name != "realname"{
		t.Fatalf("wrong name: %s", data[0].Name)
	}
}

func TestListSN(t *testing.T) {

	fakeRepl := SetRepl()
	cli, clientset := InitClients(nil, fakeRepl)

	router := api.TestRouter(cli, clientset)
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/api/list/wrongns", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("bad status code: %d", recorder.Code)
	}

	var data []api.RRInfo
	body := recorder.Body.String()
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}

	if len(data) != 0{
		t.Fatalf("list should be empty")
	}
}

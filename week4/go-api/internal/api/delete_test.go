package api_test

import (
	"Users/mariia.rubina13/Projects/cloud/week4/go-api/internal/api"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"context"
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/api/meta"
	
)

//happy path: testing deletion
func TestDelete(t *testing.T) {
	req := api.DeleteRequest{
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

	request, err := http.NewRequest("DELETE", "/delete", bytes.NewReader(reqbytes))
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
	if response.Status != "realname deleted." {
		t.Fatalf("bad status: %s", response.Status)
	}
}

//deleting a nonexisting replication
func TestDeleteSN(t *testing.T) {
	req := api.DeleteRequest{
		Namespace: "wrongns",
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

	request, err := http.NewRequest("DELETE", "/delete", bytes.NewReader(reqbytes))
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

	if !strings.Contains(response.Status, "realname not deleted") {
		t.Fatalf("bad status: %s", response.Status)
	}
}

type BadDeleteRequest struct {
	Name int  `json:"name"`
}

func TestBadDeleteRequest(t *testing.T) {

	cli, clientset := InitClients(nil, nil)

	router := api.Router(cli, clientset)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/create", strings.NewReader(`{"name": "test"`))
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
}


type NoDelClient struct {
	fakeClient client.Client // the real client you want to wrap
}


 func (w *NoDelClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
    // var err error
	err:= errors.New("testerror")
	return err
}

func (w *NoDelClient) Get(
    ctx context.Context,
    key client.ObjectKey,
    obj client.Object,
    opts ...client.GetOption,
) error {
    return w.fakeClient.Get(ctx, key, obj, opts...)
}

func (w *NoDelClient) List(
    ctx context.Context,
    list client.ObjectList,
    opts ...client.ListOption,
) error {
    return w.fakeClient.List(ctx, list, opts...)
}

func (w *NoDelClient) Create(
    ctx context.Context,
    obj client.Object,
    opts ...client.CreateOption,
) error {
    return w.fakeClient.Create(ctx, obj, opts...)
}

func (w *NoDelClient) Update(
    ctx context.Context,
    obj client.Object,
    opts ...client.UpdateOption,
) error {
    return w.fakeClient.Update(ctx, obj, opts...)
}

func (w *NoDelClient) Patch(
    ctx context.Context,
    obj client.Object,
    patch client.Patch,
    opts ...client.PatchOption,
) error {
    return w.fakeClient.Patch(ctx, obj, patch, opts...)
}

func (w *NoDelClient) DeleteAllOf(
    ctx context.Context,
    obj client.Object,
    opts ...client.DeleteAllOfOption,
) error {
    return w.fakeClient.DeleteAllOf(ctx, obj, opts...)
}

// StatusClient
func (w *NoDelClient) Status() client.StatusWriter {
    return w.fakeClient.Status()
}

// Other required methods
func (w *NoDelClient) Scheme() *runtime.Scheme {
    return w.fakeClient.Scheme()
}

func (w *NoDelClient) RESTMapper() meta.RESTMapper {
    return w.fakeClient.RESTMapper()
}

func (w *NoDelClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
    return w.fakeClient.GroupVersionKindFor(obj)
}

func (w *NoDelClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
    return w.fakeClient.IsObjectNamespaced(obj)
}

func (w *NoDelClient) SubResource(sub string) client.SubResourceClient {
    return w.fakeClient.SubResource(sub)
}


func TestDeleteFail(t *testing.T) {
	req := api.DeleteRequest{
		Namespace: "realns",
		Name:      "realname",
	}

	fakeRepl := SetRepl()
	// var cli NoDelClient

	client, clientset := InitClients(nil, fakeRepl)
	cli := &NoDelClient{fakeClient: client}
	// cli.fakeClient = client
	router := api.Router(cli, clientset)
	recorder := httptest.NewRecorder()

	reqbytes, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}

	request, err := http.NewRequest("DELETE", "/delete", bytes.NewReader(reqbytes))
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusConflict {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
}


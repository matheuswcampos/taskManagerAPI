package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"task-prioritization-api/app/api"
	"task-prioritization-api/app/repositories"
	"task-prioritization-api/app/services"
)

func newTaskAPI() *fiber.App {
	app := fiber.New()
	repo := repositories.NewTaskRepository()
	advisor := services.NewPriorityAdvisor()
	taskService := services.NewTaskService(repo, advisor)
	api.RegisterTaskRoutes(app, taskService)
	return app
}

func TestTaskRoutes_CreateReturns201(t *testing.T) {
	app := newTaskAPI()

	body := []byte(`{"title":"Nova tarefa","description":"Descricao"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestTaskRoutes_ListReturns200(t *testing.T) {
	app := newTaskAPI()

	createBody := []byte(`{"title":"Tarefa para listagem"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")

	createResp, err := app.Test(createReq)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected create status %d, got %d", http.StatusCreated, createResp.StatusCode)
	}

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("list request failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestTaskRoutes_DeleteReturns204(t *testing.T) {
	app := newTaskAPI()

	createBody := []byte(`{"title":"Tarefa para exclusao"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")

	createResp, err := app.Test(createReq)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected create status %d, got %d", http.StatusCreated, createResp.StatusCode)
	}

	var created struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(createResp.Body).Decode(&created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}
	if created.ID == "" {
		t.Fatal("expected created id")
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/tasks/"+created.ID, nil)
	deleteResp, err := app.Test(deleteReq)
	if err != nil {
		t.Fatalf("delete request failed: %v", err)
	}
	if deleteResp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, deleteResp.StatusCode)
	}
}

func TestTaskRoutes_GetByIDReturns404WhenNotFound(t *testing.T) {
	app := newTaskAPI()

	req := httptest.NewRequest(http.MethodGet, "/tasks/id-inexistente", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}

func TestTaskRoutes_GetByIDReturns200WhenFound(t *testing.T) {
	app := newTaskAPI()

	createBody := []byte(`{"title":"Tarefa para buscar por ID"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createResp, err := app.Test(createReq)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected create status %d, got %d", http.StatusCreated, createResp.StatusCode)
	}

	var created struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(createResp.Body).Decode(&created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/tasks/"+created.ID, nil)
	getResp, err := app.Test(getReq)
	if err != nil {
		t.Fatalf("get request failed: %v", err)
	}
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, getResp.StatusCode)
	}
}

func TestTaskRoutes_UpdateReturns200(t *testing.T) {
	app := newTaskAPI()

	createBody := []byte(`{"title":"Tarefa para atualizar"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createResp, err := app.Test(createReq)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected create status %d, got %d", http.StatusCreated, createResp.StatusCode)
	}

	var created struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(createResp.Body).Decode(&created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}

	updateBody := []byte(`{"title":"Titulo atualizado","status":"doing"}`)
	updateReq := httptest.NewRequest(http.MethodPut, "/tasks/"+created.ID, bytes.NewReader(updateBody))
	updateReq.Header.Set("Content-Type", "application/json")
	updateResp, err := app.Test(updateReq)
	if err != nil {
		t.Fatalf("update request failed: %v", err)
	}
	if updateResp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, updateResp.StatusCode)
	}
}

func TestTaskRoutes_CreateReturns400OnInvalidJSON(t *testing.T) {
	app := newTaskAPI()

	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader([]byte(`{"title":`)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestTaskRoutes_UpdateReturns400OnInvalidJSON(t *testing.T) {
	app := newTaskAPI()

	createBody := []byte(`{"title":"Tarefa base"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createResp, err := app.Test(createReq)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected create status %d, got %d", http.StatusCreated, createResp.StatusCode)
	}

	var created struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(createResp.Body).Decode(&created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/tasks/"+created.ID, bytes.NewReader([]byte(`{"status":`)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestTaskRoutes_DeleteReturns404WhenNotFound(t *testing.T) {
	app := newTaskAPI()

	req := httptest.NewRequest(http.MethodDelete, "/tasks/id-inexistente", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}

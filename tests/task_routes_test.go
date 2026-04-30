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


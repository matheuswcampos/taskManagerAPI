package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewApp_HealthAndTaskRoutes(t *testing.T) {
	app := newApp()

	healthReq := httptest.NewRequest(http.MethodGet, "/health", nil)
	healthResp, err := app.Test(healthReq)
	if err != nil {
		t.Fatalf("health request failed: %v", err)
	}
	if healthResp.StatusCode != http.StatusOK {
		t.Fatalf("expected /health status %d, got %d", http.StatusOK, healthResp.StatusCode)
	}

	createReq := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		bytes.NewReader([]byte(`{"title":"task from bootstrap test"}`)),
	)
	createReq.Header.Set("Content-Type", "application/json")
	createResp, err := app.Test(createReq)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected /tasks create status %d, got %d", http.StatusCreated, createResp.StatusCode)
	}
}

package tests

import (
	"testing"
	"time"

	"task-prioritization-api/app/models"
	"task-prioritization-api/app/repositories"
)

func TestTaskRepository_ListSortsByCreatedAtThenID(t *testing.T) {
	repo := repositories.NewTaskRepository()
	base := time.Now().UTC()

	repo.Create(models.Task{ID: "b", Title: "B", Status: models.StatusTodo, Priority: models.PriorityLow, CreatedAt: base})
	repo.Create(models.Task{ID: "a", Title: "A", Status: models.StatusTodo, Priority: models.PriorityLow, CreatedAt: base})
	repo.Create(models.Task{ID: "c", Title: "C", Status: models.StatusTodo, Priority: models.PriorityLow, CreatedAt: base.Add(time.Second)})

	list := repo.List()
	if len(list) != 3 {
		t.Fatalf("expected 3 tasks, got %d", len(list))
	}
	if list[0].ID != "a" || list[1].ID != "b" || list[2].ID != "c" {
		t.Fatalf("unexpected order: %s, %s, %s", list[0].ID, list[1].ID, list[2].ID)
	}
}

func TestTaskRepository_UpdatePreservesCreatedAt(t *testing.T) {
	repo := repositories.NewTaskRepository()
	createdAt := time.Now().UTC().Add(-time.Hour)

	created := repo.Create(models.Task{
		ID:        "task-1",
		Title:     "Original",
		Status:    models.StatusTodo,
		Priority:  models.PriorityLow,
		CreatedAt: createdAt,
	})
	time.Sleep(2 * time.Millisecond)

	updated, err := repo.Update(models.Task{
		ID:       created.ID,
		Title:    "Updated",
		Status:   models.StatusDoing,
		Priority: models.PriorityHigh,
	})
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}
	if !updated.CreatedAt.Equal(created.CreatedAt) {
		t.Fatalf("expected created_at to be preserved")
	}
	if !updated.UpdatedAt.After(created.UpdatedAt) {
		t.Fatalf("expected updated_at to be after previous updated_at")
	}
}

package tests

import (
	"errors"
	"testing"

	"task-prioritization-api/app/models"
	"task-prioritization-api/app/repositories"
	"task-prioritization-api/app/services"
)

type fakeAdvisor struct {
	priority models.TaskPriority
	err      error
}

func (f *fakeAdvisor) SuggestPriority(_, _ string) (models.TaskPriority, error) {
	return f.priority, f.err
}

func newTaskServiceFixture() *services.TaskService {
	repo := repositories.NewTaskRepository()
	advisor := &fakeAdvisor{priority: models.PriorityHigh}
	return services.NewTaskService(repo, advisor)
}

func TestTaskService_Create(t *testing.T) {
	svc := newTaskServiceFixture()

	created, err := svc.Create(models.TaskCreate{
		Title:       "Corrigir bug de login",
		Description: "Impacta clientes em producao",
	})
	if err != nil {
		t.Fatalf("expected no error on create, got %v", err)
	}
	if created.ID == "" {
		t.Fatal("expected generated ID")
	}
	if created.Title != "Corrigir bug de login" {
		t.Fatalf("expected title to be preserved, got %q", created.Title)
	}
	if created.Status != models.StatusTodo {
		t.Fatalf("expected default status %q, got %q", models.StatusTodo, created.Status)
	}
	if created.Priority != models.PriorityHigh {
		t.Fatalf("expected advisor priority %q, got %q", models.PriorityHigh, created.Priority)
	}
	if created.CreatedAt.IsZero() || created.UpdatedAt.IsZero() {
		t.Fatal("expected timestamps to be set")
	}
}

func TestTaskService_List(t *testing.T) {
	svc := newTaskServiceFixture()

	first, err := svc.Create(models.TaskCreate{Title: "Primeira tarefa"})
	if err != nil {
		t.Fatalf("create first task: %v", err)
	}
	second, err := svc.Create(models.TaskCreate{Title: "Segunda tarefa"})
	if err != nil {
		t.Fatalf("create second task: %v", err)
	}

	tasks := svc.List()
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	ids := map[string]bool{
		tasks[0].ID: true,
		tasks[1].ID: true,
	}
	if !ids[first.ID] || !ids[second.ID] {
		t.Fatalf("expected list to contain IDs %q and %q, got %q and %q", first.ID, second.ID, tasks[0].ID, tasks[1].ID)
	}
}

func TestTaskService_Update(t *testing.T) {
	svc := newTaskServiceFixture()

	created, err := svc.Create(models.TaskCreate{
		Title:       "Tarefa original",
		Description: "Descricao original",
		Priority:    models.PriorityLow,
	})
	if err != nil {
		t.Fatalf("create task: %v", err)
	}

	newTitle := "Tarefa atualizada"
	newDescription := "Descricao atualizada"
	newStatus := models.StatusDoing

	updated, err := svc.Update(created.ID, models.TaskUpdate{
		Title:       &newTitle,
		Description: &newDescription,
		Status:      &newStatus,
	})
	if err != nil {
		t.Fatalf("expected no error on update, got %v", err)
	}
	if updated.Title != newTitle {
		t.Fatalf("expected title %q, got %q", newTitle, updated.Title)
	}
	if updated.Description != newDescription {
		t.Fatalf("expected description %q, got %q", newDescription, updated.Description)
	}
	if updated.Status != newStatus {
		t.Fatalf("expected status %q, got %q", newStatus, updated.Status)
	}
	if updated.Priority != models.PriorityHigh {
		t.Fatalf("expected recomputed priority %q, got %q", models.PriorityHigh, updated.Priority)
	}
	if updated.UpdatedAt.IsZero() {
		t.Fatal("expected updated_at to be set")
	}
}

func TestTaskService_Delete(t *testing.T) {
	svc := newTaskServiceFixture()

	created, err := svc.Create(models.TaskCreate{Title: "Tarefa para excluir"})
	if err != nil {
		t.Fatalf("create task: %v", err)
	}

	if err := svc.Delete(created.ID); err != nil {
		t.Fatalf("expected no error on delete, got %v", err)
	}

	_, err = svc.GetByID(created.ID)
	if !errors.Is(err, repositories.ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound after delete, got %v", err)
	}
}

func TestTaskService_ReturnsNotFoundForUnknownID(t *testing.T) {
	svc := newTaskServiceFixture()

	unknownID := "id-inexistente"

	t.Run("get by id", func(t *testing.T) {
		_, err := svc.GetByID(unknownID)
		if !errors.Is(err, repositories.ErrTaskNotFound) {
			t.Fatalf("expected ErrTaskNotFound, got %v", err)
		}
	})

	t.Run("update", func(t *testing.T) {
		title := "novo titulo"
		_, err := svc.Update(unknownID, models.TaskUpdate{Title: &title})
		if !errors.Is(err, repositories.ErrTaskNotFound) {
			t.Fatalf("expected ErrTaskNotFound, got %v", err)
		}
	})

	t.Run("delete", func(t *testing.T) {
		err := svc.Delete(unknownID)
		if !errors.Is(err, repositories.ErrTaskNotFound) {
			t.Fatalf("expected ErrTaskNotFound, got %v", err)
		}
	})
}

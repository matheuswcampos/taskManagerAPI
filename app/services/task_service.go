package services

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"time"

	"task-prioritization-api/app/models"
)

// TaskRepository defines persistence operations required by TaskService.
type TaskRepository interface {
	Create(task models.Task) models.Task
	List() []models.Task
	GetByID(id string) (models.Task, error)
	Update(task models.Task) (models.Task, error)
	Delete(id string) error
}

// PriorityAdvisor defines priority suggestion behavior.
type PriorityAdvisor interface {
	SuggestPriority(title, description string) (models.TaskPriority, error)
}

// TaskService centralizes task business rules.
type TaskService struct {
	repo    TaskRepository
	advisor PriorityAdvisor
}

// NewTaskService builds a new TaskService.
func NewTaskService(repo TaskRepository, advisor PriorityAdvisor) *TaskService {
	if advisor == nil {
		advisor = NewPriorityAdvisor()
	}

	return &TaskService{
		repo:    repo,
		advisor: advisor,
	}
}

// Create creates a new task and applies default business rules.
func (s *TaskService) Create(input models.TaskCreate) (models.Task, error) {
	now := time.Now().UTC()
	task := models.Task{
		ID:          newTaskID(),
		Title:       strings.TrimSpace(input.Title),
		Description: strings.TrimSpace(input.Description),
		Status:      input.Status,
		Priority:    input.Priority,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if task.Status == "" {
		task.Status = models.StatusTodo
	}

	if task.Priority == "" {
		suggested, err := s.advisor.SuggestPriority(task.Title, task.Description)
		if err == nil && suggested != "" {
			task.Priority = suggested
		}
	}

	created := s.repo.Create(task)
	return created, nil
}

// List returns all tasks.
func (s *TaskService) List() []models.Task {
	return s.repo.List()
}

// GetByID returns a task by ID.
func (s *TaskService) GetByID(id string) (models.Task, error) {
	return s.repo.GetByID(id)
}

// Update updates a task by ID, applying partial input changes.
func (s *TaskService) Update(id string, input models.TaskUpdate) (models.Task, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return models.Task{}, err
	}

	changedText := false

	if input.Title != nil {
		task.Title = strings.TrimSpace(*input.Title)
		changedText = true
	}
	if input.Description != nil {
		task.Description = strings.TrimSpace(*input.Description)
		changedText = true
	}
	if input.Status != nil {
		task.Status = *input.Status
	}
	if input.Priority != nil {
		task.Priority = *input.Priority
	} else if changedText {
		suggested, err := s.advisor.SuggestPriority(task.Title, task.Description)
		if err == nil && suggested != "" {
			task.Priority = suggested
		}
	}

	updated, err := s.repo.Update(task)
	if err != nil {
		return models.Task{}, err
	}

	return updated, nil
}

// Delete removes a task by ID.
func (s *TaskService) Delete(id string) error {
	return s.repo.Delete(id)
}

func newTaskID() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return time.Now().UTC().Format("20060102150405.000000000")
	}
	return hex.EncodeToString(buf)
}

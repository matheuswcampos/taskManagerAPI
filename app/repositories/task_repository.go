package repositories

import (
	"errors"
	"sort"
	"sync"
	"time"

	"task-prioritization-api/app/models"
)

var (
	// ErrTaskNotFound is returned when the task does not exist.
	ErrTaskNotFound = errors.New("task not found")
)

// TaskRepository stores tasks in memory.
type TaskRepository struct {
	mu    sync.RWMutex
	tasks map[string]models.Task
}

// NewTaskRepository creates a new in-memory task repository.
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]models.Task),
	}
}

// Create inserts a new task into the repository.
func (r *TaskRepository) Create(task models.Task) models.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now().UTC()
	if task.CreatedAt.IsZero() {
		task.CreatedAt = now
	}
	task.UpdatedAt = now

	r.tasks[task.ID] = task
	return task
}

// List returns all tasks sorted by creation time and then by ID.
func (r *TaskRepository) List() []models.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		out = append(out, task)
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].CreatedAt.Equal(out[j].CreatedAt) {
			return out[i].ID < out[j].ID
		}
		return out[i].CreatedAt.Before(out[j].CreatedAt)
	})

	return out
}

// GetByID retrieves a task by its ID.
func (r *TaskRepository) GetByID(id string) (models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return models.Task{}, ErrTaskNotFound
	}

	return task, nil
}

// Update replaces an existing task by ID.
func (r *TaskRepository) Update(task models.Task) (models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	current, ok := r.tasks[task.ID]
	if !ok {
		return models.Task{}, ErrTaskNotFound
	}

	if task.CreatedAt.IsZero() {
		task.CreatedAt = current.CreatedAt
	}
	task.UpdatedAt = time.Now().UTC()

	r.tasks[task.ID] = task
	return task, nil
}

// Delete removes a task by ID.
func (r *TaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
}

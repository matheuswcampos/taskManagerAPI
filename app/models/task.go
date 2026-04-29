package models

import "time"

// TaskStatus represents the current status of a task.
type TaskStatus string

const (
	StatusTodo     TaskStatus = "todo"
	StatusDoing    TaskStatus = "doing"
	StatusDone     TaskStatus = "done"
	StatusCanceled TaskStatus = "canceled"
)

// TaskPriority represents the priority level of a task.
type TaskPriority string

const (
	PriorityLow    TaskPriority = "low"
	PriorityMedium TaskPriority = "medium"
	PriorityHigh   TaskPriority = "high"
	PriorityCritic TaskPriority = "critic"
)

// Task represents the persisted task entity.
type Task struct {
	ID          string     `json:"id" validate:"required"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status" validate:"required,oneof=todo doing done"`
	Priority    TaskPriority `json:"priority" validate:"required,oneof=low medium high critic"`
	CreatedAt   time.Time  `json:"created_at" validate:"required"`
	UpdatedAt   time.Time  `json:"updated_at" validate:"required"`
}

// TaskCreate is the payload for creating a task.
type TaskCreate struct {
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status,omitempty" validate:"omitempty,oneof=todo doing done"`
	Priority    TaskPriority `json:"priority,omitempty" validate:"omitempty,oneof=low medium high critic"`
}

// TaskUpdate is the payload for partial task updates.
type TaskUpdate struct {
	Title       *string     `json:"title,omitempty" validate:"omitempty"`
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty" validate:"omitempty,oneof=todo doing done"`
	Priority    *TaskPriority `json:"priority,omitempty" validate:"omitempty,oneof=low medium high critic"`
}

// TaskOut is the HTTP response model for tasks.
type TaskOut struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status"`
	Priority    TaskPriority `json:"priority"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// ToOut converts a Task entity to the API output model.
func (t Task) ToOut() TaskOut {
	return TaskOut{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Priority:    t.Priority,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

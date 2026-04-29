package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"task-prioritization-api/app/models"
	"task-prioritization-api/app/repositories"
	"task-prioritization-api/app/services"
)

// RegisterTaskRoutes registers HTTP routes for task CRUD operations.
func RegisterTaskRoutes(router fiber.Router, taskService *services.TaskService) {
	router.Post("/tasks", createTask(taskService))
	router.Get("/tasks", listTasks(taskService))
	router.Get("/tasks/:id", getTaskByID(taskService))
	router.Put("/tasks/:id", updateTask(taskService))
	router.Delete("/tasks/:id", deleteTask(taskService))
}

func createTask(taskService *services.TaskService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.TaskCreate
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		created, err := taskService.Create(input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create task"})
		}

		return c.Status(fiber.StatusCreated).JSON(created.ToOut())
	}
}

func listTasks(taskService *services.TaskService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks := taskService.List()
		out := make([]models.TaskOut, 0, len(tasks))
		for _, task := range tasks {
			out = append(out, task.ToOut())
		}

		return c.Status(fiber.StatusOK).JSON(out)
	}
}

func getTaskByID(taskService *services.TaskService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		task, err := taskService.GetByID(id)
		if err != nil {
			if errors.Is(err, repositories.ErrTaskNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get task"})
		}

		return c.Status(fiber.StatusOK).JSON(task.ToOut())
	}
}

func updateTask(taskService *services.TaskService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var input models.TaskUpdate
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		updated, err := taskService.Update(id, input)
		if err != nil {
			if errors.Is(err, repositories.ErrTaskNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update task"})
		}

		return c.Status(fiber.StatusOK).JSON(updated.ToOut())
	}
}

func deleteTask(taskService *services.TaskService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		err := taskService.Delete(id)
		if err != nil {
			if errors.Is(err, repositories.ErrTaskNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete task"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

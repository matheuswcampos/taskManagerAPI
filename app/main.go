package main

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"task-prioritization-api/app/api"
	"task-prioritization-api/app/repositories"
	"task-prioritization-api/app/services"
)

const (
	statusOK = "ok"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func main() {
	app := newApp()
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

func newApp() *fiber.App {
	app := fiber.New()
	repo := repositories.NewTaskRepository()
	advisor := services.NewPriorityAdvisor()
	taskService := services.NewTaskService(repo, advisor)

	app.Get("/health", func(c *fiber.Ctx) error {
		response := HealthResponse{
			Status:    statusOK,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}

		return c.Status(fiber.StatusOK).JSON(response)
	})
	api.RegisterTaskRoutes(app, taskService)

	return app
}

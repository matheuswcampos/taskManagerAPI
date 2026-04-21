package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	statusOK = "ok"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		response := HealthResponse{
			Status:    statusOK,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}

		return c.Status(fiber.StatusOK).JSON(response)
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

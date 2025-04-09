package tasks

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {
	tasks := r.Group("/tasks")

	tasks.Post("/", createTask)
	tasks.Get("/", getTasks)
	tasks.Get("/:id", getByID)
	tasks.Put("/:id", UpdateTask)
	tasks.Delete("/:id", deleteByID)
}

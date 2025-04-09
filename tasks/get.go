package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lleoserrano/fiber_mongo/db"
	"net/http"
)

func getTasks(c *fiber.Ctx) error {

	var tasks []Task

	err := db.FindAll("tasks", nil, &tasks)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(tasks)
}

func getByID(c *fiber.Ctx) error {
	var task Task

	err := db.FindByID("tasks", c.Params("id"), &task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(task)
}

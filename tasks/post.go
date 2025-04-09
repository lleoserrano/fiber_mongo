package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lleoserrano/fiber_mongo/db"
	"github.com/lleoserrano/fiber_mongo/tags"
	"net/http"
)

func createTask(c *fiber.Ctx) error {
	body := new(Task)

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	id, err := db.Insert("tasks", body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	body.ID = id

	err = tags.AddTask(body.ID.Hex(), body.Tags)
	if err != nil {
		db.DeleteByID("tasks", body.ID.Hex())
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(body)
}

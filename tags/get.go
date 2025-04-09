package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lleoserrano/fiber_mongo/db"
)

func getAll(c *fiber.Ctx) error {
	var tags []Tag

	err := db.FindAll("tags", nil, &tags)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(tags)
}

func getByID(c *fiber.Ctx) error {
	var tag Tag

	err := db.FindByID("tags", c.Params("id"), &tag)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())

	}
	return c.Status(http.StatusOK).JSON(tag)
}

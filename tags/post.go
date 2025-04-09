package tags

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lleoserrano/fiber_mongo/db"
	"net/http"
)

func addTag(c *fiber.Ctx) error {
	body := new(Tag)

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	id, err := db.Insert("tags", body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	body.ID = id
	return c.Status(http.StatusCreated).JSON(body)
}

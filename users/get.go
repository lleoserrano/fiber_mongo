package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lleoserrano/fiber_mongo/db"
)

func getAll(c *fiber.Ctx) error {
	var documents []User

	err := db.FindAll("users", nil, &documents)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(documents)
}

func getByID(c *fiber.Ctx) error {
	var document User

	err := db.FindByID("users", c.Params("id"), &document)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(document)
}

package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lleoserrano/fiber_mongo/db"
	"net/http"
)

func deleteUser(c *fiber.Ctx) error {

	err := db.DeleteByID("users", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusNoContent).JSON("")
}

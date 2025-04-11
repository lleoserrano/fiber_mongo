package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lleoserrano/fiber_mongo/db"
	"github.com/lleoserrano/fiber_mongo/tags"
	"github.com/lleoserrano/fiber_mongo/tasks"
	"github.com/lleoserrano/fiber_mongo/users"
)

func main() {
	app := fiber.New()
	v1 := app.Group("/v1")
	defer db.CloseClient()

	users.SetRoutes(v1)
	tasks.SetRoutes(v1)
	tags.SetRoutes(v1)

	err := app.Listen(":9001")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

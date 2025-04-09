package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lleoserrano/fiber_mongo/db"
	"github.com/lleoserrano/fiber_mongo/tags"
	"net/http"
)

func UpdateTask(c *fiber.Ctx) error {
	body := new(Task)
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid JSON")
	}

	var prev Task
	err := db.FindByID("tasks", c.Params("id"), &prev)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	var result Task
	err = db.UpdateByID("tasks", c.Params("id"), body, &result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = updateTagsTask(c.Params("id"), prev.Tags, result.Tags)
	if err != nil {
		err = tags.RemoveTask(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		err = db.DeleteByID("tasks", c.Params("id"))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		_, err = db.Insert("tasks", &result)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		err = tags.AddTask(result.ID.Hex(), result.Tags)
		if err != nil {
			db.DeleteByID("tasks", c.Params("id"))
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
	}

	return c.JSON(result)
}

func updateTagsTask(id string, oldTags []string, newTags []string) error {
	mapOldTags := make(map[string]int, len(oldTags))

	for k, v := range oldTags {
		mapOldTags[v] = k
	}

	var diff []string

	for _, v := range newTags {
		if _, key := mapOldTags[v]; !key {
			diff = append(diff, v)
		} else {
			delete(mapOldTags, v)
		}
	}

	if len(diff) > 0 {
		err := tags.AddTask(id, diff)
		if err != nil {
			return err
		}
	}

	if len(mapOldTags) > 0 {
		dt := make([]string, 0, len(mapOldTags))
		for k := range mapOldTags {
			dt = append(dt, k)
		}

		err := tags.RemoveTask(id, dt...)
		if err != nil {
			return err
		}
	}
	return nil
}

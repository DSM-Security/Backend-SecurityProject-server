package post

import (
	"github.com/Backend-SecurityProject-server/server/db"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "POST POST",
	})
}

func GetPost(c *fiber.Ctx) error {
	row, err := db.GetDB().Query("SELECT * FROM post")
	defer row.Close()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(row)
}

func DeletePost(c *fiber.Ctx) error {
	pid := c.Params("pid")
	if pid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to delete",
	})
}

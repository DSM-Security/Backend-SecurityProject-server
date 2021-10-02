package post

import "github.com/gofiber/fiber/v2"

func CreatePost(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "POST POST",
	})
}

func GetPost(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GET POST",
	})
}

func DeletePost(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DELETE POST",
	})
}

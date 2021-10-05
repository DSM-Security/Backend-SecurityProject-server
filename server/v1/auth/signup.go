package auth

import (
	"github.com/Backend-SecurityProject-server/server/db"
	"github.com/gofiber/fiber/v2"
)

type signupRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

func Signup(c *fiber.Ctx) error {
	var signupReq signupRequest
	err := c.BodyParser(&signupReq)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "BadRequest",
		})
	}

	_, err = db.GetDB().Query("INSERT INTO user (id, password, nickname) VALUES (?, ?, ?)", signupReq.Id, signupReq.Password, signupReq.Nickname)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User ID is already exist",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Success to create",
	})
}

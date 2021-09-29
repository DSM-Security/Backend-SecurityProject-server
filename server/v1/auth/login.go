package auth

import (
	"github.com/Backend-SecurityProject-server/server/db"
	"github.com/gofiber/fiber/v2"
)

type loginRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type user struct {
	Id       string `json:"id"`
	Password string `password:"password"`
	Nickname string `json:"nickname"`
}

func Login(c *fiber.Ctx) error {
	var loginReq loginRequest
	var userResult user
	err := c.BodyParser(&loginReq)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "BadRequest",
		})
	}

	row, err := db.GetDB().Query("SELECT * FROM user WHERE id = ?", loginReq.Id)
	defer row.Close()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "NotFound",
		})
	}

	for row.Next() {
		row.Scan(&userResult.Id, &userResult.Password, &userResult.Nickname)
	}

	if userResult.Password != loginReq.Password {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	return c.Status(200).JSON(userResult)
}

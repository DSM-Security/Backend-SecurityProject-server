package auth

import (
	"github.com/Backend-SecurityProject-server/server/db"
	"github.com/Backend-SecurityProject-server/utils"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type User struct {
	Id       string `json:"id"`
	Password string `password:"password"`
	Nickname string `json:"nickname"`
}

func Login(c *fiber.Ctx) error {
	var loginReq LoginRequest
	var userResult User
	err := c.BodyParser(&loginReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BadRequest",
		})
	}

	row, err := db.GetDB().Query("SELECT id FROM user WHERE id = ?", loginReq.Id)
	defer row.Close()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "NotFound",
		})
	}

	for row.Next() {
		row.Scan(&userResult.Id, &userResult.Password, &userResult.Nickname)
	}

	if userResult.Password != loginReq.Password {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	payload := utils.JwtPayload{
		Id:       userResult.Id,
		Nickname: userResult.Nickname,
	}

	jwt := utils.AccessToken(payload)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken": jwt,
	})
}

package middleware

import (
	"github.com/Backend-SecurityProject-server/server/db"
	"github.com/Backend-SecurityProject-server/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {

	jwt, err := utils.GetTokenString(c)
	utils.HandleErr(err)

	_, user, err := utils.ValidateToken(string(jwt))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "expired token",
		})
	}

	row, err := db.GetDB().Query("SELECT * FROM user WHERE id = ?", user.Id)
	defer row.Close()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err != nil {
		return c.SendStatus(401)
	}
	return c.Next()
}

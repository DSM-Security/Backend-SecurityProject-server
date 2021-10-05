package post

import (
	"fmt"

	"github.com/Backend-SecurityProject-server/server/db"
	"github.com/Backend-SecurityProject-server/utils"
	"github.com/gofiber/fiber/v2"
)

type post struct {
	Pid       int    `json:"pid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Writer    string `json:"writer"`
	CreatedAt string `json:"createdAt"`
}

type createReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreatePost(c *fiber.Ctx) error {
	var request createReq
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "BadRequest",
		})
	}

	jwt, _ := utils.GetTokenString(c)
	_, user, _ := utils.ValidateToken(string(jwt))

	_, err = db.GetDB().Query("INSERT INTO post (title, content, writer) VALUES (?, ?, ?)", request.Title, request.Content, user.Id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Failed to save",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success to create post",
	})
}

func GetPost(c *fiber.Ctx) error {
	var postResult post
	row := db.GetDB().QueryRow("SELECT * FROM post")

	err := row.Scan(&postResult)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(row)
}

func DeletePost(c *fiber.Ctx) error {
	pid := c.Params("pid")
	_, err := db.GetDB().Query("SELECT * FROM post WHERE pid = ?", pid)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	_, err = db.GetDB().Query("DELETE FROM post WHERE pid = ?", pid)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to delete",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to delete",
	})
}

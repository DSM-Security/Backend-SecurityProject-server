package post

import (
	"fmt"

	"github.com/Backend-SecurityProject-server/server/db"
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "POST POST",
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

	fmt.Println(pid)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to delete",
	})
}

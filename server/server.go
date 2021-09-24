package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start(port int) {
	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("api")
	v1R := api.Group("v1")

	v1auth := v1R.Group("auth")
	v1auth.Post("/login")
	v1auth.Post("/signup")

	v1post := v1R.Group("post")
	v1post.Get("/get")
	v1post.Post("/post")
	v1post.Delete("/delete")

	log.Fatal(app.Listen(":" + fmt.Sprint(port)))
}

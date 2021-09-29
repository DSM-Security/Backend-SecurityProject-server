package server

import (
	"fmt"
	"log"

	"github.com/Backend-SecurityProject-server/server/db"
	"github.com/Backend-SecurityProject-server/server/v1/auth"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start(port int) {
	app := fiber.New()

	db.Start()
	defer db.CloseDB()

	app.Use(cors.New())

	api := app.Group("api")
	v1R := api.Group("v1")

	v1auth := v1R.Group("auth")
	v1auth.Post("/login", auth.Login)
	v1auth.Post("/signup", auth.Signup)

	v1post := v1R.Group("post")
	v1post.Get("/get", auth.Login)
	v1post.Post("/post", auth.Login)
	v1post.Delete("/delete", auth.Login)

	log.Fatal(app.Listen(":" + fmt.Sprint(port)))
}

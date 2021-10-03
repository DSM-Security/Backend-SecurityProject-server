package server

import (
	"fmt"
	"log"

	"github.com/Backend-SecurityProject-server/server/db"
	"github.com/Backend-SecurityProject-server/server/middleware"
	"github.com/Backend-SecurityProject-server/server/v1/auth"
	"github.com/Backend-SecurityProject-server/server/v1/post"
	"github.com/Backend-SecurityProject-server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start(port int) {
	app := fiber.New()

	db.Start()
	defer db.CloseDB()

	app.Use(cors.New())
	app.Use(logger.New(utils.ConsoleLogger()))

	api := app.Group("api")
	v1R := api.Group("v1")

	v1auth := v1R.Group("auth")
	v1auth.Post("/login", auth.Login)
	v1auth.Post("/signup", auth.Signup)

	v1post := v1R.Group("post")
	v1post.Use(middleware.AuthMiddleware)
	v1post.Get("/get", post.GetPost)
	v1post.Post("/post", post.CreatePost)
	v1post.Delete("/delete/:pid", post.DeletePost)

	log.Fatal(app.Listen(":" + fmt.Sprint(port)))
}

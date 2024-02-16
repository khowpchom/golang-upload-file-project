package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/khowpchom/golang-upload-file-project/configs"
	"github.com/khowpchom/golang-upload-file-project/middleware"
	"github.com/khowpchom/golang-upload-file-project/routes"
)

func main() {
	configs.LoadConfig()
	configs.ConnectMongoDBClient(configs.AppConfig.MongoURI, configs.AppConfig.DBName)
	defer configs.Client.Close()
	configs.ConnectMailer(
		configs.AppConfig.MailerHost,
		configs.AppConfig.MailerUsername,
		configs.AppConfig.MailerPassword,
	)
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, SECRET-KEY",
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(middleware.SecureMiddleware)
	routes.FileRoute(app)
	routes.MailRoute(app)

	app.Listen(":" + configs.AppConfig.Port)

}

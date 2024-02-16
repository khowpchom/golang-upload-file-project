package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khowpchom/golang-upload-file-project/controllers"
)

func MailRoute(app *fiber.App) {
	app.Post("/mail/send/:id", controllers.SendFileToEmail)
}


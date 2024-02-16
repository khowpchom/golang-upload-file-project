package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khowpchom/golang-upload-file-project/controllers"
)

func FileRoute(app *fiber.App) {
	app.Post("/file", controllers.UploadFile)
	app.Get("/file", controllers.ListFile)
	app.Get("/file/:id", controllers.DownloadFile)
}

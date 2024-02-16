package controllers

import (
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gomail.v2"

	"github.com/khowpchom/golang-upload-file-project/configs"
	"github.com/khowpchom/golang-upload-file-project/models"
	"github.com/khowpchom/golang-upload-file-project/utils"
)

func SendFileToEmail(c *fiber.Ctx) error {
	email := new(models.Email)
	if err := c.BodyParser(email) ; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if isEmailValid := email.IsEmailValid() ; !isEmailValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is wrong pattern",
		})
	}
	idParam := c.Params("id")
	fileID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	query := bson.M{"_id": fileID}
	var fileCol *mongo.Collection = configs.Client.GetCollection("file")
	var file models.File
	err = fileCol.FindOne(c.Context(), query).Decode(&file)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Set(fiber.HeaderContentType, "application/octet-stream")
	c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", file.Name))

	m := utils.Mailer{}
	message := gomail.NewMessage()
	message.SetHeader("To", email.Email)
	message.SetHeader("Subject", "File attachment")
	message.SetBody("text/plain", "Please find the attached file.")
	message.Attach(file.Name, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(file.Content)
		return err
	}))

	m.Send(message)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"info": "Send File Successfully",
	})
}
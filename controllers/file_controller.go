package controllers

import (
	"context"
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/khowpchom/golang-upload-file-project/configs"
	"github.com/khowpchom/golang-upload-file-project/models"
	"github.com/khowpchom/golang-upload-file-project/responses"
)


func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fileContents, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer fileContents.Close()
	contents, err := io.ReadAll(fileContents)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var fileCol *mongo.Collection = configs.Client.GetCollection("file")
	err = saveFileMongo(fileCol, file.Filename, contents)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"info": "File uploaded successfully",
	})
}

func ListFile(c *fiber.Ctx) error {
	query := bson.D{{}}
	var fileCol *mongo.Collection = configs.Client.GetCollection("file")
	cursor, err := fileCol.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	files := make([]responses.File, 0)

	err = cursor.All(c.Context(), &files)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(files)
}

func DownloadFile(c *fiber.Ctx) error {
	idParam := c.Params("id")
	fileID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	
	query := bson.M{"_id": fileID}
	var fileCol *mongo.Collection = configs.Client.GetCollection("file")
	var file models.File
	err = fileCol.FindOne(c.Context(), query).Decode(&file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Set(fiber.HeaderContentType, "application/octet-stream")
	c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", file.Name))

	return c.Send(file.Content)
}

func saveFileMongo(collection *mongo.Collection, name string, contents []byte) error {
	file := models.File{
		Name:     name,
		Content: contents,
	}

	_, err := collection.InsertOne(context.Background(), file)
	return err
}
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khowpchom/golang-upload-file-project/configs"
)

func SecureMiddleware(c *fiber.Ctx) error {
	if secret := c.Get("SECRET-KEY", ""); secret != configs.AppConfig.Secret {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}
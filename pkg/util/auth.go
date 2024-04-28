package util

import "github.com/gofiber/fiber/v2"

func GetUserIDFromHeader(c *fiber.Ctx) string {
	if c.Locals("user_id") == nil {
		return ""
	}

	return c.Locals("user_id").(string)
}

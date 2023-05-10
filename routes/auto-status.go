package routes

import "github.com/gofiber/fiber/v2"

func r_return(c *fiber.Ctx, obj interface{}) error {
	return c.Status(200).JSON(obj)
}
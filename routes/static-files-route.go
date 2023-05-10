package routes

import (
	"testApi/database"
	"testApi/models"

	"github.com/gofiber/fiber/v2"
)

type StaticFiles struct {
	ID    uint
	Title string
	Owner User
}

func CreateResponseStaticFiles(staticFiles models.StaticFiles, user User) StaticFiles {
	return StaticFiles{ID: staticFiles.ID, Title: staticFiles.Title, Owner: user}
}

func CreateOrder(c *fiber.Ctx) error {
	var staticFiles models.StaticFiles

	if err := c.BodyParser(&staticFiles); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&staticFiles)
	return r_return(c, staticFiles)
}

func GetOrders(c *fiber.Ctx) error {
	staticFiles := []models.StaticFiles{}

	database.Database.Db.Find(&staticFiles)

	if len(staticFiles) == 0 {
		return c.Status(400).SendString("No files")
	}

	responseOrders := []StaticFiles{}

	for _, order := range staticFiles {
		user := models.User{}
		database.Database.Db.First(&user, order.Owner)
		responseUser := CreateResponseUser(user)
		responseOrder := CreateResponseStaticFiles(order, responseUser)
		responseOrders = append(responseOrders, responseOrder)
	}

	return r_return(c, responseOrders)
}

package routes

import (
	"testApi/database"
	"testApi/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID          uint
	CreatedAt   time.Time
	Name        string
	Email       string
	StaticFiles []models.StaticFiles
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Name: userModel.Name,
		Email: userModel.Email}
}

func CreateUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	u := database.Database.Db.Create(&user)
	if u.Error != nil {
		erType := ErrorValidator(u.Error)
		return c.Status(erType.Status).JSON(u.Error.Error())
	}
	responseUser := CreateResponseUser(user)

	return r_return(c, responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	u := database.Database.Db.Find(&users)
	if u.Error != nil {
		erType := ErrorValidator(u.Error)
		return c.Status(erType.Status).JSON(u.Error.Error())
	}
	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func TestGet(c *fiber.Ctx) error {
	return c.SendString("It works")
}

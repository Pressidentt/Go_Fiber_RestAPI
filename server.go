package main

import (
	"fmt"
	"log"
	"strconv"
	"testApi/config"
	"testApi/database"
	"testApi/routes"

	"github.com/gofiber/fiber/v2"
)

const V1 = "/api/v1/"

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hey")
}

func userRoutes(app *fiber.App) {
	app.Post(V1+"user", routes.CreateUser)
	app.Get(V1+"user", routes.GetUsers)
	app.Get(V1+"usertest", routes.TestGet)
}

func orderRoutes(app *fiber.App) {
	app.Post(V1+"order", routes.CreateOrder)
	app.Get(V1+"order", routes.GetOrders)
}

func main() {
	// var config env.Config
	conf := config.New()
	port := strconv.Itoa(conf.Ports.PORT)

	app := fiber.New()
	database.ConnectDb()

	app.Get("/", welcome)
	userRoutes(app)
	orderRoutes(app)

	fmt.Println("Server is running on: ", port)
	log.Fatal(app.Listen(":" + port))
}

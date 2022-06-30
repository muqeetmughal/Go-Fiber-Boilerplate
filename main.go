package main

import (
	"golang_boilerplate/database"
	"golang_boilerplate/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}

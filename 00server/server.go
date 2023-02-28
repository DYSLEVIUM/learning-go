package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// function starting with lowercase is private, and function starting with uppercase is public
func getHome(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func main() {
	// := is used for automatic type deduction
	app := fiber.New()
	app.Get("/", getHome)

	log.Fatal(app.Listen(":3000"))

	var str string = "Go server is running!"

	fmt.Println(str)
}

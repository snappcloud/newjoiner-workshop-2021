package main

import (
  "github.com/gofiber/fiber"
  "fmt"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
                fmt.Println("Someone called!!")
		c.Send("Welcome to my awesome app!\n")
	})
	app.Get("/healthz", func(c *fiber.Ctx) {
		c.Status(200)
	})

	app.Listen(3000)
}

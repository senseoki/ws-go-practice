package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()
	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})
	app.Listen(":3001")
}

package api

import (
	"log"

	"github.com/shivanshsharma13/jarvios/daemon/provider"
)

func start() {
	app := fiber.new()

	app.get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":   "running",
			"provider": provider.Current(),
		})
	})

	app.Post("/chat", func(c *fiber.Ctx) error {
		type Request struct {
			Message string `json:"message"`
		}
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "bad request"})
		}

		reply, err := provider.Chat(req.Message)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"reply": reply})
	})

	log.Println("API listening on :7777")
	log.Fatal(app.Listen(":7777"))
}
